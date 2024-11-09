package connection

type DataSender interface {
	Write(msgType int, msg []byte) error
	GetID() string
	Close()
}
