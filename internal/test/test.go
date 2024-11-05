package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/7luisfelipe/projetoleilao/configuration/logger"
	"github.com/7luisfelipe/projetoleilao/internal/infra/database/auction"
	"github.com/7luisfelipe/projetoleilao/internal/usecase/auction_usecase"
	"github.com/7luisfelipe/projetoleilao/internal/usecase/bid_usecase"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	TestAuctionClose()
}

func testGetConn(ctx context.Context) (*mongo.Database, error) {
	mongoURL := "mongodb://admin:admin@localhost:27017/auctions?authSource=admin"
	mongoDatabase := "auctions"

	client, err := mongo.Connect(
		ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		logger.Error("Error trying to connect to mongodb database", err)
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("Error trying to ping mongodb database", err)
		return nil, err
	}

	return client.Database(mongoDatabase), nil
}

// func TestAuctionClose(t *testing.T) {
func TestAuctionClose() {
	fmt.Println()
	fmt.Println("========================")
	fmt.Println("Iniciando teste...")
	fmt.Println()

	//mongoURL := "mongodb://admin:admin@mongodb:27017/auctions?authSource=admin"
	mongoCollection := "auctions"

	createAuctionURL := "http://localhost:8080/auction"
	createBidURL := "http://localhost:8080/bid"

	//Coneta no banco de dados
	// Conectar ao MongoDB
	ctx := context.Background()
	conn, err := testGetConn(ctx)
	if err != nil {
		fmt.Println("Falha ao conectar no MongoDB")
	}

	//Input leilão
	createAuctionInput := auction_usecase.AuctionInputDTO{
		ProductName: "Curso GoLang",
		Category:    "Educação",
		Description: "Do zero ao especialista em GO",
		Condition:   1,
	}

	//Input proposta
	createBidInput := bid_usecase.BidInputDTO{
		UserId:    "a6deb003-ba12-4dae-801f-e2f75a56de47",
		AuctionId: "",
		Amount:    100.00,
	}

	// Converter a estrutura para JSON
	createAuctionJson, err := json.Marshal(createAuctionInput)
	if err != nil {
		fmt.Printf("Erro ao converter input de criar leilão para JSON: %v\n", err)
		return
	}

	// 1. Chamar o serviço que cria um cadastro um leilão
	resp, err := http.Post(createAuctionURL, "application/json", bytes.NewBuffer(createAuctionJson))
	if err != nil {
		fmt.Printf("Falha ao chamar o serviço de criar leilão: %v", err)
	}
	defer resp.Body.Close()

	var auctionId string
	if err := json.NewDecoder(resp.Body).Decode(&auctionId); err != nil {
		fmt.Printf("Falha ao decodificar resposta do de criar leilão: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Esperado status 201, mas retornou status %d", resp.StatusCode)
	} else {
		fmt.Println("Leilão criado com sucesso com ID: ", auctionId)
	}

	//Seta o ID do leilão para criar uma proposta
	createBidInput.AuctionId = auctionId
	// Converter a estrutura para JSON
	createBidJson, err := json.Marshal(createBidInput)
	if err != nil {
		fmt.Printf("Erro ao converter input de criar leilão para JSON: %v\n", err)
		return
	}

	// 2. Chamar o serviço que cria uma proposta
	resp, err = http.Post(createBidURL, "application/json", bytes.NewBuffer(createBidJson))
	if err != nil {
		fmt.Printf("Falha ao chamar o serviço de cadastro de venda: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Esperado status 201, mas obteve %d", resp.StatusCode)
	} else {
		fmt.Println("Proposta cadastada")
	}

	collection := conn.Collection(mongoCollection)

	// 3. Acessar o banco de dados MongoDB filtrando pelo ID retornado do primeiro endpoint
	var auction auction.AuctionEntityMongo
	filter := bson.D{{Key: "_id", Value: auctionId}}
	err = collection.FindOne(context.TODO(), filter).Decode(&auction)
	if err != nil {
		fmt.Printf("Falha ao acessar o banco de dados MongoDB: %v", err)
	}

	// Verificar o status do produto
	fmt.Printf("Status do leilão: %v", auction.Status)

	// Esperar 20 segundos
	fmt.Println()
	fmt.Println("Esperando tempo para encerrar o leilão (20 segundos)")
	time.Sleep(20 * time.Second)

	// 4. Acessar o banco de dados MongoDB novamente para verificar o status após 20 segundos
	err = collection.FindOne(context.TODO(), filter).Decode(&auction)
	if err != nil {
		fmt.Printf("Falha ao acessar o banco de dados MongoDB após 20 segundos: %v", err)
	}

	// Verificar o status do produto
	fmt.Printf("Status do leilão: %v", auction.Status)

	fmt.Println()
	fmt.Println()
	fmt.Println("Testes finalizados")
	fmt.Println("========================")

}
