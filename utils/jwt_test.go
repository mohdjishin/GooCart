package utils

import (
	"testing"

	"gorm.io/gorm"
)

var tkn = NewToken()

func TestRefreshToken(t *testing.T) {
	type args struct {
		db          *gorm.DB
		refresh     string
		accessToken string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
		want2 string
	}{
		// TODO: Add test cases.
		{name: "test 1", args: args{accessToken: "", refresh: "23435235p238235"}, want: "no token found", want1: "", want2: ""},
		{name: "test 2", args: args{accessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzQ5MDgzMzgsInJvbGUiOiJ1c2VyIiwic3ViIjoxfQ.hf9XkMahH-koPtO0xRUqEdXXnAGNbPIGgAZWHZp8osM", refresh: "cd7739c0-e192-43da-a304-0575763c55f3"}, want: "unauthorized", want1: "", want2: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tkn.RefreshToken(tt.args.db, tt.args.refresh, tt.args.accessToken)
			if got != tt.want {

				t.Errorf("RefreshToken() got = %v,  want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("RefreshToken() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("RefreshToken() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
