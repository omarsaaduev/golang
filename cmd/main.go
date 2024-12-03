package main

import (
	"fmt"
	"golang/pkg/logger"
	"log/slog"
)

func main() {
	logg := logger.SetupLogger()
	if logg == nil {
		fmt.Println("Logger is nil")
	}
	slog.SetDefault(logg)

}
