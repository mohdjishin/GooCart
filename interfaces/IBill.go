package interfaces

import (
	"github.com/mohdjishin/GoCart/model"
)

type IBillGenerator interface {
	GenerateInvoice(model.Invoice) string
}
