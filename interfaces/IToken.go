package interfaces

import "gorm.io/gorm"

type IToken interface {
	GenJwtToken(string, uint, int) (string, string)
	RefreshToken(*gorm.DB, string, string) (string, string, string)
	AdminRefreshToken(*gorm.DB, string, string) (string, string, string)
}
