package routes

import (
	"api-go-rest/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}