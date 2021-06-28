package transactions

import "gorm.io/gorm"

type Repository interface {
	GetCampaignByID(campaignID int) ([]Transactions, error)
	GetByUserID(userID int) ([]Transactions, error)
}

type repository struct {
	db *gorm.DB
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

func (r *repository) GetByUserID(userID int) ([]Transactions, error) {
	var transaction []Transactions

	//preload untuk nyari gambar di campaign mana transksi dibuat
	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", userID).Find(&transaction).Error

	if err != nil {
		return transaction, err
	}
	return transaction, nil

}
