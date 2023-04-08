package main

import (
	"github.com/Kazukite12/e-learning/controllers/routes"
	"github.com/Kazukite12/e-learning/models"
)

func main() {
	models.ConnectDB()
	routes.UserRoutes()

}
