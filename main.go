package main

import (
	"fmt"
	"net/http"

	"github.com/pluralsight/webservice/controllers"
)

const (
	PORT = 5000
)

func main() {
	// u := models.User{
	// 	ID:        1,
	// 	FirstName: "Tricia",
	// 	LastName:  "McMillan",
	// }

	// fmt.Println(u)
	// port := 3000
	// port, err := startWebServer(port)
	// fmt.Println(port)
	// fmt.Println(err)
	controllers.RegisterControllers()

	http.ListenAndServe(":3000", nil)
	execute()

}

func execute() {
	slice := []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	for i, v := range slice {
		println(i, v)
	}
}

func startWebServer(port int) (int, error) {
	fmt.Println("Starting server at port ", port)
	fmt.Println("Server started . . . ")
	// Return an error
	// return errors.New("Something went wrong")

	return port, nil
}
