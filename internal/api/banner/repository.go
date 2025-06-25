package banner

import "github.com/jmoiron/sqlx"

type BannerRepository interface {
	GetBannersByUserID(userID string, limit, offset int) ([]*Banner, error)
	CountBannersByUserID(userID string) (int, error)
}

type bannerRepository struct {
	db *sqlx.DB
}

func NewBannerRepository(db *sqlx.DB) BannerRepository {
	return &bannerRepository{
		db: db,
	}
}

func (r *bannerRepository) GetBannersByUserID(userID string, limit, offset int) ([]*Banner, error) {
	var banners []*Banner
	query := "SELECT banner_id, user_id, title, description, image, created_at, updated_at FROM banners WHERE user_id = ? LIMIT ? OFFSET ?"
	err := r.db.Select(&banners, query, userID, limit, offset)
	return banners, err
}

func (r *bannerRepository) CountBannersByUserID(userID string) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM banners WHERE user_id = ?"
	err := r.db.Get(&count, query, userID)
	return count, err
}
