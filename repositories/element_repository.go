package repositories

import (
	"database/sql"
	"go-data-service/models"
)

type ElementRepository struct {
	db *sql.DB
}

func NewElementRepository(db *sql.DB) *ElementRepository {
	return &ElementRepository{db: db}
}

func (r *ElementRepository) GetElements(page, pageSize int) ([]models.Element, error) {
	var elements []models.Element
	query := `SELECT id, name, html_code, css_code, label, user_id, created_at FROM elements LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(query, pageSize, page*pageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var element models.Element
		if err := rows.Scan(&element.ID, &element.Name, &element.HtmlCode, &element.CssCode, &element.Label, &element.UserID, &element.CreatedAt); err != nil {
			return nil, err
		}
		elements = append(elements, element)
	}
	return elements, nil
}

func (r *ElementRepository) GetElementByID(id string) (*models.Element, error) {
	var element models.Element
	query := `SELECT id, name, html_code, css_code, label, user_id, created_at FROM elements WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&element.ID, &element.Name, &element.HtmlCode, &element.CssCode, &element.Label, &element.UserID, &element.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &element, nil
}

func (r *ElementRepository) Save(element *models.Element) error {
	query := `INSERT INTO elements (id, name, html_code, css_code, label, user_id, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(query, element.ID, element.Name, element.HtmlCode, element.CssCode, element.Label, element.UserID, element.CreatedAt)
	return err
}
