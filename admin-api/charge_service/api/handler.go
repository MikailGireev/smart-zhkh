package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"charge_service/internal/httpx"
	"charge_service/internal/logger"
)

var log logger.Log = &logger.Console{}

func readTransactions() ([]Transaction, error) {
	var txs []Transaction

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return txs, nil
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &txs)
	return txs, err
}

func writeTransactions(txs []Transaction) error {
	data, err := json.MarshalIndent(txs, "", "  ")
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}

func CreateTransaction(c *gin.Context) {
	var tx Transaction
	if err := c.ShouldBindJSON(&tx); err != nil {
		log.Errorf("Invalid JSON: %v", err)
		httpx.NewJSONError(c.Writer, http.StatusBadRequest, "Invalid JSON", err.Error())
		return
	}

	txs, err := readTransactions()
	if err != nil {
		log.Errorf("Read error: %v", err)
		httpx.NewJSONError(c.Writer, http.StatusInternalServerError, "Failed to read file", err.Error())
		return
	}

	var maxID uint = 0
	for _, existingTx := range txs {
		if existingTx.ID > maxID {
			maxID = existingTx.ID
		}
	}
	tx.ID = maxID + 1

	txs = append(txs, tx)

	if err := writeTransactions(txs); err != nil {
		log.Errorf("Write error: %v", err)
		httpx.NewJSONError(c.Writer, http.StatusInternalServerError, "Failed to write data", err.Error())
		return
	}

	log.Info("Transaction created: %+v", tx)
	c.JSON(http.StatusOK, tx)
}


func GetTransactions(c *gin.Context) {
	txs, err := readTransactions()
	if err != nil {
		log.Errorf("Read error: %v", err)
		httpx.NewJSONError(c.Writer, http.StatusInternalServerError, "Failed to read file", err.Error())
		return
	}

	var result []Transaction
	for _, tx := range txs {
		match := true

		if v := c.Query("id"); v != "" {
			id, _ := strconv.Atoi(v)
			if tx.ID != uint(id) {
				match = false
			}
		}
		if v := c.Query("user_id"); v != "" && tx.UserID != v {
			match = false
		}
		if v := c.Query("amount"); v != "" {
			a, _ := strconv.ParseFloat(v, 64)
			if tx.Amount != a {
				match = false
			}
		}
		if v := c.Query("date"); v != "" && tx.Date != v {
			match = false
		}
		if v := c.Query("category"); v != "" && tx.Category != v {
			match = false
		}
		if v := c.Query("paid"); v != "" {
			p, _ := strconv.ParseBool(v)
			if tx.Paid != p {
				match = false
			}
		}

		if match {
			result = append(result, tx)
		}
	}

	log.Info("Returned %d filtered transactions", len(result))
	c.JSON(http.StatusOK, result)
}
