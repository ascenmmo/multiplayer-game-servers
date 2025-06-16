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
type DevToolsServer interface {
	// @tg http-headers=token|Token
	// @tg summary=`server addServer`
	AddServer(ctx context.Context, token string, name string, address string) (err error)
	// @tg http-headers=token|Token
	// @tg summary=`server getServers`
	GetServers(ctx context.Context, token string) (servers []types.Server, err error)
	// @tg http-headers=token|Token
	// @tg summary=`server getServers`
	DeleteServers(ctx context.Context, token string, serverID uuid.UUID) (err error)
}
