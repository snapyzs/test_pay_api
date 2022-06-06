package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"test_project_sell/model"
	"time"
)

const (
	successStatus = "SUCCESS"
	failureStatus = "FAILURE"
	errorStatus   = "ERROR"
	abortStatus   = "ABORT"
	newStatus     = "NEW"
)

func randomStatus(t *model.Transaction) {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(10) > 7 {
		t.Status = errorStatus
	}
}

func (t *Handler) GenerateTokenForUse(c *gin.Context) {
	var psi model.PaySystem
	if err := c.BindJSON(&psi); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	token, err := t.service.GenerateToken(psi.PaySystemId)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (t *Handler) CreatePay(c *gin.Context) {
	var trans model.Transaction
	if err := c.BindJSON(&trans); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if trans.Status != newStatus {
		newErrResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	randomStatus(&trans)

	if err := t.service.Transaction.CreatePay(trans); err != nil {
		newErrResponse(c, http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "pay created",
	})

}
func (t *Handler) EditStatusPay(c *gin.Context) {
	var trans model.Transaction
	if err := c.BindJSON(&trans); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	status, err := t.service.Transaction.CheckStatus(trans)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}
	if status == successStatus || status == failureStatus || status == errorStatus || status == abortStatus {
		newErrResponse(c, http.StatusInternalServerError, fmt.Sprintf("status can't change he is %s", status))
		return
	}

	if err := t.service.Transaction.EditStatusPay(trans); err != nil {
		newErrResponse(c, http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "status edit on " + trans.Status,
	})

}
func (t *Handler) CheckPay(c *gin.Context) {
	var trans model.Transaction
	if err := c.BindJSON(&trans); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	status, err := t.service.Transaction.CheckStatus(trans)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})

}
func (t *Handler) GetAllPayUserById(c *gin.Context) {
	var trans model.Transaction
	if err := c.BindJSON(&trans); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	status, err := t.service.Transaction.GetAllPayUserById(trans)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}
	if len(status) < 1 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "empty",
		})
	} else {
		m := make(map[uint]string)
		for _, v := range status {
			m[v.IdTransaction] = v.Status
		}
		c.JSON(http.StatusOK, m)
	}
}
func (t *Handler) GetAllPayUserByEmail(c *gin.Context) {
	var trans model.Transaction
	if err := c.BindJSON(&trans); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	status, err := t.service.Transaction.GetAllPayUserByEmail(trans)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}
	if len(status) < 1 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "empty",
		})
	} else {
		m := make(map[uint]string)
		for _, v := range status {
			m[v.IdTransaction] = v.Status
		}
		c.JSON(http.StatusOK, m)
	}

}
func (t *Handler) CancelPayById(c *gin.Context) {
	var trans model.Transaction
	if err := c.BindJSON(&trans); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	status, err := t.service.Transaction.CheckStatus(trans)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}
	if status == successStatus || status == failureStatus || status == errorStatus || status == abortStatus {
		newErrResponse(c, http.StatusInternalServerError, fmt.Sprintf("status can't cancel he is %s", status))
		return
	}
	trans.Status = "ABORT"
	if err := t.service.Transaction.EditStatusPay(trans); err != nil {
		newErrResponse(c, http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "status " + trans.Status,
	})
}
