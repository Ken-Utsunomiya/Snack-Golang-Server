package validators

import "time"

type SnackRegisterRequest struct {
	SnackName	string	`json:"snack_name" binding:"required"`
	SnackTypeID	int	`json:"snack_type_id" binding:"required"`
	Description	string	`json:"description" binding:"required"`
	ImageURI	string	`json:"image_uri" binding:"required"`
	Price	int	`json:"price" binding:"required"`
	IsActive	bool	`json:"is_active" binding:"required"`
	OrderThreshold	*int	`json:"order_threshold"`
	LastUpdatedDTM	time.Time	`json:"last_updated_dtm" binding:"required"`
	LastUpdatedBy	string	`json:"last_updated_by" binding:"required"`
}
