package service

import (
	"toolbox/internal/errors"
	"toolbox/internal/model"
	"toolbox/internal/repository"
	rmodel "toolbox/internal/repository/model"
	"toolbox/internal/storage"
)

var (
	Environments = []string{
		"研发环境", "预发环境", "生产环境",
	}
)

type IToolsService interface {
	Create(req *model.CreateToolRequest) error
	Update(id string, req *model.ToolBasicInfo) error
	Delete(id string) error

	Get(id string) (*model.Tool, error)
	List(env string, columnID string) ([]*model.Tool, error)
}

type toolsService struct{}

func (t *toolsService) Create(req *model.CreateToolRequest) error {
	if err := t.checkBasicInfo(&req.ToolBasicInfo); err != nil {
		return err
	}
	if req.Environment == "" || req.ColumnName == "" {
		return errors.ErrEmptyToolParams
	}
	if !t.containsEnv(req.Environment) {
		return errors.ErrEnvironmentNotExists
	}
	_, err := repository.GetColumnsRepository().Get(req.ColumnName)
	if err != nil {
		return err
	}

	err = repository.GetToolsRepository().Set(&rmodel.Tool{
		Name:        req.Name,
		URL:         req.URL,
		Profile:     req.Profile,
		Description: req.Description,
		Environment: req.Environment,
		ColumnName:  req.ColumnName,
	})
	if err != nil {
		return err
	}
	return nil
}

func (t *toolsService) Update(id string, req *model.ToolBasicInfo) error {
	if err := t.checkBasicInfo(req); err != nil {
		return err
	}

	tool, err := repository.GetToolsRepository().Get(id)
	if err != nil {
		return err
	}

	tool.Name = req.Name
	tool.URL = req.URL
	if req.Profile != "" {
		tool.Profile = req.Profile
	}
	tool.Description = req.Description

	err = repository.GetToolsRepository().Set(tool)
	if err != nil {
		return err
	}
	return nil
}

func (t *toolsService) Delete(id string) error {
	return repository.GetToolsRepository().Delete(id)
}

func (t *toolsService) Get(id string) (*model.Tool, error) {
	tool, err := repository.GetToolsRepository().Get(id)
	if err != nil {
		return nil, err
	}

	return &model.Tool{
		ID:          tool.ID,
		Name:        tool.Name,
		URL:         tool.URL,
		Profile:     tool.Profile,
		Description: tool.Description,
		Environment: tool.Environment,
		ColumnName:  tool.ColumnName,
	}, nil
}

func (t *toolsService) List(env string, columnName string) ([]*model.Tool, error) {
	tools, err := repository.GetToolsRepository().List()
	if err != nil {
		return nil, err
	}

	output := make([]*model.Tool, 0)
	for _, tool := range tools {
		if env != "" && tool.Environment != env {
			continue
		}
		if columnName != "" && tool.ColumnName != columnName {
			continue
		}

		output = append(output, &model.Tool{
			ID:          tool.ID,
			Name:        tool.Name,
			URL:         tool.URL,
			Profile:     tool.Profile,
			Description: tool.Description,
			Environment: tool.Environment,
			ColumnName:  tool.ColumnName,
		})
	}
	return output, nil
}

func (t *toolsService) checkBasicInfo(req *model.ToolBasicInfo) error {
	if req.Name == "" || req.URL == "" {
		return errors.ErrEmptyToolParams
	}
	if req.Profile != "" && !storage.Get().Contains(req.Profile) {
		return errors.ErrProfileNotExists
	}
	return nil
}

func (t *toolsService) containsEnv(env string) bool {
	for i := range Environments {
		if Environments[i] == env {
			return true
		}
	}
	return false
}

func NewToolsService() IToolsService {
	return &toolsService{}
}
