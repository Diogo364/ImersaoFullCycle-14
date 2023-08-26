package usecase

import "github.com/Diogo364/ImersaoFullCycle-14/go/internal/routes/entity"

type CreateRouteInput struct {
	Name        string
	Source      entity.JSONString
	Destination entity.JSONString
}


type CreateRouteUseCase struct {
	Repository entity.RouteRepository
}

func NewCreateRouteUseCase(repository entity.RouteRepository) *CreateRouteUseCase {
	return &CreateRouteUseCase{
		Repository: repository,
	}
}

func (u *CreateRouteUseCase) Create(input CreateRouteInput) error {
	route := entity.NewRoute("", input.Name, input.Source, input.Destination)

	err := u.Repository.Create(route)
	if err != nil {
		return err
	}
	return nil
}
