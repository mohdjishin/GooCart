package controller

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var u = NewUserFunc()

func TestHome(t *testing.T) {

	server := fiber.New()
	// routes
	server.Get("/home", u.Home)

	req, _ := http.NewRequest(http.MethodGet, "/home", nil)
	resp, _ := server.Test(req, -1)

	t.Log(resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
	assert.Equal(t, 200, resp.StatusCode)
}
