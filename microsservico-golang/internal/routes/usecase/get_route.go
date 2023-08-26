package usecase

import "github.com/Diogo364/ImersaoFullCycle-14/go/internal/routes/entity"

type GetRouteUseCase struct {
	Repository entity.RouteRepository
}

func NewGetRouteUseCase(repository entity.RouteRepository) *GetRouteUseCase {
	return &GetRouteUseCase{
		Repository: repository,
	}
}

func (u *GetRouteUseCase) FindAll() ([]entity.Route, error) {
	routesList, err := u.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	return routesList, nil
}
