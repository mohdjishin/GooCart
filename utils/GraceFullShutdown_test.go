package utils

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestListenAndShutdown(t *testing.T) {
	type args struct {
		app *fiber.App
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "TestListenAndShutdown",
			args: args{
				app: fiber.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ListenAndShutdown(tt.args.app)
		})
	}
}
