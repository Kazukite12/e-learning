package routes

import (
	"net/http"

	//coursecontroller "github.com/Kazukite12/e-learning/controllers/courseController"
	coursecontroller "github.com/Kazukite12/e-learning/controllers/courseController"
	"github.com/Kazukite12/e-learning/controllers/userController"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func UserRoutes() {

	router := gin.Default()

	//User Api

	router.GET("api/user", userController.Index)
	router.GET("api/data/user", userController.Show)
	router.POST("api/user", userController.Create)
	router.PUT("api/user/:id", userController.Update)
	router.DELETE("api/user", userController.Delete)
	router.POST("api/user/register", userController.Register)
	router.POST("api/user/login", userController.Login)
	router.POST("api/user/logout", userController.Logout)
	router.PUT("api/join", userController.JoinCourse)

	//Course API

	router.GET("api/course", coursecontroller.Index)
	router.GET("api/course/:id", coursecontroller.Show)
	router.POST("api/course/create", coursecontroller.Create)
	router.PUT("api/course/update", coursecontroller.Update)
	router.POST("api/course/delete", coursecontroller.Delete)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5174"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	http.ListenAndServe("localhost:8080", handler)
}

//func corsMiddleware() gin.HandlerFunc {
//c := cors.New(cors.Options{
//	AllowedOrigins: []string{"http://localhost:8080"},
//	AllowCredentials: true,
//	})
//	return func(c *gin.Context) {
//		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
//		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
//		if c.Request.Method == "OPTIONS" {
//			c.AbortWithStatus(204)
//			return
//		}
//		c.Next()
//	}
// }

//exmaple use

// router.GET("api/user", corsMiddleware(), userController.Index)
//	router.GET("api/:id", userController.Show)
//	router.POST("api/user", corsMiddleware(), userController.Create)
//	router.PUT("api/user/:id", corsMiddleware(), userController.Update)
//	router.DELETE("api/user", corsMiddleware(), userController.Delete)
//	router.POST("api/user/register", userController.Register)
//	router.POST("api/user/login", userController.Login)

//router.Use(cors.New(cors.Config{
//	AllowOrigins:     []string{"http://localhost:5173"},
//	AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "HEAD", "OPTIONS"},
//	AllowHeaders:     []string{"Content-Type", "application/json", "Authorization", "Access-Control-Allow-Origin"},
//	ExposeHeaders:    []string{"Content-Length"},
//	AllowCredentials: true,
// }))
