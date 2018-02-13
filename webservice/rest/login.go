package rest

import (
	"encoding/json"
	"net/http"

	"github.com/jeremyhahn/tradebot/common"
	"github.com/jeremyhahn/tradebot/service"
)

type LoginRestService interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type LoginRestServiceImpl struct {
	ctx         *common.Context
	authService service.AuthService
	jsonWriter  common.HttpWriter
}

func NewLoginRestService(ctx *common.Context, authService service.AuthService, jsonWriter common.HttpWriter) LoginRestService {
	return &LoginRestServiceImpl{
		ctx:         ctx,
		authService: authService,
		jsonWriter:  jsonWriter}
}

func (service *LoginRestServiceImpl) Login(w http.ResponseWriter, r *http.Request) {
	service.ctx.Logger.Debug("[LoginRestService.Login]")
	var loginRequest LoginRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&loginRequest); err != nil {
		service.ctx.Logger.Errorf("[LoginRestService.Login] %s", err.Error())
		respondWithJSON(w, http.StatusBadRequest, LoginResponse{
			Error: err.Error()})
		return
	}

	service.ctx.Logger.Debugf("[LoginRestService.Login] username: %s, password: %s",
		loginRequest.Username, loginRequest.Password)

	_, err := service.authService.Login(loginRequest.Username, loginRequest.Password)
	var response LoginResponse
	if err != nil {
		response.Error = "Invalid username / password"
		response.Success = false
	} else {
		response.Success = true
	}

	service.jsonWriter.Write(w, http.StatusOK, response)
}