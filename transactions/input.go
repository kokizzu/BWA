package transactions

import "BWA/user"

type GetTransactionCampaignDetailsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
