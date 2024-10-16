package adapter

import (
	"context"
	"modapilab1/internal/domain/entities"
)

type IFindZipCode interface {
	FindData(ctx context.Context, zipcode string) (*entities.ZipCode, error)
}
