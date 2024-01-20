package usecase

import (
    "system-gift-point/domain/entity"
    "system-gift-point/domain/repository"
)

type ActionUsecase interface {
    GetActions() ([]entity.Action, int, error)
    SaveAction(action entity.Action) error
}

type actionUsecaseImpl struct {
    databaseRepository repository.DatabaseRepository
}

func NewActionUsecase(databaseRepository repository.DatabaseRepository) actionUsecaseImpl {
    return actionUsecaseImpl{
        databaseRepository: databaseRepository,
    }
}

func (this *actionUsecaseImpl) GetActions() ([]entity.Action, int, error) {
    var actions []entity.Action
    if err := this.databaseRepository.Find(&actions); err != nil {
        return nil, 0, err
    }

    totalPoint := this.calculatePoint(actions)

    return actions, totalPoint, nil
}

func (this *actionUsecaseImpl) calculatePoint(actions []entity.Action) int {
    var totalPoints int

    for _, action := range actions {
        totalPoints += action.Point
    }

    return totalPoints
}

func (this *actionUsecaseImpl) SaveAction(action entity.Action) error {
    if err := this.databaseRepository.Create(&action); err != nil {
        return err
    }

    return nil
}
