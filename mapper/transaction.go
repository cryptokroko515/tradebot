package mapper

import (
	"strings"

	"github.com/jeremyhahn/tradebot/common"
	"github.com/jeremyhahn/tradebot/dto"
	"github.com/jeremyhahn/tradebot/entity"
	"github.com/jeremyhahn/tradebot/viewmodel"
)

type TransactionMapper interface {
	MapTransactionEntityToDto(entity entity.TransactionEntity) common.Transaction
	MapTransactionDtoToEntity(dto common.Transaction) entity.TransactionEntity
	MapTransactionDtoToViewModel(dto common.Transaction) viewmodel.Transaction
}

type DefaultTransactionMapper struct {
	ctx common.Context
	TransactionMapper
}

func NewTransactionMapper(ctx common.Context) TransactionMapper {
	return &DefaultTransactionMapper{ctx: ctx}
}

func (mapper *DefaultTransactionMapper) MapTransactionEntityToDto(entity entity.TransactionEntity) common.Transaction {
	marketPair, _ := common.NewCurrencyPair(entity.GetMarketPair(), mapper.ctx.GetUser().GetLocalCurrency())
	currencyPair, _ := common.NewCurrencyPair(entity.GetCurrencyPair(), mapper.ctx.GetUser().GetLocalCurrency())
	return &dto.TransactionDTO{
		Id:                     entity.GetId(),
		Network:                entity.GetNetwork(),
		NetworkDisplayName:     entity.GetNetworkDisplayName(),
		Date:                   entity.GetDate(),
		Type:                   entity.GetType(),
		Category:               entity.GetCategory(),
		MarketPair:             marketPair,
		CurrencyPair:           currencyPair,
		Quantity:               entity.GetQuantity(),
		QuantityCurrency:       entity.GetQuantityCurrency(),
		FiatQuantity:           entity.GetFiatQuantity(),
		FiatQuantityCurrency:   entity.GetFiatQuantityCurrency(),
		Price:                  entity.GetPrice(),
		PriceCurrency:          entity.GetPriceCurrency(),
		FiatPrice:              entity.GetFiatPrice(),
		FiatPriceCurrency:      entity.GetFiatPriceCurrency(),
		QuoteFiatPrice:         entity.GetQuoteFiatPrice(),
		QuoteFiatPriceCurrency: entity.GetQuoteFiatPriceCurrency(),
		Fee:               entity.GetFee(),
		FeeCurrency:       entity.GetFeeCurrency(),
		FiatFee:           entity.GetFiatFee(),
		FiatFeeCurrency:   entity.GetFiatFeeCurrency(),
		Total:             entity.GetTotal(),
		TotalCurrency:     entity.GetTotalCurrency(),
		FiatTotal:         entity.GetFiatTotal(),
		FiatTotalCurrency: entity.GetFiatTotalCurrency(),
		Deleted:           entity.IsDeleted()}
}

func (mapper *DefaultTransactionMapper) MapTransactionDtoToEntity(dto common.Transaction) entity.TransactionEntity {
	userId := mapper.ctx.GetUser().GetId()
	deleted := 0
	if dto.IsDeleted() {
		deleted = 1
	}
	return &entity.Transaction{
		Id:                     dto.GetId(),
		UserId:                 userId,
		Date:                   dto.GetDate(),
		MarketPair:             dto.GetMarketPair().String(),
		CurrencyPair:           dto.GetCurrencyPair().String(),
		Type:                   dto.GetType(),
		Category:               dto.GetCategory(),
		Network:                dto.GetNetwork(),
		NetworkDisplayName:     dto.GetNetworkDisplayName(),
		Quantity:               dto.GetQuantity(),
		QuantityCurrency:       dto.GetQuantityCurrency(),
		FiatQuantity:           dto.GetFiatQuantity(),
		FiatQuantityCurrency:   dto.GetFiatQuantityCurrency(),
		Price:                  dto.GetPrice(),
		PriceCurrency:          dto.GetPriceCurrency(),
		FiatPrice:              dto.GetFiatPrice(),
		FiatPriceCurrency:      dto.GetFiatPriceCurrency(),
		QuoteFiatPrice:         dto.GetQuoteFiatPrice(),
		QuoteFiatPriceCurrency: dto.GetQuoteFiatPriceCurrency(),
		Fee:               dto.GetFee(),
		FeeCurrency:       dto.GetFeeCurrency(),
		FiatFee:           dto.GetFiatFee(),
		FiatFeeCurrency:   dto.GetFiatFeeCurrency(),
		Total:             dto.GetTotal(),
		TotalCurrency:     dto.GetTotalCurrency(),
		FiatTotal:         dto.GetFiatTotal(),
		FiatTotalCurrency: dto.GetFiatTotalCurrency(),
		Deleted:           deleted}
}

func (mapper *DefaultTransactionMapper) MapTransactionDtoToViewModel(dto common.Transaction) viewmodel.Transaction {
	return viewmodel.Transaction{
		Id:                     dto.GetId(),
		Network:                dto.GetNetwork(),
		NetworkDisplayName:     dto.GetNetworkDisplayName(),
		Date:                   dto.GetDate().Format(common.TIME_DISPLAY_FORMAT),
		Type:                   strings.Title(dto.GetType()),
		Category:               strings.Title(dto.GetCategory()),
		CurrencyPair:           dto.GetMarketPair(),
		Quantity:               dto.GetQuantity(),
		QuantityCurrency:       dto.GetQuantityCurrency(),
		FiatQuantity:           dto.GetFiatQuantity(),
		FiatQuantityCurrency:   dto.GetFiatQuantityCurrency(),
		Price:                  dto.GetPrice(),
		PriceCurrency:          dto.GetPriceCurrency(),
		FiatPrice:              dto.GetFiatPrice(),
		FiatPriceCurrency:      dto.GetFiatPriceCurrency(),
		QuoteFiatPrice:         dto.GetQuoteFiatPrice(),
		QuoteFiatPriceCurrency: dto.GetQuoteFiatPriceCurrency(),
		Fee:               dto.GetFee(),
		FeeCurrency:       dto.GetFeeCurrency(),
		FiatFee:           dto.GetFiatFee(),
		FiatFeeCurrency:   dto.GetFiatFeeCurrency(),
		Total:             dto.GetTotal(),
		TotalCurrency:     dto.GetTotalCurrency(),
		FiatTotal:         dto.GetFiatTotal(),
		FiatTotalCurrency: dto.GetFiatTotalCurrency()}
}
