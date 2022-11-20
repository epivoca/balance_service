package api

import (
	"fmt"
	"net/http"

	db "github.com/epivoca/balance_service/db/sqlc"
	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) creatеTransfer(ctx *gin.Context) {

	var req transferRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !server.sameAccountCurrency(ctx, req.FromAccountID, req.Currency) {
		return
	}

	if !server.sameAccountCurrency(ctx, req.ToAccountID, req.Currency) {
		return
	}

	if !server.positiveBalance(ctx, req.FromAccountID, req.Amount) {
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) positiveBalance(ctx *gin.Context, accountID int64, amount int64) bool {

	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if account.Balance-amount < 0 {
		err := fmt.Errorf("account [%d] doesn't have enough money to complete this transaction", accountID)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}

	return true
}

// TODO: Valid account
func (server *Server) sameAccountCurrency(ctx *gin.Context, accountID int64, currency string) bool {

	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", account.ID, account.Currency, currency)
		// TODO: Maybe have to implement currency exchange
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}

	return true
}
