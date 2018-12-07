package handlers

import (
	"github.com/gin-gonic/gin"
	"../../../ethereum"
	"../../../config"
	"net/http"
)

func EthereumBalance(c *gin.Context) {
	response:= ethereum.GetBalance(config.GetVerifier().VerifierPublicKey)
	c.JSON(http.StatusOK, gin.H{
		"balance": response,
	})
}

func PlasmaBalance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"balance": "0",
	})
}

func PlasmaContractAddress(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"address": config.GetVerifier().PlasmaContractAddress,
	})
}


func DepositHandler(c *gin.Context) {
	result := ethereum.Deposit(c.Param("sum"))
	c.JSON(http.StatusOK, gin.H{
		"txHash": result,
	})
}

func TransferHandler(c *gin.Context) {
	address := c.Param("address")
	sum := c.Param("sum")

	c.JSON(http.StatusOK, gin.H{
		"status": address + sum,
	})
}

func ExitHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

var id = 0
var users = 11
var verifiers_amount = 999

func LatestBlockHandler(c *gin.Context) {
	id = id + 1
	c.JSON(http.StatusOK, gin.H{
		"latest_block_number":  id,
	})
}

func VerifiersAmountHandler(c *gin.Context) {
	verifiers_amount = verifiers_amount + 1
	c.JSON(http.StatusOK, gin.H{
		"verifiers_amount": verifiers_amount,
	})
}

func UsersCountHandler(c *gin.Context) {
	users = users + 99
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func TotalBalanceHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"balance": "0",
	})
}

func HistoryAllHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"test": "0",
	})
}

func HistoryTxHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"verifiers_amount": "0",
	})
}
