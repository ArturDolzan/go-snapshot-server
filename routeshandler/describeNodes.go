package routerhandler

import (
	"net/http"

	"github.com/arturdolzan/go-snapshot-server/domain/dtos"
	"github.com/arturdolzan/go-snapshot-server/presentation"

	"github.com/gin-gonic/gin"
)

func DescribeNodes(ctx *gin.Context) {
	profileRegionDto := dtos.ProfileRegionDto{}

	if err := ctx.BindJSON(&profileRegionDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := profileRegionDto.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	nodes, err := presentation.DescribeNodes(profileRegionDto.Profile, profileRegionDto.Region)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	sendSuccess(ctx, "Request ok", nodes)
}
