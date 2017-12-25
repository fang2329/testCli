package msg

import (
	"time"

	"math/rand"

	mylog "../common/common_log"
	netmsg "../common/net"
	"github.com/golang/protobuf/proto"
)

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func LoginMsg(uid int, key string) []byte {
	deviceid := GetRandomString(10)
	k := []byte(key)
	d := []byte(deviceid)
	loginMsg := &netmsg.MsgLoginReq{
		Uid:      proto.Uint32(uint32(uid)),
		Key:      k,
		Deviceid: d,
	}

	data, err := proto.Marshal(loginMsg)
	if nil != err {
		mylog.LOG_ERROR("marshal error :%d", err)
	}
	return data
}

func RewardLoginMsg() []byte {
	rewardlogMsg := &netmsg.MsgGetLoginRewardReq{
		RewardFlag: proto.Uint32(0),
	}

	data, err := proto.Marshal(rewardlogMsg)
	if nil != err {
		mylog.LOG_ERROR("reward marshal error[%d]", err)
	}
	return data
}

func GetSrvInfo() []byte {
	return nil
}

func EnterSrv(srvID uint32) []byte {
	enterMsg := &netmsg.MsgEnterGamesvrReq{
		Svrid: proto.Uint32(srvID),
	}
	data, err := proto.Marshal(enterMsg)
	if nil != err {
		mylog.LOG_ERROR("neter gamesrv error[%d]", err)
	}
	return data
}

func RquestSrvRoomList(srvid, servType uint32) []byte {
	roomList := &netmsg.MsgRoomsInfoReq{
		GameType: proto.Uint32(servType),
		SvrId:    proto.Uint32(srvid),
	}
	data, err := proto.Marshal(roomList)
	if nil != err {
		mylog.LOG_ERROR("neter gamesrv error[%d]", err)
	}
	return data
}

func EnterRoom(roomId uint32) []byte {
	enterRoom := &netmsg.MsgEnterRoomReq{
		RoomId: proto.Uint32(roomId),
	}
	data, err := proto.Marshal(enterRoom)
	if nil != err {
		mylog.LOG_ERROR("enter room msg error", err.Error())
	}
	return data
}

func EnterTable() []byte {
	enterTable := &netmsg.MsgEnterTableReq{
		TableId: proto.Uint32(0),
	}
	data, err := proto.Marshal(enterTable)
	if nil != err {
		mylog.LOG_ERROR("enter table Marshal error", err.Error())
	}
	return data
}

func HeartBeatMsg() []byte {
	heartBeat := &netmsg.MsgHeartTest{
		SvrTime: proto.Uint32(uint32(time.Now().Unix())),
	}
	data, err := proto.Marshal(heartBeat)
	if nil != err {
		mylog.LOG_ERROR("heartBeatMsg Marshal error", err.Error())
	}
	return data
}

func TableReadyMsg() []byte {
	readyMsg := &netmsg.MsgTableReadyReq{
		Ready: proto.Uint32(1),
	}
	data, err := proto.Marshal(readyMsg)
	if nil != err {
		mylog.LOG_ERROR("ready msg Marshal error", err.Error())
	}
	return data
}

/////////////////////////////////////炸金花/////////////////////////////////////////////////

func JinhuaAddScore(addScore int64, is_allin uint32) []byte {
	addScoreMsg := &netmsg.MsgZajinhuaAddscoreReq{
		AddScore: proto.Int64(addScore),
		IsAllin:  proto.Uint32(is_allin),
	}
	data, err := proto.Marshal(addScoreMsg)
	if nil != err {
		mylog.LOG_ERROR("JjnhuaAddScore  error", err.Error())
	}
	return data
}

func JinhuaGiveup() []byte {
	giveupMsg := &netmsg.MsgZajinhuaGiveupReq{
		GiveupReason: proto.Uint32(1),
	}
	data, err := proto.Marshal(giveupMsg)
	if err != nil {
		mylog.LOG_DEBUG("JinhuaGiveup error", err.Error())
	}
	return data
}
