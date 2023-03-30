package model

type Tool struct {
	ID          string
	Name        string
	URL         string
	Profile     string
	Description string

	Environment string
	ColumnName  string

	CreatedAt int64
	UpdatedAt int64
}
