// @tg version=1.0.3
// @tg backend="Asenmmo"
// @tg title=`Ascenmmo Rest API`
// @tg servers=`http://stage.ascenmmo.com;stage cluster`

package api

import (
	"context"

	"github.com/google/uuid"

	"github.com/ascenmmo/tcp-server/pkg/api/types"
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
