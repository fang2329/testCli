package msg

import (
	"fmt"
	"reflect"

	mylog "../common/common_log"
	"../common/myconobj"
	netmsg "../common/net"
	"github.com/golang/protobuf/proto"
)

func OnReceivedMsg(pPlayer *myconobj.PlayerObj, cmd netmsg.CLIENT_MSG_ID, msg []byte) {
	switch cmd {
	case netmsg.CLIENT_MSG_ID_C2S_MSG_HEART:
		{
			handleHearbeat(pPlayer, msg)
		}
	case netmsg.CLIENT_MSG_ID_S2C_MSG_LOGIN_REP:
		{
			handleMsgLoginBack(pPlayer, msg)
		}
	case netmsg.CLIENT_MSG_ID_S2C_MSG_PLAYER_INFO:
		{
			handleMsgPlayerBase(pPlayer, msg)
		}
	case netmsg.CLIENT_MSG_ID_S2C_MSG_UPDATA_GAME_INFO:
		{
			handleMsgPlayerUpdateGameInfo(pPlayer, msg)
		}
	case netmsg.CLIENT_MSG_ID_S2C_MSG_SEND_ALL_MISSION_REP:
		{
			handleMsgMission(pPlayer, msg)
		}
	case netmsg.CLIENT_MSG_ID_S2C_MSG_ENTER_GAME:
		{
			handleEnterGame(pPlayer, msg)
		}
	case netmsg.CLIENT_MSG_ID_S2C_MSG_GET_LOGIN_REWARD_REP:
		{
			handleMsgRewardLogMsg(pPlayer, msg)
		}

	case netmsg.CLIENT_MSG_ID_S2C_MSG_SVRS_INFO:
		{
			handleMsgSrvInfo(pPlayer, msg)
		}
	case netmsg.CLIENT_MSG_ID_S2C_MSG_ENTER_SVR_REP:
		{
			handleMsgEnterSrv(pPlayer, msg)
		}
	case netmsg.CLIENT_MSG_ID_S2C_MSG_ROOMS_INFO:
		{
			handleMsgRoomList(pPlayer, msg)
		}
	case netmsg.CLIENT_MSG_ID_S2C_MSG_ENTER_ROOM_REP:
		{
			handleMsgEnterRoom(pPlayer, msg)
		}
	case netmsg.CLIENT_MSG_ID_S2C_MSG_TABLE_INFO:
		{
			handleMsgTableInfo(pPlayer, msg)
		}
	case netmsg.CLIENT_MSG_ID_S2C_MSG_ENTER_TABLE:
		{
			handleMsgEnterTable(pPlayer, msg)
		}
	case netmsg.CLIENT_MSG_ID_S2C_MSG_SEATS_INFO:
		{
			handleMsgSeatInfo(pPlayer, msg)
		}
	case netmsg.CLIENT_MSG_ID_S2C_MSG_TABLE_READY_REP:
		{
			handleMsgTableReady(pPlayer, msg)
		}
		/////////////////////////////////////炸金花//////////////////////////////////////////////////////
	case netmsg.CLIENT_MSG_ID_S2C_MSG_ZAJINHUA_GAME_PLAY_INFO:
		{
			handleZhajinhuaStart(pPlayer, msg)
		}
	case netmsg.CLIENT_MSG_ID_S2C_MSG_ZAJINHUA_ADD_SCORE:
		{
			handleZhajinhuaAddScore(pPlayer, msg)
		}
	case netmsg.CLIENT_MSG_ID_S2C_MSG_ZAJINHUA_GIVEUP:
		{
			handleZhajinhuaGiveUp(pPlayer, msg)
		}
	}
}

func handleHearbeat(pPlayer *myconobj.PlayerObj, msg []byte) {
	retmsg := &netmsg.MsgHeartTest{}
	err := proto.Unmarshal(msg, retmsg)
	if nil != err {
		mylog.LOG_ERROR("heart beat err Unmarshal ", err.Error())
	}
	mylog.LOG_INFO("heatBeat result %d", retmsg.GetSvrTime())
}

func handleMsgLoginBack(pPlayer *myconobj.PlayerObj, msg []byte) {
	retmsg := &netmsg.MsgLoginRep{}
	err := proto.Unmarshal(msg, retmsg)
	if nil != err {
		mylog.LOG_ERROR("unmarshling error:", err)
	}
	mylog.LOG_INFO("result [%d] time [%d] uid[%d]", retmsg.GetResult(), retmsg.GetServerTime(), pPlayer.GetPlayerUid())
}

func handleMsgPlayerBase(pPlayer *myconobj.PlayerObj, msg []byte) {
	retmsg := &netmsg.MsgPlayerDataRep{}
	err := proto.Unmarshal(msg, retmsg)
	if nil != err {
		mylog.LOG_ERROR("msg unmarsh error:", err.Error())
	}

	baseInfo := retmsg.GetBaseData()
	mylog.LOG_INFO("userbaseInfo:uid[%d],name[%s],sex[%d]", baseInfo.GetUid(), baseInfo.GetName(), baseInfo.GetSex())
}

func handleMsgPlayerUpdateGameInfo(pPlayer *myconobj.PlayerObj, msg []byte) {
	retmsg := &netmsg.MsgUpdateGameInfo{}
	err := proto.Unmarshal(msg, retmsg)
	if nil != err {
		mylog.LOG_ERROR("msg unmarsh error", err.Error())
	}

	updateInfo := retmsg.GetData()
	mylog.LOG_INFO("gamedate :uid[%d],gameType[%d],win[%d],lose[%d]", pPlayer.GetPlayerUid(), updateInfo.GetGameType(), updateInfo.GetWin(), updateInfo.GetLose())
}

func handleMsgMission(pPlayer *myconobj.PlayerObj, msg []byte) {
	retmsg := &netmsg.MsgSendAllMissionRep{}
	err := proto.Unmarshal(msg, retmsg)
	if nil != err {
		mylog.LOG_ERROR("msg unmarsh error", err.Error())
	}

	mission := retmsg.GetMissions()
	for i := 0; i < len(mission); i++ {
		missData := mission[i]
		fmt.Println(reflect.TypeOf(missData))
	}

}

func handleEnterGame(pPlayer *myconobj.PlayerObj, msg []byte) {
	retmsg := &netmsg.MsgEnterGameRep{}
	err := proto.Unmarshal(msg, retmsg)
	if nil != err {
		mylog.LOG_ERROR("msg unmarsh error", err.Error())
	}
	mylog.LOG_INFO("result [%d],uid[%d]", retmsg.GetResult(), pPlayer.GetPlayerUid())
	srvInfoMsg := GetSrvInfo()
	pPlayer.SendMsgToSvr(netmsg.CLIENT_MSG_ID_C2S_MSG_REQ_SVRS_INFO, srvInfoMsg)
}
func handleMsgRewardLogMsg(pPlayer *myconobj.PlayerObj, msg []byte) {
	retmsg := &netmsg.MsgGetLoginRewardRep{}
	err := proto.Unmarshal(msg, retmsg)
	if nil != err {
		mylog.LOG_ERROR("unmarshling error", err)
	}

	flag := retmsg.GetRewardFlag()
	result := retmsg.GetResult()
	mylog.LOG_INFO("flag is [%d],result is [%d] uid[%d]", flag, result, pPlayer.GetPlayerUid())
}

func handleMsgSrvInfo(pPlayer *myconobj.PlayerObj, msg []byte) {
	retmsg := &netmsg.MsgSvrsInfoRep{}
	err := proto.Unmarshal(msg, retmsg)
	if nil != err {
		mylog.LOG_ERROR("unmarshling error", err)
	}
	srvCount := retmsg.GetSvrs()

	for i := 0; i < len(srvCount); i++ {
		srv := srvCount[i]
		mylog.LOG_INFO("srvinfo : [SrvID:%d,State:%d,GameType:%d,SubType:%d]", srv.GetSvrid(), srv.GetState(), srv.GetGameType(), srv.GetGameSubtype())
	}

	entersrvMsg := EnterSrv(51)
	pPlayer.SendMsgToSvr(netmsg.CLIENT_MSG_ID_C2S_MSG_ENTER_SVR, entersrvMsg)
}

func handleMsgEnterSrv(pPlayer *myconobj.PlayerObj, msg []byte) {
	retmsg := &netmsg.MsgEnterGamesvrRep{}
	err := proto.Unmarshal(msg, retmsg)
	if nil != err {
		mylog.LOG_ERROR("unmarshal error", err)
	}
	mylog.LOG_DEBUG("result [%d],Srvid [%d],uid[%d]", retmsg.GetResult(), retmsg.GetSvrid(), pPlayer.GetPlayerUid())
	roomlist := RquestSrvRoomList(1, retmsg.GetSvrid())
	pPlayer.SendMsgToSvr(netmsg.CLIENT_MSG_ID_C2S_MSG_REQ_ROOMS_INFO, roomlist)
}

func handleMsgRoomList(pPlayer *myconobj.PlayerObj, msg []byte) {
	retmsg := &netmsg.MsgRoomsInfoRep{}
	proto.Unmarshal(msg, retmsg)
	rooms := retmsg.GetRooms()
	for i := 0; i < len(rooms); i++ {
		roomInfo := rooms[i]
		mylog.LOG_INFO("roomid[%d],srvid[%d]-->roomInfo id[%d],consume[%d],deal[%d],EnterMin[%d]", retmsg.GetCurRoomid(), retmsg.GetSvrId(), roomInfo.GetId(), roomInfo.GetConsume(),
			roomInfo.GetDeal(), roomInfo.GetEnterMin())
	}
	enterRoom := EnterRoom(1)
	pPlayer.SendMsgToSvr(netmsg.CLIENT_MSG_ID_C2S_MSG_ENTER_ROOM, enterRoom)
}

func handleMsgEnterRoom(pPlayer *myconobj.PlayerObj, msg []byte) {
	retmsg := &netmsg.MsgEnterRoomRep{}
	proto.Unmarshal(msg, retmsg)
	roomInfo := retmsg.GetRoom()
	mylog.LOG_INFO("result[%d],roominfo-->id[%d],consume[%d],deal[%d],table[%d]", retmsg.GetResult(), roomInfo.GetId(), roomInfo.GetConsume(), roomInfo.GetDeal(), retmsg.GetCurTable())
	enterTable := EnterTable()
	pPlayer.SendMsgToSvr(netmsg.CLIENT_MSG_ID_C2S_MSG_ENTER_TABLE_REQ, enterTable)
}

func handleMsgEnterTable(pPlayer *myconobj.PlayerObj, msg []byte) {
	retmsg := &netmsg.MsgEnterTableRep{}
	proto.Unmarshal(msg, retmsg)
	mylog.LOG_INFO("enterTable tableID[%d],Result[%d]", retmsg.GetTableId(), retmsg.GetResult())
}

func handleMsgTableInfo(pPlayer *myconobj.PlayerObj, msg []byte) {
	retmsg := &netmsg.MsgTableInfoRep{}
	proto.Unmarshal(msg, retmsg)
	tableInfo := retmsg.GetTableInfo()
	zhajinhua := tableInfo.GetZajinhua()
	mylog.LOG_INFO("jinhuatable tid[%d],uid[%d]", zhajinhua.GetTableid(), pPlayer.GetPlayerUid())
	readyMsg := TableReadyMsg()
	pPlayer.SendMsgToSvr(netmsg.CLIENT_MSG_ID_C2S_MSG_TABLE_READY, readyMsg)
}

func handleMsgSeatInfo(pPlayer *myconobj.PlayerObj, msg []byte) {
	retmsg := &netmsg.MsgSeatInfoRep{}
	proto.Unmarshal(msg, retmsg)
	seatInfo := retmsg.GetPlayers()
	for i := 0; i < len(seatInfo); i++ {
		mylog.LOG_INFO("seatinfo uid[%d],name[%s],sex[%d]", seatInfo[i].GetUid(), string(seatInfo[i].GetName()), seatInfo[i].GetSex())
	}

}

func handleMsgTableReady(pPlayer *myconobj.PlayerObj, msg []byte) {
	retmsg := &netmsg.MsgTableReadyRep{}
	proto.Unmarshal(msg, retmsg)
	mylog.LOG_INFO("uid[%d],readys[%d],auto_state[%d]", pPlayer.GetPlayerUid(), retmsg.GetReadys(), retmsg.GetAutoStates())
}

//////////////////////////////////////////////////////////////////jinhua//////////////////////////////////////////////////////////////////
func handleZhajinhuaStart(pPlayer *myconobj.PlayerObj, msg []byte) {
	mylog.LOG_DEBUG("......zhajinhua start......")
	retmsg := &netmsg.MsgZajinhuaStartRep{}
	proto.Unmarshal(msg, retmsg)
	mylog.LOG_INFO("banker[%d],curUser[%d],maxScore[%d],provateRound[%d]", retmsg.GetBankerUser(), retmsg.GetCurrentUser(), retmsg.GetUserMaxScore(), retmsg.GetCurPrivateRound())

	addScore := JinhuaAddScore(1, 0)
	pPlayer.SendMsgToSvr(netmsg.CLIENT_MSG_ID_C2S_MSG_ZAJINHUA_ADD_SCORE, addScore)
}

func handleZhajinhuaAddScore(pPlayer *myconobj.PlayerObj, msg []byte) {
	mylog.LOG_DEBUG(".......zhajinhua addscore........")
	retmsg := &netmsg.MsgZajinhuaAddscoreRep{}
	proto.Unmarshal(msg, retmsg)
	mylog.LOG_INFO("curUser[%d],addUser[%d],compareState[%d],addScoreCount[%d],curTimes[%d]", retmsg.GetCurrentUser(), retmsg.GetAddScoreUser(), retmsg.GetCompareState(),
		retmsg.GetAddScoreCount(), retmsg.GetCurrentTimes())

	giveUp := JinhuaGiveup()
	pPlayer.SendMsgToSvr(netmsg.CLIENT_MSG_ID_C2S_MSG_ZAJINHUA_GIVEUP, giveUp)
}

func handleZhajinhuaGiveUp(pPlayer *myconobj.PlayerObj, msg []byte) {
	mylog.LOG_DEBUG("..........zhajinhua giveup.........")
	retmsg := &netmsg.MsgZajinhuaGiveupRep{}
	proto.Unmarshal(msg, retmsg)
	mylog.LOG_INFO("user[%d]", retmsg.GetGiveupUser())
}
