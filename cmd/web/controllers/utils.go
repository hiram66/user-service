package controllers

import (
	"github.com/hiram66/user-service/cmd/web/internal"
	"net/http"
)

func writeResp(w http.ResponseWriter, resp internal.WebResp) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	if resp.JsonBody != nil {
		w.Write(resp.JsonBody)
	}
}
