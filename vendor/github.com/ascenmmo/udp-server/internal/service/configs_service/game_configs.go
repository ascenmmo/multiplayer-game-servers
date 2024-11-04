package configsService

import (
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	tokentype "github.com/ascenmmo/token-generator/token_type"
	memoryDB "github.com/ascenmmo/udp-server/internal/storage"
	"github.com/ascenmmo/udp-server/internal/utils"
	"github.com/ascenmmo/udp-server/pkg/restconnection/types"
	"github.com/google/uuid"
)

type GameConfigsService interface {
	Do(token string, clientInfo tokentype.Info, gameConfig types.GameConfigs, data interface{})
	GetDeletedRoomsResults(clientInfo tokentype.Info, onlinePlayersTokens []string) (results []types.GameConfigResults, ok bool)
	SetServerExecuteToGameConfig(clientInfo tokentype.Info, gameConfig types.GameConfigs) (newGameConfig types.GameConfigs)
}

type gameConfig struct {
	allGameConfigExecutor []types.GameConfigExecutor
	token                 tokengenerator.TokenGenerator
	storage               memoryDB.IMemoryDB
}

func (g *gameConfig) Do(token string, clientInfo tokentype.Info, gameConfig types.GameConfigs, data interface{}) {
	if !gameConfig.IsExists {
		return
	}
	results, _ := g.getOldResults(clientInfo)
	for _, sorting := range gameConfig.SortingConfig {
		if sorting.Executor == nil {
			continue
		}
		newResult, err := sorting.Executor.Execute(clientInfo, sorting, data, results)
		if err != nil {
			continue
		}
		results = newResult
	}
	g.storage.AddConnection(token)
	g.serOldResults(clientInfo, results)
}

func (g *gameConfig) GetDeletedRoomsResults(_ tokentype.Info, onlinePlayersTokens []string) (results []types.GameConfigResults, ok bool) {
	ids := g.storage.GetAllConnection()

	uniqueTokens := make(map[string]struct{})
	for _, token := range onlinePlayersTokens {
		uniqueTokens[token] = struct{}{}
	}

	var offlinePlayers []string
	for _, token := range ids {
		if _, exists := uniqueTokens[token]; !exists {
			offlinePlayers = append(offlinePlayers, token)
		}
	}

	clientsInfo := make(map[string]tokentype.Info)
	for _, token := range offlinePlayers {
		info, _ := g.token.ParseToken(token)
		if info.RoomID != uuid.Nil {
			clientsInfo[info.RoomID.String()] = info
		}

	}

	for _, info := range clientsInfo {
		configResults, ok := g.getOldResults(info)
		if !ok {
			continue
		}
		results = append(results, configResults)
	}

	return results, len(results) > 0
}

func (g *gameConfig) SetServerExecuteToGameConfig(_ tokentype.Info, gameConfig types.GameConfigs) (newGameConfig types.GameConfigs) {
	isConfigExecutorFound := false
	for i, conf := range gameConfig.SortingConfig {
		for _, executor := range g.allGameConfigExecutor {
			if executor.Name() == conf.Name {
				gameConfig.SortingConfig[i].Executor = executor
				isConfigExecutorFound = true
			}
		}
	}
	gameConfig.IsExists = isConfigExecutorFound
	return gameConfig
}

func (g *gameConfig) getOldResults(clientInfo tokentype.Info) (configResults types.GameConfigResults, ok bool) {
	key := utils.GenerateRoomKey(clientInfo)
	data, ok := g.storage.GetData(key)
	if ok {
		if configResults, ok = data.(types.GameConfigResults); ok {
			return configResults, true
		}
	}
	return types.GameConfigResults{
		GameID: clientInfo.GameID,
		RoomID: clientInfo.RoomID,
		Result: make(map[string]interface{}),
	}, false
}

func (g *gameConfig) serOldResults(clientInfo tokentype.Info, configResults types.GameConfigResults) {
	key := utils.GenerateRoomKey(clientInfo)
	g.storage.SetData(key, configResults)
}

func getAllFunctions() []types.GameConfigExecutor {
	return GetCountingFunctions()
}

func NewGameConfigsService(storage memoryDB.IMemoryDB, token tokengenerator.TokenGenerator) GameConfigsService {
	return &gameConfig{allGameConfigExecutor: getAllFunctions(), storage: storage, token: token}
}
