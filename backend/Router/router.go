package Router

import (
	castController "backend/Controllers"
	genreController "backend/Controllers"
	movieController "backend/Controllers"
	ratingController "backend/Controllers"
	userController "backend/Controllers"
	watchedListController "backend/Controllers"
	jwt "backend/Middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Main() {
	router := gin.Default()

	// Customer Router, put all customer router in this.
	customerGroup := router.Group("/customer")
	{
		customerGroup.POST("/add", userController.LogIn)      // when create sth
		customerGroup.GET("/query", userController.LogIn)     // when query sth
		customerGroup.PUT("/update", userController.LogIn)    // when update sth
		customerGroup.DELETE("/delete", userController.LogIn) // when delete sth
	}

	// Admin Router, put all customer router in this. // there is a Middleware that authentic admin
	adminGroup := router.Group("/admin", gin.BasicAuth(gin.Accounts{"admin": "123"}))
	{
		adminGroup.GET("/", userController.LogIn)
		adminGroup.GET("/query", userController.LogIn)     // when query sth
		adminGroup.PUT("/update", userController.LogIn)    // when update sth
		adminGroup.DELETE("/delete", userController.LogIn) // when delete sth
		castGroup := router.Group("/admin/cast")
		{
			castGroup.POST("/addCast")
			castGroup.POST("/UpdateCast")
			castGroup.POST("/DeleteCast")
			castGroup.POST("/SearchCast")
		}
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
		}
		castGroup := router.Group("/user/cast")
		{
			castGroup.POST("/searchCastById", castController.SearchCastById)
			castGroup.POST("/searchCastByMovieId", castController.SearchCastByMovieId)
			castGroup.POST("/searchRelativeMovieByCastId", castController.SearchRelativeMovieByCastId)

		}
		reviewGroup := router.Group("/user/review")
		{
			reviewGroup.POST("/addReview", movieController.AddReview)
			reviewGroup.POST("/updateReview", movieController.UpdateReview)
			reviewGroup.POST("/deleteReview", movieController.DeleteReview)
			reviewGroup.POST("/readReview", movieController.ReadReview)
			reviewGroup.POST("/readReviewByMovieId", movieController.ReadReviewByMovieId)
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
		}
	}
	userGroup.Use(jwt.JWTAuth())
	{
		userGroup.GET("/jwt", userController.GetDataByTime)
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
