package engudp

import "github.com/ascenmmo/multiplayer-game-servers/pkg/admin_client/dev_doc/types"

func GetUDPConnectionsStruct() []types.DocStruct {
	udp := types.DocStruct{
		Title: "UDP Connection",
		Info: "This service is designed for transmitting small data in real time. " +
			"If your clients need to receive data as quickly as possible, and speed is more important than guaranteed delivery, this service is perfect. " +
			"UDP is a very fast solution, but it comes with drawbacks: some packets may be lost or delivered out of order.",
		DockLists: []types.DocStructList{
			{
				Title: "Connection Phase",
				Description: "To connect to the service, you need to know the IP address and port. " +
					"The maximum number of messages per second is 200. " +
					"You also need to wait for the server to respond with your client ID. The UserID in the response confirms that the connection is authenticated. " +
					"Example connection command:",
				RequestPath:      "udp://ascenmmo.com:4500",
				Method:           "udp write",
				RequestBody:      `eyJhbGciOiJIUz11NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzA3NjA4ODYsIkluZm8iOnsiZ2FtZV9pZCI61jJhZDIyNDNhLThhN2UtMzkzMC1iNjEzLTg5YzY2NDk2YWFjZCIsInJvb21faWQiOiIwMDAwMDAwMC0wMDAwLTAwMDAtMDAwMC0wMDAwMDAwMDAwMDAiLCJ1c2VyX2lkIjoiN2RkZmNiNzEtOWQ3OC0zMTczLWExNGQtNDVjZTc1ODIyNmM3IiwidHRsIjo5MDAwMDAwMDAwMDB9fQ.wmq1VH88zlws_tSAdvJjfXcdcoaxL8vT8W9gsRZrw_o$_$j`,
				RequestBodyInfo:  "You need to send a string without any wrapper several times until the server returns the userID. The server must recognize that your connection is valid; otherwise, you are not considered a connected user.",
				ResponseBody:     `3fa85f64-5717-4562-b3fc-2c963f66afa6`,
				ResponseBodyInfo: "Proceed to the Reading Phase. Before moving on to the Sending Phase, you must receive a userID, which confirms that your connection is authenticated.",
			},
			{
				Title: "Reading Phase",
				Description: "Use the same connection channel; do not create a new one. " +
					"The connection channel for reading and sending must be the same. " +
					"Reading command:",
				RequestPath:     "udp://ascenmmo.com:4500",
				Method:          "udp read",
				RequestBody:     `Belongs to the sending phase, during parallel reading no data needs to be sent.`,
				RequestBodyInfo: "Here we need to listen for the server's response.",
				ResponseBody:    `data1 data2 data2   or   dat data2 d$#3 dt9   or  data`,
				ResponseBodyInfo: "The response may contain one or multiple messages. UDP does not guarantee message integrity or order. " +
					"The server operates on a broadcast model, where a message from ClientOne is duplicated to ClientTwo and ClientThree.",
			},
			{
				Title: "Sending Phase",
				Description: "Use the same connection channel; do not create a new one. " +
					"The connection channel for reading and sending must be the same. " +
					"Since you authenticated on the server in the first phase, you can now send packets. " +
					"Example sending command:",
				RequestPath:     "udp://ascenmmo.com:4500",
				Method:          "udp write",
				RequestBody:     `Client 1 sent message 101, or sent JSON {"name": "MyUser1"}`,
				RequestBodyInfo: "The server only works with []byte data. The byte array is distributed to all clients.",
				ResponseBody:    `Client 2 sent message 15, or sent JSON {"name": "MyUser2"} Client 3 sent message 53, or sent JSON {"name": "MyUser3"}`,
				ResponseBodyInfo: "The response may contain one or multiple messages. UDP does not guarantee message integrity or order. " +
					"If the messages are unreadable, try adjusting the sender's timing. " +
					"The server does not validate or split messages; it broadcasts them as they are.",
			},
		},
	}

	return []types.DocStruct{udp}
}
