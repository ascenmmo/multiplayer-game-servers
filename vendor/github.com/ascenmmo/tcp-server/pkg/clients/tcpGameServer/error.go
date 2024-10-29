// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package tcpGameServer

import "encoding/json"

type errorJsonRPC struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (err errorJsonRPC) Error() string {
	return err.Message
}

type ErrorDecoder func(errData json.RawMessage) error

func defaultErrorDecoder(errData json.RawMessage) (err error) {

	var jsonrpcError errorJsonRPC
	if err = json.Unmarshal(errData, &jsonrpcError); err != nil {
		return
	}
	return jsonrpcError
}
