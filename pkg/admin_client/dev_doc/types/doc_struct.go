package types

type DocCategory struct {
	CategoryTitle string `json:"title"`
	DocStruct     []DocStruct
}

type DocStruct struct {
	Title     string          `json:"title"`
	Info      string          `json:"info"`
	DockLists []DocStructList `json:"dock_lists"`
}
type DocStructList struct {
	Title            string `json:"title"`
	Description      string `json:"description"`
	RequestPath      string `json:"requestPath"`
	Method           string `json:"method"`
	RequestHeader    string `json:"requestHeader"`
	RequestBody      string `json:"requestBody"`
	RequestBodyInfo  string `json:"RequestBodyInfo"`
	ResponseBody     string `json:"ResponseBody"`
	ResponseBodyInfo string `json:"ResponseBodyInfo"`
	ErrorBody        string `bson:"errorBody"`
	DocErrorList     []DocErrorList
}

type DocErrorList struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Body        string `json:"body"`
}
