package banner

import (
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/errs"
	"golang.org/x/sync/errgroup"
)

type BannerUsecase interface {
	GetBannersByUserID(userID string, limit, offset int) ([]*BannerResponse, int, error)
}

type bannerUsecase struct {
	banner BannerRepository
}

func NewBannerUsecase(banner BannerRepository) BannerUsecase {
	return &bannerUsecase{
		banner: banner,
	}
}

func (u *bannerUsecase) GetBannersByUserID(userID string, limit, offset int) ([]*BannerResponse, int, error) {
	var g errgroup.Group
	var banners []*Banner
	var total int

	g.Go(func() error {
		var err error
		banners, err = u.banner.GetBannersByUserID(userID, limit, offset)
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		var err error
		total, err = u.banner.CountBannersByUserID(userID)
		if err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, 0, errs.Internal(err)
	}

	bannerResponses := make([]*BannerResponse, 0, len(banners))
	for _, b := range banners {
		bannerResponses = append(bannerResponses, NewBannerResponse(b))
	}

	return bannerResponses, total, nil
}
