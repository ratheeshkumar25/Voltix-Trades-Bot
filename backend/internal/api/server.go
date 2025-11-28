package api

import (
	"trading_bot/internal/db"
	"trading_bot/internal/exchange"
	"trading_bot/internal/logic"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New())

	api := app.Group("/api")

	api.Post("/login", func(c *fiber.Ctx) error {
		type LoginRequest struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		var req LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}
		// Mock login
		if req.Username == "admin" && req.Password == "password" {
			return c.JSON(fiber.Map{"token": "mock-token"})
		}
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	})

	api.Get("/balance/:exchange", func(c *fiber.Ctx) error {
		exName := c.Params("exchange")
		var ex exchange.Exchange

		switch exName {
		case "binance":
			ex = exchange.NewBinance("key", "secret")
		case "mt5":
			ex = exchange.NewMT5()
		case "ctrader":
			ex = exchange.NewCTrader()
		default:
			return c.Status(400).JSON(fiber.Map{"error": "Unknown exchange"})
		}

		balance, _ := ex.GetBalance("USDT")
		return c.JSON(fiber.Map{"exchange": exName, "balance": balance})
	})

	api.Post("/trade", func(c *fiber.Ctx) error {
		type TradeRequest struct {
			Exchange string  `json:"exchange"`
			Symbol   string  `json:"symbol"`
			Side     string  `json:"side"`
			Quantity float64 `json:"quantity"`
		}
		var req TradeRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		// 1. Get Exchange
		var ex exchange.Exchange
		switch req.Exchange {
		case "binance":
			ex = exchange.NewBinance("key", "secret")
		case "mt5":
			ex = exchange.NewMT5()
		case "ctrader":
			ex = exchange.NewCTrader()
		default:
			return c.Status(400).JSON(fiber.Map{"error": "Unknown exchange"})
		}

		// 2. Get Current Price
		ticker, err := ex.GetTicker(req.Symbol)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to get ticker"})
		}

		// 3. Execute Trade on Exchange
		side := exchange.Buy
		if req.Side == "SELL" {
			side = exchange.Sell
		}

		order, err := ex.PlaceOrder(req.Symbol, side, exchange.Market, req.Quantity, ticker.Price)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		// 4. Get Prediction
		predService := logic.NewPredictionService()
		profitPct, _ := predService.PredictProfit(req.Symbol, ticker.Price)
		predictedProfit := (req.Quantity * ticker.Price) * (profitPct / 100)

		// 5. Save to DB
		trade := db.Trade{
			Exchange: req.Exchange,
			Symbol:   req.Symbol,
			Side:     req.Side,
			Price:    ticker.Price,
			Quantity: req.Quantity,
			Profit:   predictedProfit, // Storing predicted profit for now
			Status:   "FILLED",
		}
		if db.DB != nil {
			db.DB.Create(&trade)
		}

		return c.JSON(fiber.Map{
			"status":            "success",
			"orderId":           order.ID,
			"profit_prediction": predictedProfit,
			"price":             ticker.Price,
		})
	})
}
