package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "zoo-backend/config"
    "zoo-backend/controllers"
    "zoo-backend/migrations"
    "zoo-backend/repositories"
    "zoo-backend/services"
    "zoo-backend/middleware" 
)

func main() {
    config.Connect()
    migrations.Migrate()

    zooRepo := &repositories.ZooRepository{DB: config.DB}
    zooService := &services.ZooService{Repo: zooRepo}
    zooController := &controllers.ZooController{Service: zooService}

    
    router := mux.NewRouter()

    router.Use(middleware.LoggerMiddleware)

    router.HandleFunc("/zoos", zooController.GetAllZoos).Methods(http.MethodGet)
    router.HandleFunc("/zoos", zooController.CreateZoo).Methods(http.MethodPost)
    router.HandleFunc("/zoos/{id}", zooController.GetZooByID).Methods(http.MethodGet) // Mendapatkan berdasarkan ID
    router.HandleFunc("/zoos/{id}", zooController.UpdateZoo).Methods(http.MethodPut) // Memperbarui berdasarkan ID
    router.HandleFunc("/zoos/{id}", zooController.DeleteZoo).Methods(http.MethodDelete) // Menghapus berdasarkan ID

    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
