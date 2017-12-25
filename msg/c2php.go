package msg

import (
	mylog "../common/common_log"
	//"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	url = "http://192.168.2.241/kd2/api/flashapi.php?"
)

type UserBaseInfo struct {
	Uid         int    `json:"uid"`
	Nickname    string `json:"nickname"`
	Version     string `json:"version"`
	Sex         int    `json:"sex"`
	Rtime       int    `json:"rtime"`
	Logintime   int    `json:"logintime"`
	Offinetime  int    `json:"offinetime"`
	Clogin      int    `json:"clogin"`
	Weeklogin   int    `json:"weeklogin"`
	Alllogin    int    `json:"alllogin"`
	Onlinetime  int    `json:"onlinetime"`
	Reward      int    `json:"reward"`
	Bankrupt    int    `json:"bankrupt"`
	Dgcount     int    `json:"dgcount"`
	Setting     string `json:"setting"`
	Safepriblem string `json:"safepriblem"`
	Safeanswer  string `json:"safeanswer"`
	Ispay       int    `json:"ispay"`
	Payall      int    `json:"payall"`
	Paytime     int    `json:"paytime"`
	Fpaytime    int    `json:"fpaytime"`
	Lpaytime    int    `json:"lpaytime"`
	Ip          string `json:"ip"`
	Loginip     string `json:"loginip"`
	Mac         string `json:"mac"`
	Signtype    []byte `json:"signtype"`
	Imagetype   int    `json:"imagetype"`
	Imageurl    string `json:"imageurl"`
	Deviceid    string `json:"deviceid"`
	Ldeviceid   string `json:"ldeviceid"`
	Parentuid   int    `json:"parentuid"`
	Oldname     string `json:"oldname"`
	Uptime      int    `json:"uptime"`
	Forbid      int    `json:"forbid"`
	Rewardmsg   string `json:"rewardmsg"`
	Regreward   int    `json:"regreward"`
	Diamond     int    `json:diamond`
	Coin        int    `json:"coin"`
	Ingot       int    `json:"ingot"`
	Score       int    `josn :"score"`
	Cvalue      int    `json:"cvalue"`
	Vip         int    `json:"vip"`
	Safecoin    int    `json:"safecoin"`
	Feewin      int    `json:"feewin"`
	Feelose     int    `json:"feelose"`
	Sysreward   int    `json:"sysreward"`
	Recharge    int    `json:"Recharge"`
	Convert     int    `json:"convert"`
	Win         int    `json:"win"`
}

type PassportInfo struct {
	Uid     int    `json:"uid"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Ano     string `json:"ano"`
	Rtype   int    `json:"rtype"`
	Sitemid string `json:"sitemid"`
	Pwd     string `json:"pwd"`
	Siteid  int    `json:"siteid"`
	Unionid string `json:"unionid"`
}

type Result struct {
	Ret         int          `json:"ret"`
	User        UserBaseInfo `json:"user"`
	Passport    PassportInfo `json:"passport"`
	Session_key string       `json:"session_key"`
	Paytype     int          `json:"paytype"`
}

func GetUserLoginInfo(openid int) (uid int, sessionKey string) {
	valueString := fmt.Sprintf("{\"deviceid\":\"8378832527\",\"method\":\"GameMember.otherLogin\",\"version\":\"2.3.1\",\"devicetype\":0,\"siteid\":1,\"model\":\"unknown\",\"param\":{\"logintime\":0,\"access_token\":\"\",\"openid\":%d,\"rtype\":9}}", openid)
	keystring := "win_param="
	reqData := keystring + valueString
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(reqData))
	if err != nil {
		mylog.LOG_ERROR("post error:", err)
	}
	defer resp.Body.Close()

	//fmt.Println("status", resp.Status)
	//fmt.Println("header", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		mylog.LOG_ERROR("handle resp eror", err)
	}
	var res Result
	json.Unmarshal(body, &res)
	uid = res.User.Uid
	sessionKey = res.Session_key
	return uid, sessionKey
}
