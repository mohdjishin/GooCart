package controller

import (
	"github.com/mohdjishin/GoCart/model"
)

type IBillGenerator interface {
	GenerateInvoice(model.Invoice) string
}
