package movies

import (
	"github.com/bootcamp-go/storage/internal/domain"
	// "api-movies/pkg/utils"
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var (
	ERRORFORZADO = errors.New("error al ejecutar la consulta - error")
)

var pain = 0
var movie_test = domain.Movie{
	ID:           1,
	Created_at:   time.Now(),
	Updated_at:   time.Now(),
	Title:        "Cars 1",
	Rating:       4,
	Awards:       2,
	Release_date: time.Layout,
	Length:       &pain,
	Genre_id:     &pain,
}


func TestSave(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("Store Ok", func(t *testing.T) {

		mock.ExpectPrepare(regexp.QuoteMeta(STORE_MOVIE))
		mock.ExpectExec(regexp.QuoteMeta(STORE_MOVIE)).WillReturnResult(sqlmock.NewResult(1, 1))

		columns := []string{"id", "title", "rating", "awards", "length", "genre_id"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow(movie_test.ID, movie_test.Title, movie_test.Rating, movie_test.Awards, movie_test.Length, movie_test.Genre_id)
		mock.ExpectQuery(regexp.QuoteMeta(GET_MOVIE)).WithArgs(1).WillReturnRows(rows)

		repository := NewRepository(db)
		ctx := context.TODO()

		newID, err := repository.Store(ctx, movie_test)
		assert.NoError(t, err)

		movieResult, err := repository.GetMovieByID(ctx, int(newID))
		assert.NoError(t, err)

		assert.NotNil(t, movieResult)
		assert.Equal(t, movie_test.ID, movieResult.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Store Fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectPrepare(regexp.QuoteMeta(STORE_MOVIE))
		mock.ExpectExec(regexp.QuoteMeta(STORE_MOVIE)).WillReturnError(ERRORFORZADO)

		repository := NewRepository(db)
		ctx := context.TODO()

		id, err := repository.Store(ctx, movie_test)

		assert.EqualError(t, err, ERRORFORZADO.Error())
		assert.Empty(t, id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

