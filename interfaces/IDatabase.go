package interfaces

import (
	"gorm.io/gorm"
)

type IDatabase interface {
	CloseDb(*gorm.DB)
	OpenDb() *gorm.DB
}
