// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package multiplayer

import (
	"context"
	"fmt"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/clients/multiplayer/hasher"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/clients/multiplayer/jsonrpc"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
)

type ClientDevToolsConnections struct {
	*ClientJsonRPC
}

type retDevToolsConnectionsCreateRoom = func(newToken string, err error)
type retDevToolsConnectionsGetRoomsAll = func(rooms []types.Room, err error)
type retDevToolsConnectionsJoinRoomByID = func(newToken string, err error)
type retDevToolsConnectionsRemoveRoomByID = func(err error)
type retDevToolsConnectionsGetRoomsConnectionUrls = func(connectionsServer []types.ConnectionServer, err error)

func (cli *ClientDevToolsConnections) CreateRoom(ctx context.Context, token string, gameID uuid.UUID, serverType string) (newToken string, err error) {

	request := requestDevToolsConnectionsCreateRoom{
		GameID:     gameID,
		ServerType: serverType,
		Token:      token,
	}
	var response responseDevToolsConnectionsCreateRoom
	var rpcResponse *jsonrpc.ResponseRPC
	cacheKey, _ := hasher.Hash(request)
	rpcResponse, err = cli.rpc.Call(ctx, "devtoolsconnections.createroom", request)
	var fallbackCheck func(error) bool
	if cli.fallbackDevToolsConnections != nil {
		fallbackCheck = cli.fallbackDevToolsConnections.CreateRoom
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
	return response.NewToken, err
}

func (cli *ClientDevToolsConnections) ReqCreateRoom(ctx context.Context, callback retDevToolsConnectionsCreateRoom, token string, gameID uuid.UUID, serverType string) (request RequestRPC) {

	request = RequestRPC{rpcRequest: &jsonrpc.RequestRPC{
		ID:      jsonrpc.NewID(),
		JSONRPC: jsonrpc.Version,
		Method:  "devtoolsconnections.createroom",
		Params: requestDevToolsConnectionsCreateRoom{
			GameID:     gameID,
			ServerType: serverType,
			Token:      token,
		},
	}}
	if callback != nil {
		var response responseDevToolsConnectionsCreateRoom
		request.retHandler = func(err error, rpcResponse *jsonrpc.ResponseRPC) {
			cacheKey, _ := hasher.Hash(request.rpcRequest.Params)
			var fallbackCheck func(error) bool
			if cli.fallbackDevToolsConnections != nil {
				fallbackCheck = cli.fallbackDevToolsConnections.CreateRoom
			}
			if rpcResponse != nil && rpcResponse.Error != nil {
				if cli.errorDecoder != nil {
					err = cli.errorDecoder(rpcResponse.Error.Raw())
				} else {
					err = fmt.Errorf(rpcResponse.Error.Message)
				}
			}
			callback(response.NewToken, cli.proceedResponse(ctx, err, cacheKey, fallbackCheck, rpcResponse, &response))
		}
	}
	return
}

func (cli *ClientDevToolsConnections) GetRoomsAll(ctx context.Context, token string, gameID uuid.UUID) (rooms []types.Room, err error) {

	request := requestDevToolsConnectionsGetRoomsAll{
		GameID: gameID,
		Token:  token,
	}
	var response responseDevToolsConnectionsGetRoomsAll
	var rpcResponse *jsonrpc.ResponseRPC
	cacheKey, _ := hasher.Hash(request)
	rpcResponse, err = cli.rpc.Call(ctx, "devtoolsconnections.getroomsall", request)
	var fallbackCheck func(error) bool
	if cli.fallbackDevToolsConnections != nil {
		fallbackCheck = cli.fallbackDevToolsConnections.GetRoomsAll
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
	return response.Rooms, err
}

func (cli *ClientDevToolsConnections) ReqGetRoomsAll(ctx context.Context, callback retDevToolsConnectionsGetRoomsAll, token string, gameID uuid.UUID) (request RequestRPC) {

	request = RequestRPC{rpcRequest: &jsonrpc.RequestRPC{
		ID:      jsonrpc.NewID(),
		JSONRPC: jsonrpc.Version,
		Method:  "devtoolsconnections.getroomsall",
		Params: requestDevToolsConnectionsGetRoomsAll{
			GameID: gameID,
			Token:  token,
		},
	}}
	if callback != nil {
		var response responseDevToolsConnectionsGetRoomsAll
		request.retHandler = func(err error, rpcResponse *jsonrpc.ResponseRPC) {
			cacheKey, _ := hasher.Hash(request.rpcRequest.Params)
			var fallbackCheck func(error) bool
			if cli.fallbackDevToolsConnections != nil {
				fallbackCheck = cli.fallbackDevToolsConnections.GetRoomsAll
			}
			if rpcResponse != nil && rpcResponse.Error != nil {
				if cli.errorDecoder != nil {
					err = cli.errorDecoder(rpcResponse.Error.Raw())
				} else {
					err = fmt.Errorf(rpcResponse.Error.Message)
				}
			}
			callback(response.Rooms, cli.proceedResponse(ctx, err, cacheKey, fallbackCheck, rpcResponse, &response))
		}
	}
	return
}

func (cli *ClientDevToolsConnections) JoinRoomByID(ctx context.Context, token string, gameID uuid.UUID, roomID uuid.UUID) (newToken string, err error) {

	request := requestDevToolsConnectionsJoinRoomByID{
		GameID: gameID,
		RoomID: roomID,
		Token:  token,
	}
	var response responseDevToolsConnectionsJoinRoomByID
	var rpcResponse *jsonrpc.ResponseRPC
	cacheKey, _ := hasher.Hash(request)
	rpcResponse, err = cli.rpc.Call(ctx, "devtoolsconnections.joinroombyid", request)
	var fallbackCheck func(error) bool
	if cli.fallbackDevToolsConnections != nil {
		fallbackCheck = cli.fallbackDevToolsConnections.JoinRoomByID
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
	return response.NewToken, err
}

func (cli *ClientDevToolsConnections) ReqJoinRoomByID(ctx context.Context, callback retDevToolsConnectionsJoinRoomByID, token string, gameID uuid.UUID, roomID uuid.UUID) (request RequestRPC) {

	request = RequestRPC{rpcRequest: &jsonrpc.RequestRPC{
		ID:      jsonrpc.NewID(),
		JSONRPC: jsonrpc.Version,
		Method:  "devtoolsconnections.joinroombyid",
		Params: requestDevToolsConnectionsJoinRoomByID{
			GameID: gameID,
			RoomID: roomID,
			Token:  token,
		},
	}}
	if callback != nil {
		var response responseDevToolsConnectionsJoinRoomByID
		request.retHandler = func(err error, rpcResponse *jsonrpc.ResponseRPC) {
			cacheKey, _ := hasher.Hash(request.rpcRequest.Params)
			var fallbackCheck func(error) bool
			if cli.fallbackDevToolsConnections != nil {
				fallbackCheck = cli.fallbackDevToolsConnections.JoinRoomByID
			}
			if rpcResponse != nil && rpcResponse.Error != nil {
				if cli.errorDecoder != nil {
					err = cli.errorDecoder(rpcResponse.Error.Raw())
				} else {
					err = fmt.Errorf(rpcResponse.Error.Message)
				}
			}
			callback(response.NewToken, cli.proceedResponse(ctx, err, cacheKey, fallbackCheck, rpcResponse, &response))
		}
	}
	return
}

func (cli *ClientDevToolsConnections) RemoveRoomByID(ctx context.Context, token string, gameID uuid.UUID, roomID uuid.UUID) (err error) {

	request := requestDevToolsConnectionsRemoveRoomByID{
		GameID: gameID,
		RoomID: roomID,
		Token:  token,
	}
	var response responseDevToolsConnectionsRemoveRoomByID
	var rpcResponse *jsonrpc.ResponseRPC
	cacheKey, _ := hasher.Hash(request)
	rpcResponse, err = cli.rpc.Call(ctx, "devtoolsconnections.removeroombyid", request)
	var fallbackCheck func(error) bool
	if cli.fallbackDevToolsConnections != nil {
		fallbackCheck = cli.fallbackDevToolsConnections.RemoveRoomByID
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

func (cli *ClientDevToolsConnections) ReqRemoveRoomByID(ctx context.Context, callback retDevToolsConnectionsRemoveRoomByID, token string, gameID uuid.UUID, roomID uuid.UUID) (request RequestRPC) {

	request = RequestRPC{rpcRequest: &jsonrpc.RequestRPC{
		ID:      jsonrpc.NewID(),
		JSONRPC: jsonrpc.Version,
		Method:  "devtoolsconnections.removeroombyid",
		Params: requestDevToolsConnectionsRemoveRoomByID{
			GameID: gameID,
			RoomID: roomID,
			Token:  token,
		},
	}}
	if callback != nil {
		var response responseDevToolsConnectionsRemoveRoomByID
		request.retHandler = func(err error, rpcResponse *jsonrpc.ResponseRPC) {
			cacheKey, _ := hasher.Hash(request.rpcRequest.Params)
			var fallbackCheck func(error) bool
			if cli.fallbackDevToolsConnections != nil {
				fallbackCheck = cli.fallbackDevToolsConnections.RemoveRoomByID
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

func (cli *ClientDevToolsConnections) GetRoomsConnectionUrls(ctx context.Context, token string) (connectionsServer []types.ConnectionServer, err error) {

	request := requestDevToolsConnectionsGetRoomsConnectionUrls{Token: token}
	var response responseDevToolsConnectionsGetRoomsConnectionUrls
	var rpcResponse *jsonrpc.ResponseRPC
	cacheKey, _ := hasher.Hash(request)
	rpcResponse, err = cli.rpc.Call(ctx, "devtoolsconnections.getroomsconnectionurls", request)
	var fallbackCheck func(error) bool
	if cli.fallbackDevToolsConnections != nil {
		fallbackCheck = cli.fallbackDevToolsConnections.GetRoomsConnectionUrls
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
	return response.ConnectionsServer, err
}

func (cli *ClientDevToolsConnections) ReqGetRoomsConnectionUrls(ctx context.Context, callback retDevToolsConnectionsGetRoomsConnectionUrls, token string) (request RequestRPC) {

	request = RequestRPC{rpcRequest: &jsonrpc.RequestRPC{
		ID:      jsonrpc.NewID(),
		JSONRPC: jsonrpc.Version,
		Method:  "devtoolsconnections.getroomsconnectionurls",
		Params:  requestDevToolsConnectionsGetRoomsConnectionUrls{Token: token},
	}}
	if callback != nil {
		var response responseDevToolsConnectionsGetRoomsConnectionUrls
		request.retHandler = func(err error, rpcResponse *jsonrpc.ResponseRPC) {
			cacheKey, _ := hasher.Hash(request.rpcRequest.Params)
			var fallbackCheck func(error) bool
			if cli.fallbackDevToolsConnections != nil {
				fallbackCheck = cli.fallbackDevToolsConnections.GetRoomsConnectionUrls
			}
			if rpcResponse != nil && rpcResponse.Error != nil {
				if cli.errorDecoder != nil {
					err = cli.errorDecoder(rpcResponse.Error.Raw())
				} else {
					err = fmt.Errorf(rpcResponse.Error.Message)
				}
			}
			callback(response.ConnectionsServer, cli.proceedResponse(ctx, err, cacheKey, fallbackCheck, rpcResponse, &response))
		}
	}
	return
}
