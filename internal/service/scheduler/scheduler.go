package scheduler

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/internal/storage"
	"github.com/google/uuid"
)

type Scheduler interface {
}

type scheduler struct {
	gameStorage       storage.GameStorage
	roomStorage       storage.RoomsStorage
	serverStorage     storage.ServersStorage
	confResultStorage storage.GameConfigsResultsStorage
}

func (s scheduler) Run(ctx context.Context) {

}

func (s scheduler) getConfigResultsFromServer(ctx context.Context) (err error) {
	serverIDs, err := s.serverStorage.FindAllServerIDs()
	if err != nil {
		return
	}

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
			data, err := server.IsExists()

		}

	}
}

func NewScheduler() Scheduler {

}
