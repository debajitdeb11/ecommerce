package main

import (
	"ecommerce/logger"
	"fmt"
	"os"
	"time"
)

func main() {
	logger.Log("***** Executing Main *****")
	Start()
	logger.Log("***** Finishing Main *****")
}

func init() {
	dir, _ := os.Getwd()
	fileName := logger.Path(fmt.Sprint(time.Now().Unix()))
	logger.SetUpLog(logger.Path(dir), fileName, logger.MIXED)
}
