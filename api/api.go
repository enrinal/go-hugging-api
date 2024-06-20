package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hupe1980/go-huggingface"
	"go-hugging-api/entitiy"
	"go-hugging-api/hugging"
	"net/http"
	"time"
)

type Api struct {
	Hugging hugging.Hugging
}

func (a *Api) TextClassificationHandler(c *gin.Context) {
	var req huggingface.TextClassificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, entitiy.TextClassificationResponse{
			Status: "error",
			Error:  err.Error(),
		})

		return
	}

	res, err := a.Hugging.TextClassification(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entitiy.TextClassificationResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, entitiy.TextClassificationResponse{
		Data:       *res,
		ServerTime: time.Now().String(),
		Status:     "success",
	})

	return
}
