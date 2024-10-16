package viacep

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"modapilab1/internal/domain/entities"
	"modapilab1/internal/infrastructure/config"
	"net/http"
)

// Implements IFindZipCode
type ViaCep struct {
}

func (v *ViaCep) FindData(ctx context.Context, zipcode string) (*entities.ZipCode, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "http://viacep.com.br/ws/"+zipcode+"/json", nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data entities.ZipCode
	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}

	if data.Cep == "" {
		return nil, errors.New(config.ZIP_CODE_NOT_FOUND)
	}

	return &data, nil
}
