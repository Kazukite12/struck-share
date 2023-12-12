package api

import (
	"database/sql"
	"net/http"

	db "github.com/Kazukite12/StruckShare/db/sqlc"
	"github.com/Kazukite12/StruckShare/token"
	"github.com/gin-gonic/gin"
)

type createPostLikesReq struct {
	PostID int64 `json:"post_id" binding:"required,min=1"`
}

func (server *Server) createPostLikes(ctx *gin.Context) {

	var req createPostLikesReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreatePostLikesParams{
		PostID:    req.PostID,
		PostLiker: authPayload.Username,
	}

	GetArg := db.GetPostLikeParams{
		PostLiker: authPayload.Username,
		PostID:    req.PostID,
	}

	_, err := server.store.GetPostLike(ctx, GetArg)

	if err != nil {
		if err == sql.ErrNoRows {
			postLikes, err := server.store.CreatePostLikes(ctx, arg)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}

			ctx.JSON(http.StatusOK, postLikes)
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	} else {

		server.store.CreatePostUnlikes(ctx, authPayload.Username)
		ctx.JSON(http.StatusOK, "unliked")
	}

}

type createGetPostLikesReq struct {
	PostID int64 `json:"post_id" binding:"required,min=1"`
}

func (server *Server) getPostLikes(ctx *gin.Context) {

	var req createGetPostLikesReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	likes, err := server.store.CountPostTotalLikes(ctx, req.PostID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, likes)
}
