package service

import (
	"toolbox/internal/errors"
	"toolbox/internal/model"
	"toolbox/internal/repository"
	rmodel "toolbox/internal/repository/model"
)

type IColumnsService interface {
	Put(name string) error
	Delete(name string) error

	Get(name string) (*model.Column, error)

	List() ([]*model.Column, error)
}

type columnService struct{}

func (c *columnService) Put(name string) error {
	if name == "" {
		return errors.ErrEmptyColumnParams
	}
	err := repository.GetColumnsRepository().Set(&rmodel.Column{
		Name: name,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *columnService) Delete(name string) error {
	tools, err := repository.GetToolsRepository().List()
	if err != nil {
		return err
	}
	for _, tool := range tools {
		if tool.ColumnName == name {
			return errors.ErrColumnIsNotEmpty
		}
	}

	err = repository.GetColumnsRepository().Delete(name)
	if err != nil {
		return err
	}
	return nil
}

func (c *columnService) Get(name string) (*model.Column, error) {
	column, err := repository.GetColumnsRepository().Get(name)
	if err != nil {
		return nil, err
	}
	return &model.Column{
		Name: column.Name,
	}, nil
}

func (c *columnService) List() ([]*model.Column, error) {
	columns, err := repository.GetColumnsRepository().List()
	if err != nil {
		return nil, err
	}
	ans := make([]*model.Column, len(columns))
	for i := range columns {
		ans[i] = &model.Column{
			Name: columns[i].Name,
		}
	}
	return ans, nil
}

func NewColumnsService() IColumnsService {
	return &columnService{}
}
