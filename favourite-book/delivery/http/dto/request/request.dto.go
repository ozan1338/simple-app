package request

import "favourite-book/domain/entity"

type RequestBookDTO struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func (this RequestBookDTO) ToBookEntity() entity.Book {
    return entity.Book{
        Author: this.Author,
        Title: this.Title,
        Rating: this.Rating,
    }
}
