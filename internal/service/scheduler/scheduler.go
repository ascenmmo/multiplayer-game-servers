package scheduler

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/internal/storage"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	tokentype "github.com/ascenmmo/token-generator/token_type"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"time"
)

type Scheduler interface {
	Run(ctx context.Context)
}

type scheduler struct {
	gameStorage       storage.GameStorage
	roomStorage       storage.RoomsStorage
	serverStorage     storage.ServersStorage
	confResultStorage storage.GameConfigsResultsStorage
	tokenGenerator    tokengenerator.TokenGenerator
}

func (s scheduler) Run(ctx context.Context) {
	ticker := time.NewTicker(time.Second * 1)
	for range ticker.C {
		err := s.getConfigResultsFromServer(ctx)
		if err != nil {
			log.Error().Err(err).Msg("error getting config results from server")
		}
	}
}

func (s scheduler) getConfigResultsFromServer(ctx context.Context) (err error) {
	serverIDs, err := s.serverStorage.FindAllServerIDs()
	if err != nil {
		return
	}

	token, _ := s.tokenGenerator.GenerateToken(tokentype.Info{
		GameID: uuid.New(),
		RoomID: uuid.New(),
		UserID: uuid.New(),
	}, tokengenerator.JWT)

	var batchServerIDs [][]uuid.UUID
	batchSize := 10

	for i := 0; i < len(serverIDs); i += batchSize {
		end := i + batchSize

		if end > len(serverIDs) {
			end = len(serverIDs)
		}

		batchServerIDs = append(batchServerIDs, serverIDs[i:end])
	}

	for _, batch := range batchServerIDs {
		servers, err := s.serverStorage.FindByIDs(batch)
		if err != nil {
			continue
		}

		for _, server := range servers {
			if server.IsActive {
				continue
			}
			results, err := server.GetGameConfigResults(ctx, token)
			if err != nil {
				continue
			}

			if len(results) == 0 {
				continue
			}

			err = s.confResultStorage.CreateMany(results)
			if err != nil {
				return err
			}

			for _, result := range results {
				err := s.roomStorage.Delete(result.RoomID)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func NewScheduler(gameStorage storage.GameStorage, roomStorage storage.RoomsStorage, serverStorage storage.ServersStorage, confResultStorage storage.GameConfigsResultsStorage, tokenGenerator tokengenerator.TokenGenerator) *scheduler {
	return &scheduler{gameStorage: gameStorage, roomStorage: roomStorage, serverStorage: serverStorage, confResultStorage: confResultStorage, tokenGenerator: tokenGenerator}
}
