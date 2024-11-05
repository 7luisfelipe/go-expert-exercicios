package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"modapilab1/internal/domain/usecase"
	"modapilab1/internal/infrastructure/config"
	"modapilab1/internal/infrastructure/viacep"
	"modapilab1/internal/infrastructure/weather"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type FindDataController struct {
	FindDataUseCase usecase.IFindDataUseCase
}

func (c *FindDataController) FindData(w http.ResponseWriter, r *http.Request) {
	// Criar um novo span para a requisição
	log.Println("Iniciando span para FindData")
	tr := otel.Tracer("FindDataTracer")
	ctx, span := tr.Start(r.Context(), "FindData")
	//_, span := tr.Start(r.Context(), "FindData")
	//defer span.End()
	defer func() {
		log.Println("Encerrando span para FindData")
		span.End()
	}()

	//Parâmetro
	zipCodeParam := r.URL.Query().Get("cep")
	span.SetAttributes(attribute.String("zipCode", zipCodeParam))
	if zipCodeParam == "" || len(zipCodeParam) != 8 {
		w.WriteHeader(422)
		w.Write([]byte(config.INVALID_ZIP_CODE))
		return
	}

	c.FindDataUseCase = &usecase.FindData{
		ZipCode: &viacep.ViaCep{},
		Weather: &weather.Weather{},
	}

	result, err := c.FindDataUseCase.FindData(ctx, zipCodeParam)
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
