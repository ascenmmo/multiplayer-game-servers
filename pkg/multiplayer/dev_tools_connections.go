// @tg version=1.0.3
// @tg backend="Asenmmo multiplayer"
// @tg title=`Ascenmmo API for connection`
// @tg servers=`http://stage.ascenmmo.com;stage cluster`
//
//go:generate tg transport --services . --out ../../pkg/transport --outSwagger ../../pkg/swagger.yaml
//go:generate tg client -go --services . --outPath ../../pkg/clients/multiplayer

package multiplayer

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
)

// @tg http-prefix=api/v1
// @tg jsonRPC-server log trace metrics
// @tg tagNoOmitempty
type DevToolsConnections interface {
	// @tg http-headers=token|Token
	// @tg summary=`server createRoom`
	CreateRoom(ctx context.Context, token string, gameID uuid.UUID) (newToken string, err error)
	// @tg http-headers=token|Token
	// @tg summary=`server getRoomsAll`
	GetRoomsAll(ctx context.Context, token string, gameID uuid.UUID) (rooms []types.Room, err error)
	// @tg http-headers=token|Token
	// @tg summary=`server joinRoomByID`
	JoinRoomByID(ctx context.Context, token string, gameID uuid.UUID, roomID uuid.UUID) (newToken string, err error)
	// @tg http-headers=token|Token
	// @tg summary=`server removeRoomByID`
	RemoveRoomByID(ctx context.Context, token string, gameID uuid.UUID, roomID uuid.UUID) (err error)
	// @tg http-headers=token|Token
	// @tg summary=`server getRoomsConnectionUrls`
	GetRoomsConnectionUrls(ctx context.Context, token string) (connectionsServer []types.ConnectionServer, err error)
}
