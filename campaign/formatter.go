package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormattCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}

	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.Description
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.ImageURL = ""
	campaignFormatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormattCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormattCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	ID               int                       `json:"id"`
	Name             string                    `json:"name"`
	ShortDescription string                    `json:"short_description"`
	Description      string                    `json:"description"`
	ImageURL         string                    `json:"image_url"`
	GoalAmount       int                       `json:"goal_amount"`
	CurrentAmount    int                       `json:"current_amount"`
	UserID           int                       `json:"user_id"`
	Slug             string                    `json:"slug"`
	Perks            []string                  `json:"perks"`
	User             CampaignUserFormatter     `json:"user"`
	Images           []campaignImagesFormatter `json:"images_url"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type campaignImagesFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignFormat := CampaignDetailFormatter{}
	campaignFormat.ID = campaign.ID
	campaignFormat.Name = campaign.Name
	campaignFormat.ShortDescription = campaign.ShortDescription
	campaignFormat.Description = campaign.Description
	campaignFormat.UserID = campaign.UserID
	campaignFormat.GoalAmount = campaign.GoalAmount
	campaignFormat.CurrentAmount = campaign.CurrentAmount
	campaignFormat.Slug = campaign.Slug
	campaignFormat.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormat.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	campaignFormat.Perks = perks

	user := campaign.User

	campaignUser := CampaignUserFormatter{}
	campaignUser.Name = user.Name
	campaignUser.ImageURL = user.AvatarFileName

	campaignFormat.User = campaignUser

	images := []campaignImagesFormatter{}

	for _, image := range campaign.CampaignImages {
		campaignImageFormatter := campaignImagesFormatter{}
		campaignImageFormatter.ImageURL = image.FileName

		isPrimary := false

		if image.IsPrimary == 1 {
			isPrimary = true
		}
		campaignImageFormatter.IsPrimary = isPrimary
		images = append(images, campaignImageFormatter)
	}

	campaignFormat.Images = images

	return campaignFormat
}
