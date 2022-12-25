package store

import "database/sql"

type Store struct {
	db                *sql.DB
	usersRepository   *UsersRepository
	answersRepository *AnswersRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Users() *UsersRepository {
	if s.usersRepository == nil {
		s.usersRepository = &UsersRepository{
			store: s,
		}
	}

	return s.usersRepository
}

func (s *Store) Answers() *AnswersRepository {
	if s.answersRepository == nil {
		s.answersRepository = &AnswersRepository{
			store: s,
		}
	}

	return s.answersRepository
}
