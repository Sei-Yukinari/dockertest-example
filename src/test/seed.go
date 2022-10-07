package test

import (
	"dockertest-example/src/domain/model"

	"gorm.io/gorm"
)

type Model interface {
	model.User | int
}

func Seeds(db *gorm.DB, seeds []interface{}) error {
	if db == nil {
		return nil
	}
	if seeds == nil {
		return nil
	}
	for _, s := range seeds {
		if err := db.Create(s).Error; err != nil {
			return err
		}
	}
	return nil
}
