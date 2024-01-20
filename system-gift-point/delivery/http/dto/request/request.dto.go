package request

import "system-gift-point/domain/entity"

type RequestsystemPointDTO struct {
    Name  string `json:"name"`
    Point int    `json:"point"`
}

func (this RequestsystemPointDTO) ToActionEntity() entity.Action {
    return entity.Action{
        Name:  this.Name,
        Point: this.Point,
    }
}
