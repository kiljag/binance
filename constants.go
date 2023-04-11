package binance

// global enums
type SideType string
type PositionSideType string
type OrderType string
type TimeInForceType string
type NewOrderRespType string
type OrderExecutionType string
type OrderStatusType string
type SymbolType string
type SymbolStatusType string
type SymbolFilterType string
type SideEffectType string
type WorkingType string
type MarginType string
type ContractType string
type UserDataEventType string
type UserDataEventReasonType string
type ForceOrderCloseType string

const (
	SideTypeBuy  SideType = "BUY"
	SideTypeSell SideType = "SELL"

	PositionSideTypeBoth  PositionSideType = "BOTH"
	PositionSideTypeLong  PositionSideType = "LONG"
	PositionSideTypeShort PositionSideType = "SHORT"

	OrderTypeLimit              OrderType = "LIMIT"
	OrderTypeMarket             OrderType = "MARKET"
	OrderTypeStop               OrderType = "STOP"
	OrderTypeStopMarket         OrderType = "STOP_MARKET"
	OrderTypeTakeProfit         OrderType = "TAKE_PROFIT"
	OrderTypeTakeProfitMarket   OrderType = "TAKE_PROFIT_MARKET"
	OrderTypeTrailingStopMarket OrderType = "TRAILING_STOP_MARKET"

	TimeInForceTypeGTC TimeInForceType = "GTC" // Good Till Cancel
	TimeInForceTypeIOC TimeInForceType = "IOC" // Immediate or Cancel
	TimeInForceTypeFOK TimeInForceType = "FOK" // Fill or Kill
	TimeInForceTypeGTX TimeInForceType = "GTX" // Good Till Crossing (Post Only)

	NewOrderRespTypeACK    NewOrderRespType = "ACK"
	NewOrderRespTypeRESULT NewOrderRespType = "RESULT"

	OrderExecutionTypeNew         OrderExecutionType = "NEW"
	OrderExecutionTypePartialFill OrderExecutionType = "PARTIAL_FILL"
	OrderExecutionTypeFill        OrderExecutionType = "FILL"
	OrderExecutionTypeCanceled    OrderExecutionType = "CANCELED"
	OrderExecutionTypeCalculated  OrderExecutionType = "CALCULATED"
	OrderExecutionTypeExpired     OrderExecutionType = "EXPIRED"
	OrderExecutionTypeTrade       OrderExecutionType = "TRADE"

	OrderStatusTypeNew             OrderStatusType = "NEW"
	OrderStatusTypePartiallyFilled OrderStatusType = "PARTIALLY_FILLED"
	OrderStatusTypeFilled          OrderStatusType = "FILLED"
	OrderStatusTypeCanceled        OrderStatusType = "CANCELED"
	OrderStatusTypeRejected        OrderStatusType = "REJECTED"
	OrderStatusTypeExpired         OrderStatusType = "EXPIRED"
	OrderStatusTypeNewInsurance    OrderStatusType = "NEW_INSURANCE"
	OrderStatusTypeNewADL          OrderStatusType = "NEW_ADL"

	SymbolTypeFuture SymbolType = "FUTURE"

	WorkingTypeMarkPrice     WorkingType = "MARK_PRICE"
	WorkingTypeContractPrice WorkingType = "CONTRACT_PRICE"

	SymbolStatusTypePreTrading   SymbolStatusType = "PRE_TRADING"
	SymbolStatusTypeTrading      SymbolStatusType = "TRADING"
	SymbolStatusTypePostTrading  SymbolStatusType = "POST_TRADING"
	SymbolStatusTypeEndOfDay     SymbolStatusType = "END_OF_DAY"
	SymbolStatusTypeHalt         SymbolStatusType = "HALT"
	SymbolStatusTypeAuctionMatch SymbolStatusType = "AUCTION_MATCH"
	SymbolStatusTypeBreak        SymbolStatusType = "BREAK"

	SymbolFilterTypeLotSize          SymbolFilterType = "LOT_SIZE"
	SymbolFilterTypePrice            SymbolFilterType = "PRICE_FILTER"
	SymbolFilterTypePercentPrice     SymbolFilterType = "PERCENT_PRICE"
	SymbolFilterTypeMarketLotSize    SymbolFilterType = "MARKET_LOT_SIZE"
	SymbolFilterTypeMaxNumOrders     SymbolFilterType = "MAX_NUM_ORDERS"
	SymbolFilterTypeMaxNumAlgoOrders SymbolFilterType = "MAX_NUM_ALGO_ORDERS"
	SymbolFilterTypeMinNotional      SymbolFilterType = "MIN_NOTIONAL"

	SideEffectTypeNoSideEffect SideEffectType = "NO_SIDE_EFFECT"
	SideEffectTypeMarginBuy    SideEffectType = "MARGIN_BUY"
	SideEffectTypeAutoRepay    SideEffectType = "AUTO_REPAY"

	MarginTypeIsolated MarginType = "ISOLATED"
	MarginTypeCrossed  MarginType = "CROSSED"

	ContractTypePerpetual ContractType = "PERPETUAL"

	UserDataEventTypeListenKeyExpired    UserDataEventType = "listenKeyExpired"
	UserDataEventTypeMarginCall          UserDataEventType = "MARGIN_CALL"
	UserDataEventTypeAccountUpdate       UserDataEventType = "ACCOUNT_UPDATE"
	UserDataEventTypeOrderTradeUpdate    UserDataEventType = "ORDER_TRADE_UPDATE"
	UserDataEventTypeAccountConfigUpdate UserDataEventType = "ACCOUNT_CONFIG_UPDATE"

	UserDataEventReasonTypeDeposit             UserDataEventReasonType = "DEPOSIT"
	UserDataEventReasonTypeWithdraw            UserDataEventReasonType = "WITHDRAW"
	UserDataEventReasonTypeOrder               UserDataEventReasonType = "ORDER"
	UserDataEventReasonTypeFundingFee          UserDataEventReasonType = "FUNDING_FEE"
	UserDataEventReasonTypeWithdrawReject      UserDataEventReasonType = "WITHDRAW_REJECT"
	UserDataEventReasonTypeAdjustment          UserDataEventReasonType = "ADJUSTMENT"
	UserDataEventReasonTypeInsuranceClear      UserDataEventReasonType = "INSURANCE_CLEAR"
	UserDataEventReasonTypeAdminDeposit        UserDataEventReasonType = "ADMIN_DEPOSIT"
	UserDataEventReasonTypeAdminWithdraw       UserDataEventReasonType = "ADMIN_WITHDRAW"
	UserDataEventReasonTypeMarginTransfer      UserDataEventReasonType = "MARGIN_TRANSFER"
	UserDataEventReasonTypeMarginTypeChange    UserDataEventReasonType = "MARGIN_TYPE_CHANGE"
	UserDataEventReasonTypeAssetTransfer       UserDataEventReasonType = "ASSET_TRANSFER"
	UserDataEventReasonTypeOptionsPremiumFee   UserDataEventReasonType = "OPTIONS_PREMIUM_FEE"
	UserDataEventReasonTypeOptionsSettleProfit UserDataEventReasonType = "OPTIONS_SETTLE_PROFIT"

	ForceOrderCloseTypeLiquidation ForceOrderCloseType = "LIQUIDATION"
	ForceOrderCloseTypeADL         ForceOrderCloseType = "ADL"

	key_SYMBOL      = "symbol"
	key_INTERVAL    = "interval"
	key_LIMIT       = "limit"
	key_STARTTIME   = "startTime"
	key_ENDTIME     = "endTime"
	key_LEVERAGE    = "leverage"
	key_MARGIN_TYPE = "marginType"

	key_TIMESTAMP  = "timestamp"
	key_SIGNATURE  = "signature"
	key_RECVWINDOW = "recvWindow"
	X_MBX_APIKEY   = "X-MBX-APIKEY"

	// future pairs
	BTCUSDT   = "BTCUSDT"
	ETHUSDT   = "ETHUSDT"
	ADAUSDT   = "ADAUSDT"
	MATICUSDT = "MATICUSDT"
	WAVESUSDT = "WAVESUSDT"

	// intervals
	INTERVAL_1MIN  = "1m"
	INTERVAL_5MIN  = "5m"
	INTERVAL_15MIN = "15m"

	// order
	SIDE          = "side"
	POSITION_SIDE = "positionSide"
	ORDER_TYPE    = "type"
	TIME_IN_FORCE = "timeInForce"
	QUANTITY      = "quantity"
	PRICE         = "price"
	STOP_PRICE    = "stopPrice"

	// user data event types
	EVENT_ListenKeyExpired      = "listentKeyExpired"
	EVENT_MARGIN_CALL           = "MARGIN_CALL"
	Event_ACCOUNT_UPDATE        = "ACCOUNT_UPDATE"
	Event_ORDER_TRADE_UPDATE    = "ORDER_TRADE_UPDATE"
	EVENT_ACCOUNT_CONFIG_UPDATE = "ACCOUNT_CONFIG_UPDATE"
)
