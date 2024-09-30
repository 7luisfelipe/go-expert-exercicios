package adapter

import "modapilab1/internal/domain/entities"

type IFindWeather interface {
	FindData(cityName string) (*entities.Weather, error)
}
