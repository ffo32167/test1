package main

import (
	"fmt"
	"os"

	"github.com/ffo32167/test1/internal/http"
	"github.com/ffo32167/test1/internal/slice"
	"go.uber.org/zap"
)

func main() {
	log, err := zap.NewProduction()
	if err != nil {
		fmt.Println(fmt.Errorf("cant start logger: %w", err))
	}
	defer func() {
		err = log.Sync()
		if err != nil {
			fmt.Println(fmt.Errorf("cant sync logger: %w", err))
		}
	}()
	storage := slice.New()

	apiServer := http.New(os.Getenv(":80"), log)
	err = apiServer.Run(storage)
	if err != nil {
		log.Error("cant start api server:", zap.Error(err))
	}
}
