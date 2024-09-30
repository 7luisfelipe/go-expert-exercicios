package adapter

import "modapilab1/internal/domain/entities"

type IFindZipCode interface {
	FindData(zipcode string) (*entities.ZipCode, error)
}
