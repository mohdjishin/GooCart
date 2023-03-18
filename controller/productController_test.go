package controller

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

var p = NewProduct()

func TestSearchProduct(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := p.SearchProduct(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SearchProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
