package query

import "github.com/dwarukira/findcare/internal/entity"

type ProviderResult struct {
}

func Providers(limit, offset int) (results entity.Providers, err error) {
	err = UnscopedDb().Debug().Table("providers").Select("*").Offset(offset).Limit(limit).Find(&results).Error
	return results, err
}
