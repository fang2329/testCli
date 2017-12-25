package net

import (
	"bytes"
	"encoding/binary"
	"unsafe"

	mylog "../common_log"
)

type PacketHeader struct {
	Uin     uint32
	Cmd     uint16
	Datalen uint16
}

func newMsgHeader(u uint32, c uint16, mLen uint16) *PacketHeader {
	return &PacketHeader{
		Uin:     u,
		Cmd:     c,
		Datalen: mLen,
	}
}

func GetMsgHeaderLen() uintptr {
	tmpHeader := newMsgHeader(0, 0, 0)
	return unsafe.Sizeof(tmpHeader)
}

type PacketMsg struct {
	PacketHeader
	protobuf []byte
}

func (pMsg *PacketHeader) GetCmd() uint16 {
	return pMsg.Cmd
}

func (pMsg *PacketHeader) GetDataLen() uint16 {
	return pMsg.Datalen
}
func (p *PacketHeader) Encode() []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, p)
	if nil != err {
		mylog.LOG_ERROR("binary.Write failed", err)
	}
	return buf.Bytes()
}

func (p *PacketHeader) Decode(b []byte) {
	buf := bytes.NewBuffer(b)
	err := binary.Read(buf, binary.LittleEndian, p)
	if nil != err {
		mylog.LOG_ERROR("binary,Read failed", err)
	}
}

func FormMsg(cmd CLIENT_MSG_ID, data []byte, datalen int) []byte {
	tmpHeader := newMsgHeader(0, uint16(cmd), uint16(datalen))
	header := tmpHeader.Encode()

	headerlen := len(header)

	if headerlen == 0 && datalen == 0 {
		mylog.LOG_ERROR("do not need form new msg")
	}

	newdata := make([]byte, (headerlen)+datalen)
	copy(newdata, header)
	copy(newdata[headerlen:], data[0:])
	return newdata
}
