package api

import (
	"net/http"

	db "github.com/Kazukite12/StruckShare/db/sqlc"
	"github.com/Kazukite12/StruckShare/token"
	"github.com/gin-gonic/gin"
)

type createCommentReq struct {
	PostID  int64  `json:"post_id" binding:"required,min=1"`
	Comment string `json:"comment" binding:"required"`
}

func (server *Server) createComment(ctx *gin.Context) {

	var req createCommentReq

	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreateCommentParams{
		CreatedByUserID: int64(authPayload.UserID),
		PostID:          req.PostID,
		Comment:         req.Comment,
	}

	comment, err := server.store.CreateComment(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return

	}
	ctx.JSON(http.StatusOK, comment)
}

type listCommentReq struct {
	PostID   int64 `json:"post_id" binding:"required,min=1"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listComment(ctx *gin.Context) {
	var req listCommentReq

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCommentParams{
		PostID: req.PostID,
		Limit:  req.PageID,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	comments, err := server.store.ListComment(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, comments)
}
