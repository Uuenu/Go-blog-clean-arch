package service

type AuthorStorage interface {
	Create() error
	Update() error
}

type authorService struct {
	Storage AuthorStorage
}

func NewAuthorService(storage AuthorStorage) *authorService {
	return &authorService{
		Storage: storage,
	}
}
