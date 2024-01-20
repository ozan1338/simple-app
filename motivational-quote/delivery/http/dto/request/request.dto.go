package request

import "motivational-reminder/domain/entity"

type RequestQuoteDTO struct {
    Text string `json:"text"`
}

func (this RequestQuoteDTO) ToQuoteEntity() entity.Quote {
    return entity.Quote{
        Text: this.Text,
    }
}
