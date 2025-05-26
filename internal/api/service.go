package api

import (
	"github.com/biryanim/leadgenmarket_tz/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Implementation struct {
	encodeService service.EncodeService
}

func NewImplementation(encodeService service.EncodeService) *Implementation {
	return &Implementation{encodeService: encodeService}
}

func (i *Implementation) EncodeText(c *gin.Context) {
	var req TextRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	encodedText, err := i.encodeService.EncodeText(c.Request.Context(), req.ToModel())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var resp EncodedTextResponse
	resp.FromModel(encodedText)

	c.JSON(http.StatusOK, resp)
}
