package main

import (
	"fmt"
	"net/http"
	"testing"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, client")
// }

// func TestHandler(t *testing.T) {
// req := httptest.NewRequest("GET", "/users", nil)
// w := httptest.NewRecorder()
// getUsers(w, req)

// res := w.Result()
// if res.StatusCode != http.StatusOK {
// 	t.Errorf("expected status 200, got %v", res.StatusCode)
// }
// }

func Test_getUsers(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getUsers(tt.args.w, tt.args.r)
			fmt.Print(t)
		})
	}
}
