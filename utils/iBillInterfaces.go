package utils

import (
	"github.com/mohdjishin/GoCart/model"
)

type IBillGenerator interface {
	GenerateInvoice(model.Invoice) string
}
