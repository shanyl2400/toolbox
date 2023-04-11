package model

type Tool struct {
	ID          string
	Name        string
	URL         string
	Profile     string
	Description string

	Environment string
	ColumnName  string
}

type CreateToolRequest struct {
	ToolBasicInfo

	Environment string
	ColumnName  string
}

type ToolBasicInfo struct {
	Name        string
	URL         string
	Profile     string
	Description string
}

type Column struct {
	Name     string
	Priority uint
}
