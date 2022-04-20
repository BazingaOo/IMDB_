package Test

import (
	genreController "backend/Controllers"
	movieController "backend/Controllers"
	jwt "backend/Middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

//func TestAddGenre(t *testing.T) {
//	var genre = Models.Genre{Genre_id: 3, Genre_name: "3"}
//	if Models.AddGenre(genre) == 0 {
//		t.Error("result is wrong!")
//	} else {
//		t.Log("result is right")
//	}
//}

func TestAddGenre(t *testing.T) {
	router := gin.New()
	adminGroup := router.Group("/admin")
	adminGroup.Use(jwt.JWTAuth())
	{
		adminGroup.POST("/addGenre", genreController.AddGenre)
		adminGroup.POST("/updateGenre", genreController.UpdateGenre)
		adminGroup.POST("/deleteGenre", genreController.DeleteGenre)
		adminGroup.POST("/addMovie", movieController.AddMovie)
		adminGroup.POST("/updateMovie", movieController.UpdateMovie)
		adminGroup.DELETE("/deleteMovie", movieController.DeleteMovie)
		adminGroup.POST("/addMovieGenre", movieController.AddMovieGenre)
		adminGroup.POST("/addMovieCast", movieController.AddMovieCast)

	}

	params := url.Values{}
	params.Add("genreId", "3")
	params.Add("genreName", "5555")

	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/addGenre", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestUpdateGenre(t *testing.T) {
	router := gin.New()
	adminGroup := router.Group("/admin")
	adminGroup.Use(jwt.JWTAuth())
	{
		adminGroup.POST("/addGenre", genreController.AddGenre)
		adminGroup.POST("/updateGenre", genreController.UpdateGenre)
		adminGroup.POST("/deleteGenre", genreController.DeleteGenre)
		adminGroup.POST("/addMovie", movieController.AddMovie)
		adminGroup.POST("/updateMovie", movieController.UpdateMovie)
		adminGroup.DELETE("/deleteMovie", movieController.DeleteMovie)
		adminGroup.POST("/addMovieGenre", movieController.AddMovieGenre)
		adminGroup.POST("/addMovieCast", movieController.AddMovieCast)

	}

	params := url.Values{}
	params.Add("genreId", "3")
	params.Add("genreName", "333")

	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/updateGenre", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}
func TestDeleteGenre(t *testing.T) {
	router := gin.New()
	adminGroup := router.Group("/admin")
	adminGroup.Use(jwt.JWTAuth())
	{
		adminGroup.POST("/addGenre", genreController.AddGenre)
		adminGroup.POST("/updateGenre", genreController.UpdateGenre)
		adminGroup.POST("/deleteGenre", genreController.DeleteGenre)
		adminGroup.POST("/addMovie", movieController.AddMovie)
		adminGroup.POST("/updateMovie", movieController.UpdateMovie)
		adminGroup.DELETE("/deleteMovie", movieController.DeleteMovie)
		adminGroup.POST("/addMovieGenre", movieController.AddMovieGenre)
		adminGroup.POST("/addMovieCast", movieController.AddMovieCast)

	}

	params := url.Values{}
	params.Add("genreId", "3")
	//params.Add("genreName", "5555")

	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/deleteGenre", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}
func TestAddMovie(t *testing.T) {
	router := gin.New()
	adminGroup := router.Group("/admin")
	adminGroup.Use(jwt.JWTAuth())
	{
		adminGroup.POST("/addGenre", genreController.AddGenre)
		adminGroup.POST("/updateGenre", genreController.UpdateGenre)
		adminGroup.POST("/deleteGenre", genreController.DeleteGenre)
		adminGroup.POST("/addMovie", movieController.AddMovie)
		adminGroup.POST("/updateMovie", movieController.UpdateMovie)
		adminGroup.DELETE("/deleteMovie", movieController.DeleteMovie)
		adminGroup.POST("/addMovieGenre", movieController.AddMovieGenre)
		adminGroup.POST("/addMovieCast", movieController.AddMovieCast)

	}

	params := url.Values{}
	params.Add("movieId", "3")
	params.Add("movieName", "tao")
	params.Add("year", "2022")
	params.Add("desciption", "5555")
	//params.Add("genreName", "5555")
	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/addMovie", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestUpdateMovie(t *testing.T) {
	router := gin.New()
	adminGroup := router.Group("/admin")
	adminGroup.Use(jwt.JWTAuth())
	{
		adminGroup.POST("/addGenre", genreController.AddGenre)
		adminGroup.POST("/updateGenre", genreController.UpdateGenre)
		adminGroup.POST("/deleteGenre", genreController.DeleteGenre)
		adminGroup.POST("/addMovie", movieController.AddMovie)
		adminGroup.POST("/updateMovie", movieController.UpdateMovie)
		adminGroup.DELETE("/deleteMovie", movieController.DeleteMovie)
		adminGroup.POST("/addMovieGenre", movieController.AddMovieGenre)
		adminGroup.POST("/addMovieCast", movieController.AddMovieCast)

	}

	params := url.Values{}
	params.Add("movieId", "3")
	params.Add("movieName", "tao")
	params.Add("year", "2022")
	params.Add("desciption", "5555")
	//params.Add("genreName", "5555")
	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/updateMovie", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestDeleteMovie(t *testing.T) {
	router := gin.New()
	adminGroup := router.Group("/admin")
	adminGroup.Use(jwt.JWTAuth())
	{
		adminGroup.POST("/addGenre", genreController.AddGenre)
		adminGroup.POST("/updateGenre", genreController.UpdateGenre)
		adminGroup.POST("/deleteGenre", genreController.DeleteGenre)
		adminGroup.POST("/addMovie", movieController.AddMovie)
		adminGroup.POST("/updateMovie", movieController.UpdateMovie)
		adminGroup.DELETE("/deleteMovie", movieController.DeleteMovie)
		adminGroup.POST("/addMovieGenre", movieController.AddMovieGenre)
		adminGroup.POST("/addMovieCast", movieController.AddMovieCast)

	}

	params := url.Values{}
	params.Add("movieId", "3")
	//params.Add("genreName", "5555")
	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/deleteMovie", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestAddMovieGenre(t *testing.T) {
	router := gin.New()
	adminGroup := router.Group("/admin")
	adminGroup.Use(jwt.JWTAuth())
	{
		adminGroup.POST("/addGenre", genreController.AddGenre)
		adminGroup.POST("/updateGenre", genreController.UpdateGenre)
		adminGroup.POST("/deleteGenre", genreController.DeleteGenre)
		adminGroup.POST("/addMovie", movieController.AddMovie)
		adminGroup.POST("/updateMovie", movieController.UpdateMovie)
		adminGroup.DELETE("/deleteMovie", movieController.DeleteMovie)
		adminGroup.POST("/addMovieGenre", movieController.AddMovieGenre)
		adminGroup.POST("/addMovieCast", movieController.AddMovieCast)

	}

	params := url.Values{}
	params.Add("movieId", "3")
	params.Add("genreId", "1")
	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/addMovieGenre", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestAddMovieCast(t *testing.T) {
	router := gin.New()
	adminGroup := router.Group("/admin")
	adminGroup.Use(jwt.JWTAuth())
	{
		adminGroup.POST("/addGenre", genreController.AddGenre)
		adminGroup.POST("/updateGenre", genreController.UpdateGenre)
		adminGroup.POST("/deleteGenre", genreController.DeleteGenre)
		adminGroup.POST("/addMovie", movieController.AddMovie)
		adminGroup.POST("/updateMovie", movieController.UpdateMovie)
		adminGroup.DELETE("/deleteMovie", movieController.DeleteMovie)
		adminGroup.POST("/addMovieGenre", movieController.AddMovieGenre)
		adminGroup.POST("/addMovieCast", movieController.AddMovieCast)

	}

	params := url.Values{}
	params.Add("movieId", "3")
	params.Add("castId", "1")
	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/addMovieCast", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}
