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

func Test_getUser(t *testing.T) {
	tests := []struct {
		name   string
		method string
		id     string
		data   string
		status int
	}{
		{
			name:   "get user by id=1, status OK",
			method: "GET",
			id:     "1",
			data:   "{\"id\":1,\"name\":\"Aleksey\",\"age\":56,\"address\":\"Tver\"}\n",
			status: 200,
		},
		{
			name:   "get user by id=2, status OK",
			method: "GET",
			id:     "2",
			data:   "{\"id\":2,\"name\":\"Irina\",\"age\":60,\"address\":\"Tver\"}\n",
			status: 200,
		},
		{
			name:   "get user by id=0, status 406",
			method: "GET",
			id:     "0",
			data:   "",
			status: 406,
		},
		{
			name:   "get POST, status 405",
			method: "POST",
			id:     "1",
			data:   "",
			status: 405,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/user", nil)
			w := httptest.NewRecorder()
			q := req.URL.Query()
			q.Add("id", tt.id)
			req.URL.RawQuery = q.Encode()
			getUser(w, req)
			res := w.Result()
			defer res.Body.Close()
			body := w.Body.String()
			assert.Equal(t, tt.status, res.StatusCode)
			assert.Equal(t, tt.data, body)
		})
	}
}
