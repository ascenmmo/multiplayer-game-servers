// @tg version=1.0.3
// @tg backend="Asenmmo"
// @tg title=`Ascenmmo Rest API`
// @tg servers=`http://stage.ascenmmo.com;stage cluster`
//
//go:generate tg transport --services . --out ../../pkg/transport --outSwagger ../../pkg/swagger.yaml
//go:generate tg client -go --services . --outPath ../../pkg/clients/tcpGameServer

package api

import (
	"context"
	"github.com/ascenmmo/tcp-server/pkg/api/types"
	"github.com/google/uuid"
)

// @tg http-prefix=api/v1/rest/
// @tg jsonRPC-server log trace
// @tg tagNoOmitempty
type GameConnections interface {
	// @tg http-headers=token|Token
	// @tg summary=`SetUsersAndMessage`
	SetSendMessage(ctx context.Context, token string, message types.RequestSetMessage) (err error)
	// @tg http-headers=token|Token
	// @tg summary=`GetUserMessage`
	GetMessage(ctx context.Context, token string) (messages types.ResponseGetMessage, err error)
	// @tg http-headers=token|Token
	// @tg summary=`RemoveUser`
	RemoveUser(ctx context.Context, token string, userID uuid.UUID) (err error)
}
