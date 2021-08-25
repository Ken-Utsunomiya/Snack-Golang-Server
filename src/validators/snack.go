package validators

import (
	"Snack-Golang-Server/src/models"
)

type SnackRegisterRequest struct {
	SnackName	string	`json:"snack_name" binding:"required"`
	SnackTypeID	int	`json:"snack_type_id" binding:"required"`
	Description	string	`json:"description" binding:"required"`
	ImageURI	string	`json:"image_uri" binding:"required"`
	Price	int	`json:"price" binding:"required"`
	IsActive	bool	`json:"is_active" binding:"required"`
	OrderThreshold	*int	`json:"order_threshold"`
	LastUpdatedBy	string	`json:"last_updated_by" binding:"required"`
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
	snack.OrderThreshold = request.OrderThreshold
	if request.OrderThreshold != nil {
		snack.OrderThreshold = request.OrderThreshold
	}
	return snack
}
