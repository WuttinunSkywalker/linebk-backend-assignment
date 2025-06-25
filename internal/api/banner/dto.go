package banner

import "time"

type BannerResponse struct {
	BannerID    string    `json:"banner_id" example:"banner_27ce63f9-63ac-4815-8681-64d1218e625f"`
	UserID      string    `json:"user_id" example:"0cc8b473-df92-47ec-9b85-57e28aca4adf"`
	Title       string    `json:"title" example:"Want some money?"`
	Description string    `json:"description" example:"You can start applying"`
	Image       *string   `json:"image" example:"https://example.com/banner.jpg"`
	CreatedAt   time.Time `json:"created_at" example:"2024-01-15T10:30:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2024-01-15T10:30:00Z"`
}

func NewBannerResponse(banner *Banner) *BannerResponse {
	return &BannerResponse{
		BannerID:    banner.BannerID,
		UserID:      banner.UserID,
		Title:       banner.Title,
		Description: banner.Description,
		Image:       banner.Image,
		CreatedAt:   banner.CreatedAt,
		UpdatedAt:   banner.UpdatedAt,
	}
}
