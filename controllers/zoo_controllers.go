package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/gorilla/mux" // Tambahkan import ini
    "zoo-backend/models"
    "zoo-backend/services"
)

type ZooController struct {
    Service *services.ZooService
}

func (c *ZooController) CreateZoo(w http.ResponseWriter, r *http.Request) {
    var zoo models.Zoo
    err := json.NewDecoder(r.Body).Decode(&zoo)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    id, err := c.Service.CreateZoo(zoo)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

func (c *ZooController) GetAllZoos(w http.ResponseWriter, r *http.Request) {
    zoos, err := c.Service.GetAllZoos()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(zoos)
}

func (c *ZooController) GetZooByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])
    zoo, err := c.Service.GetZooByID(id)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(zoo)
}

func (c *ZooController) UpdateZoo(w http.ResponseWriter, r *http.Request) {
    var zoo models.Zoo
    err := json.NewDecoder(r.Body).Decode(&zoo)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    err = c.Service.UpdateZoo(zoo)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func (c *ZooController) DeleteZoo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])
    err := c.Service.DeleteZoo(id)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
