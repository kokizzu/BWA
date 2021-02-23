package transactions

type service struct {
	repository Repository
}

type Service interface {
	GetTrasactionByCampaignID(input GetTransactionCampaignDetailsInput) ([]Transactions, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTrasactionByCampaignID(input GetTransactionCampaignDetailsInput) ([]Transactions, error) {

	trans, err := s.repository.GetCampaignByID(input.ID)
	if err != nil {
		return trans, err
	}
	return trans, nil

}
