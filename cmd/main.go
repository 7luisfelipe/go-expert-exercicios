package main

import (
	"context"
	"encoding/json"
	"fmt"
	"modratelimiter/config"
	"modratelimiter/database"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type RedisData struct {
	Quantity  int       `json:"quantity"`
	Timestamp time.Time `json:"timestamp"`
	IsBlocked bool      `json:"isblocked"`
}

func getLimit(ipLimit int, tokenLimit int, hasToken bool) int {
	if tokenLimit > ipLimit && hasToken {
		return tokenLimit
	}
	return ipLimit
}

func Limiter(db database.StorageContext, ctx context.Context, ipLimit int, tokenLimit int, blockTime int) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ip := r.RemoteAddr

			hasToken := true
			apiKey := r.Header.Get("api_key")

			if apiKey != "" {
				hasToken = false
			}

			result := &RedisData{
				Quantity:  0,
				Timestamp: time.Now(),
			}

			val, err := db.Get(ctx, ip)
			//Se não existir registro ou estiver no intervalo de 10 requisições
			if err == nil {
				err = json.Unmarshal([]byte(val), &result)
				if err != nil {
					panic(err)
				}

				//Se estiver bloqueado e estiver dentro do tempo de bloqueio
				if result.IsBlocked && (start.Sub(result.Timestamp).Seconds()) <= float64(blockTime) {
					http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", 429)
					return
				}

				//Se passou o período para analisar requisições
				if (start.Sub(result.Timestamp).Seconds()) > 10 {
					result = &RedisData{
						Quantity:  0,
						Timestamp: time.Now(),
						IsBlocked: false,
					}
				}
			}

			result.Quantity = result.Quantity + 1
			jsonData, err := json.Marshal(result)
			if err != nil {
				panic(err)
			}

			err = db.Set(ctx, ip, jsonData)
			if err != nil {
				http.Error(w, "Não foi possível salvar dados no redis", http.StatusInternalServerError)
				return
			}

			if result.Quantity >= getLimit(ipLimit, tokenLimit, hasToken) {
				http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", 429)

				result.IsBlocked = true
				jsonData, err := json.Marshal(result)
				if err != nil {
					panic(err)
				}

				err = db.Set(ctx, ip, jsonData)
				if err != nil {
					http.Error(w, "Não foi possível salvar dados no redis", http.StatusInternalServerError)
					return
				}

				return
			}

			// Chama o próximo handler
			next.ServeHTTP(w, r)

			// Registra o tempo que levou para processar a requisição
			fmt.Printf("Método: %s, Caminho: %s, Tempo: %s\n", r.Method, r.URL.Path, time.Since(start))
		})
	}
}

func main() {
	ctx := context.Background()
	conf, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	ipLimit := conf.LimitIp
	tokenLimit := conf.LimitToken
	blockTime := conf.BlockTime

	fmt.Println(conf)

	//rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	rdb := database.NewRedisStorage("localhost:6379")
	storageContext := database.NewStorageContext(rdb)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(Limiter(*storageContext, ctx, ipLimit, tokenLimit, blockTime))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		hasToken := true
		apiKey := r.Header.Get("api_key")

		if apiKey != "" {
			hasToken = false
		}
		val, err := storageContext.Get(ctx, r.RemoteAddr)
		if err != nil {
			http.Error(w, "Não foi possível consultar dados", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Total de chamadas: " + val + "/" + strconv.Itoa(getLimit(conf.LimitIp, conf.LimitToken, hasToken))))
	})

	http.ListenAndServe(":8080", r)
}
