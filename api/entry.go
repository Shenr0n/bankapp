package api

import (
	"net/http"

	db "github.com/Shenr0n/bankapp/db/sqlc"
	"github.com/gin-gonic/gin"
)

func (server *Server) getEntry(ctx *gin.Context) {
	var req getIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	entry, err := server.store.GetEntry(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, entry)
}

type listEntriesRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAccountEntries(ctx *gin.Context) {
	var reqID getAccountRequest
	var req listEntriesRequest
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListEntriesParams{
		AccountID: reqID.ID,
		Limit:     req.PageSize,
		Offset:    (req.PageID - 1) * req.PageSize,
	}
	entries, err := server.store.ListEntries(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, entries)
}
