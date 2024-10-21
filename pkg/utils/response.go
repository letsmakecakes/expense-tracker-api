package utils

import "github.com/gin-gonic/gin"

func RespondWithError(ctx *gin.Context, code int, message string) {
	RespondWithJSON(ctx, code, gin.H{"error": message})
}

func RespondWithJSON(ctx *gin.Context, code int, payload interface{}) {
	ctx.JSON(code, payload)
}
