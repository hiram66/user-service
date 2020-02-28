package controllers

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/hiram66/user-service/cmd/web/internal"
	"github.com/hiram66/user-service/configs"
	userService "github.com/hiram66/user-service/pkg/user"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	user := new(userService.User)
	e := jsoniter.NewDecoder(r.Body).Decode(user)
	if e != nil {
		jsonError := internal.GetJsonError(e)
		writeResp(w, internal.WebResp{JsonBody: jsonError, StatusCode: http.StatusInternalServerError})
		return
	}
	ctx, _ := context.WithTimeout(r.Context(), configs.HttpTimeout)
	webResp := internal.AddUser(ctx, *user)
	writeResp(w, webResp)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["userId"]
	ctx, _ := context.WithTimeout(r.Context(), configs.HttpTimeout)
	webResp := internal.GetUserById(ctx, id)
	writeResp(w, webResp)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	q := new(userService.UserQuery)
	e := jsoniter.NewDecoder(r.Body).Decode(q)
	if e != nil {
		jsonError := internal.GetJsonError(e)
		writeResp(w, internal.WebResp{JsonBody: jsonError, StatusCode: http.StatusInternalServerError})
		return
	}
	ctx, _ := context.WithTimeout(r.Context(), configs.HttpTimeout)
	webResp := internal.GetUser(ctx, *q)
	writeResp(w, webResp)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := new(userService.User)
	e := jsoniter.NewDecoder(r.Body).Decode(user)
	if e != nil {
		jsonError := internal.GetJsonError(e)
		writeResp(w, internal.WebResp{JsonBody: jsonError, StatusCode: http.StatusInternalServerError})
		return
	}
	id := mux.Vars(r)["userId"]
	ctx, _ := context.WithTimeout(r.Context(), configs.HttpTimeout)
	webResp := internal.UpdateUser(ctx, id, *user)
	writeResp(w, webResp)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["userId"]
	ctx, _ := context.WithTimeout(r.Context(), configs.HttpTimeout)
	webResp := internal.DeleteUser(ctx, id)
	writeResp(w, webResp)
}
