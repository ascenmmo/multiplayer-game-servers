package types

type RequestSetMessage struct {
	Data any `json:"data"`
}

type ResponseGetMessage struct {
	DataArray []any `json:"dataArray"`
}
