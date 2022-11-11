package usecase

import (
	"agrokan-backend/src/business/domain"
	"agrokan-backend/src/business/usecase/user"
)

type Usecase struct {
	User user.Interface
}

func Init(d *domain.Domains) *Usecase {
	uc := &Usecase{
		User: user.Init(d.User),
	}

	return uc
}
