package comm

//订单状态
const (
	//OrderTypeOrdered 已下单
	OrderTypeOrdered = 0

	//OrderTypeOrdering 交易中
	OrderTypeOrdering = 1

	//OrderTypeSuccess 完成
	OrderTypeSuccess = 2

	//OrderTypeFailed 失败
	OrderTypeFailed = 3

	//OrderTypeCancel 取消
	OrderTypeCancel = 4
)

const (
	PubStrategy    = "strategyPub"
	CancelStrategy = "strategyCancel"
	PubOrdering    = "ordering"
)
const (
	OrderSideBuy  = "BUY"
	OrderSideSell = "SELL"

	OrderTypeLIMIT  = "LIMIT"
	OrderSideMARKET = "MARKET"

	OrderTimeInForceGTC         = "GTC"
	OrderNewOrderRespTypeRESULT = "RESULT"
)

const (
	OrderStatusNEW             = "NEW"
	OrderStatusPARTIALLYFILLED = "PARTIALLY_FILLED"
	OrderStatusFILLED          = "FILLED"
	OrderStatusCANCELED        = "CANCELED"
	OrderStatusREJECTED        = "REJECTED"
	OrderStatusEXPIRED         = "EXPIRED"
)
