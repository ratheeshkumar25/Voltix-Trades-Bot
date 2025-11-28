package exchange

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type CTrader struct {
	balance float64
}

func NewCTrader() *CTrader {
	return &CTrader{
		balance: 7500.0,
	}
}

func (c *CTrader) Name() string {
	return "cTrader"
}

func (c *CTrader) GetBalance(asset string) (float64, error) {
	return c.balance, nil
}

func (c *CTrader) GetTicker(symbol string) (*Ticker, error) {
	price := 150.0 + rand.Float64()*5.0
	return &Ticker{
		Symbol: symbol,
		Price:  price,
	}, nil
}

func (c *CTrader) PlaceOrder(symbol string, side OrderSide, orderType OrderType, quantity float64, price float64) (*Order, error) {
	cost := quantity * price
	if side == Buy {
		if c.balance < cost {
			return nil, fmt.Errorf("insufficient funds")
		}
		c.balance -= cost
	} else {
		c.balance += cost
	}

	return &Order{
		ID:        uuid.New().String(),
		Symbol:    symbol,
		Side:      side,
		Type:      orderType,
		Price:     price,
		Quantity:  quantity,
		Status:    "FILLED",
		Timestamp: time.Now().Unix(),
	}, nil
}
