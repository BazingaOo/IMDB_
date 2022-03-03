package Controllers

import (
	"backend/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SearchMovieByName(c *gin.Context) {
	var movieName string
	//user.Username = c.PostForm("username")
	movieName = c.PostForm("movieName")
	res := Models.SearchMovieByName(movieName)
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"user":    res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"user":    res,
		})
	}

}

func SearchMovieByCast(c *gin.Context) {
	var castName string
	//user.Username = c.PostForm("username")

	castName = c.PostForm("castName")
	res := Models.SearchMovieByCast(castName)
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"user":    res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"user":    res,
		})
	}

}

func SearchMovieByMovieId(c *gin.Context) {
	var movieId int
	movieId, _ = strconv.Atoi(c.PostForm("movieId"))
	//movieId = c.PostForm("movieId")
	res := Models.SearchMovieByMovieId(movieId)
	movie := Models.Movie{}
	if res == movie {
		//if res == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"user":    res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"user":    res,
		})
	}

}

func AddMovie(c *gin.Context) {
	var movie Models.Movie
	movie.Movie_id, _ = strconv.Atoi(c.PostForm("movieId"))
	movie.Movie_name = c.PostForm("movieName")
	movie.Year, _ = strconv.Atoi(c.PostForm("year"))
	movie.Desciption = c.PostForm("desciption")
	res := Models.AddMovie(movie)
	if res == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"user":    res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"user":    res,
		})
	}
}

func UpdateMovie(c *gin.Context) {
	var movie Models.Movie
	movie.Movie_id, _ = strconv.Atoi(c.PostForm("movieId"))
	movie.Movie_name = c.PostForm("movieName")
	movie.Year, _ = strconv.Atoi(c.PostForm("year"))
	movie.Desciption = c.PostForm("desciption")
	//res := Models.AddMovie(movie)
	res := Models.UpdateMovie(movie)
	if res == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"user":    res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"user":    res,
		})
	}
}

func DeleteMovie(c *gin.Context) {
	var movie Models.Movie
	movie.Movie_id, _ = strconv.Atoi(c.PostForm("movieId"))
	res := Models.DeleteMovie(movie.Movie_id)
	if res == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"user":    res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"user":    res,
		})
	}
}
