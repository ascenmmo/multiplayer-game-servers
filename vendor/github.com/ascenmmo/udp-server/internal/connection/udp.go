package connection

import "net"

type UDPConnection struct {
	ClientAddr *net.UDPAddr
	Conn       *net.UDPConn
}

func (u *UDPConnection) GetID() string {
	add := u.ClientAddr.String()
	return add
}

func (u *UDPConnection) Write(msg []byte) error {
	_, err := u.Conn.WriteToUDP(msg, u.ClientAddr)
	if err != nil {
		return err
	}
	return nil
}
