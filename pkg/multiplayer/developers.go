// @tg version=1.0.3
// @tg backend="Asenmmo multiplayer"
// @tg title=`Ascenmmo API for devToolsServer`
// @tg servers=`http://stage.ascenmmo.com;stage cluster`

package multiplayer

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
)

// @tg http-prefix=api/v1
// @tg jsonRPC-server log trace metrics
// @tg tagNoOmitempty
type Developers interface {
	// @tg summary=`server signUp`
	SignUp(ctx context.Context, developer types.Developer) (token string, refresh string, err error)
	// @tg summary=`server signIn`
	SignIn(ctx context.Context, developer types.Developer) (token string, refresh string, err error)
	// @tg http-headers=token|Token
	// @tg http-headers=refresh|RefreshToken
	// @tg summary=`server refreshToken`
	RefreshToken(ctx context.Context, token string, refresh string) (newToken, newRefresh string, err error)
	// @tg http-headers=token|Token
	// @tg summary=`server getDeveloper`
	GetDeveloper(ctx context.Context, token string) (developer types.Developer, err error)
	// @tg http-headers=token|Token
	// @tg summary=`server updateDeveloper`
	UpdateDeveloper(ctx context.Context, token string, developer types.Developer) (err error)
}
