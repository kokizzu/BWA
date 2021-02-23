package transactions

type GetTransactionCampaignDetailsInput struct {
	ID int `uri:"id" binding:"required"`
}
