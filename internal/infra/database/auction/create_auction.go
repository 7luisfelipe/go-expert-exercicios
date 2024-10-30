package auction

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/7luisfelipe/projetoleilao/configuration/logger"
	"github.com/7luisfelipe/projetoleilao/internal/entity/auction_entity"
	"github.com/7luisfelipe/projetoleilao/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	Id          string                          `bson:"_id"`
	ProductName string                          `bson:"product_name"`
	Category    string                          `bson:"category"`
	Description string                          `bson:"description"`
	Condition   auction_entity.ProductCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus    `bson:"status"`
	Timestamp   int64                           `bson:"timestamp"`
}
type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}

func (ar *AuctionRepository) CreateAuction(
	ctx context.Context,
	auctionEntity *auction_entity.Auction) (any, *internal_error.InternalError) {
	auctionEntityMongo := &AuctionEntityMongo{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.Condition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp.Unix(),
	}
	res, err := ar.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		logger.Error("Error trying to insert auction", err)
		return nil, internal_error.NewInternalServerError("Error trying to insert auction")
	}

	// Inicia a goroutine para fechar o leilão
	go ar.CloseAuction(ctx, auctionEntityMongo.Timestamp, res.InsertedID)

	return res.InsertedID, nil
}

// CloseAuction encerra o leilão após a duração especificada.
func (ar *AuctionRepository) CloseAuction(ctx context.Context, timestamp int64, auctionID any) {
	// Parse a duração do leilão
	auctionDuration := os.Getenv("AUCTION_INTERVAL")
	duration, err := time.ParseDuration(auctionDuration)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Calcula o tempo de encerramento
	closingTime := time.Unix(timestamp, 0).Add(duration)

	// Aguarda até o horário de encerramento
	time.Sleep(time.Until(closingTime))

	// Realiza a atualização no banco de dados para encerrar as propostas
	_, err = ar.Collection.UpdateMany(ctx,
		bson.M{"_id": auctionID, "status": 0},
		bson.M{"$set": bson.M{"status": 1}})

	if err != nil {
		logger.Error("Error trying to close auction", err)
		return
	}

	fmt.Println("Auction closed successfully.")
}
