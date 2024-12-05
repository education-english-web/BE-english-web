package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/education-english-web/BE-english-web/app/errors"
	"github.com/education-english-web/BE-english-web/app/interface/api/context"
)

// BaseHandler help us respond to client
type BaseHandler struct{}

// GetUserEmail extracts user email from Context
func (h *BaseHandler) GetUserEmail(ctx *gin.Context) (string, error) {
	email := context.GetUserEmail(ctx)
	if email == "" {
		return "", errors.NewErrorUnauthorized()
	}

	return email, nil
}

// GetPhoneNumber extracts phone number from Context
func (h *BaseHandler) GetPhoneNumber(ctx *gin.Context) (string, error) {
	phoneNumber := context.GetPhoneNumber(ctx)
	if phoneNumber == "" {
		return "", errors.NewErrorUnauthorized()
	}

	return phoneNumber, nil
}

// GetNumberPhone extracts phone number from Context
func (h *BaseHandler) GetNumberPhone(ctx *gin.Context) (string, error) {
	phoneNumber := context.GetPhoneNumber(ctx)
	if phoneNumber == "" {
		return "", errors.NewErrorUnauthorized()
	}

	return phoneNumber, nil
}

// GetUserID extracts user id from Context
func (h *BaseHandler) GetUserID(ctx *gin.Context) (uuid.UUID, error) {
	userID := context.GetUserID(ctx)
	if userID == uuid.Nil {
		return uuid.Nil, errors.NewErrorUnauthorized()
	}

	return userID, nil
}
