package main

import "fmt"

// App - the struct which contains things like pointers to database connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting Up our App.")
	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error Starting Up.")
		fmt.Println(err)
	}
	fmt.Println("Go server")
}
