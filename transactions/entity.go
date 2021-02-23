package transactions

import (
	"BWA/user"
	"time"
)

type Transactions struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	User       user.User
	Status     string
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
