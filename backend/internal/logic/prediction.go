package logic

import (
	"math/rand"
	"time"
)

type PredictionService struct{}

func NewPredictionService() *PredictionService {
	return &PredictionService{}
}

// PredictProfit returns a predicted profit percentage and confidence score
func (p *PredictionService) PredictProfit(symbol string, currentPrice float64) (float64, float64) {
	rand.Seed(time.Now().UnixNano())

	// Mock logic: Random prediction between -5% and +15%
	profitPct := -5.0 + rand.Float64()*(15.0-(-5.0))

	// Mock confidence: Random between 0.5 and 0.99
	confidence := 0.5 + rand.Float64()*(0.99-0.5)

	return profitPct, confidence
}
