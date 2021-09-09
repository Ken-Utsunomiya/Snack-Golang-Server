package validators

import (
	"Snack-Golang-Server/src/models"
	"time"
)

type SnackRegisterRequest struct {
	SnackName	string	`json:"snack_name" binding:"required"`
	SnackTypeID	int	`json:"snack_type_id" binding:"required"`
	Description	string	`json:"description" binding:"required"`
	ImageURI	string	`json:"image_uri" binding:"required"`
	Quantity int `json:"quantity" binding:"min=0"`
	Price	int	`json:"price" binding:"required"`
	IsActive	bool	`json:"is_active" binding:"required"`
	OrderThreshold	*int	`json:"order_threshold"`
	ExpirationDTM 	*time.Time 	`json:"expiration_dtm"`
	LastUpdatedBy	string	`json:"last_updated_by" binding:"required"`
}

type SnackUpdateRequest struct {
	SnackName	*string	`json:"snack_name"`
	SnackTypeID	*int	`json:"snack_type_id"`
	Description	*string	`json:"description"`
	ImageURI	*string	`json:"image_uri"`
	Price	*int	`json:"price"`
	IsActive	*bool	`json:"is_active"`
	OrderThreshold	*int	`json:"order_threshold"`
	LastUpdatedBy	*string	`json:"last_updated_by"`
}

func RegisterRequestToSnackModel(request SnackRegisterRequest) models.Snack {
	snack := models.Snack{}
	snack.Name = request.SnackName
	snack.SnackTypeID = request.SnackTypeID
	snack.Description = request.Description
	snack.ImageURI = request.ImageURI
	snack.Price = request.Price
	snack.IsActive = request.IsActive
	snack.LastUpdatedBy = request.LastUpdatedBy
	if request.OrderThreshold != nil {
		snack.OrderThreshold = request.OrderThreshold
	}
	snack.LastUpdatedDTM = time.Now()
	return snack
}

func UpdateRequestToSnackModel(request SnackUpdateRequest, snack *models.Snack) {
	if request.SnackName != nil {
		snack.Name = *request.SnackName
	}
	if request.SnackTypeID != nil {
		snack.SnackTypeID = *request.SnackTypeID
	}
	if request.Description != nil {
		snack.Description = *request.Description
	}
	if request.ImageURI != nil {
		snack.ImageURI = *request.ImageURI
	}
	if request.Price != nil {
		snack.Price = *request.Price
	}
	if request.IsActive != nil {
		snack.IsActive = *request.IsActive
	}
	if request.OrderThreshold != nil {
		snack.OrderThreshold = request.OrderThreshold
	}
	if request.LastUpdatedBy != nil {
		snack.LastUpdatedBy = *request.LastUpdatedBy
	}
	snack.LastUpdatedDTM = time.Now()
}
