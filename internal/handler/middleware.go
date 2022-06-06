package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	auth     = "Authorization"
	paySysId = "pay_system_id"
)

func (h *Handler) checkAuth(c *gin.Context) {
	header := c.GetHeader(auth)
	if header == "" {
		newErrResponse(c, http.StatusUnauthorized, "empty header auth")
		return
	}
	headerPair := strings.Split(header, " ")
	if len(headerPair) != 2 {
		newErrResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	paySystemId, err := h.service.ParseToken(headerPair[1])
	if err != nil {
		newErrResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(paySysId, paySystemId)

}
