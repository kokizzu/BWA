package transactions

import (
	"BWA/campaign"
	"BWA/user"
	"time"
)

type Transactions struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
