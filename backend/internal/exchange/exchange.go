package exchange

type OrderSide string

const (
	Buy  OrderSide = "BUY"
	Sell OrderSide = "SELL"
)

type OrderType string

const (
	Market OrderType = "MARKET"
	Limit  OrderType = "LIMIT"
)

type Order struct {
	ID        string
	Symbol    string
	Side      OrderSide
	Type      OrderType
	Price     float64
	Quantity  float64
	Status    string
	Timestamp int64
}

type Ticker struct {
	Symbol string
	Price  float64
}

type Exchange interface {
	Name() string
	GetBalance(asset string) (float64, error)
	GetTicker(symbol string) (*Ticker, error)
	PlaceOrder(symbol string, side OrderSide, orderType OrderType, quantity float64, price float64) (*Order, error)
}
