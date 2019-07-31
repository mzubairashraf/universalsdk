package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
	"universalsdk/controller"
	"universalsdk/service"
)

func main() {

	router := mux.NewRouter()

	var sessionKeyMap sync.Map

	usdkService := service.NewUsdkService(&sessionKeyMap)
	usdkController := controller.NewUsdkController(usdkService)

	router.HandleFunc("/isgood", usdkController.DeviceCheck).Methods("POST")

	log.Println("##  Starting Server on Port 8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Error while initializing server", err)
	}
}
