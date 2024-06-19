package main

// migrate için ayrı olarak çalıştır. go run migrate/migrate.go
import (
	"go-crud-example/initializers"
	"go-crud-example/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
