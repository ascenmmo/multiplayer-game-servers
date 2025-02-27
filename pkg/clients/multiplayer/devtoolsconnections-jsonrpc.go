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
type retDevToolsConnectionsJoinRoomByRoomCode = func(newToken string, err error)
type retDevToolsConnectionsGetMyRoom = func(room types.Room, err error)
type retDevToolsConnectionsLeaveRoom = func(err error)
type retDevToolsConnectionsRemoveRoomByID = func(err error)
type retDevToolsConnectionsGetRoomsConnectionUrls = func(connectionsServer []types.ConnectionServer, err error)

func (cli *ClientDevToolsConnections) CreateRoom(ctx context.Context, token string, name string) (newToken string, err error) {

	request := requestDevToolsConnectionsCreateRoom{
		Name:  name,
		Token: token,
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

func (cli *ClientDevToolsConnections) ReqCreateRoom(ctx context.Context, callback retDevToolsConnectionsCreateRoom, token string, name string) (request RequestRPC) {

	request = RequestRPC{rpcRequest: &jsonrpc.RequestRPC{
		ID:      jsonrpc.NewID(),
		JSONRPC: jsonrpc.Version,
		Method:  "devtoolsconnections.createroom",
		Params: requestDevToolsConnectionsCreateRoom{
			Name:  name,
			Token: token,
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

func (cli *ClientDevToolsConnections) GetRoomsAll(ctx context.Context, token string) (rooms []types.Room, err error) {

	request := requestDevToolsConnectionsGetRoomsAll{Token: token}
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

func (cli *ClientDevToolsConnections) ReqGetRoomsAll(ctx context.Context, callback retDevToolsConnectionsGetRoomsAll, token string) (request RequestRPC) {

	request = RequestRPC{rpcRequest: &jsonrpc.RequestRPC{
		ID:      jsonrpc.NewID(),
		JSONRPC: jsonrpc.Version,
		Method:  "devtoolsconnections.getroomsall",
		Params:  requestDevToolsConnectionsGetRoomsAll{Token: token},
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

func (cli *ClientDevToolsConnections) JoinRoomByID(ctx context.Context, token string, roomID uuid.UUID) (newToken string, err error) {

	request := requestDevToolsConnectionsJoinRoomByID{
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

func (cli *ClientDevToolsConnections) ReqJoinRoomByID(ctx context.Context, callback retDevToolsConnectionsJoinRoomByID, token string, roomID uuid.UUID) (request RequestRPC) {

	request = RequestRPC{rpcRequest: &jsonrpc.RequestRPC{
		ID:      jsonrpc.NewID(),
		JSONRPC: jsonrpc.Version,
		Method:  "devtoolsconnections.joinroombyid",
		Params: requestDevToolsConnectionsJoinRoomByID{
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

func (cli *ClientDevToolsConnections) JoinRoomByRoomCode(ctx context.Context, token string, roomCode string) (newToken string, err error) {

	request := requestDevToolsConnectionsJoinRoomByRoomCode{
		RoomCode: roomCode,
		Token:    token,
	}
	var response responseDevToolsConnectionsJoinRoomByRoomCode
	var rpcResponse *jsonrpc.ResponseRPC
	cacheKey, _ := hasher.Hash(request)
	rpcResponse, err = cli.rpc.Call(ctx, "devtoolsconnections.joinroombyroomcode", request)
	var fallbackCheck func(error) bool
	if cli.fallbackDevToolsConnections != nil {
		fallbackCheck = cli.fallbackDevToolsConnections.JoinRoomByRoomCode
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

func (cli *ClientDevToolsConnections) ReqJoinRoomByRoomCode(ctx context.Context, callback retDevToolsConnectionsJoinRoomByRoomCode, token string, roomCode string) (request RequestRPC) {

	request = RequestRPC{rpcRequest: &jsonrpc.RequestRPC{
		ID:      jsonrpc.NewID(),
		JSONRPC: jsonrpc.Version,
		Method:  "devtoolsconnections.joinroombyroomcode",
		Params: requestDevToolsConnectionsJoinRoomByRoomCode{
			RoomCode: roomCode,
			Token:    token,
		},
	}}
	if callback != nil {
		var response responseDevToolsConnectionsJoinRoomByRoomCode
		request.retHandler = func(err error, rpcResponse *jsonrpc.ResponseRPC) {
			cacheKey, _ := hasher.Hash(request.rpcRequest.Params)
			var fallbackCheck func(error) bool
			if cli.fallbackDevToolsConnections != nil {
				fallbackCheck = cli.fallbackDevToolsConnections.JoinRoomByRoomCode
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

func (cli *ClientDevToolsConnections) GetMyRoom(ctx context.Context, token string) (room types.Room, err error) {

	request := requestDevToolsConnectionsGetMyRoom{Token: token}
	var response responseDevToolsConnectionsGetMyRoom
	var rpcResponse *jsonrpc.ResponseRPC
	cacheKey, _ := hasher.Hash(request)
	rpcResponse, err = cli.rpc.Call(ctx, "devtoolsconnections.getmyroom", request)
	var fallbackCheck func(error) bool
	if cli.fallbackDevToolsConnections != nil {
		fallbackCheck = cli.fallbackDevToolsConnections.GetMyRoom
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
	return response.Room, err
}

func (cli *ClientDevToolsConnections) ReqGetMyRoom(ctx context.Context, callback retDevToolsConnectionsGetMyRoom, token string) (request RequestRPC) {

	request = RequestRPC{rpcRequest: &jsonrpc.RequestRPC{
		ID:      jsonrpc.NewID(),
		JSONRPC: jsonrpc.Version,
		Method:  "devtoolsconnections.getmyroom",
		Params:  requestDevToolsConnectionsGetMyRoom{Token: token},
	}}
	if callback != nil {
		var response responseDevToolsConnectionsGetMyRoom
		request.retHandler = func(err error, rpcResponse *jsonrpc.ResponseRPC) {
			cacheKey, _ := hasher.Hash(request.rpcRequest.Params)
			var fallbackCheck func(error) bool
			if cli.fallbackDevToolsConnections != nil {
				fallbackCheck = cli.fallbackDevToolsConnections.GetMyRoom
			}
			if rpcResponse != nil && rpcResponse.Error != nil {
				if cli.errorDecoder != nil {
					err = cli.errorDecoder(rpcResponse.Error.Raw())
				} else {
					err = fmt.Errorf(rpcResponse.Error.Message)
				}
			}
			callback(response.Room, cli.proceedResponse(ctx, err, cacheKey, fallbackCheck, rpcResponse, &response))
		}
	}
	return
}

func (cli *ClientDevToolsConnections) LeaveRoom(ctx context.Context, token string, roomID uuid.UUID) (err error) {

	request := requestDevToolsConnectionsLeaveRoom{
		RoomID: roomID,
		Token:  token,
	}
	var response responseDevToolsConnectionsLeaveRoom
	var rpcResponse *jsonrpc.ResponseRPC
	cacheKey, _ := hasher.Hash(request)
	rpcResponse, err = cli.rpc.Call(ctx, "devtoolsconnections.leaveroom", request)
	var fallbackCheck func(error) bool
	if cli.fallbackDevToolsConnections != nil {
		fallbackCheck = cli.fallbackDevToolsConnections.LeaveRoom
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

func (cli *ClientDevToolsConnections) ReqLeaveRoom(ctx context.Context, callback retDevToolsConnectionsLeaveRoom, token string, roomID uuid.UUID) (request RequestRPC) {

	request = RequestRPC{rpcRequest: &jsonrpc.RequestRPC{
		ID:      jsonrpc.NewID(),
		JSONRPC: jsonrpc.Version,
		Method:  "devtoolsconnections.leaveroom",
		Params: requestDevToolsConnectionsLeaveRoom{
			RoomID: roomID,
			Token:  token,
		},
	}}
	if callback != nil {
		var response responseDevToolsConnectionsLeaveRoom
		request.retHandler = func(err error, rpcResponse *jsonrpc.ResponseRPC) {
			cacheKey, _ := hasher.Hash(request.rpcRequest.Params)
			var fallbackCheck func(error) bool
			if cli.fallbackDevToolsConnections != nil {
				fallbackCheck = cli.fallbackDevToolsConnections.LeaveRoom
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

func (cli *ClientDevToolsConnections) RemoveRoomByID(ctx context.Context, token string, roomID uuid.UUID) (err error) {

	request := requestDevToolsConnectionsRemoveRoomByID{
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

func (cli *ClientDevToolsConnections) ReqRemoveRoomByID(ctx context.Context, callback retDevToolsConnectionsRemoveRoomByID, token string, roomID uuid.UUID) (request RequestRPC) {

	request = RequestRPC{rpcRequest: &jsonrpc.RequestRPC{
		ID:      jsonrpc.NewID(),
		JSONRPC: jsonrpc.Version,
		Method:  "devtoolsconnections.removeroombyid",
		Params: requestDevToolsConnectionsRemoveRoomByID{
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
