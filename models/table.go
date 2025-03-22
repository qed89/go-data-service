package models

type Table struct {
	Name    string        `db:"table_name"`
	Columns []TableColumn `json:"columns"`
}

type TableColumn struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type TableData struct {
	Columns []string        `json:"columns"`
	Rows    [][]interface{} `json:"rows"`
}
