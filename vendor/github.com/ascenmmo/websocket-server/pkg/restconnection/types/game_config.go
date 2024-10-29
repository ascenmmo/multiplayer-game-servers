package types

import (
	tokentype "github.com/ascenmmo/token-generator/token_type"
	"github.com/google/uuid"
)

type GameConfigExecutor interface {
	Name() string
	Execute(clientInfo tokentype.Info, config SortingConfig, newUserData interface{}, oldResult GameConfigResults) (newResult GameConfigResults, err error)
}

type GameConfigResults struct {
	GameID uuid.UUID              `json:"game_id"`
	RoomID uuid.UUID              `json:"room_id"`
	Result map[string]interface{} `json:"result"`
}

type GameConfigs struct {
	GameID        uuid.UUID       `json:"game_id" bson:"_id"`
	SortingConfig []SortingConfig `json:"sorting_config" bson:"sorting_config"`
	IsExists      bool            `json:"-"`
}

type SortingConfig struct {
	Name            string             `json:"name" bson:"name"`
	Params          []ParamMetadata    `json:"params" bson:"params"`
	UseOnServerType string             `json:"use_on_server_type" bson:"use_on_server_type"`
	ResultName      string             `json:"result_name" bson:"result_name"`
	ResultType      string             `json:"result_type" bson:"result_type"`
	Executor        GameConfigExecutor `json:"-"`
}

type ParamMetadata struct {
	ColumnName string `json:"column_name" bson:"column_name"`
	ValueType  string `json:"value_type" bson:"value_type"`
}

func (g *GameConfigs) RemoveSortingConfig(name string) {
	for i, v := range g.SortingConfig {
		if v.Name == name {
			g.SortingConfig = append(g.SortingConfig[:i], g.SortingConfig[i+1:]...)
		}
	}
}
