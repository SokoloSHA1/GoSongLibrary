package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type signInInput struct {
	Ip    string `json:"ip" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func (h *Handler) singIn(c *gin.Context) {
	var input signInInput

	guid := c.Query("guid")

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Aunthorization.GenerateToken(input.Ip)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	token_refresh, err := h.services.GenerateRefreshToken(guid, input.Email, input.Ip)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token":         token,
		"refresh_token": token_refresh,
	})
}

type refreshInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
	Ip           string `json:"ip" binding:"required"`
}

func (h *Handler) refresh(c *gin.Context) {
	var input refreshInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user_token, err := h.services.Aunthorization.CheckRefreshToken(input.RefreshToken)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if user_token.Ip != input.Ip {
		NewErrorResponse(c, http.StatusInternalServerError, "invalid ip, check email")

		err := h.services.Aunthorization.SendMail()
		if err != nil {
			NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		return
	}

	token, err := h.services.Aunthorization.GenerateToken(user_token.Ip)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
