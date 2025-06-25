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

// func (u *userUsecase) GetUserGreetings(userID string, limit, offset int) ([]*UserGreetingResponse, int, error) {
// 	var g errgroup.Group
// 	var greetings []*UserGreeting
// 	var total int

// 	g.Go(func() error {
// 		var err error
// 		greetings, err = u.user.GetUserGreetings(userID, 0, 0)
// 		if err != nil {
// 			logger.Error(err)
// 			return fmt.Errorf("failed to get user greetings: %w", err)
// 		}
// 		return nil
// 	})

// 	g.Go(func() error {
// 		var err error
// 		total, err = u.user.CountUserGreetings(userID)
// 		if err != nil {
// 			logger.Error(err)
// 			return fmt.Errorf("failed to count user greetings: %w", err)
// 		}
// 		return nil
// 	})

// 	if err := g.Wait(); err != nil {
// 		return nil, 0, err
// 	}

// 	greetingResponses := make([]*UserGreetingResponse, 0, len(greetings))
// 	for _, g := range greetings {
// 		greetingResponses = append(greetingResponses, NewUserGreetingResponse(g))
// 	}

// 	return greetingResponses, total, nil
// }
