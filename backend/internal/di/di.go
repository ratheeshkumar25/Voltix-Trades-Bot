package di

import "github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/config"

func Init() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
}
