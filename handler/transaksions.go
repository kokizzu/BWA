package handler

import (
	"BWA/helper"
	"BWA/transactions"
	"BWA/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionsHandler struct {
	service transactions.Service
}

func NewTransaction(service transactions.Service) *transactionsHandler {
	return &transactionsHandler{service}
}

func (h *transactionsHandler) GetTransactions(c *gin.Context) {
	var input transactions.GetTransactionCampaignDetailsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	trans, err := h.service.GetTrasactionByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("get transaction account success", http.StatusOK, "success", transactions.FormatTransactions(trans))
	c.JSON(http.StatusBadRequest, response)
	return

}

//GetUSerTRansaction
//handler
//ambil parameter dari jwt
//keservice
//repo manggil database, jgn lupa preload campaign yang pernah user bayar
