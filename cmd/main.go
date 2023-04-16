package main

import (
	"context"
	"database/sql"
	"log"
	"sample-api/db"
	"sample-api/infrastructure/persistence/postgres"
	"sample-api/presentation/handler"
	"sample-api/usecase/interactor"
	"strings"

	"net/http"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"

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

	// Initialize Firebase
	app := initFirebase()

	r := gin.Default()

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			// 認証が不要なエンドポイント
			users := v1.Group("/users")
			{
				users.GET("/", userHandler.GetAll)
				users.GET("/:id", userHandler.GetByID)
				users.POST("/", userHandler.Create)
				users.PUT("/:id", userHandler.Update)
				users.DELETE("/:id", userHandler.Delete)
			}

			// 認証が必要なエンドポイント
			authUsers := v1.Group("/auth-users")
			authUsers.Use(FirebaseAuthMiddleware(app))
			{
				authUsers.GET("/", userHandler.GetAll)
				authUsers.GET("/:id", userHandler.GetByID)
				authUsers.POST("/", userHandler.Create)
				authUsers.PUT("/:id", userHandler.Update)
				authUsers.DELETE("/:id", userHandler.Delete)
			}
		}
	}

	r.Run(":8080")
}

// Initialize Firebase
func initFirebase() *firebase.App {
	saKeyPath := "serviceAccountKey.json"
	opt := option.WithCredentialsFile(saKeyPath)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase: %v", err)
	}

	return app
}

// FirebaseAuthMiddleware is a middleware for Firebase authentication
func FirebaseAuthMiddleware(app *firebase.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		log.Printf(authHeader)
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not found"})
			c.Abort()
			return
		}

		idToken := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))

		log.Printf(idToken)

		client, err := app.Auth(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error initializing Auth client"})
			c.Abort()
			return
		}

		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid ID token"})
			c.Abort()
			return
		}

		c.Set("uid", token.UID)

		c.Next()
	}
}
