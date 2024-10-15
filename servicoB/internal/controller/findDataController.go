package controller

import (
	"encoding/json"
	"fmt"
	"modapilab1/internal/domain/usecase"
	"modapilab1/internal/infrastructure/config"
	"modapilab1/internal/infrastructure/viacep"
	"modapilab1/internal/infrastructure/weather"
	"net/http"
)

type FindDataController struct {
	FindDataUseCase usecase.IFindDataUseCase
}

func (c *FindDataController) FindData(w http.ResponseWriter, r *http.Request) {
	//Parâmetro
	zipCodeParam := r.URL.Query().Get("cep")
	if zipCodeParam == "" || len(zipCodeParam) != 8 {
		w.WriteHeader(422)
		w.Write([]byte(config.INVALID_ZIP_CODE))
		return
	}

	c.FindDataUseCase = &usecase.FindData{
		ZipCode: &viacep.ViaCep{},
		Weather: &weather.Weather{},
	}

	result, err := c.FindDataUseCase.FindData(zipCodeParam)
	if err != nil {
		if err.Error() == config.ZIP_CODE_NOT_FOUND {
			w.WriteHeader(404)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err.Error())

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
