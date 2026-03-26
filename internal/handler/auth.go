package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {

}

type signInInput struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash" binding:"required"`
	Username     string `json:"username"`
}

func (h *Handler) signIn(c *gin.Context) {}

func (h *Handler) refreshHandler(c *gin.Context) {}

func (h *Handler) logout(c *gin.Context) {}

func (h *Handler) logoutAll(c *gin.Context) {}
