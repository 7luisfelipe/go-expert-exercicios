package auction

import (
	"context"
	"testing"
	"time"

	"github.com/7luisfelipe/projetoleilao/internal/entity/auction_entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Estrutura de mock
type MockDatabase struct {
	mock.Mock
}

// Mock para inserir um registro
func (m *MockDatabase) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

// Mock para atualizar registros
func (m *MockDatabase) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (m *MockDatabase) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.SingleResult)
}

func (m *MockDatabase) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}

// Teste funcionalidade para cadastrar anúncio retorna sucesso
func TestCreateAuctionSuccess(t *testing.T) {
	mockRepository := new(MockDatabase)
	auctionRepository := &AuctionRepository{Collection: mockRepository}

	ctx := context.TODO()
	auctionId := "hash_1"

	auction := &auction_entity.Auction{
		Id:          auctionId,
		ProductName: "Curso GoLang - Pós",
		Category:    "Educação",
		Description: "Do zero ao especialista em GO",
		Condition:   1,
		Status:      0,
		Timestamp:   time.Now(),
	}

	auctionEntityMongo := &AuctionEntityMongo{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   auction.Condition,
		Status:      auction.Status,
		Timestamp:   auction.Timestamp.Unix(),
	}

	insertedID := auctionId
	mockRepository.On("InsertOne", ctx, auctionEntityMongo).Return(&mongo.InsertOneResult{InsertedID: insertedID}, nil)

	id, err := auctionRepository.CreateAuction(ctx, auction)
	if err != nil {
		assert.Fail(t, "Erro deveria ser nil")
	}
	assert.Equal(t, insertedID, id)

	mockRepository.AssertExpectations(t)
}

// Testa funcionalidade para fechar o leilão
func TestCloseAuction(t *testing.T) {
	mockRepository := new(MockDatabase)
	auctionRepository := &AuctionRepository{Collection: mockRepository}

	ctx := context.TODO()
	timestamp := time.Now().Unix()
	auctionId := "hash_1"

	mockRepository.On("UpdateMany", ctx, bson.M{"_id": auctionId, "status": 0}, bson.M{"$set": bson.M{"status": 1}}).Return(&mongo.UpdateResult{}, nil)

	go auctionRepository.CloseAuction(ctx, timestamp, auctionId)

	// sleep
	time.Sleep(2 * time.Second)

	mockRepository.AssertCalled(t, "UpdateMany", ctx, bson.M{"_id": auctionId, "status": 0}, bson.M{"$set": bson.M{"status": 1}})
}
