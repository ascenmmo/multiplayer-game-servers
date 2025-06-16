// @tg version=1.0.3
// @tg backend="Asenmmo multiplayer"
// @tg title=`Ascenmmo API for devToolsServer`
// @tg servers=`http://stage.ascenmmo.com;stage cluster`

package multiplayer

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
)

// @tg http-prefix=api/v1
// @tg jsonRPC-server log trace metrics
// @tg tagNoOmitempty
type DevTools interface {
	// @tg http-headers=token|Token
	// @tg summary=`server createGame`
	CreateGame(ctx context.Context, token string, game types.Game) (id uuid.UUID, err error)
	// @tg http-headers=token|Token
	// @tg summary=`server gameAddOwner`
	GameAddOwner(ctx context.Context, token string, gameID, userID uuid.UUID) (err error)
	// @tg http-headers=token|Token
	// @tg summary=`server иameRemoveOwner`
	GameRemoveOwner(ctx context.Context, token string, gameID, userID uuid.UUID) (err error)
	// @tg http-headers=token|Token
	// @tg summary=`server updateGame`
	UpdateGame(ctx context.Context, token string, gameID uuid.UUID, newGame types.Game) (id uuid.UUID, err error)
	// @tg http-headers=token|Token
	// @tg summary=`server deleteGame`
	DeleteGame(ctx context.Context, token string, gameID uuid.UUID) (err error)
	// @tg http-headers=token|Token
	// @tg summary=`server getMyGames`
	GetMyGames(ctx context.Context, token string) (games []types.Game, err error)
	// @tg http-headers=token|Token
	// @tg summary=`server getMyGame`
	GetGameByGameID(ctx context.Context, token string, gameID uuid.UUID) (game types.Game, err error)
	// @tg http-headers=token|Token
	// @tg summary=`server turnOnServerInGame`
	TurnOnServerInGame(ctx context.Context, token string, serverID uuid.UUID, gameId uuid.UUID) (err error)
	// @tg http-headers=token|Token
	// @tg summary=`server turnOnServerInGame`
	TurnOffServerInGame(ctx context.Context, token string, serverID uuid.UUID, gameId uuid.UUID) (err error)
}
