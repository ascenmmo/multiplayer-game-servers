// @tg version=1.0.3
// @tg backend="Asenmmo multiplayer"
// @tg title=`Ascenmmo API for connection`
// @tg servers=`http://stage.ascenmmo.com;stage cluster`

package multiplayer

import (
	"context"

	"github.com/google/uuid"

	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
)

// DevToolsGameConfigs
// @tg http-prefix=api/v1
// @tg jsonRPC-server log trace metrics
// @tg tagNoOmitempty
type DevToolsGameConfigs interface {
	// @tg http-headers=token|Token
	// @tg summary=`CreateOrUpdateConfig`
	CreateOrUpdateConfig(ctx context.Context, token string, configs types.GameConfigs) (err error)
	// @tg http-headers=token|Token
	// @tg summary=`GetGameConfig`
	GetGameConfig(ctx context.Context, token string, gameID uuid.UUID) (configs types.GameConfigs, err error)
	// @tg http-headers=token|Token
	// @tg summary=`GetGameResultConfigPreview`
	GetGameResultConfigPreview(ctx context.Context, token string, gameID uuid.UUID) (gameResult types.GameConfigResults, err error)
}
