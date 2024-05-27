package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Brasilapi struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type ApiResponse struct {
	Cep     string `json:"cep"`
	Rua     string `json:"rua"`
	Bairro  string `json:"bairro"`
	UF      string `json:"uf"`
	Servico string `json:"servico"`
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	cep := "89031401"

	resultChan := make(chan ApiResponse)

	go ApiViaCep(ctx, resultChan, cep)
	go ApiBrasil(ctx, resultChan, cep)

	select {
	case result := <-resultChan:
		fmt.Println("API: ", result.Servico)
		fmt.Println("CEP: ", result.Cep)
		fmt.Println("Rua: ", result.Rua)
		fmt.Println("Bairro: ", result.Bairro)
		fmt.Println("UF: ", result.UF)
	case <-ctx.Done():
		log.Println("Done...")
	}
}

func ApiViaCep(ctx context.Context, resultChan chan<- ApiResponse, cep string) {
	//time.Sleep(5 * time.Second)
	url := "http://viacep.com.br/ws/" + cep + "/json"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Println("Falha a criar requisição")
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		//Valida se o erro é o timeout do context
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("=======> TimeOut -> Falha ao consultar API, timeout da API excedido - ViaCep")
			println()
		}
		//return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Falha ao ler response")
	}

	var c ViaCep
	err = json.Unmarshal(body, &c)
	if err != nil {
		fmt.Println("Falha decodificar JSON: ", err)
	}

	var cepResponse ApiResponse
	cepResponse.Cep = c.Cep
	cepResponse.Rua = c.Logradouro
	cepResponse.Bairro = c.Bairro
	cepResponse.UF = c.Uf
	cepResponse.Servico = "ViaCep"

	resultChan <- cepResponse

	ctx.Done()
}

func ApiBrasil(ctx context.Context, resultChan chan<- ApiResponse, cep string) {
	//time.Sleep(5 * time.Second)
	url := "https://brasilapi.com.br/api/cep/v1/" + cep
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Println("Falha a criar requisição")
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		//Valida se o erro é o timeout do context
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("=======> TimeOut -> Falha ao consultar API, timeout da API excedido - BrasilApi")
			println()
		}
		//return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Falha ao ler response")
	}

	var c Brasilapi
	err = json.Unmarshal(body, &c)
	if err != nil {
		fmt.Println("Falha decodificar JSON: ", err)
	}

	var cepResponse ApiResponse
	cepResponse.Cep = c.Cep
	cepResponse.Rua = c.Street
	cepResponse.Bairro = c.Neighborhood
	cepResponse.UF = c.State
	cepResponse.Servico = "brasilapi"

	resultChan <- cepResponse

	ctx.Done()
}
