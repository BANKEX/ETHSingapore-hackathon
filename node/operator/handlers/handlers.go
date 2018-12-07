package handlers

import (
	"../../blockchain"
	"../../transactionManager"
	"github.com/gin-gonic/gin"
	"net/http"
)

// todo refactor this
var (
	Manager *transactionManager.TransactionManager
)

func PostTransaction(c *gin.Context) {
	var t blockchain.Transaction
	err := c.BindJSON(&t)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err = Manager.SubmitTransaction(&t)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func GetUtxos(c *gin.Context) {
	addr := c.Param("address")
	utxos, err := Manager.GetUtxosForAddress(addr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, utxos)
}
