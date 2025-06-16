// @tg version=1.0.3
// @tg backend="Asenmmo multiplayer"
// @tg title=`Ascenmmo API for clients`
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
type DevToolsClient interface {
	// @tg summary=`server signUp`
	SignUp(ctx context.Context, client types.Client) (token string, refresh string, err error)
	// @tg summary=`server signIn`
	SignIn(ctx context.Context, client types.Client) (token, refresh string, err error)
	// @tg http-headers=token|Token
	// @tg http-headers=refresh|RefreshToken
	// @tg summary=`server refreshToken`
	RefreshToken(ctx context.Context, token string, refresh string) (newToken, newRefresh string, err error)
	// @tg http-headers=token|Token
	// @tg summary=`server getClient`
	GetClient(ctx context.Context, token string, gameID uuid.UUID) (client types.Client, err error)
	// @tg http-headers=token|Token
	// @tg summary=`server updateClient`
	UpdateClient(ctx context.Context, token string, client types.Client) (err error)
	// @tg http-headers=token|Token
	// @tg summary=`server getGameSaves`
	GetGameSaves(ctx context.Context, token string) (gameSaves types.GameSaves, err error)
	// @tg http-headers=token|Token
	// @tg summary=`server setGameSaves`
	SetGameSaves(ctx context.Context, token string, gameSaves types.GameSaves) (err error)
	// @tg http-headers=token|Token
	// @tg summary=`server deleteGameSaves`
	DeleteGameSaves(ctx context.Context, token string) (err error)
}
