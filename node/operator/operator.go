package main

import (
	"../config"
	"../transactionManager"
	"./handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Assemble block ~ each second
	manager := transactionManager.NewTransactionManager()
	transactionManager.NewBlockPublisher(manager)
	transactionManager.NewEventMonitor(manager)

	r := gin.New()
	r.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)
	handlers.Manager = manager // todo refactor this

	r.POST("/tx", handlers.PostTransaction)
	r.GET("/config", handlers.GetConfig)
	r.GET("/status", handlers.GetStatus)
	r.GET("/utxo/:address", handlers.GetUtxos)

	// debug hendlers
	r.GET("/fund/:address", handlers.FundAddress)

	err := r.Run(fmt.Sprintf(":%d", config.GetOperator().OperatorPort))
	if err != nil {
		log.Fatal(err)
	}

	println("Operator started")
}
