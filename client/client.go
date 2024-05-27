package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Dolar struct {
	Dolar string `json:"dolar"`
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Println("Falha ao criar requisição")
		fmt.Println(err)
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		//Valida se o erro é o timeout do context
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("=======> TimeOut -> Falha ao consultar API, timeout do client excedido")
			println()
		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Falha ao ler dados")
		fmt.Println(err)
	}

	var d Dolar
	err = json.Unmarshal(body, &d)
	if err != nil {
		fmt.Println("Falha ao ler dados")
		fmt.Println(err)
	}

	//Salva o arquivo
	//"cotacao.txt" no formato: Dólar: {valor}
	f, err := os.Create("cotacao.txt")
	if err != nil {
		fmt.Println("Falha ao criar arquivo txt")
		fmt.Println(err)
	}

	data := "Dólar: " + d.Dolar
	_, err = f.Write([]byte(data))
	if err != nil {
		fmt.Println("Falha escrever dados no arquivo.txt")
		fmt.Println(err)
	}

	fmt.Println("Resultado da cotação do dolar:")
	fmt.Println(d)
}
