package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (server *Server) profileUpload(ctx *gin.Context) {

	file, err := ctx.FormFile("media")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
	}

	filename := time.Now().Format("20060102-150405") + "-" + file.Filename

	err = ctx.SaveUploadedFile(file, "assets/uploads/profilePicture/"+filename)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"image": filename})

}
