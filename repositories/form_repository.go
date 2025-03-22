package repositories

import (
	"database/sql"
	"go-data-service/models"
	"log"
)

type FormRepository struct {
	db *sql.DB
}

func NewFormRepository(db *sql.DB) *FormRepository {
	return &FormRepository{db: db}
}

func (r *FormRepository) GetForms(page, pageSize int) ([]models.Form, error) {
	log.Printf("Fetching forms from database. Page: %d, PageSize: %d\n", page, pageSize)

	var forms []models.Form
	query := `SELECT id, name, created_date, user_id FROM forms LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(query, pageSize, page*pageSize)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var form models.Form
		if err := rows.Scan(&form.ID, &form.Name, &form.CreatedDate, &form.UserID); err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		forms = append(forms, form)
	}

	log.Printf("Successfully fetched %d forms\n", len(forms))
	return forms, nil
}

func (r *FormRepository) GetFormByID(id string) (*models.Form, error) {
	var form models.Form
	query := `SELECT id, name, created_date, user_id FROM forms WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&form.ID, &form.Name, &form.CreatedDate, &form.UserID)
	if err != nil {
		return nil, err
	}
	return &form, nil
}

func (r *FormRepository) Save(form *models.Form) error {
	query := `INSERT INTO forms (id, name, created_date, user_id) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, form.ID, form.Name, form.CreatedDate, form.UserID)
	return err
}

func (r *FormRepository) Update(form *models.Form) error {
	query := `UPDATE forms SET name = $1, created_date = $2, user_id = $3 WHERE id = $4`
	_, err := r.db.Exec(query, form.Name, form.CreatedDate, form.UserID, form.ID)
	return err
}

func (r *FormRepository) Delete(id string) error {
	query := `DELETE FROM forms WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
