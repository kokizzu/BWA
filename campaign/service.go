package campaign

type Service interface {
	FindCampaign(userID int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindCampaign(userID int) ([]Campaign, error) {

	if userID != 0 {
		campaign, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaign, err
		}
		return campaign, nil
	}

	campaign, err := s.repository.FindAll()
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}
