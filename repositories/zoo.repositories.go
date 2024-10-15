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
    err := r.DB.QueryRow("SELECT id FROM animal WHERE id = $1", zoo.ID).Scan(&existingID)
    if err == nil {
        return 0, fmt.Errorf("zoo with ID '%d' already exists", zoo.ID)
    } else if err != sql.ErrNoRows {
        log.Printf("SQL error during check for existing zoo: %v", err) 
        return 0, err
    }

    // Menggunakan RETURNING untuk mendapatkan ID yang dihasilkan
    var id int64
    err = r.DB.QueryRow("INSERT INTO animal (name, class, legs) VALUES ($1, $2, $3) RETURNING id", zoo.Name, zoo.Class, zoo.Legs).Scan(&id)
    if err != nil {
        log.Printf("SQL Exec error: %v", err) 
        return 0, err
    }

    return id, nil
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
    var existingID int64
    err := r.DB.QueryRow("SELECT id FROM animal WHERE id = $1", zoo.ID).Scan(&existingID)
    
    if err == nil {
        // Jika ID sudah ada, lakukan update
        result, err := r.DB.Exec("UPDATE animal SET name = $1, class = $2, legs = $3 WHERE id = $4", zoo.Name, zoo.Class, zoo.Legs, zoo.ID)
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
        return true, nil 
    } else if err != sql.ErrNoRows {
        log.Printf("ZooRepository Upsert: Error while checking for existing zoo: %v", err)
        return false, err
    }

    // Jika ID tidak ada, lakukan insert
    _, err = r.DB.Exec("INSERT INTO animal (id, name, class, legs) VALUES ($1, $2, $3, $4)", zoo.ID, zoo.Name, zoo.Class, zoo.Legs)
    if err != nil {
        log.Printf("ZooRepository Upsert: Insert failed: %v", err)
        return false, err
    }

    log.Printf("ZooRepository Upsert: Inserted new zoo with ID %d", zoo.ID)
    return false, nil 
}



func (r *ZooRepository) Delete(id int) error {
    result, err := r.DB.Exec("DELETE FROM animal WHERE id = $1", id)
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected == 0 {
        return fmt.Errorf("animal with ID '%d' not found", id)
    }

    return nil
}



