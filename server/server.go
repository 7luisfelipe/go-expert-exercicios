package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

//# ############## #
//# Banco de dados #
//# ############## #

// Abre a conexão com o banco de dados
func getConn() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./cotacaoDataBase.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

// Configuração de banco de dados, se essa etapa falhar a aplicação não inicia
func configDataBase() error {
	//Abre conexão com o banco de dados, não existir a base de dados ela será criada
	db, err := getConn()
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()

	//Cria a tabela onde as cotações vão ser aramzenadas
	sqlStm := `
		CREATE TABLE IF NOT EXISTS usdbrl (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			code TEXT,
			codein TEXT,
			name TEXT,
			high TEXT,
			low TEXT,
			varBid TEXT,
			pctChange TEXT,
			bid TEXT,
			ask TEXT,
			timestamp TEXT,
			create_date TEXT
		);
	`
	_, err = db.Exec(sqlStm)
	if err != nil {
		log.Fatalf("Falha ao criar banco de dados: %s", err)
		return err
	}
	return nil
}

func salvarCotacaoRepository(ctxDb context.Context, u UsdBrl) error {
	db, err := getConn()
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()

	//Salva a cotação no banco de dados
	stmt, err := db.Prepare(`
		INSERT INTO usdbrl (
				code, 
				codein, 
				name, 
				high, 
				low, 
				varBid, 
				pctChange, 
				bid, 
				ask, 
				timestamp, 
				create_date
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctxDb,
		u.Code,
		u.Codein,
		u.Name,
		u.High,
		u.Low,
		u.VarBid,
		u.PctChange,
		u.Bid,
		u.Ask,
		u.Timestamp,
		u.CreateDate,
	)
	if err != nil {
		//Valida se o erro é o timeout do context
		if ctxDb.Err() == context.DeadlineExceeded {
			fmt.Println("=======> TimeOut -> Falha ao salvar dados, timeout do Banco de Dados excedido")
			println()
		}
		return err
	}
	println("Registro salvo com sucesso")
	return nil
}

// # ############## #
// # Aplicação      #
// # ############## #
// obs: A API retorna todos os dados como string, manter assim facilita o processo para este exercício
type UsdBrl struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type Cotacao struct {
	USDBRL UsdBrl `json:"USDBRL"`
}

type CotacaoResponse struct {
	Dolar string `json:"dolar"`
}

func main() {
	err := configDataBase()
	if err != nil {
		log.Fatal("Falha ao iniciar aplicação, não foi possível inicial o banco de dados")
		fmt.Println(err)
	} else {
		println("Iniciando aplicação...")
		http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
			println()

			//Context
			ctx := context.Background()
			ctxAPI, cancelAPI := context.WithTimeout(ctx, 200*time.Millisecond)
			//ctxAPI, cancelAPI := context.WithTimeout(ctx, 1*time.Microsecond)
			ctxDB, cancelDB := context.WithTimeout(ctx, 10*time.Millisecond)
			defer cancelAPI()
			defer cancelDB()

			cotacao, err := buscarCotacaoDolar(ctxAPI)
			if err != nil {
				log.Fatalf("Falha ao consultar API: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			/*
				//Salva a cotação no banco de dados
				err = salvarCotacaoRepository(ctxDB, cotacao.USDBRL)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			*/
			//Sem tratar o erro da função "salvarCotacaoRepository" vamos retornar o valor da consulta da API igual
			_ = salvarCotacaoRepository(ctxDB, cotacao.USDBRL)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			//Retorna apenas o valor da cotação
			response := CotacaoResponse{Dolar: cotacao.USDBRL.Bid}
			json.NewEncoder(w).Encode(response)
		})
		println("Aplicação server.go iniciada...")
		http.ListenAndServe(":8080", nil)
	}
}

// Consulta a cotação no dolar via API e retorna uma struct com os dados encontrados
func buscarCotacaoDolar(ctxApi context.Context) (*Cotacao, error) {
	req, err := http.NewRequestWithContext(ctxApi, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		//Valida se o erro é o timeout do context
		if ctxApi.Err() == context.DeadlineExceeded {
			fmt.Println("=======> TimeOut -> Falha ao consultar API, timeout da API excedido")
			println()
		}
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var c Cotacao
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
