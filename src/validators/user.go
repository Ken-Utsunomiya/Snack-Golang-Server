package validators

type UserRegisterRequest struct {
	FirstName	string `json:"first_name" binding:"required"`
	LastName	string `json:"last_name" binding:"required"`
	Email			string `json:"email_address" binding:"required"`
	ImageURI	string `json:"image_uri" binding:"required"`
}
