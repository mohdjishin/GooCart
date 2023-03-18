package utils

import (
	"testing"

	"github.com/mohdjishin/GoCart/model"
)

var bill = NewBillGenerator()

func TestGenerateInvoice(t *testing.T) {
	type args struct {
		bill model.Invoice
	}

	tests := []struct {
		name string
		args args
	}{
		struct {
			name string
			args args
		}{name: "test1", args: args{bill: model.Invoice{OrderId: "1", Name: "Jishin", Phone: "1234567890", ProductName: "test", Quantity: "1", Price: "100", Total: "100"}}}, {
			name: "test2", args: args{bill: model.Invoice{OrderId: "2", Name: "Jiseehin", Phone: "1234567890", ProductName: "test", Quantity: "1", Price: "100", Total: "10"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			bill.GenerateInvoice(tt.args.bill)
		})
	}

}
