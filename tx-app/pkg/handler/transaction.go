package handler

import (
	"net/http"
	"strconv"

	"github.com/4halavandala/transactions/tx-app/domain/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTransaction(c *gin.Context) {
	var input model.Transaction

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Transaction.Create(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type GetAllTransactionsResponse struct {
	Data []model.Transaction `json:"data"`
}

func (h *Handler) GetAllTransactions(c *gin.Context) {
	txLists, err := h.services.Transaction.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetAllTransactionsResponse{
		Data: txLists,
	})
}

func (h *Handler) GetTransactionById(c *gin.Context) {
	txId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	tx, err := h.services.Transaction.GetById(txId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tx)

}
