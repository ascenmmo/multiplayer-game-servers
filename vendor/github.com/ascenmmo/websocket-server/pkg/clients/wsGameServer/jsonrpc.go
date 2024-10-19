// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package wsGameServer

import (
	"context"
	"github.com/ascenmmo/websocket-server/pkg/clients/wsGameServer/cb"
	"github.com/ascenmmo/websocket-server/pkg/clients/wsGameServer/jsonrpc"
	"os"
	"strconv"
	"time"
)

type ClientJsonRPC struct {
	name string

	rpc     *jsonrpc.ClientRPC
	rpcOpts []jsonrpc.Option

	cache cache

	cbCfg cb.Settings
	cb    *cb.CircuitBreaker

	fallbackTTL            time.Duration
	fallbackServerSettings fallbackServerSettings

	errorDecoder ErrorDecoder
}

func New(endpoint string, opts ...Option) (cli *ClientJsonRPC) {

	hostname, _ := os.Hostname()
	cli = &ClientJsonRPC{
		errorDecoder: defaultErrorDecoder,
		fallbackTTL:  time.Hour * 24,
		name:         hostname + "_" + "github.com/ascenmmo/websocket-server",
	}
	cli.applyOpts(opts)
	cli.rpc = jsonrpc.NewClient(endpoint, cli.rpcOpts...)
	cli.cb = cb.NewCircuitBreaker("github.com/ascenmmo/websocket-server", cli.cbCfg)
	return
}

func (cli *ClientJsonRPC) ServerSettings() *ClientServerSettings {
	return &ClientServerSettings{ClientJsonRPC: cli}
}

func (cli *ClientJsonRPC) proceedResponse(ctx context.Context, httpErr error, cacheKey uint64, fallbackCheck func(error) bool, rpcResponse *jsonrpc.ResponseRPC, methodResponse interface{}) (err error) {

	err = cli.cb.Execute(func() (err error) {
		if httpErr != nil {
			return httpErr
		}
		return rpcResponse.GetObject(&methodResponse)
	}, cb.IsSuccessful(func(err error) (success bool) {
		if fallbackCheck != nil {
			return fallbackCheck(err)
		}
		if success = err == nil; success {
			if cli.cache != nil && cacheKey != 0 {
				_ = cli.cache.SetTTL(ctx, strconv.FormatUint(cacheKey, 10), methodResponse, cli.fallbackTTL)
			}
		}
		return
	}), cb.Fallback(func(err error) error {
		if cli.cache != nil && cacheKey != 0 {
			_, _, err = cli.cache.GetTTL(ctx, strconv.FormatUint(cacheKey, 10), &methodResponse)
		}
		return err
	}))
	return
}
