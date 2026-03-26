package handler

import (
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {

}

func getUserId(c *gin.Context) (int, error) {

	return 0, nil
}
