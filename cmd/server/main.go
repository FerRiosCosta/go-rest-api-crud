package main

import (
	"fmt"
	"net/http"

	"github.com/FerRiosCosta/go-rest-api-crud/internal/comment"
	"github.com/FerRiosCosta/go-rest-api-crud/internal/database"
	transportHTTP "github.com/FerRiosCosta/go-rest-api-crud/internal/transport/http"
)

// App - the struct which contains things like pointers to database connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting Up our App.")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println(db)
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server.")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go Rest API CRUD.")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error Starting Up the Rest API.")
		fmt.Println(err)
	}

}
