package handlers

import (
	"../../blockchain"
	"../../plasmautils/slice"
	"../../transactionManager"
	"encoding/hex"
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

// returns a list of utxos for an address
func GetUtxos(c *gin.Context) {
	addr := c.Param("address")
	utxos, err := Manager.GetUtxosForAddress(addr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, utxos)
}

// returns last plasma block number etc.
func GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"lastBlock": Manager.GetLastBlockNumber(),
	})
}

// returns contract address and abi
func GetConfig(c *gin.Context) {
	// todo not implemented
	c.JSON(http.StatusOK, nil)
}

// returns a list of utxos for an address
func FundAddress(c *gin.Context) {
	addr, _ := hex.DecodeString(c.Param("address")[2:])
	out := blockchain.Output{
		Owner: addr,
		Slice: slice.Slice{Begin: 10, End: 20},
	}
	_, err := Manager.AssembleDepositBlock(out)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, nil)
}
