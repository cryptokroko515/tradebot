package service

import (
	"github.com/jeremyhahn/tradebot/common"
	"github.com/jeremyhahn/tradebot/dao"
)

type TradeServiceImpl struct {
	ctx      *common.Context
	tradeDAO dao.TradeDAO
	TradeService
}

func NewTradeService(ctx *common.Context, tradeDAO dao.TradeDAO) TradeService {
	return &TradeServiceImpl{
		ctx:      ctx,
		tradeDAO: tradeDAO}
}

func (ts *TradeServiceImpl) Buy(exchange common.Exchange, trade common.Trade) {
	ts.ctx.Logger.Debugf("[TradeServiceImpl.Buy] %+v\n", trade)
}

func (ts *TradeServiceImpl) Sell(exchange common.Exchange, trade common.Trade) {
	ts.ctx.Logger.Debugf("[TradeServiceImpl.Sell] %+v\n", trade)
}

func (ts *TradeServiceImpl) Save(trade *common.Trade) {
	ts.tradeDAO.Create(&dao.Trade{
		Id:        trade.Id,
		ChartId:   trade.ChartId,
		Date:      trade.Date,
		Exchange:  trade.Exchange,
		Base:      trade.Base,
		Quote:     trade.Quote,
		Type:      trade.Type,
		Price:     trade.Price,
		Amount:    trade.Amount,
		ChartData: trade.ChartData})
}

func (ts *TradeServiceImpl) GetLastTrade(chart *common.Chart) *common.Trade {
	daoChart := &dao.Chart{Id: chart.Id}
	entity := ts.tradeDAO.GetLastTrade(daoChart)
	return &common.Trade{
		Id:        entity.Id,
		UserId:    ts.ctx.User.Id,
		ChartId:   entity.ChartId,
		Date:      entity.Date,
		Exchange:  entity.Exchange,
		Type:      entity.Type,
		Base:      entity.Base,
		Quote:     entity.Quote,
		Amount:    entity.Amount,
		Price:     entity.Price,
		ChartData: entity.ChartData}
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