package api

import (
	"net/http"
	"toolbox/internal/config"
	"toolbox/internal/errors"
	"toolbox/internal/model"
	"toolbox/internal/service"

	"github.com/gin-gonic/gin"
)

type API struct {
	usersService    service.IUsersService
	toolsService    service.IToolsService
	columnsService  service.IColumnsService
	profilesService service.IProfilesService
}

func (api *API) register(c *gin.Context) {
	req := new(RegisterUserRequest)
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, GeneralResponse{Message: "bad request"})
		return
	}
	err = api.usersService.Register(req.Name, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GeneralResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, GeneralResponse{Message: "success"})
}

func (api *API) login(c *gin.Context) {
	req := new(LoginRequest)
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, GeneralResponse{Message: "bad request"})
		return
	}
	token, err := api.usersService.Login(req.Name, req.Password)
	switch err {
	case errors.ErrGetUserFailed:
		c.JSON(http.StatusNotFound, GeneralResponse{Message: errors.ErrNoSuchUser.Error()})
		return
	case errors.ErrIncorrectPassword:
		c.JSON(http.StatusForbidden, GeneralResponse{Message: err.Error()})
		return
	case nil:
		c.JSON(http.StatusOK, GeneralResponse{
			Message: "success",
			Data:    token,
		})
	default:
		c.JSON(http.StatusInternalServerError, GeneralResponse{Message: err.Error()})
		return
	}

}

func (api *API) getUser(c *gin.Context) {
	name := c.Param("name")
	user, err := api.usersService.Get(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GeneralResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, GeneralResponse{
		Message: "success",
		Data:    UserInfo{Name: user.UserName},
	})
}

func (api *API) listUsers(c *gin.Context) {
	users, err := api.usersService.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, GeneralResponse{Message: err.Error()})
		return
	}
	ans := make([]*UserInfo, 0, len(users))
	for _, user := range users {
		ans = append(ans, &UserInfo{Name: user.UserName})
	}
	c.JSON(http.StatusOK, GeneralResponse{
		Message: "success",
		Data:    ans,
	})
}

func (api *API) createTool(c *gin.Context) {
	req := new(CreateToolRequest)
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, GeneralResponse{Message: "bad request"})
		return
	}
	if req.Profile == "" {
		req.Profile = "default.png"
	}
	err = api.toolsService.Create(&model.CreateToolRequest{
		ToolBasicInfo: model.ToolBasicInfo{
			Name:        req.Name,
			URL:         req.URL,
			Profile:     req.Profile,
			Description: req.Description,
		},
		Environment: req.Environment,
		ColumnName:  req.ColumnName,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, GeneralResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, GeneralResponse{
		Message: "success",
	})
}

func (api *API) updateTool(c *gin.Context) {
	id := c.Param("id")
	req := new(UpdateToolRequest)
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, GeneralResponse{Message: "bad request"})
		return
	}
	err = api.toolsService.Update(id, &model.ToolBasicInfo{
		Name:        req.Name,
		URL:         req.URL,
		Profile:     req.Profile,
		Description: req.Description,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, GeneralResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, GeneralResponse{
		Message: "success",
	})
}

func (api *API) deleteTool(c *gin.Context) {
	id := c.Param("id")
	err := api.toolsService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GeneralResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, GeneralResponse{
		Message: "success",
	})
}

func (api *API) getTool(c *gin.Context) {
	id := c.Param("id")
	tool, err := api.toolsService.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GeneralResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, GeneralResponse{
		Message: "success",
		Data: ToolInfo{
			ID:          tool.ID,
			Name:        tool.Name,
			URL:         tool.URL,
			Description: tool.Description,
			Profile:     tool.Profile,
			Environment: tool.Environment,
			ColumnName:  tool.ColumnName,
		},
	})
}

func (api *API) listTools(c *gin.Context) {
	env := c.Query("env")
	column := c.Query("column")
	groupBy := c.Query("groupBy")
	tools, err := api.toolsService.List(env, column)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GeneralResponse{Message: err.Error()})
		return
	}
	switch groupBy {
	case "":
		api.groupByTools(c, tools)
	case "tool":
		api.groupByTools(c, tools)
	case "environment":
		api.groupByEnvironments(c, tools)
	}
}

func (api *API) putColumn(c *gin.Context) {
	name := c.Param("name")
	err := api.columnsService.Put(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GeneralResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, GeneralResponse{
		Message: "success",
	})
}

func (api *API) deleteColumn(c *gin.Context) {
	name := c.Param("name")
	err := api.columnsService.Delete(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GeneralResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, GeneralResponse{
		Message: "success",
	})
}

func (api *API) listColumns(c *gin.Context) {
	columns, err := api.columnsService.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, GeneralResponse{Message: err.Error()})
		return
	}
	output := make([]*ColumnInfo, 0, len(columns))
	for _, column := range columns {
		output = append(output, &ColumnInfo{Name: column.Name})
	}

	c.JSON(http.StatusOK, GeneralResponse{
		Message: "success",
		Data:    output,
	})
}

func (api *API) listEnvironments(c *gin.Context) {
	environments := config.Get().Environments
	output := make([]*EnvironmentInfo, 0, len(environments))
	for _, env := range environments {
		output = append(output, &EnvironmentInfo{Name: env})
	}

	c.JSON(http.StatusOK, GeneralResponse{
		Message: "success",
		Data:    output,
	})
}

func (api *API) uploadProfile(c *gin.Context) {
	//获取文件
	file, header, err := c.Request.FormFile("file")
	//处理获取文件错误
	if err != nil {
		c.JSON(http.StatusBadRequest, GeneralResponse{Message: err.Error()})
		return
	}

	name, err := api.profilesService.Upload(header.Filename, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GeneralResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, GeneralResponse{
		Message: "success",
		Data:    name,
	})
}

func (api *API) getProfile(c *gin.Context) {
	name := c.Param("name")

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+name)
	c.Header("Content-Transfer-Encoding", "binary")
	err := api.profilesService.Download(name, c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GeneralResponse{Message: err.Error()})
	}
}

func (api *API) groupByTools(c *gin.Context, tools []*model.Tool) {
	output := make([]*ToolInfo, 0, len(tools))
	for _, tool := range tools {
		output = append(output, &ToolInfo{
			ID:          tool.ID,
			Name:        tool.Name,
			URL:         tool.URL,
			Description: tool.Description,
			Profile:     tool.Profile,
			Environment: tool.Environment,
			ColumnName:  tool.ColumnName,
		})
	}
	c.JSON(http.StatusOK, GeneralResponse{
		Message: "success",
		Data:    output,
	})
}

func (api *API) groupByEnvironments(c *gin.Context, tools []*model.Tool) {
	output := make([]*Environment, 0, len(tools))
	envColumnMap := make(map[string]map[string][]*ToolInfo)
	columnsSet := make([]string, 0)
	columnsMap := make(map[string]struct{})
	for _, tool := range tools {
		envMap := envColumnMap[tool.Environment]
		if envMap == nil {
			envMap = make(map[string][]*ToolInfo)
		}
		envMap[tool.ColumnName] = append(envMap[tool.ColumnName], &ToolInfo{
			ID:          tool.ID,
			Name:        tool.Name,
			URL:         tool.URL,
			Profile:     tool.Profile,
			Description: tool.Description,
			Environment: tool.Environment,
			ColumnName:  tool.ColumnName,
		})
		envColumnMap[tool.Environment] = envMap

		if _, ok := columnsMap[tool.ColumnName]; !ok {
			columnsSet = append(columnsSet, tool.ColumnName)
			columnsMap[tool.ColumnName] = struct{}{}
		}
	}

	environments := config.Get().Environments
	for _, envName := range environments {
		env := &Environment{
			Name: envName,
		}
		envEntity := envColumnMap[envName]

		for _, columnName := range columnsSet {
			column := &Column{
				Name:  columnName,
				Tools: envEntity[columnName],
			}
			env.Columns = append(env.Columns, column)
		}
		output = append(output, env)
	}
	c.JSON(http.StatusOK, GeneralResponse{
		Message: "success",
		Data:    output,
	})
}

func newAPI() *API {
	return &API{
		usersService:    service.NewUsersService(),
		toolsService:    service.NewToolsService(),
		columnsService:  service.NewColumnsService(),
		profilesService: service.NewProfilesService(),
	}
}
