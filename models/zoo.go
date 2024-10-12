package models

type Zoo struct {
    ID    int    `json:"id" gorm:"primaryKey"`
    Name  string `json:"name"`
    Class string `json:"class"`
    Legs  int    `json:"legs"` 
}
