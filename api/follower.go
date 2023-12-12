package api

import (
	"database/sql"
	"net/http"

	db "github.com/Kazukite12/StruckShare/db/sqlc"
	"github.com/Kazukite12/StruckShare/token"
	"github.com/gin-gonic/gin"
)

type createFollowerReq struct {
	FollowedUserID int `json:"followed_user_id" binding:"required,min=1"`
}

func (server *Server) createFollower(ctx *gin.Context) {

	var req createFollowerReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreateFollowerParams{
		UserID:         int64(authPayload.UserID),
		FollowedUserID: int64(req.FollowedUserID),
	}

	getArg := db.GetUserFollowerParams{
		UserID:         int64(authPayload.UserID),
		FollowedUserID: int64(req.FollowedUserID),
	}

	_, err := server.store.GetUserFollower(ctx, getArg)

	if err != nil {
		if err == sql.ErrNoRows {
			follower, err := server.store.CreateFollower(ctx, arg)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}

			ctx.JSON(http.StatusOK, follower)
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	} else {
		server.store.DeleteFollower(ctx, int64(req.FollowedUserID))
		ctx.JSON(http.StatusOK, "berhasil di unfollow")
	}

}

type listFollowerReq struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

type getFollowerReq struct {
	UserID int64 `uri:"id" binding:"required",min=1`
}

func (server *Server) listFollower(ctx *gin.Context) {

	var req listFollowerReq

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var getReq getFollowerReq

	if err := ctx.ShouldBindUri(&getReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListFollowerParams{
		UserID: getReq.UserID,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	follower, err := server.store.ListFollower(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))

	}

	ctx.JSON(http.StatusOK, follower)
}

func (server *Server) countUserTotalFollower(ctx *gin.Context) {

	var req getFollowerReq

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	totalFollower, err := server.store.CountUserTotalFollower(ctx, req.UserID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, totalFollower)
}
