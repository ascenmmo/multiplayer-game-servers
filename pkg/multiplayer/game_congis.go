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
type DevToolsGameConfigs interface {
	// @tg http-headers=token|Token
	// @tg summary=`CreateOrUpdateConfig`
	CreateOrUpdateConfig(ctx context.Context, token string, configs types.GameConfigs) (err error)
	// @tg http-headers=token|Token
	// @tg summary=`CreateOrUpdateConfig`
	GetGameConfig(ctx context.Context, token string, gameID uuid.UUID) (configs types.GameConfigs, err error)
}
