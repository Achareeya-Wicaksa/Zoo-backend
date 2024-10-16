package main

import (
    "log"
    "net/http"
    "os"
    "fmt"

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

    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Backend berhasil berjalan")
    }).Methods(http.MethodGet)

    router.HandleFunc("/zoos", zooController.GetAllZoos).Methods(http.MethodGet)
    router.HandleFunc("/zoos", zooController.CreateZoo).Methods(http.MethodPost)
    router.HandleFunc("/zoos/{id}", zooController.GetZooByID).Methods(http.MethodGet)
    router.HandleFunc("/zoos/{id}", zooController.UpdateZoo).Methods(http.MethodPut)
    router.HandleFunc("/zoos/{id}", zooController.DeleteZoo).Methods(http.MethodDelete)

    port := os.Getenv("PORT") // Ambil port dari variabel lingkungan
    if port == "" {
        port = "8080" // Default ke 8080 jika tidak ada variabel lingkungan
    }

    log.Printf("Server starting on :%s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}
