package movies

import (
	"fmt"
        "log"
        "context"
	"database/sql"

	"github.com/bootcamp-go/storage/internal/domain"
)
const (
        STORE_MOVIE = "INSERT INTO movies (title, rating, awards, length, release_date) VALUES (?,?,?,?,?)"
        GET_MOVIE_BY_TITLE = "SELECT id, title, rating, awards, length, genre_id FROM Movies WHERE title = ?;"
	GET_ALL_MOVIES = "SELECT m.id ,m.title, m.rating, m.awards, m.length, m.genre_id FROM movies m;"
)

type Repository interface {
	Store(p domain.Movie) (int, error)
	GetByName(name string) (domain.Movie, error)
        GetAll(c context.Context) ([]domain.Movie, error)
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

type repository struct {
	db *sql.DB
}

/* Ejercicio 2 - Replicar Store() */
func (r *repository) Store(p domain.Movie) (int, error) {

	stmt, err := r.db.Prepare(STORE_MOVIE)
	if err != nil {
		return 0, fmt.Errorf("error al preparar la consulta - error %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(p.Title, p.Rating, p.Awards, p.Length, p.Release_date)
	if err != nil {
		return 0, fmt.Errorf("error al ejecutar la consulta - error %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error al obtener último id - error %v", err)
	}

	return int(id), nil
}


func (r *repository) GetAll(c context.Context) ([]domain.Movie, error) {
	var movies []domain.Movie
	rows, err := r.db.Query(GET_ALL_MOVIES)
	if err != nil {
		return []domain.Movie{}, err
	}

	for rows.Next() {
		var movie domain.Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Rating, &movie.Awards, &movie.Length, &movie.Genre_id)
		if err != nil {
			return []domain.Movie{}, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

// Ejercicio 1 - Implementar GetByName()
func (r *repository) GetByName(title string) (domain.Movie, error) {

	stmt, err := r.db.Prepare(GET_MOVIE_BY_TITLE)
	if err != nil {
                log.Printf("My error is %+v\n", err)
		return domain.Movie{}, fmt.Errorf("error al preparar la consulta - error %v", err)
	}

	defer stmt.Close()

	var movie domain.Movie
	err = stmt.QueryRow(title).Scan(&movie.ID, &movie.Title, &movie.Rating, &movie.Awards, &movie.Length, &movie.Genre_id)
	if err != nil {
		return domain.Movie{}, fmt.Errorf("no registros para %s - error %v", title, err)
	}

	return movie, nil
}

