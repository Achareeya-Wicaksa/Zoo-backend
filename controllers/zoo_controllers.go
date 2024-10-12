package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "zoo-backend/models"
    "zoo-backend/services"
    "log" // Tambahkan import ini untuk logging
)

type ZooController struct {
    Service *services.ZooService
}

// CreateZoo - Untuk membuat zoo baru
func (c *ZooController) CreateZoo(w http.ResponseWriter, r *http.Request) {
    var zoo models.Zoo

    // Decode body menjadi struct Zoo
    err := json.NewDecoder(r.Body).Decode(&zoo)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid input format"})
        log.Printf("CreateZoo failed: Invalid input format. Error: %v", err) // Logging
        return
    }

    // Pastikan data yang diterima valid
    if zoo.Name == "" || zoo.Class == "" || zoo.Legs <= 0 {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid data provided"})
        log.Printf("CreateZoo failed: Invalid data provided. Received: %+v", zoo) // Logging
        return
    }

    // Memanggil service untuk membuat zoo baru
    id, err := c.Service.CreateZoo(zoo)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create zoo"})
        log.Printf("CreateZoo failed: %v", err) // Logging
        return
    }

    // Jika berhasil, kirim response success
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Successfully created zoo",
        "id":      id,
    })
    log.Printf("CreateZoo succeeded: Zoo created with ID %d", id) // Logging
}

// GetAllZoos - Untuk mendapatkan semua zoo
func (c *ZooController) GetAllZoos(w http.ResponseWriter, r *http.Request) {
    zoos, err := c.Service.GetAllZoos()
    if err != nil {
        http.Error(w, "Failed to get zoos", http.StatusInternalServerError)
        log.Printf("GetAllZoos failed: %v", err) // Logging error
        return
    }

    // Jika data kosong, kirim array kosong []
    if len(zoos) == 0 {
        zoos = []models.Zoo{}
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(zoos)

    // Logging status dan jumlah data yang dikembalikan
    log.Printf("GetAllZoos succeeded: Returned %d zoos with status %d", len(zoos), http.StatusOK)
}

// GetZooByID - Untuk mendapatkan zoo berdasarkan ID
func (c *ZooController) GetZooByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    zoo, err := c.Service.GetZooByID(id)
    if err != nil {
        http.Error(w, "Zoo not found", http.StatusNotFound)
        log.Printf("GetZooByID failed: Zoo with ID %d not found", id) // Logging
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(zoo)
    log.Printf("GetZooByID succeeded: Returned zoo with ID %d", id) // Logging
}

// UpdateZoo - Untuk memperbarui zoo
func (c *ZooController) UpdateZoo(w http.ResponseWriter, r *http.Request) {
    var zoo models.Zoo
    err := json.NewDecoder(r.Body).Decode(&zoo)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        log.Printf("UpdateZoo failed: Invalid input format. Error: %v", err) // Logging
        return
    }

    err = c.Service.UpdateZoo(zoo)
    if err != nil {
        http.Error(w, "Failed to update zoo", http.StatusInternalServerError)
        log.Printf("UpdateZoo failed: %v", err) // Logging
        return
    }

    // Berikan notifikasi sukses
    response := map[string]string{
        "message": "Zoo updated successfully",
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
    log.Printf("UpdateZoo succeeded: Updated zoo with ID %d", zoo.ID) // Logging
}

// DeleteZoo - Untuk menghapus zoo
func (c *ZooController) DeleteZoo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        log.Printf("DeleteZoo failed: Invalid ID format. Status: %d", http.StatusBadRequest)
        return
    }

    err = c.Service.DeleteZoo(id)
    if err != nil {
        http.Error(w, "Zoo not found", http.StatusNotFound)
        log.Printf("DeleteZoo failed: Zoo with ID %d not found. Status: %d", id, http.StatusNotFound)
        return
    }

    // Berikan notifikasi sukses
    response := map[string]string{
        "message": "Zoo deleted successfully",
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)

    // Logging status untuk keberhasilan penghapusan
    log.Printf("DeleteZoo succeeded: Zoo with ID %d deleted successfully. Status: %d", id, http.StatusOK)
}
