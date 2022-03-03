package Controllers

import (
	"backend/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddWatched(c *gin.Context) {
	var watchedList Models.Watched_list
	watchedList.User_id, _ = strconv.Atoi(c.PostForm("userId"))
	watchedList.Movie_id, _ = strconv.Atoi(c.PostForm("movieId"))
	res := Models.AddWatched(watchedList)
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

func DeleteWatched(c *gin.Context) {
	//var Review_id int
	var watchedList Models.Watched_list
	watchedList.User_id, _ = strconv.Atoi(c.PostForm("userId"))
	watchedList.Movie_id, _ = strconv.Atoi(c.PostForm("movieId"))
	res := Models.DeleteWatched(watchedList.User_id, watchedList.Movie_id)
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

func ReadWatched(c *gin.Context) {
	//var Review_id int
	var watchedList Models.Watched_list
	watchedList.User_id, _ = strconv.Atoi(c.PostForm("userId"))
	res := Models.ReadWatched(watchedList.User_id)
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
