package Router

import (
	movieController "backend/Controllers"
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

	movieGroup := router.Group("/movie")
	{
		movieGroup.POST("/addMovie", movieController.AddMovie)
		movieGroup.POST("/deleteMovie", movieController.DeleteMovie)
		movieGroup.POST("/UpdateMovie", movieController.UpdateMovie)
	}
	// User Router, put all customer router in this.
	userGroup := router.Group("/user")
	{
		userGroup.POST("/logIn", userController.LogIn)
		userGroup.POST("/signUp", userController.SignUp) // when query sth
		movieGroup := router.Group("user/movie")
		{
			movieGroup.POST("/searchMovieByName", movieController.SearchMovieByName)
			movieGroup.POST("/searchMovieByCast", movieController.SearchMovieByCast)
		}
		reviewGroup := router.Group("/user/review")
		{
			reviewGroup.POST("/addReview", movieController.AddReview)
			reviewGroup.POST("/updateReview", movieController.UpdateReview)
			reviewGroup.POST("/deleteReview", movieController.DeleteReview)
			reviewGroup.POST("/readReview", movieController.ReadReview)
		}
		watchedListGroup := router.Group("/user/watchedList")
		{
			watchedListGroup.POST("/addWatchedList", watchedListController.AddWatched)
			watchedListGroup.POST("/deleteWatchedList", watchedListController.DeleteWatched)
			watchedListGroup.POST("/readWatchedList", watchedListController.ReadWatched)
		}
	}
	userGroup.Use(jwt.JWTAuth())
	{
		userGroup.GET("/jwt", userController.GetDataByTime)
		userGroup.POST("/UpdateProfile", userController.UpdateUserProfile)
		userGroup.GET("/DeleteUser", userController.DeleteUser)
		userGroup.POST("/CheckUsername", userController.CheckUsername) // when update sth
	}
	// Admin Router, put all customer router in this. // there is a Middleware that authentic admin
	adminGroup := router.Group("/admin", gin.BasicAuth(gin.Accounts{"admin": "123"}))
	{
		adminGroup.GET("/", userController.LogIn)
		adminGroup.GET("/query", userController.LogIn)     // when query sth
		adminGroup.PUT("/update", userController.LogIn)    // when update sth
		adminGroup.DELETE("/delete", userController.LogIn) // when delete sth
	}

	// if user input invalid address, it'll send below msg
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "please use correct address!"})
	})
	// router port
	router.Run(":8000")
}
