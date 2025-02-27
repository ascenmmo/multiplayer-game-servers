// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package udpGameServer

import (
	"context"
	"fmt"
	"github.com/ascenmmo/udp-server/pkg/api/types"
	"github.com/ascenmmo/udp-server/pkg/clients/udpGameServer/hasher"
	"github.com/ascenmmo/udp-server/pkg/clients/udpGameServer/jsonrpc"
)

type ClientServerSettings struct {
	*ClientJsonRPC
}

type retServerSettingsGetConnectionsNum = func(countConn int, exists bool, err error)
type retServerSettingsHealthCheck = func(exists bool, err error)
type retServerSettingsGetServerSettings = func(settings types.Settings, err error)
type retServerSettingsCreateRoom = func(err error)
type retServerSettingsGetDeletedRooms = func(deletedIds []types.GetDeletedRooms, err error)

func (cli *ClientServerSettings) GetConnectionsNum(ctx context.Context, token string) (countConn int, exists bool, err error) {

	request := requestServerSettingsGetConnectionsNum{Token: token}
	var response responseServerSettingsGetConnectionsNum
	var rpcResponse *jsonrpc.ResponseRPC
	cacheKey, _ := hasher.Hash(request)
	rpcResponse, err = cli.rpc.Call(ctx, "serversettings.getconnectionsnum", request)
	var fallbackCheck func(error) bool
	if cli.fallbackServerSettings != nil {
		fallbackCheck = cli.fallbackServerSettings.GetConnectionsNum
	}
	if rpcResponse != nil && rpcResponse.Error != nil {
		if cli.errorDecoder != nil {
			err = cli.errorDecoder(rpcResponse.Error.Raw())
		} else {
			err = fmt.Errorf(rpcResponse.Error.Message)
		}
	}
	if err = cli.proceedResponse(ctx, err, cacheKey, fallbackCheck, rpcResponse, &response); err != nil {
		return
	}
	return response.CountConn, response.Exists, err
}

func (cli *ClientServerSettings) ReqGetConnectionsNum(ctx context.Context, callback retServerSettingsGetConnectionsNum, token string) (request RequestRPC) {

	request = RequestRPC{rpcRequest: &jsonrpc.RequestRPC{
		ID:      jsonrpc.NewID(),
		JSONRPC: jsonrpc.Version,
		Method:  "serversettings.getconnectionsnum",
		Params:  requestServerSettingsGetConnectionsNum{Token: token},
	}}
	if callback != nil {
		var response responseServerSettingsGetConnectionsNum
		request.retHandler = func(err error, rpcResponse *jsonrpc.ResponseRPC) {
			cacheKey, _ := hasher.Hash(request.rpcRequest.Params)
			var fallbackCheck func(error) bool
			if cli.fallbackServerSettings != nil {
				fallbackCheck = cli.fallbackServerSettings.GetConnectionsNum
			}
			if rpcResponse != nil && rpcResponse.Error != nil {
				if cli.errorDecoder != nil {
					err = cli.errorDecoder(rpcResponse.Error.Raw())
				} else {
					err = fmt.Errorf(rpcResponse.Error.Message)
				}
			}
			callback(response.CountConn, response.Exists, cli.proceedResponse(ctx, err, cacheKey, fallbackCheck, rpcResponse, &response))
		}
	}
	return
}

func (cli *ClientServerSettings) HealthCheck(ctx context.Context, token string) (exists bool, err error) {

	request := requestServerSettingsHealthCheck{Token: token}
	var response responseServerSettingsHealthCheck
	var rpcResponse *jsonrpc.ResponseRPC
	cacheKey, _ := hasher.Hash(request)
	rpcResponse, err = cli.rpc.Call(ctx, "serversettings.healthcheck", request)
	var fallbackCheck func(error) bool
	if cli.fallbackServerSettings != nil {
		fallbackCheck = cli.fallbackServerSettings.HealthCheck
	}
	if rpcResponse != nil && rpcResponse.Error != nil {
		if cli.errorDecoder != nil {
			err = cli.errorDecoder(rpcResponse.Error.Raw())
		} else {
			err = fmt.Errorf(rpcResponse.Error.Message)
		}
	}
	if err = cli.proceedResponse(ctx, err, cacheKey, fallbackCheck, rpcResponse, &response); err != nil {
		return
	}
	return response.Exists, err
}

func (cli *ClientServerSettings) ReqHealthCheck(ctx context.Context, callback retServerSettingsHealthCheck, token string) (request RequestRPC) {

	request = RequestRPC{rpcRequest: &jsonrpc.RequestRPC{
		ID:      jsonrpc.NewID(),
		JSONRPC: jsonrpc.Version,
		Method:  "serversettings.healthcheck",
		Params:  requestServerSettingsHealthCheck{Token: token},
	}}
	if callback != nil {
		var response responseServerSettingsHealthCheck
		request.retHandler = func(err error, rpcResponse *jsonrpc.ResponseRPC) {
			cacheKey, _ := hasher.Hash(request.rpcRequest.Params)
			var fallbackCheck func(error) bool
			if cli.fallbackServerSettings != nil {
				fallbackCheck = cli.fallbackServerSettings.HealthCheck
			}
			if rpcResponse != nil && rpcResponse.Error != nil {
				if cli.errorDecoder != nil {
					err = cli.errorDecoder(rpcResponse.Error.Raw())
				} else {
					err = fmt.Errorf(rpcResponse.Error.Message)
				}
			}
			callback(response.Exists, cli.proceedResponse(ctx, err, cacheKey, fallbackCheck, rpcResponse, &response))
		}
	}
	return
}

func (cli *ClientServerSettings) GetServerSettings(ctx context.Context, token string) (settings types.Settings, err error) {

	request := requestServerSettingsGetServerSettings{Token: token}
	var response responseServerSettingsGetServerSettings
	var rpcResponse *jsonrpc.ResponseRPC
	cacheKey, _ := hasher.Hash(request)
	rpcResponse, err = cli.rpc.Call(ctx, "serversettings.getserversettings", request)
	var fallbackCheck func(error) bool
	if cli.fallbackServerSettings != nil {
		fallbackCheck = cli.fallbackServerSettings.GetServerSettings
	}
	if rpcResponse != nil && rpcResponse.Error != nil {
		if cli.errorDecoder != nil {
			err = cli.errorDecoder(rpcResponse.Error.Raw())
		} else {
			err = fmt.Errorf(rpcResponse.Error.Message)
		}
	}
	if err = cli.proceedResponse(ctx, err, cacheKey, fallbackCheck, rpcResponse, &response); err != nil {
		return
	}
	return response.Settings, err
}

func (cli *ClientServerSettings) ReqGetServerSettings(ctx context.Context, callback retServerSettingsGetServerSettings, token string) (request RequestRPC) {

	request = RequestRPC{rpcRequest: &jsonrpc.RequestRPC{
		ID:      jsonrpc.NewID(),
		JSONRPC: jsonrpc.Version,
		Method:  "serversettings.getserversettings",
		Params:  requestServerSettingsGetServerSettings{Token: token},
	}}
	if callback != nil {
		var response responseServerSettingsGetServerSettings
		request.retHandler = func(err error, rpcResponse *jsonrpc.ResponseRPC) {
			cacheKey, _ := hasher.Hash(request.rpcRequest.Params)
			var fallbackCheck func(error) bool
			if cli.fallbackServerSettings != nil {
				fallbackCheck = cli.fallbackServerSettings.GetServerSettings
			}
			if rpcResponse != nil && rpcResponse.Error != nil {
				if cli.errorDecoder != nil {
					err = cli.errorDecoder(rpcResponse.Error.Raw())
				} else {
					err = fmt.Errorf(rpcResponse.Error.Message)
				}
			}
			callback(response.Settings, cli.proceedResponse(ctx, err, cacheKey, fallbackCheck, rpcResponse, &response))
		}
	}
	return
}

func (cli *ClientServerSettings) CreateRoom(ctx context.Context, token string, createRoom types.CreateRoomRequest) (err error) {

	request := requestServerSettingsCreateRoom{
		CreateRoom: createRoom,
		Token:      token,
	}
	var response responseServerSettingsCreateRoom
	var rpcResponse *jsonrpc.ResponseRPC
	cacheKey, _ := hasher.Hash(request)
	rpcResponse, err = cli.rpc.Call(ctx, "serversettings.createroom", request)
	var fallbackCheck func(error) bool
	if cli.fallbackServerSettings != nil {
		fallbackCheck = cli.fallbackServerSettings.CreateRoom
	}
	if rpcResponse != nil && rpcResponse.Error != nil {
		if cli.errorDecoder != nil {
			err = cli.errorDecoder(rpcResponse.Error.Raw())
		} else {
			err = fmt.Errorf(rpcResponse.Error.Message)
		}
	}
	if err = cli.proceedResponse(ctx, err, cacheKey, fallbackCheck, rpcResponse, &response); err != nil {
		return
	}
	return err
}

func (cli *ClientServerSettings) ReqCreateRoom(ctx context.Context, callback retServerSettingsCreateRoom, token string, createRoom types.CreateRoomRequest) (request RequestRPC) {

	request = RequestRPC{rpcRequest: &jsonrpc.RequestRPC{
		ID:      jsonrpc.NewID(),
		JSONRPC: jsonrpc.Version,
		Method:  "serversettings.createroom",
		Params: requestServerSettingsCreateRoom{
			CreateRoom: createRoom,
			Token:      token,
		},
	}}
	if callback != nil {
		var response responseServerSettingsCreateRoom
		request.retHandler = func(err error, rpcResponse *jsonrpc.ResponseRPC) {
			cacheKey, _ := hasher.Hash(request.rpcRequest.Params)
			var fallbackCheck func(error) bool
			if cli.fallbackServerSettings != nil {
				fallbackCheck = cli.fallbackServerSettings.CreateRoom
			}
			if rpcResponse != nil && rpcResponse.Error != nil {
				if cli.errorDecoder != nil {
					err = cli.errorDecoder(rpcResponse.Error.Raw())
				} else {
					err = fmt.Errorf(rpcResponse.Error.Message)
				}
			}
			callback(cli.proceedResponse(ctx, err, cacheKey, fallbackCheck, rpcResponse, &response))
		}
	}
	return
}

func (cli *ClientServerSettings) GetDeletedRooms(ctx context.Context, token string, ids []types.GetDeletedRooms) (deletedIds []types.GetDeletedRooms, err error) {

	request := requestServerSettingsGetDeletedRooms{
		Ids:   ids,
		Token: token,
	}
	var response responseServerSettingsGetDeletedRooms
	var rpcResponse *jsonrpc.ResponseRPC
	cacheKey, _ := hasher.Hash(request)
	rpcResponse, err = cli.rpc.Call(ctx, "serversettings.getdeletedrooms", request)
	var fallbackCheck func(error) bool
	if cli.fallbackServerSettings != nil {
		fallbackCheck = cli.fallbackServerSettings.GetDeletedRooms
	}
	if rpcResponse != nil && rpcResponse.Error != nil {
		if cli.errorDecoder != nil {
			err = cli.errorDecoder(rpcResponse.Error.Raw())
		} else {
			err = fmt.Errorf(rpcResponse.Error.Message)
		}
	}
	if err = cli.proceedResponse(ctx, err, cacheKey, fallbackCheck, rpcResponse, &response); err != nil {
		return
	}
	return response.DeletedIds, err
}

func (cli *ClientServerSettings) ReqGetDeletedRooms(ctx context.Context, callback retServerSettingsGetDeletedRooms, token string, ids []types.GetDeletedRooms) (request RequestRPC) {

	request = RequestRPC{rpcRequest: &jsonrpc.RequestRPC{
		ID:      jsonrpc.NewID(),
		JSONRPC: jsonrpc.Version,
		Method:  "serversettings.getdeletedrooms",
		Params: requestServerSettingsGetDeletedRooms{
			Ids:   ids,
			Token: token,
		},
	}}
	if callback != nil {
		var response responseServerSettingsGetDeletedRooms
		request.retHandler = func(err error, rpcResponse *jsonrpc.ResponseRPC) {
			cacheKey, _ := hasher.Hash(request.rpcRequest.Params)
			var fallbackCheck func(error) bool
			if cli.fallbackServerSettings != nil {
				fallbackCheck = cli.fallbackServerSettings.GetDeletedRooms
			}
			if rpcResponse != nil && rpcResponse.Error != nil {
				if cli.errorDecoder != nil {
					err = cli.errorDecoder(rpcResponse.Error.Raw())
				} else {
					err = fmt.Errorf(rpcResponse.Error.Message)
				}
			}
			callback(response.DeletedIds, cli.proceedResponse(ctx, err, cacheKey, fallbackCheck, rpcResponse, &response))
		}
	}
	return
}
