package exchange

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Binance struct {
	apiKey    string
	apiSecret string
	balance   float64
}

func NewBinance(apiKey, apiSecret string) *Binance {
	return &Binance{
		apiKey:    apiKey,
		apiSecret: apiSecret,
		balance:   10000.0, // Initial mock balance
	}
}

func (b *Binance) Name() string {
	return "Binance"
}

func (b *Binance) GetBalance(asset string) (float64, error) {
	return b.balance, nil
}

func (b *Binance) GetTicker(symbol string) (*Ticker, error) {
	// Mock ticker
	price := 50000.0 + rand.Float64()*1000
	return &Ticker{
		Symbol: symbol,
		Price:  price,
	}, nil
}

func (b *Binance) PlaceOrder(symbol string, side OrderSide, orderType OrderType, quantity float64, price float64) (*Order, error) {
	cost := quantity * price
	if side == Buy {
		if b.balance < cost {
			return nil, fmt.Errorf("insufficient funds")
		}
		b.balance -= cost
	} else {
		// For sell, we'd check asset balance, but for simplicity in this mock we just add USDT
		b.balance += cost
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
