package viacep

import (
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

func (v *ViaCep) FindData(zipcode string) (*entities.ZipCode, error) {
	req, err := http.Get("http://viacep.com.br/ws/" + zipcode + "/json")

	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
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
