package adapter

import (
	"context"
	"modapilab1/internal/domain/entities"
)

type IFindWeather interface {
	FindData(ctx context.Context, cityName string) (*entities.Weather, error)
}
