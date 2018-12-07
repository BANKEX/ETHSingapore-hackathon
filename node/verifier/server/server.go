package server

import (
	"./handlers"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func GinServer() {

	r := gin.Default()
	//r := gin.New()
	//r.Use(gin.Recovery())
	r.Use(cors.Default())
	//gin.SetMode(gin.ReleaseMode)

	r.GET("/etherBalance", handlers.EthereumBalance)
	r.GET("/plasmaBalance", handlers.PlasmaBalance) //
	r.GET("/contractAddress", handlers.PlasmaContractAddress)
	r.GET("/deposit/:sum", handlers.DepositHandler)
	r.GET("/transfer/:address/:sum", handlers.TransferHandler) //
	r.GET("/exit", handlers.ExitHandler) //
	r.GET("/latestBlock", handlers.LatestBlockHandler) //
	r.GET("/verifiersAmount", handlers.VerifiersAmountHandler) //
	r.GET("/usersCount", handlers.UsersCountHandler) //
	r.GET("/totalBalance", handlers.TotalBalanceHandler) //
	r.GET("/history", handlers.HistoryAllHandler) //
	r.GET("/history/:txHash", handlers.HistoryTxHandler) //

	r.Run(":8080")
}
