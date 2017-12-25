package myconobj

import (
	netmsg "../net"
	"net"
)

type PlayerObj struct {
	uid        int
	SessionKey string
	conn       net.Conn
}

func NewPlayerObj() *PlayerObj {
	tmpPlayer := &PlayerObj{
		uid:        0,
		SessionKey: "",
	}
	return tmpPlayer
}

func (p *PlayerObj) SetPlayerInfo(u int, skey string, con net.Conn) {
	p.uid = u
	p.SessionKey = skey
	p.conn = con
}

func (p *PlayerObj) GetPlayerUid() int {
	return p.uid
}

func (p *PlayerObj) GetPlayerSession() string {
	return p.SessionKey
}

func (p *PlayerObj) GetPlayerConn() net.Conn {
	return p.conn
}

func (p *PlayerObj) SendMsgToSvr(cmd netmsg.CLIENT_MSG_ID, data []byte) {
	var finalNeedSendMsg []byte
	if data != nil {
		finalNeedSendMsg = netmsg.FormMsg(cmd, data, len(data))
	} else {
		finalNeedSendMsg = netmsg.FormMsg(cmd, data, 0)
	}
	p.conn.Write(finalNeedSendMsg)
}
