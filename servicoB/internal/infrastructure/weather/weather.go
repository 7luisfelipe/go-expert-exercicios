package weather

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"modapilab1/internal/domain/entities"
	"modapilab1/internal/infrastructure/config"
	"net/http"
	"net/url"
)

// implements  IFindWeather
type Weather struct {
}

func (f *Weather) FindData(ctx context.Context, cityName string) (*entities.Weather, error) {
	baseURL := "http://api.weatherapi.com/v1/current.json"

	// Codificar o valor do parâmetro de consulta
	cityEscaped := url.QueryEscape(cityName)

	// Construir a URL completa com o parâmetro de consulta
	fullURL := fmt.Sprintf("%s?q=%s", baseURL, cityEscaped)

	req, err := http.NewRequestWithContext(ctx, "GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("key", "b84036d06a91471babe04432243009")

	client := &http.Client{}

	// Fazer a requisição
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data entities.Weather
	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}

	if data.Location.Name == "" {
		return nil, errors.New(config.CITY_NOT_FOUND)
	}

	return &data, nil
}
