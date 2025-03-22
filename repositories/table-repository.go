package repositories

import (
	"database/sql"
	"go-data-service/models"
)

type TableRepository struct {
	db *sql.DB
}

func NewTableRepository(db *sql.DB) *TableRepository {
	return &TableRepository{db: db}
}

func (r *TableRepository) GetTables() ([]string, error) {
	var tables []string
	query := `SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, nil
}

func (r *TableRepository) GetTableData(tableName string) (*models.TableData, error) {
	var tableData models.TableData
	query := `SELECT * FROM ` + tableName
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	tableData.Columns = columns

	for rows.Next() {
		values := make([]interface{}, len(columns))
		for i := range values {
			values[i] = new(interface{})
		}
		if err := rows.Scan(values...); err != nil {
			return nil, err
		}
		row := make([]interface{}, len(columns))
		for i, v := range values {
			row[i] = *(v.(*interface{}))
		}
		tableData.Rows = append(tableData.Rows, row)
	}
	return &tableData, nil
}

func (r *TableRepository) CreateTable(table models.Table) error {
	query := `CREATE TABLE ` + table.Name + ` (id SERIAL PRIMARY KEY`
	for _, column := range table.Columns {
		query += `, ` + column.Name + ` ` + column.Type
	}
	query += `)`
	_, err := r.db.Exec(query)
	return err
}
