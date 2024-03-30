package endpoint_tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matheusgomes28/urchin/app"
	"github.com/matheusgomes28/urchin/common"
	"github.com/matheusgomes28/urchin/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostSuccess(t *testing.T) {
	app_settings := common.AppSettings{
		DatabaseAddress:  "localhost",
		DatabasePort:     3006,
		DatabaseUser:     "root",
		DatabasePassword: "root",
		DatabaseName:     "urchin",
		WebserverPort:    8080,
	}

	database_mock := mocks.DatabaseMock{
		GetPostHandler: func(post_id int) (common.Post, error) {
			return common.Post{
				Title:   "TestPost",
				Content: "TestContent",
				Excerpt: "TestExcerpt",
				Id:      post_id,
			}, nil
		},
	}

	r := app.SetupRoutes(app_settings, database_mock)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/post/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "TestPost")
}

func TestPostFailureStringKey(t *testing.T) {

	app_settings := common.AppSettings{
		DatabaseAddress:  "localhost",
		DatabasePort:     3006,
		DatabaseUser:     "root",
		DatabasePassword: "root",
		DatabaseName:     "urchin",
		WebserverPort:    8080,
	}

	database_mock := mocks.DatabaseMock{}

	router := app.SetupRoutes(app_settings, database_mock)
	responseRecorder := httptest.NewRecorder()

	request, err := http.NewRequest("GET", "/post/sampleString", nil)

	require.Nil(t, err)

	router.ServeHTTP(responseRecorder, request)

<<<<<<< HEAD
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
=======
	require.Equal(t, http.StatusInternalServerError, responseRecorder.Code)
>>>>>>> e28a20d04683848c00887850547c494c16315819

}

func TestPostFailurePostDoesntExist(t *testing.T) {

	app_settings := common.AppSettings{
		DatabaseAddress:  "localhost",
		DatabasePort:     3006,
		DatabaseUser:     "root",
		DatabasePassword: "root",
		DatabaseName:     "urchin",
		WebserverPort:    8080,
	}

<<<<<<< HEAD
	database_mock := mocks.DatabaseMock{}
=======
	database_mock := mocks.DatabaseMock{
		GetPostHandler: func(int) (common.Post, error) {
			return common.Post{}, fmt.Errorf("post doesn't exist")
		},
	}
>>>>>>> e28a20d04683848c00887850547c494c16315819

	router := app.SetupRoutes(app_settings, database_mock)
	responseRecorder := httptest.NewRecorder()

<<<<<<< HEAD
	request, err := http.NewRequest("GET", "/post/-1", nil)
=======
	request, err := http.NewRequest("GET", "/post/10000", nil)
>>>>>>> e28a20d04683848c00887850547c494c16315819

	require.Nil(t, err)

	router.ServeHTTP(responseRecorder, request)

	require.Equal(t, http.StatusInternalServerError, responseRecorder.Code)
<<<<<<< HEAD

=======
>>>>>>> e28a20d04683848c00887850547c494c16315819
}

func TestPostFailureNegativeInvalidKey(t *testing.T) {

	app_settings := common.AppSettings{
		DatabaseAddress:  "localhost",
		DatabasePort:     3006,
		DatabaseUser:     "root",
		DatabasePassword: "root",
		DatabaseName:     "urchin",
		WebserverPort:    8080,
	}

<<<<<<< HEAD
	database_mock := mocks.DatabaseMock{}
=======
	database_mock := mocks.DatabaseMock{
		GetPostHandler: func(int) (common.Post, error) {
			return common.Post{}, fmt.Errorf("post doesn't exist")
		},
	}
>>>>>>> e28a20d04683848c00887850547c494c16315819

	router := app.SetupRoutes(app_settings, database_mock)
	responseRecorder := httptest.NewRecorder()

<<<<<<< HEAD
	request, err := http.NewRequest("GET", "/post/10000", nil)
=======
	request, err := http.NewRequest("GET", "/post/-1", nil)
>>>>>>> e28a20d04683848c00887850547c494c16315819

	require.Nil(t, err)

	router.ServeHTTP(responseRecorder, request)

	require.Equal(t, http.StatusInternalServerError, responseRecorder.Code)
<<<<<<< HEAD

=======
>>>>>>> e28a20d04683848c00887850547c494c16315819
}
