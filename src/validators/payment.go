package validators

type PaymentRegisterRequest struct {
	UserID						int `json:"user_id" binding:"required"`
	CreatedBy	string `json:"created_by" binding:"required"`
	PaymentAmount	int `json:"payment_amount" binding:"required"`
	TransactionIDs	[]int `json:"transaction_ids" binding:"required"`
}
