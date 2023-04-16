package main

import (
	"database/sql"
	"sample-api/db"
	"sample-api/infrastructure/persistence/postgres"
	"sample-api/presentation/handler"
	"sample-api/usecase/interactor"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=your_db_user password=your_db_password dbname=your_db_name sslmode=disable"

	// Run database migrations
	db.RunMigrations(connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepository := postgres.NewUserRepository(db)
	userInteractor := interactor.UserInteractor{
		UserRepository: userRepository,
	}
	userHandler := handler.UserHandler{
		UserInteractor: userInteractor,
	}

	r := gin.Default()

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			users := v1.Group("/users")
			{
				users.GET("/", userHandler.GetAll)
				users.GET("/:id", userHandler.GetByID)
				users.POST("/", userHandler.Create)
				users.PUT("/:id", userHandler.Update)
				users.DELETE("/:id", userHandler.Delete)
			}
		}
	}

	r.Run(":8080")
}
