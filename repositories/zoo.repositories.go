package repositories

import (
	"database/sql"
	"zoo-backend/models"
	"log"
	"fmt"
)

type ZooRepository struct {
	DB *sql.DB
}

func (r *ZooRepository) Create(zoo models.Zoo) (int64, error) {
    var existingID int64
    // Check if the ID already exists in the database
    err := r.DB.QueryRow("SELECT id FROM animal WHERE id = ?", zoo.ID).Scan(&existingID)
    if err == nil {
        // ID already exists, return a conflict error
        return 0, fmt.Errorf("zoo with ID '%d' already exists", zoo.ID)
    } else if err != sql.ErrNoRows {
        // Another error occurred, return it
        log.Printf("SQL error during check for existing zoo: %v", err) // Logging
        return 0, err
    }

    // Insert the new zoo into the database
    result, err := r.DB.Exec("INSERT INTO animal (id, name, class, legs) VALUES (?, ?, ?, ?)", zoo.ID, zoo.Name, zoo.Class, zoo.Legs)
    if err != nil {
        log.Printf("SQL Exec error: %v", err) // Log the SQL error
        return 0, err
    }
    return result.LastInsertId()
}



func (r *ZooRepository) GetAll() ([]models.Zoo, error) {
	rows, err := r.DB.Query("SELECT id, name, class, legs FROM animal")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var zoos []models.Zoo
	for rows.Next() {
		var zoo models.Zoo
		if err := rows.Scan(&zoo.ID, &zoo.Name, &zoo.Class, &zoo.Legs); err != nil {
			return nil, err
		}
		zoos = append(zoos, zoo)
	}
	return zoos, nil
}

func (r *ZooRepository) GetByID(id int) (models.Zoo, error) {
	var zoo models.Zoo
	err := r.DB.QueryRow("SELECT id, name, class, legs FROM animal WHERE id = ?", id).Scan(&zoo.ID, &zoo.Name, &zoo.Class, &zoo.Legs)
	if err != nil {
		return zoo, err
	}
	return zoo, nil
}

func (r *ZooRepository) Upsert(zoo models.Zoo) (bool, error) {
    // First, check if a zoo with the given ID exists
    var existingID int64
    err := r.DB.QueryRow("SELECT id FROM animal WHERE id = ?", zoo.ID).Scan(&existingID)
    if err == nil {
        // If the ID exists, perform an update
        result, err := r.DB.Exec("UPDATE animal SET name = ?, class = ?, legs = ? WHERE id = ?", zoo.Name, zoo.Class, zoo.Legs, zoo.ID)
        if err != nil {
            log.Printf("ZooRepository Upsert: Update failed: %v", err)
            return false, err
        }

        rowsAffected, err := result.RowsAffected()
        if err != nil {
            log.Printf("ZooRepository Upsert: Failed to retrieve rows affected during update: %v", err)
            return false, err
        }

        log.Printf("ZooRepository Upsert: Updated zoo with ID %d. Rows affected: %d", zoo.ID, rowsAffected)
        return true, nil // Return true indicating that an update occurred
    } else if err != sql.ErrNoRows {
        log.Printf("ZooRepository Upsert: Error while checking for existing zoo: %v", err)
        return false, err
    }

    // If the ID does not exist, insert a new record
    _, err = r.DB.Exec("INSERT INTO animal (id, name, class, legs) VALUES (?, ?, ?, ?)", zoo.ID, zoo.Name, zoo.Class, zoo.Legs)
    if err != nil {
        log.Printf("ZooRepository Upsert: Insert failed: %v", err)
        return false, err
    }

    log.Printf("ZooRepository Upsert: Inserted new zoo with ID %d", zoo.ID)
    return false, nil // Return false indicating that an insert occurred
}



// In repositories/zoo_repository.go
func (r *ZooRepository) Delete(id int) error {
    _, err := r.DB.Exec("DELETE FROM animal WHERE id = ?", id)
    return err
}

