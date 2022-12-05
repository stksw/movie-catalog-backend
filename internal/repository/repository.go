package repository

import (
	"backend/internal/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies(genre ...int) ([]*models.Movie, error)
	AllGenres() ([]*models.Genre, error)
	GetUserByEmail(string) (*models.User, error)
	GetUserByID(int) (*models.User, error)

	OneMovieForEdit(id int) (*models.Movie, []*models.Genre, error)
	OneMovie(id int) (*models.Movie, error)
	InsertMovie(movie models.Movie) (int, error)
	UpdateMovie(movie models.Movie) error
	UpdateMovieGenres(id int, genreIDs []int) error
	DeleteMovie(id int) error
}
