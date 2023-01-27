package utils

import (
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestCheckComplexityOFPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "short", want: false, args: args{password: "123qwe"}},
		{name: "strong", want: true, args: args{password: "123Awe!@#~"}},
		{name: "long but not complex", want: false, args: args{password: "12345678910"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckComplexityOFPassword(tt.args.password); got != tt.want {
				t.Errorf("CheckComplexityOFPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
