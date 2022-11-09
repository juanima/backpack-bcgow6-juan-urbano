package movies

import (
	"database/sql"
	"fmt"
        "log"

	"github.com/bootcamp-go/storage/internal/domain"
)

type Repository interface {
	Store(p domain.Movie) (int, error)
	GetByName(name string) (domain.Movie, error)
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

type repository struct {
	db *sql.DB
}

/* Ejercicio 2 - Replicar Store() */
func (r *repository) Store(p domain.Movie) (int, error) {

	stmt, err := r.db.Prepare("INSERT INTO movies (title, rating, awards, length, release_date) VALUES (?,?,?,?,?)")
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

// Ejercicio 1 - Implementar GetByName()
func (r *repository) GetByName(title string) (domain.Movie, error) {

	stmt, err := r.db.Prepare("SELECT id, title, rating, awards, length, genre_id FROM Movies WHERE title = ?;")
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
