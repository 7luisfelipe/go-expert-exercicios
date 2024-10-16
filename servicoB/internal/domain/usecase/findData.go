package usecase

import (
	"context"
	"modapilab1/internal/adapter"
	"modapilab1/internal/domain/dto"
)

type IFindDataUseCase interface {
	FindData(ctx context.Context, zipcode string) (*dto.ResultOutpurDto, error)
}

type FindData struct {
	ZipCode adapter.IFindZipCode
	Weather adapter.IFindWeather
}

func (uc *FindData) FindData(ctx context.Context, zipcode string) (*dto.ResultOutpurDto, error) {
	z, err := uc.ZipCode.FindData(ctx, zipcode)
	if err != nil {
		return nil, err
	}

	w, err := uc.Weather.FindData(ctx, z.Localidade)
	if err != nil {
		return nil, err
	}

	k := w.Current.TempC + 273

	result := &dto.ResultOutpurDto{
		Temp_C: w.Current.TempC,
		Temp_F: w.Current.TempF,
		Temp_K: k,
	}

	return result, nil
}
