package api

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type CreateToolRequest struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Profile     string `json:"profile"`
	Description string `json:"description"`

	Environment string `json:"environment"`
	ColumnName  string `json:"columnName"`
}

type UpdateToolRequest struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Profile     string `json:"profile"`
	Description string `json:"description"`
}

type UserInfo struct {
	Name string `json:"name"`
}

type ToolInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Profile     string `json:"profile,omitempty"`
	Description string `json:"description"`

	Environment string `json:"environment"`
	ColumnName  string `json:"columnName"`
}

type Column struct {
	Name  string      `json:"name"`
	Tools []*ToolInfo `json:"tools,omitempty"`
}

type Environment struct {
	Name    string    `json:"name"`
	Columns []*Column `json:"columns,omitempty"`
}

type ColumnInfo struct {
	Name string `json:"name"`
}

type EnvironmentInfo struct {
	Name string `json:"name"`
}

type GeneralResponse struct {
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}
