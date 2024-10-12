package repositories

import (
	"database/sql"
	"zoo-backend/models"
)

type ZooRepository struct {
	DB *sql.DB
}

func (r *ZooRepository) Create(zoo models.Zoo) (int64, error) {
	result, err := r.DB.Exec("INSERT INTO  (name, class, legs) VALUES (?, ?, ?)", zoo.Name, zoo.Class, zoo.Legs)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *ZooRepository) GetAll() ([]models.Zoo, error) {
	rows, err := r.DB.Query("SELECT id, name, class, legs FROM testbackendzoo")
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
	err := r.DB.QueryRow("SELECT id, name, class, legs FROM testbackendzoo WHERE id = ?", id).Scan(&zoo.ID, &zoo.Name, &zoo.Class, &zoo.Legs)
	if err != nil {
		return zoo, err
	}
	return zoo, nil
}

func (r *ZooRepository) Update(zoo models.Zoo) error {
	_, err := r.DB.Exec("UPDATE testbackendzoo SET name = ?, class = ?, legs = ? WHERE id = ?", zoo.Name, zoo.Class, zoo.Legs, zoo.ID)
	return err
}

func (r *ZooRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM testbackendzoo WHERE id = ?", id)
	return err
}
