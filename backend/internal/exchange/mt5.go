package exchange

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type MT5 struct {
	balance float64
}

func NewMT5() *MT5 {
	return &MT5{
		balance: 5000.0,
	}
}

func (m *MT5) Name() string {
	return "MetaTrader 5"
}

func (m *MT5) GetBalance(asset string) (float64, error) {
	return m.balance, nil
}

func (m *MT5) GetTicker(symbol string) (*Ticker, error) {
	price := 1.1000 + rand.Float64()*0.0100
	return &Ticker{
		Symbol: symbol,
		Price:  price,
	}, nil
}

func (m *MT5) PlaceOrder(symbol string, side OrderSide, orderType OrderType, quantity float64, price float64) (*Order, error) {
	cost := quantity * price
	if side == Buy {
		if m.balance < cost {
			return nil, fmt.Errorf("insufficient funds")
		}
		m.balance -= cost
	} else {
		m.balance += cost
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
