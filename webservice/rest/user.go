package rest

import (
	"net/http"

	"github.com/jeremyhahn/tradebot/common"
	"github.com/jeremyhahn/tradebot/dao"
	"github.com/jeremyhahn/tradebot/mapper"
	"github.com/jeremyhahn/tradebot/service"
)

type UserRestService interface {
	GetExchanges(w http.ResponseWriter, r *http.Request)
}

type UserRestServiceImpl struct {
	middlewareService service.Middleware
	jsonWriter        common.HttpWriter
}

func NewUserRestService(middlewareService service.Middleware, jsonWriter common.HttpWriter) UserRestService {
	return &UserRestServiceImpl{
		middlewareService: middlewareService,
		jsonWriter:        jsonWriter}
}

func (restService *UserRestServiceImpl) createUserService(ctx common.Context) service.UserService {
	pluginDAO := dao.NewPluginDAO(ctx)
	userDAO := dao.NewUserDAO(ctx)
	userMapper := mapper.NewUserMapper()
	pluginMapper := mapper.NewPluginMapper()
	userExchangeMapper := mapper.NewUserExchangeMapper()
	marketcapService := service.NewMarketCapService(ctx)
	pluginService := service.NewPluginService(ctx, pluginDAO, pluginMapper)
	exchangeService := service.NewExchangeService(ctx, userDAO, userMapper, userExchangeMapper, pluginService)
	ethereumService, _ := service.NewEthereumService(ctx, userDAO, userMapper, marketcapService, exchangeService)
	return service.NewUserService(ctx, userDAO, userMapper, userExchangeMapper, marketcapService, ethereumService, pluginService)
}

func (restService *UserRestServiceImpl) GetExchanges(w http.ResponseWriter, r *http.Request) {
	ctx, err := restService.middlewareService.CreateContext(w, r)
	if err != nil {
		RestError(w, r, err, restService.jsonWriter)
		return
	}
	defer ctx.Close()
	ctx.GetLogger().Debugf("[UserRestService.GetExchanges]")
	restService.jsonWriter.Write(w, http.StatusOK, common.JsonResponse{
		Success: true,
		Payload: restService.createUserService(ctx).GetConfiguredExchanges()})
}