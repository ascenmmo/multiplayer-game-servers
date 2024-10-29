package configsService

import (
	"github.com/ascenmmo/tcp-server/pkg/errors"
	"github.com/ascenmmo/tcp-server/pkg/restconnection/types"
	tokentype "github.com/ascenmmo/token-generator/token_type"
)

func GetCountingFunctions() (functions []types.GameConfigExecutor) {
	functions = []types.GameConfigExecutor{
		&IncrementResult{},
		&DecrementResult{},
		&AdditionDataResultToOld{},
		&SubtractDataResultToOld{},
	}
	return functions
}

type IncrementResult struct {
}

func (c *IncrementResult) Name() string {
	return "IncrementResult"
}

func (c *IncrementResult) Execute(clientInfo tokentype.Info, config types.SortingConfig, newUserData interface{}, oldResult types.GameConfigResults) (newResult types.GameConfigResults, err error) {
	data, err := parseUserData(newUserData)
	if err != nil {
		return oldResult, errors.ErrGameConfigMarshalUserData
	}

	if len(oldResult.Result) == 0 {
		oldResult.Result = make(map[string]interface{})
	}

	var values []interface{}
	for _, v := range config.Params {
		value, ok := data[v.ColumnName]
		if !ok {
			continue
		}
		values = append(values, value)
	}

	if len(values) != len(oldResult.Result) {
		return oldResult, nil
	}

	oldValue, ok := oldResult.Result[config.ResultName]
	if !ok {
		oldValue = 0
	}

	oldResult.Result[config.ResultName] = incrementValue(oldValue)

	return oldResult, nil
}

type DecrementResult struct {
}

func (c *DecrementResult) Name() string {
	return "DecrementResult"
}

func (c *DecrementResult) Execute(clientInfo tokentype.Info, config types.SortingConfig, newUserData interface{}, oldResult types.GameConfigResults) (newResult types.GameConfigResults, err error) {
	data, err := parseUserData(newUserData)
	if err != nil {
		return oldResult, errors.ErrGameConfigMarshalUserData
	}

	if len(oldResult.Result) == 0 {
		oldResult.Result = make(map[string]interface{})
	}

	var values []interface{}
	for _, v := range config.Params {
		value, ok := data[v.ColumnName]
		if !ok {
			continue
		}
		values = append(values, value)
	}

	if len(values) != len(oldResult.Result) {
		return oldResult, nil
	}

	oldValue, ok := oldResult.Result[config.ResultName]
	if !ok {
		oldValue = 0
	}

	oldResult.Result[config.ResultName] = decrementValue(oldValue)

	return oldResult, nil
}

type AdditionDataResultToOld struct{}

func (c *AdditionDataResultToOld) Name() string {
	return "AdditionDataResultToOld"
}

func (c *AdditionDataResultToOld) Execute(clientInfo tokentype.Info, config types.SortingConfig, newUserData interface{}, oldResult types.GameConfigResults) (newResult types.GameConfigResults, err error) {
	data, err := parseUserData(newUserData)
	if err != nil {
		return oldResult, errors.ErrGameConfigMarshalUserData
	}

	if len(oldResult.Result) == 0 {
		oldResult.Result = make(map[string]interface{})
	}

	var values []interface{}
	for _, v := range config.Params {
		value, ok := data[v.ColumnName]
		if !ok {
			continue
		}
		values = append(values, value)
	}

	if len(values) != len(config.Params) {
		return oldResult, nil
	}

	oldValue, ok := oldResult.Result[config.ResultName]
	if !ok {
		oldValue = 0
	}

	for _, newValue := range values {
		oldValue = additionValues(oldValue, newValue, config.ResultType)
	}

	oldResult.Result[config.ResultName] = oldValue

	return oldResult, nil
}

type SubtractDataResultToOld struct{}

func (c *SubtractDataResultToOld) Name() string {
	return "SubtractDataResultToOld"
}

func (c *SubtractDataResultToOld) Execute(clientInfo tokentype.Info, config types.SortingConfig, newUserData interface{}, oldResult types.GameConfigResults) (newResult types.GameConfigResults, err error) {
	data, err := parseUserData(newUserData)
	if err != nil {
		return oldResult, errors.ErrGameConfigMarshalUserData
	}

	if len(oldResult.Result) == 0 {
		oldResult.Result = make(map[string]interface{})
	}

	var values []interface{}
	for _, v := range config.Params {
		value, ok := data[v.ColumnName]
		if !ok {
			continue
		}
		values = append(values, value)
	}

	if len(values) != len(config.Params) {
		return oldResult, nil
	}

	oldValue, ok := oldResult.Result[config.ResultName]
	if !ok {
		oldValue = 0
	}

	for _, newValue := range values {
		oldValue = subtractValues(oldValue, newValue, config.ResultType)
	}

	oldResult.Result[config.ResultName] = oldValue

	return oldResult, nil
}

func parseUserData(newUserData interface{}) (map[string]interface{}, error) {
	data, ok := newUserData.(map[string]interface{})
	if !ok {
		return nil, errors.ErrGameConfigMarshalUserData
	}
	return data, nil
}
