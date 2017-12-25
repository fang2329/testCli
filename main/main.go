package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"sync"
	"time"

	mylog "../common/common_log"
	"../common/myconobj"
	"../common/mydb"
	netmsg "../common/net"
	"../msg"
)

var (
	ip      = mydb.MapSrvCfg[1].GetSvrip()
	port    = mydb.MapSrvCfg[1].GetSvrPort()
	server  = fmt.Sprintf("%s:%d", ip, port)
	ReadMsg = make(chan *myconobj.PlayerObj, 50)
)

func init() {
	if *mylog.Mydaemon {
		args := os.Args[1:]
		i := 0
		for ; i < len(args); i++ {
			if args[i] == "-d=true" {
				args[i] = "-d=false"
				break
			}
		}
		cmd := exec.Command(os.Args[0], args...)
		cmd.Start()
		os.Exit(0)
	}
}

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp", server)
	if nil != err {
		mylog.LOG_ERROR("tcpAddr server[%s] resolve error:[%s]", server, err.Error())
		os.Exit(1)
	}

	var wg sync.WaitGroup
	wg.Add(*(mylog.Maxnum))

	for i := (*mylog.Begin); i < (*mylog.End); i++ {
		pPlayer := loginMsg(i, tcpAddr)
		go HearBeating(pPlayer, &wg)
		go handleMsg(pPlayer, &wg)
	}
	wg.Wait()
}

func loginMsg(num int, tcpAddr *net.TCPAddr) *myconobj.PlayerObj {
	var pPlayer = myconobj.NewPlayerObj()
	uid, key := msg.GetUserLoginInfo(num)
	if uid == 0 {
		mylog.LOG_ERROR("uid is not correct %d", uid)
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if nil != err {
		mylog.LOG_ERROR("connect to tcpAddr error :[%s]", err.Error())
		os.Exit(1)
	}
	pPlayer.SetPlayerInfo(uid, key, conn)
	data := msg.LoginMsg(uid, key)
	pPlayer.SendMsgToSvr(netmsg.CLIENT_MSG_ID_C2S_MSG_LOGIN, data)
	return pPlayer
}

func handleMsg(pPlayer *myconobj.PlayerObj, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		/*读头*/
		var headerlen uint16 = uint16(netmsg.GetMsgHeaderLen())
		recvdataHead := make([]byte, headerlen)
		_, err := pPlayer.GetPlayerConn().Read(recvdataHead)
		if err != nil {
			continue
		}
		var tmpHead netmsg.PacketHeader
		tmpHead.Decode(recvdataHead)
		datalen := tmpHead.GetDataLen()
		cmd := netmsg.CLIENT_MSG_ID(tmpHead.GetCmd())

		/*读数据*/
		recvdata := make([]byte, datalen)
		pPlayer.GetPlayerConn().Read(recvdata)
		msg.OnReceivedMsg(pPlayer, cmd, recvdata)
	}
	defer pPlayer.GetPlayerConn().Close()
}

func HearBeating(pPlayer *myconobj.PlayerObj, wg *sync.WaitGroup) {
	defer wg.Done()
	t := time.NewTimer(5 * time.Second)
	for {
		select {
		case <-t.C:
			{
				heartMsg := msg.HeartBeatMsg()
				pPlayer.SendMsgToSvr(netmsg.CLIENT_MSG_ID_C2S_MSG_HEART, heartMsg)
				t.Reset(5 * time.Second)
			}
		}
	}
}
