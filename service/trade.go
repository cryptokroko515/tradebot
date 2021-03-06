package service

import (
	"github.com/jeremyhahn/tradebot/common"
	"github.com/jeremyhahn/tradebot/dao"
	"github.com/jeremyhahn/tradebot/dto"
	"github.com/jeremyhahn/tradebot/entity"
	"github.com/jeremyhahn/tradebot/mapper"
	"github.com/shopspring/decimal"
)

type DefaultTradeService struct {
	ctx         common.Context
	tradeDAO    dao.TradeDAO
	tradeMapper mapper.TradeMapper
	TradeService
}

func NewTradeService(ctx common.Context, tradeDAO dao.TradeDAO, tradeMapper mapper.TradeMapper) TradeService {
	return &DefaultTradeService{
		ctx:         ctx,
		tradeDAO:    tradeDAO,
		tradeMapper: tradeMapper}
}

func (ts *DefaultTradeService) Buy(exchange common.Exchange, trade common.Trade) {
	ts.ctx.GetLogger().Debugf("[DefaultTradeService.Buy] %+v\n", trade)
}

func (ts *DefaultTradeService) Sell(exchange common.Exchange, trade common.Trade) {
	ts.ctx.GetLogger().Debugf("[DefaultTradeService.Sell] %+v\n", trade)
}

func (ts *DefaultTradeService) Save(dto common.Trade) {
	ts.tradeDAO.Create(ts.tradeMapper.MapTradeDtoToEntity(dto))
}

func (ts *DefaultTradeService) GetLastTrade(chart common.Chart) common.Trade {
	daoChart := &entity.Chart{Id: chart.GetId()}
	entity := ts.tradeDAO.GetLastTrade(daoChart)
	amount, err := decimal.NewFromString(entity.GetAmount())
	if err != nil {
		ts.ctx.GetLogger().Errorf("[TradeService.GetLastTrade] Error parsing amount decimal into string. Amount: %s",
			entity.GetAmount())
	}
	price, err := decimal.NewFromString(entity.GetPrice())
	if err != nil {
		ts.ctx.GetLogger().Errorf("[TradeService.GetLastTrade] Error parsing price decimal into string. Price: %s",
			entity.GetPrice())
	}
	return &dto.TradeDTO{
		Id:        entity.GetId(),
		UserId:    ts.ctx.GetUser().GetId(),
		ChartId:   entity.GetChartId(),
		Date:      entity.GetDate(),
		Exchange:  entity.GetExchangeName(),
		Type:      entity.GetType(),
		Base:      entity.GetBase(),
		Quote:     entity.GetQuote(),
		Amount:    amount,
		Price:     price,
		ChartData: entity.GetChartData()}
}

/*
	var trades []dao.Trade
	var indicators []dao.Indicator
	for _, trade := range chart.Trades {
		trades = append(trades, dao.Trade{
			Id:        trade.Id,
			UserId:    trade.UserId,
			ChartId:   chart.Id,
			Date:      trade.Date,
			Exchange:  trade.Exchange,
			Type:      trade.Type,
			Base:      trade.Base,
			Quote:     trade.Quote,
			Amount:    trade.Amount,
			Price:     trade.Price,
			ChartData: trade.ChartData})
	}
	for _, indicator := range chart.Indicators {
		indicators = append(indicators, dao.Indicator{
			Id:         indicator.Id,
			ChartId:    indicator.ChartId,
			Name:       indicator.Name,
			Parameters: indicator.Parameters})
	}
	daoChart := &dao.Chart{
		Id:         chart.Id,
		Base:       chart.Base,
		Exchange:   chart.Exchange,
		Period:     chart.Period,
		Trades:     trades,
		Indicators: indicators}
*/
