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
	_, err := transactionManager.NewEventMonitor(manager)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)
	handlers.Manager = manager // todo refactor this

	r.POST("/tx", handlers.PostTransaction)
	r.GET("/config", handlers.GetConfig)
	r.GET("/status", handlers.GetStatus)
	r.GET("/utxo/:address", handlers.GetUtxos)

	// debug handlers
	r.GET("/test/fund/:address", handlers.FundAddress)
	r.GET("/test/transfer/:block/:tx/:out/:address/:key", handlers.Transact)

	err = r.Run(fmt.Sprintf(":%d", config.GetOperator().OperatorPort))
	if err != nil {
		log.Fatal(err)
	}

	println("Operator started")
}
