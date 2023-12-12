package api

import (
	"database/sql"
	"net/http"

	db "github.com/Kazukite12/StruckShare/db/sqlc"
	"github.com/Kazukite12/StruckShare/token"
	"github.com/gin-gonic/gin"
)

type createPostReq struct {
	Caption string `json:"caption" binding:"required"`
}

func (server *Server) createPost(ctx *gin.Context) {
	var req createPostReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreatePostParams{
		CreatedByUserID: int64(authPayload.UserID),
		Caption:         req.Caption,
	}

	post, err := server.store.CreatePost(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, post)
}

type getPostReq struct {
	ID int64 `uri:"id" binding:"required",min=1`
}

func (server *Server) getPost(ctx *gin.Context) {
	var req getPostReq

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	post, err := server.store.GetPost(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, post)
}

type listPostReq struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listPost(ctx *gin.Context) {
	var req listPostReq

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.ListPostParams{
		CreatedByUserID: int64(authPayload.UserID),
		Limit:           req.PageSize,
		Offset:          (req.PageID - 1) * req.PageSize,
	}

	posts, err := server.store.ListPost(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, posts)
}
