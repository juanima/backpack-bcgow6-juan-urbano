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

func Test_RepositoryGetAll(t *testing.T) {

        length, genre_id := 0, 1
        length2, genre_id2 := 2, 2

	t.Run("GetAll Ok", func(t *testing.T) {

	        db, mock, err := sqlmock.New()
	        assert.NoError(t, err)
	        defer db.Close()
	        // Columns

	        columns := []string{"id", "title", "rating", "awards", "length", "genre_id"}
	        rows := sqlmock.NewRows(columns)
	        movies := []domain.Movie{{ID: 1, Title: "Avatar", Rating: 22, Awards: 99, Length: &length, Genre_id: &genre_id}, {ID: 2, Title: "Simpson", Rating: 33, Awards: 11, Length: &length2, Genre_id: &genre_id2}}

	        for _, movie := range movies {
	        	rows.AddRow(movie.ID, movie.Title, movie.Rating, movie.Awards, movie.Length, movie.Genre_id)
	        }

	        mock.ExpectQuery(regexp.QuoteMeta(GET_ALL_MOVIES)).WillReturnRows(rows)

	        repo := NewRepository(db)
	        resultMovies, err := repo.GetAll(context.TODO())

	        assert.NoError(t, err)
	        assert.Equal(t, movies, resultMovies)
	        assert.NoError(t, mock.ExpectationsWereMet())
        })

        t.Run("GetAll Fail", func(t *testing.T) {

                db, mock, err := sqlmock.New()
	        assert.NoError(t, err)
	        defer db.Close()
	        // Columns
	        columns := []string{"id", "title", "rating", "awards", "length", "genre_id"}
	        rows := sqlmock.NewRows(columns)
	        movies := []domain.Movie{{ID: 1, Title: "Avatar", Rating: 22, Awards: 99, Length: &length, Genre_id: &genre_id}, {ID: 2, Title: "Simpson", Rating: 33, Awards: 11, Length: &length2, Genre_id: &genre_id2}}

	        for _, movie := range movies {
	        	rows.AddRow(movie.ID, movie.Title, movie.Rating, movie.Awards, movie.Length, movie.Genre_id)
	        }

	        mock.ExpectQuery(regexp.QuoteMeta(GET_ALL_MOVIES)).WillReturnError(ERRORFORZADO)

	        repo := NewRepository(db)
	        resultMovies, err := repo.GetAll(context.TODO())

	        assert.EqualError(t, err, ERRORFORZADO.Error())
	        assert.Empty(t, resultMovies)
	        assert.NoError(t, mock.ExpectationsWereMet())
        })
}


func Test_Delete(t *testing.T) {

        t.Run("DeleteOK", func(t *testing.T) {

	        // Arrange
	        db, mock, err := sqlmock.New()
	        assert.NoError(t, err)
	        defer db.Close()

	        mock.ExpectPrepare(regexp.QuoteMeta(DELETE_MOVIE))
	        mock.ExpectExec(regexp.QuoteMeta(DELETE_MOVIE)).WithArgs(movie_test.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	        repo := NewRepository(db)
	        // Act
	        err = repo.Delete(context.TODO(), 1)

	        // Assert
	        assert.NoError(t, err)
	        assert.NoError(t, mock.ExpectationsWereMet())

        })

        t.Run("Delete Fail", func(t *testing.T) {

	        // arrange
	        db, mock, err := sqlmock.New()
	        assert.NoError(t, err)
	        defer db.Close()

	        mock.ExpectPrepare(regexp.QuoteMeta(DELETE_MOVIE)).ExpectExec().WithArgs(movie_test.ID).WillReturnError(ERRORFORZADO)

	        repo := NewRepository(db)

	        // act
	        err = repo.Delete(context.TODO(), int64(movie_test.ID))

	        // assert
	        assert.EqualError(t, err, ERRORFORZADO.Error())
	        assert.NoError(t, mock.ExpectationsWereMet())
        })
}
