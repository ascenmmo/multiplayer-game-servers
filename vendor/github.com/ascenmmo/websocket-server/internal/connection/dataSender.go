package connection

type DataSender interface {
	Write([]byte) error
	GetID() string
}
