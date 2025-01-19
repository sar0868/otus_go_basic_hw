package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	data := string("[{\"id\":1,\"name\":\"Aleksey\",\"age\":56,\"address\":\"Tver\"}," +
		"{\"id\":2,\"name\":\"Irina\",\"age\":60,\"address\":\"Tver\"}," +
		"{\"id\":3,\"name\":\"Maria\",\"age\":27,\"address\":\"Tver\"}]\n")

	getUsers(w, req)

	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK, got %v", res.StatusCode)
	}
	body := w.Body.String()

	assert.Equal(t, data, body)
}
