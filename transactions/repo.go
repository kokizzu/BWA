package transactions

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetCampaignByID(campaignID int) ([]Transactions, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetCampaignByID(campaignID int) ([]Transactions, error) {
	var transactions []Transactions
	err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
