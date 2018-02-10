package service

import (
	"strconv"

	"github.com/jeremyhahn/tradebot/common"
	"github.com/jeremyhahn/tradebot/dao"
	"github.com/jeremyhahn/tradebot/dto"
	"github.com/jeremyhahn/tradebot/entity"
	"github.com/jeremyhahn/tradebot/mapper"
)

type DefaultUserService struct {
	ctx              *common.Context
	dao              dao.UserDAO
	marketcapService *MarketCapService
	userMapper       mapper.UserMapper
	UserService
}

func NewUserService(ctx *common.Context, dao dao.UserDAO,
	marketcapService *MarketCapService, userMapper mapper.UserMapper) *DefaultUserService {
	return &DefaultUserService{
		ctx:              ctx,
		dao:              dao,
		marketcapService: marketcapService,
		userMapper:       userMapper}
}

func (service *DefaultUserService) CreateUser(user common.User) {
	service.dao.Create(&entity.User{
		Username: user.GetUsername()})
}

func (service *DefaultUserService) GetCurrentUser() (common.User, error) {
	entity, err := service.dao.GetById(service.ctx.GetUser().GetId())
	if err != nil {
		return nil, err
	}
	return service.userMapper.MapUserEntityToDto(entity), nil
}

func (service *DefaultUserService) GetUserById(userId uint) (common.User, error) {
	entity, err := service.dao.GetById(userId)
	if err != nil {
		return nil, err
	}
	return service.userMapper.MapUserEntityToDto(entity), nil
}

func (service *DefaultUserService) GetUserByName(username string) (common.User, error) {
	entity, err := service.dao.GetByName(username)
	if err != nil {
		return nil, err
	}
	return service.userMapper.MapUserEntityToDto(entity), nil
}

func (service *DefaultUserService) GetExchange(user common.User, name string, currencyPair *common.CurrencyPair) common.Exchange {
	daoUser := &entity.User{Id: user.GetId()}
	exchanges := service.dao.GetExchanges(daoUser)
	for _, ex := range exchanges {
		if ex.Name == name {
			return NewExchangeService(service.ctx, dao.NewExchangeDAO(service.ctx)).CreateExchange(user, ex.Name)
		}
	}
	return nil
}

func (service *DefaultUserService) GetExchanges(user common.User, currencyPair *common.CurrencyPair) []common.CryptoExchange {
	var exchangeList []common.CryptoExchange
	var chans []chan common.CryptoExchange
	daoUser := &entity.User{Id: user.GetId()}
	exchanges := service.dao.GetExchanges(daoUser)
	for _, ex := range exchanges {
		c := make(chan common.CryptoExchange, 1)
		chans = append(chans, c)
		exchangeDAO := dao.NewExchangeDAO(service.ctx)
		exchangeService := NewExchangeService(service.ctx, exchangeDAO)
		exchange := exchangeService.CreateExchange(user, ex.Name)
		go func() { c <- exchange.GetExchange() }()
	}
	for i := 0; i < len(exchanges); i++ {
		exchangeList = append(exchangeList, <-chans[i])
	}
	return exchangeList
}

func (service *DefaultUserService) GetWallets(user common.User) []common.CryptoWallet {
	var walletList []common.CryptoWallet
	daoUser := &entity.User{Id: user.GetId()}
	wallets := service.dao.GetWallets(daoUser)
	var chans []chan common.CryptoWallet
	for _, _wallet := range wallets {
		wallet := _wallet
		c := make(chan common.CryptoWallet, 1)
		chans = append(chans, c)
		go func() {
			balance := service.getBalance(wallet.Currency, wallet.Address)
			c <- &dto.CryptoWalletDTO{
				Address:  wallet.Address,
				Currency: wallet.Currency,
				Balance:  balance,
				NetWorth: balance * service.getPrice(wallet.Currency, balance)}
		}()
	}
	for i := 0; i < len(wallets); i++ {
		walletList = append(walletList, <-chans[i])
	}
	return walletList
}

func (service *DefaultUserService) GetWallet(user common.User, currency string) common.CryptoWallet {
	daoUser := &entity.User{Id: user.GetId()}
	wallet := service.dao.GetWallet(daoUser, currency)
	balance := service.getBalance(wallet.GetCurrency(), wallet.GetAddress())
	return &dto.CryptoWalletDTO{
		Address:  wallet.GetAddress(),
		Currency: wallet.GetCurrency(),
		Balance:  balance,
		NetWorth: service.getPrice(wallet.GetCurrency(), balance)}
}

func (service *DefaultUserService) getBalance(currency, address string) float64 {
	service.ctx.Logger.Debugf("[DefaultUserService.getBalance] currency=%s, address=%s", currency, address)
	if currency == "XRP" {
		return NewRipple(service.ctx, service.marketcapService).GetBalance(address).GetBalance()
	} else if currency == "BTC" {
		return NewBlockchainInfo(service.ctx).GetBalance(address).GetBalance()
	}
	return 0.0
}

func (service *DefaultUserService) getPrice(currency string, amt float64) float64 {
	service.ctx.Logger.Debugf("[DefaultUserService.getPrice] currency=%s, amt=%.8f", currency, amt)
	/*
		if currency == "BTC" {
			return util.TruncateFloat(NewBlockchainInfo(service.ctx).GetPrice()*amt, 8)
		}*/
	f, _ := strconv.ParseFloat(service.marketcapService.GetMarket(currency).PriceUSD, 64)
	return f
}
