package Router

import (
	castController "backend/Controllers"
	genreController "backend/Controllers"
	movieController "backend/Controllers"
	ratingController "backend/Controllers"
	reviewController "backend/Controllers"
	userController "backend/Controllers"
	watchedListController "backend/Controllers"
	jwt "backend/Middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Main() {
	router := gin.Default()

	// Admin Router, put all customer router in this. // there is a Middleware that authentic admin
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

	// User Router, put all customer router in this.
	userGroup := router.Group("/user")
	{
		userGroup.POST("/logIn", userController.LogIn)
		userGroup.POST("/signUp", userController.SignUp)               // when query sth
		userGroup.POST("/checkUsername", userController.CheckUsername) // when update sth
		movieGroup := router.Group("/user/movie")
		{
			movieGroup.POST("/searchMovieByName", movieController.SearchMovieByName)
			movieGroup.POST("/searchMovieByCast", movieController.SearchMovieByCast)
			movieGroup.POST("/searchMovieByMovieId", movieController.SearchMovieByMovieId)
			movieGroup.POST("/searchMovieWithGenre", genreController.SearchMovieWithGenre)
			movieGroup.POST("/searchMovieByYear", movieController.SearchMovieByYear)
			movieGroup.POST("/topMovie", movieController.TopMovie)
		}
		castGroup := router.Group("/user/cast")
		{
			castGroup.POST("/searchCastById", castController.SearchCastById)
			castGroup.POST("/searchCastByMovieId", castController.SearchCastByMovieId)
			castGroup.POST("/searchRelativeMovieByCastId", castController.SearchRelativeMovieByCastId)
			castGroup.POST("/searchCastByName", castController.SearchCastByName)

		}
		reviewGroup := router.Group("/user/review")
		{
			reviewGroup.POST("/addReview", reviewController.AddReview)
			reviewGroup.POST("/updateReview", reviewController.UpdateReview)
			reviewGroup.POST("/deleteReview", reviewController.DeleteReview)
			reviewGroup.POST("/readReview", reviewController.ReadReview)
			reviewGroup.POST("/readReviewByMovieId", reviewController.ReadReviewByMovieId)
		}
		watchedListGroup := router.Group("/user/watchedList")
		{
			watchedListGroup.POST("/addWatchedList", watchedListController.AddWatched)
			watchedListGroup.POST("/deleteWatchedList", watchedListController.DeleteWatched)
			watchedListGroup.POST("/readWatchedList", watchedListController.ReadWatched)
		}
		ratingGroup := router.Group("/user/rating")
		{
			ratingGroup.POST("/readRating", ratingController.ReadRating)
			ratingGroup.POST("/addRating", ratingController.AddRating)
			ratingGroup.POST("/updateRating", ratingController.UpdateRating)
			ratingGroup.POST("/computeAvgGrade", ratingController.ComputeAvgGrade)
			ratingGroup.POST("/deleteRating", ratingController.DeleteRating)
		}
		genreGroup := router.Group("/user/genre")
		{
			genreGroup.POST("/searchGenreName", genreController.SearchGenreName)
			genreGroup.POST("/searchGenreByGenreName", genreController.SearchGenreByGenreName)
		}
	}
	userGroup.Use(jwt.JWTAuth())
	{
		userGroup.POST("/UpdateProfile", userController.UpdateUserProfile)
		userGroup.POST("/DeleteUser", userController.DeleteUser)
		userGroup.POST("/queryUserInfoById", userController.QueryUserInfoById)
	}

	// if user input invalid address, it'll send below msg
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "please use correct address!"})
	})
	// router port
	router.Run(":8000")
}
