package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "zoo-backend/models"
    "zoo-backend/services"
)

type ZooController struct {
    Service *services.ZooService
}

// CreateZoo - Untuk membuat zoo baru
func (c *ZooController) CreateZoo(w http.ResponseWriter, r *http.Request) {
    var zoo models.Zoo
    err := json.NewDecoder(r.Body).Decode(&zoo)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    id, err := c.Service.CreateZoo(zoo)
    if err != nil {
        http.Error(w, "Failed to create zoo", http.StatusInternalServerError)
        return
    }

    // Berikan notifikasi sukses
    response := map[string]interface{}{
        "message": "Zoo created successfully",
        "id":      id,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}

// GetAllZoos - Untuk mendapatkan semua zoo
func (c *ZooController) GetAllZoos(w http.ResponseWriter, r *http.Request) {
    zoos, err := c.Service.GetAllZoos()
    if err != nil {
        http.Error(w, "Failed to get zoos", http.StatusInternalServerError)
        return
    }

    // Jika data kosong, kirim array kosong []
    if len(zoos) == 0 {
        zoos = []models.Zoo{}
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(zoos)
}


// GetZooByID - Untuk mendapatkan zoo berdasarkan ID
func (c *ZooController) GetZooByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    zoo, err := c.Service.GetZooByID(id)
    if err != nil {
        http.Error(w, "Zoo not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(zoo)
}

// UpdateZoo - Untuk memperbarui zoo
func (c *ZooController) UpdateZoo(w http.ResponseWriter, r *http.Request) {
    var zoo models.Zoo
    err := json.NewDecoder(r.Body).Decode(&zoo)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    err = c.Service.UpdateZoo(zoo)
    if err != nil {
        http.Error(w, "Failed to update zoo", http.StatusInternalServerError)
        return
    }

    // Berikan notifikasi sukses
    response := map[string]string{
        "message": "Zoo updated successfully",
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

// DeleteZoo - Untuk menghapus zoo
func (c *ZooController) DeleteZoo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    err := c.Service.DeleteZoo(id)
    if err != nil {
        http.Error(w, "Zoo not found", http.StatusNotFound)
        return
    }

    // Berikan notifikasi sukses
    response := map[string]string{
        "message": "Zoo deleted successfully",
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}
