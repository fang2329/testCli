// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zajinhua_logic_msg.proto

package net

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 游戏状态
type MsgZajinhuaGameInfoFreeRep struct {
	CurPrivateRound  *uint32 `protobuf:"varint,1,opt,name=cur_private_round" json:"cur_private_round,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgZajinhuaGameInfoFreeRep) Reset()                    { *m = MsgZajinhuaGameInfoFreeRep{} }
func (m *MsgZajinhuaGameInfoFreeRep) String() string            { return proto.CompactTextString(m) }
func (*MsgZajinhuaGameInfoFreeRep) ProtoMessage()               {}
func (*MsgZajinhuaGameInfoFreeRep) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *MsgZajinhuaGameInfoFreeRep) GetCurPrivateRound() uint32 {
	if m != nil && m.CurPrivateRound != nil {
		return *m.CurPrivateRound
	}
	return 0
}

type MsgZajinhuaGameInfoPlayRep struct {
	// 加注信息
	CurrentTimes *uint32 `protobuf:"varint,3,opt,name=current_times" json:"current_times,omitempty"`
	UserMaxScore *int64  `protobuf:"varint,4,opt,name=user_max_score" json:"user_max_score,omitempty"`
	// 状态信息
	BankerUser  *uint32  `protobuf:"varint,5,opt,name=banker_user" json:"banker_user,omitempty"`
	CurrentUser *uint32  `protobuf:"varint,6,opt,name=current_user" json:"current_user,omitempty"`
	PlayStatus  []uint32 `protobuf:"varint,7,rep,name=play_status" json:"play_status,omitempty"`
	CardStat    []uint32 `protobuf:"varint,8,rep,name=card_stat" json:"card_stat,omitempty"`
	TableScore  []int64  `protobuf:"varint,9,rep,name=table_score" json:"table_score,omitempty"`
	// 扑克信息
	HandCardData []uint32 `protobuf:"varint,10,rep,name=hand_card_data" json:"hand_card_data,omitempty"`
	// 状态信息
	CompareState     *uint32 `protobuf:"varint,11,opt,name=compare_state" json:"compare_state,omitempty"`
	CurrentRound     *uint32 `protobuf:"varint,12,opt,name=current_round" json:"current_round,omitempty"`
	OperTime         *uint32 `protobuf:"varint,13,opt,name=oper_time" json:"oper_time,omitempty"`
	IsMingpai        *uint32 `protobuf:"varint,14,opt,name=is_mingpai" json:"is_mingpai,omitempty"`
	CurPrivateRound  *uint32 `protobuf:"varint,15,opt,name=cur_private_round" json:"cur_private_round,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgZajinhuaGameInfoPlayRep) Reset()                    { *m = MsgZajinhuaGameInfoPlayRep{} }
func (m *MsgZajinhuaGameInfoPlayRep) String() string            { return proto.CompactTextString(m) }
func (*MsgZajinhuaGameInfoPlayRep) ProtoMessage()               {}
func (*MsgZajinhuaGameInfoPlayRep) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *MsgZajinhuaGameInfoPlayRep) GetCurrentTimes() uint32 {
	if m != nil && m.CurrentTimes != nil {
		return *m.CurrentTimes
	}
	return 0
}

func (m *MsgZajinhuaGameInfoPlayRep) GetUserMaxScore() int64 {
	if m != nil && m.UserMaxScore != nil {
		return *m.UserMaxScore
	}
	return 0
}

func (m *MsgZajinhuaGameInfoPlayRep) GetBankerUser() uint32 {
	if m != nil && m.BankerUser != nil {
		return *m.BankerUser
	}
	return 0
}

func (m *MsgZajinhuaGameInfoPlayRep) GetCurrentUser() uint32 {
	if m != nil && m.CurrentUser != nil {
		return *m.CurrentUser
	}
	return 0
}

func (m *MsgZajinhuaGameInfoPlayRep) GetPlayStatus() []uint32 {
	if m != nil {
		return m.PlayStatus
	}
	return nil
}

func (m *MsgZajinhuaGameInfoPlayRep) GetCardStat() []uint32 {
	if m != nil {
		return m.CardStat
	}
	return nil
}

func (m *MsgZajinhuaGameInfoPlayRep) GetTableScore() []int64 {
	if m != nil {
		return m.TableScore
	}
	return nil
}

func (m *MsgZajinhuaGameInfoPlayRep) GetHandCardData() []uint32 {
	if m != nil {
		return m.HandCardData
	}
	return nil
}

func (m *MsgZajinhuaGameInfoPlayRep) GetCompareState() uint32 {
	if m != nil && m.CompareState != nil {
		return *m.CompareState
	}
	return 0
}

func (m *MsgZajinhuaGameInfoPlayRep) GetCurrentRound() uint32 {
	if m != nil && m.CurrentRound != nil {
		return *m.CurrentRound
	}
	return 0
}

func (m *MsgZajinhuaGameInfoPlayRep) GetOperTime() uint32 {
	if m != nil && m.OperTime != nil {
		return *m.OperTime
	}
	return 0
}

func (m *MsgZajinhuaGameInfoPlayRep) GetIsMingpai() uint32 {
	if m != nil && m.IsMingpai != nil {
		return *m.IsMingpai
	}
	return 0
}

func (m *MsgZajinhuaGameInfoPlayRep) GetCurPrivateRound() uint32 {
	if m != nil && m.CurPrivateRound != nil {
		return *m.CurPrivateRound
	}
	return 0
}

// 游戏开始
type MsgZajinhuaStartRep struct {
	// 用户信息
	BankerUser       *uint32 `protobuf:"varint,1,opt,name=banker_user" json:"banker_user,omitempty"`
	CurrentUser      *uint32 `protobuf:"varint,2,opt,name=current_user" json:"current_user,omitempty"`
	UserMaxScore     *int64  `protobuf:"varint,3,opt,name=user_max_score" json:"user_max_score,omitempty"`
	CurPrivateRound  *uint32 `protobuf:"varint,4,opt,name=cur_private_round" json:"cur_private_round,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgZajinhuaStartRep) Reset()                    { *m = MsgZajinhuaStartRep{} }
func (m *MsgZajinhuaStartRep) String() string            { return proto.CompactTextString(m) }
func (*MsgZajinhuaStartRep) ProtoMessage()               {}
func (*MsgZajinhuaStartRep) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *MsgZajinhuaStartRep) GetBankerUser() uint32 {
	if m != nil && m.BankerUser != nil {
		return *m.BankerUser
	}
	return 0
}

func (m *MsgZajinhuaStartRep) GetCurrentUser() uint32 {
	if m != nil && m.CurrentUser != nil {
		return *m.CurrentUser
	}
	return 0
}

func (m *MsgZajinhuaStartRep) GetUserMaxScore() int64 {
	if m != nil && m.UserMaxScore != nil {
		return *m.UserMaxScore
	}
	return 0
}

func (m *MsgZajinhuaStartRep) GetCurPrivateRound() uint32 {
	if m != nil && m.CurPrivateRound != nil {
		return *m.CurPrivateRound
	}
	return 0
}

// 用户下注
type MsgZajinhuaAddscoreReq struct {
	AddScore         *int64  `protobuf:"varint,1,opt,name=add_score" json:"add_score,omitempty"`
	IsAllin          *uint32 `protobuf:"varint,2,opt,name=is_allin" json:"is_allin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgZajinhuaAddscoreReq) Reset()                    { *m = MsgZajinhuaAddscoreReq{} }
func (m *MsgZajinhuaAddscoreReq) String() string            { return proto.CompactTextString(m) }
func (*MsgZajinhuaAddscoreReq) ProtoMessage()               {}
func (*MsgZajinhuaAddscoreReq) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func (m *MsgZajinhuaAddscoreReq) GetAddScore() int64 {
	if m != nil && m.AddScore != nil {
		return *m.AddScore
	}
	return 0
}

func (m *MsgZajinhuaAddscoreReq) GetIsAllin() uint32 {
	if m != nil && m.IsAllin != nil {
		return *m.IsAllin
	}
	return 0
}

type MsgZajinhuaAddscoreRep struct {
	CurrentUser      *uint32 `protobuf:"varint,1,opt,name=current_user" json:"current_user,omitempty"`
	AddScoreUser     *uint32 `protobuf:"varint,2,opt,name=add_score_user" json:"add_score_user,omitempty"`
	CompareState     *uint32 `protobuf:"varint,3,opt,name=compare_state" json:"compare_state,omitempty"`
	AddScoreCount    *int64  `protobuf:"varint,4,opt,name=add_score_count" json:"add_score_count,omitempty"`
	CurrentTimes     *uint32 `protobuf:"varint,5,opt,name=current_times" json:"current_times,omitempty"`
	CurrentRound     *uint32 `protobuf:"varint,6,opt,name=current_round" json:"current_round,omitempty"`
	OnlyCompare      *uint32 `protobuf:"varint,7,opt,name=only_compare" json:"only_compare,omitempty"`
	IsAllin          *uint32 `protobuf:"varint,8,opt,name=is_allin" json:"is_allin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgZajinhuaAddscoreRep) Reset()                    { *m = MsgZajinhuaAddscoreRep{} }
func (m *MsgZajinhuaAddscoreRep) String() string            { return proto.CompactTextString(m) }
func (*MsgZajinhuaAddscoreRep) ProtoMessage()               {}
func (*MsgZajinhuaAddscoreRep) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{4} }

func (m *MsgZajinhuaAddscoreRep) GetCurrentUser() uint32 {
	if m != nil && m.CurrentUser != nil {
		return *m.CurrentUser
	}
	return 0
}

func (m *MsgZajinhuaAddscoreRep) GetAddScoreUser() uint32 {
	if m != nil && m.AddScoreUser != nil {
		return *m.AddScoreUser
	}
	return 0
}

func (m *MsgZajinhuaAddscoreRep) GetCompareState() uint32 {
	if m != nil && m.CompareState != nil {
		return *m.CompareState
	}
	return 0
}

func (m *MsgZajinhuaAddscoreRep) GetAddScoreCount() int64 {
	if m != nil && m.AddScoreCount != nil {
		return *m.AddScoreCount
	}
	return 0
}

func (m *MsgZajinhuaAddscoreRep) GetCurrentTimes() uint32 {
	if m != nil && m.CurrentTimes != nil {
		return *m.CurrentTimes
	}
	return 0
}

func (m *MsgZajinhuaAddscoreRep) GetCurrentRound() uint32 {
	if m != nil && m.CurrentRound != nil {
		return *m.CurrentRound
	}
	return 0
}

func (m *MsgZajinhuaAddscoreRep) GetOnlyCompare() uint32 {
	if m != nil && m.OnlyCompare != nil {
		return *m.OnlyCompare
	}
	return 0
}

func (m *MsgZajinhuaAddscoreRep) GetIsAllin() uint32 {
	if m != nil && m.IsAllin != nil {
		return *m.IsAllin
	}
	return 0
}

// 用户放弃
type MsgZajinhuaGiveupReq struct {
	GiveupReason     *uint32 `protobuf:"varint,1,opt,name=giveup_reason" json:"giveup_reason,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgZajinhuaGiveupReq) Reset()                    { *m = MsgZajinhuaGiveupReq{} }
func (m *MsgZajinhuaGiveupReq) String() string            { return proto.CompactTextString(m) }
func (*MsgZajinhuaGiveupReq) ProtoMessage()               {}
func (*MsgZajinhuaGiveupReq) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{5} }

func (m *MsgZajinhuaGiveupReq) GetGiveupReason() uint32 {
	if m != nil && m.GiveupReason != nil {
		return *m.GiveupReason
	}
	return 0
}

type MsgZajinhuaGiveupRep struct {
	GiveupUser       *uint32 `protobuf:"varint,1,opt,name=giveup_user" json:"giveup_user,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgZajinhuaGiveupRep) Reset()                    { *m = MsgZajinhuaGiveupRep{} }
func (m *MsgZajinhuaGiveupRep) String() string            { return proto.CompactTextString(m) }
func (*MsgZajinhuaGiveupRep) ProtoMessage()               {}
func (*MsgZajinhuaGiveupRep) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{6} }

func (m *MsgZajinhuaGiveupRep) GetGiveupUser() uint32 {
	if m != nil && m.GiveupUser != nil {
		return *m.GiveupUser
	}
	return 0
}

// 比牌数据包
type MsgZajinhuaCompareCardReq struct {
	CompareUser      *uint32 `protobuf:"varint,1,opt,name=compare_user" json:"compare_user,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgZajinhuaCompareCardReq) Reset()                    { *m = MsgZajinhuaCompareCardReq{} }
func (m *MsgZajinhuaCompareCardReq) String() string            { return proto.CompactTextString(m) }
func (*MsgZajinhuaCompareCardReq) ProtoMessage()               {}
func (*MsgZajinhuaCompareCardReq) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{7} }

func (m *MsgZajinhuaCompareCardReq) GetCompareUser() uint32 {
	if m != nil && m.CompareUser != nil {
		return *m.CompareUser
	}
	return 0
}

type MsgZajinhuaCompareCardRep struct {
	CurrentUser      *uint32  `protobuf:"varint,1,opt,name=current_user" json:"current_user,omitempty"`
	CompareUser      []uint32 `protobuf:"varint,2,rep,name=compare_user" json:"compare_user,omitempty"`
	LostUser         *uint32  `protobuf:"varint,3,opt,name=lost_user" json:"lost_user,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *MsgZajinhuaCompareCardRep) Reset()                    { *m = MsgZajinhuaCompareCardRep{} }
func (m *MsgZajinhuaCompareCardRep) String() string            { return proto.CompactTextString(m) }
func (*MsgZajinhuaCompareCardRep) ProtoMessage()               {}
func (*MsgZajinhuaCompareCardRep) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{8} }

func (m *MsgZajinhuaCompareCardRep) GetCurrentUser() uint32 {
	if m != nil && m.CurrentUser != nil {
		return *m.CurrentUser
	}
	return 0
}

func (m *MsgZajinhuaCompareCardRep) GetCompareUser() []uint32 {
	if m != nil {
		return m.CompareUser
	}
	return nil
}

func (m *MsgZajinhuaCompareCardRep) GetLostUser() uint32 {
	if m != nil && m.LostUser != nil {
		return *m.LostUser
	}
	return 0
}

type MsgZajinhuaLookCardReq struct {
	Chairid          *uint32 `protobuf:"varint,1,opt,name=chairid" json:"chairid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgZajinhuaLookCardReq) Reset()                    { *m = MsgZajinhuaLookCardReq{} }
func (m *MsgZajinhuaLookCardReq) String() string            { return proto.CompactTextString(m) }
func (*MsgZajinhuaLookCardReq) ProtoMessage()               {}
func (*MsgZajinhuaLookCardReq) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{9} }

func (m *MsgZajinhuaLookCardReq) GetChairid() uint32 {
	if m != nil && m.Chairid != nil {
		return *m.Chairid
	}
	return 0
}

// 看牌数据包
type MsgZajinhuaLookCardRep struct {
	LookCardUser     *uint32  `protobuf:"varint,1,opt,name=look_card_user" json:"look_card_user,omitempty"`
	CardData         []uint32 `protobuf:"varint,2,rep,name=card_data" json:"card_data,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *MsgZajinhuaLookCardRep) Reset()                    { *m = MsgZajinhuaLookCardRep{} }
func (m *MsgZajinhuaLookCardRep) String() string            { return proto.CompactTextString(m) }
func (*MsgZajinhuaLookCardRep) ProtoMessage()               {}
func (*MsgZajinhuaLookCardRep) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{10} }

func (m *MsgZajinhuaLookCardRep) GetLookCardUser() uint32 {
	if m != nil && m.LookCardUser != nil {
		return *m.LookCardUser
	}
	return 0
}

func (m *MsgZajinhuaLookCardRep) GetCardData() []uint32 {
	if m != nil {
		return m.CardData
	}
	return nil
}

// 开牌
type MsgZajinhuaOpenCardReq struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgZajinhuaOpenCardReq) Reset()                    { *m = MsgZajinhuaOpenCardReq{} }
func (m *MsgZajinhuaOpenCardReq) String() string            { return proto.CompactTextString(m) }
func (*MsgZajinhuaOpenCardReq) ProtoMessage()               {}
func (*MsgZajinhuaOpenCardReq) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{11} }

// 开牌数据包
type MsgZajinhuaOpenCard struct {
	Winner           *uint32 `protobuf:"varint,1,opt,name=winner" json:"winner,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgZajinhuaOpenCard) Reset()                    { *m = MsgZajinhuaOpenCard{} }
func (m *MsgZajinhuaOpenCard) String() string            { return proto.CompactTextString(m) }
func (*MsgZajinhuaOpenCard) ProtoMessage()               {}
func (*MsgZajinhuaOpenCard) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{12} }

func (m *MsgZajinhuaOpenCard) GetWinner() uint32 {
	if m != nil && m.Winner != nil {
		return *m.Winner
	}
	return 0
}

// 游戏结束
type MsgZajinhuaGameEndRep struct {
	GameScore        []int64     `protobuf:"varint,1,rep,name=game_score" json:"game_score,omitempty"`
	CardData         []*MsgCards `protobuf:"bytes,2,rep,name=card_data" json:"card_data,omitempty"`
	CompareUser      []uint32    `protobuf:"varint,3,rep,name=compare_user" json:"compare_user,omitempty"`
	EndState         *uint32     `protobuf:"varint,4,opt,name=end_state" json:"end_state,omitempty"`
	SpecialScore     []int64     `protobuf:"varint,5,rep,name=special_score" json:"special_score,omitempty"`
	CurPrivateRound  *uint32     `protobuf:"varint,6,opt,name=cur_private_round" json:"cur_private_round,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *MsgZajinhuaGameEndRep) Reset()                    { *m = MsgZajinhuaGameEndRep{} }
func (m *MsgZajinhuaGameEndRep) String() string            { return proto.CompactTextString(m) }
func (*MsgZajinhuaGameEndRep) ProtoMessage()               {}
func (*MsgZajinhuaGameEndRep) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{13} }

func (m *MsgZajinhuaGameEndRep) GetGameScore() []int64 {
	if m != nil {
		return m.GameScore
	}
	return nil
}

func (m *MsgZajinhuaGameEndRep) GetCardData() []*MsgCards {
	if m != nil {
		return m.CardData
	}
	return nil
}

func (m *MsgZajinhuaGameEndRep) GetCompareUser() []uint32 {
	if m != nil {
		return m.CompareUser
	}
	return nil
}

func (m *MsgZajinhuaGameEndRep) GetEndState() uint32 {
	if m != nil && m.EndState != nil {
		return *m.EndState
	}
	return 0
}

func (m *MsgZajinhuaGameEndRep) GetSpecialScore() []int64 {
	if m != nil {
		return m.SpecialScore
	}
	return nil
}

func (m *MsgZajinhuaGameEndRep) GetCurPrivateRound() uint32 {
	if m != nil && m.CurPrivateRound != nil {
		return *m.CurPrivateRound
	}
	return 0
}

// 亮牌
type MsgZajinhuaShowCardReq struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgZajinhuaShowCardReq) Reset()                    { *m = MsgZajinhuaShowCardReq{} }
func (m *MsgZajinhuaShowCardReq) String() string            { return proto.CompactTextString(m) }
func (*MsgZajinhuaShowCardReq) ProtoMessage()               {}
func (*MsgZajinhuaShowCardReq) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{14} }

type MsgZajinhuaShowCardRep struct {
	ShowChairid      *uint32  `protobuf:"varint,1,opt,name=show_chairid" json:"show_chairid,omitempty"`
	Cards            []uint32 `protobuf:"varint,2,rep,name=cards" json:"cards,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *MsgZajinhuaShowCardRep) Reset()                    { *m = MsgZajinhuaShowCardRep{} }
func (m *MsgZajinhuaShowCardRep) String() string            { return proto.CompactTextString(m) }
func (*MsgZajinhuaShowCardRep) ProtoMessage()               {}
func (*MsgZajinhuaShowCardRep) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{15} }

func (m *MsgZajinhuaShowCardRep) GetShowChairid() uint32 {
	if m != nil && m.ShowChairid != nil {
		return *m.ShowChairid
	}
	return 0
}

func (m *MsgZajinhuaShowCardRep) GetCards() []uint32 {
	if m != nil {
		return m.Cards
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgZajinhuaGameInfoFreeRep)(nil), "net.msg_zajinhua_game_info_free_rep")
	proto.RegisterType((*MsgZajinhuaGameInfoPlayRep)(nil), "net.msg_zajinhua_game_info_play_rep")
	proto.RegisterType((*MsgZajinhuaStartRep)(nil), "net.msg_zajinhua_start_rep")
	proto.RegisterType((*MsgZajinhuaAddscoreReq)(nil), "net.msg_zajinhua_addscore_req")
	proto.RegisterType((*MsgZajinhuaAddscoreRep)(nil), "net.msg_zajinhua_addscore_rep")
	proto.RegisterType((*MsgZajinhuaGiveupReq)(nil), "net.msg_zajinhua_giveup_req")
	proto.RegisterType((*MsgZajinhuaGiveupRep)(nil), "net.msg_zajinhua_giveup_rep")
	proto.RegisterType((*MsgZajinhuaCompareCardReq)(nil), "net.msg_zajinhua_compare_card_req")
	proto.RegisterType((*MsgZajinhuaCompareCardRep)(nil), "net.msg_zajinhua_compare_card_rep")
	proto.RegisterType((*MsgZajinhuaLookCardReq)(nil), "net.msg_zajinhua_look_card_req")
	proto.RegisterType((*MsgZajinhuaLookCardRep)(nil), "net.msg_zajinhua_look_card_rep")
	proto.RegisterType((*MsgZajinhuaOpenCardReq)(nil), "net.msg_zajinhua_open_card_req")
	proto.RegisterType((*MsgZajinhuaOpenCard)(nil), "net.msg_zajinhua_open_card")
	proto.RegisterType((*MsgZajinhuaGameEndRep)(nil), "net.msg_zajinhua_game_end_rep")
	proto.RegisterType((*MsgZajinhuaShowCardReq)(nil), "net.msg_zajinhua_show_card_req")
	proto.RegisterType((*MsgZajinhuaShowCardRep)(nil), "net.msg_zajinhua_show_card_rep")
}

func init() { proto.RegisterFile("zajinhua_logic_msg.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 621 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x54, 0xcd, 0x72, 0xd3, 0x30,
	0x10, 0x1e, 0xd7, 0x4d, 0x7f, 0x36, 0x4d, 0xda, 0xba, 0xb4, 0x55, 0x3b, 0x30, 0x04, 0x9f, 0x72,
	0x21, 0xc3, 0x30, 0xc3, 0x8d, 0x03, 0x9c, 0x78, 0x0b, 0x34, 0xaa, 0xbd, 0x4d, 0x44, 0x1d, 0x49,
	0x48, 0x72, 0x4b, 0x79, 0x27, 0xde, 0x81, 0x37, 0xe0, 0x95, 0x18, 0xad, 0x7f, 0xe2, 0x38, 0x49,
	0x8f, 0xfe, 0xb4, 0xfb, 0x7d, 0xbb, 0xdf, 0xee, 0x1a, 0xd8, 0x6f, 0xf1, 0x43, 0xaa, 0x45, 0x29,
	0x78, 0xa1, 0xe7, 0x32, 0xe3, 0x4b, 0x37, 0x9f, 0x19, 0xab, 0xbd, 0x4e, 0x62, 0x85, 0xfe, 0xf6,
	0xf2, 0x4e, 0x38, 0x0c, 0x20, 0x77, 0xde, 0x96, 0x99, 0xaf, 0xde, 0x6e, 0xcf, 0xd0, 0x5a, 0x6d,
	0x79, 0xa6, 0x73, 0xac, 0x90, 0xf4, 0x33, 0xbc, 0x0d, 0x51, 0x2d, 0xdb, 0x5c, 0x2c, 0x91, 0x4b,
	0x75, 0xaf, 0xf9, 0xbd, 0x45, 0xe4, 0x16, 0x4d, 0x72, 0x03, 0xe7, 0x59, 0x69, 0xb9, 0xb1, 0xf2,
	0x51, 0x78, 0xe4, 0x56, 0x97, 0x2a, 0x67, 0xd1, 0x24, 0x9a, 0x8e, 0xd2, 0xbf, 0x7b, 0x3b, 0xd3,
	0x4d, 0x21, 0x9e, 0x29, 0xfd, 0x12, 0x46, 0x59, 0x69, 0x2d, 0x2a, 0xcf, 0xbd, 0x5c, 0xa2, 0x63,
	0x71, 0x48, 0x4d, 0xae, 0x60, 0x5c, 0x3a, 0xb4, 0x7c, 0x29, 0x7e, 0x71, 0x97, 0x69, 0x8b, 0x6c,
	0x7f, 0x12, 0x4d, 0xe3, 0xe4, 0x02, 0x86, 0x77, 0x42, 0x3d, 0xa0, 0xe5, 0xe1, 0x99, 0x0d, 0x28,
	0xf8, 0x15, 0x9c, 0x34, 0x1c, 0x84, 0x1e, 0x10, 0x7a, 0x01, 0x43, 0x52, 0x71, 0x5e, 0xf8, 0xd2,
	0xb1, 0xc3, 0x49, 0x3c, 0x1d, 0x25, 0xe7, 0x70, 0x9c, 0x09, 0x9b, 0x13, 0xc8, 0x8e, 0x08, 0xba,
	0x80, 0xa1, 0x17, 0x77, 0x05, 0xd6, 0x3a, 0xc7, 0x93, 0x78, 0x1a, 0x07, 0xfd, 0x85, 0x50, 0x39,
	0xa7, 0xe0, 0x5c, 0x78, 0xc1, 0x80, 0x82, 0x43, 0xb9, 0x7a, 0x69, 0x84, 0x45, 0xa2, 0x40, 0x36,
	0x24, 0xad, 0x4e, 0x17, 0x95, 0x01, 0x27, 0x04, 0x9f, 0xc3, 0xb1, 0x36, 0x68, 0xa9, 0x33, 0x36,
	0x22, 0x28, 0x01, 0x90, 0x8e, 0x2f, 0xa5, 0x9a, 0x1b, 0x21, 0xd9, 0x98, 0xb0, 0xad, 0x16, 0x9e,
	0x92, 0x85, 0x8f, 0x70, 0xb5, 0xe6, 0xa0, 0xf3, 0xc2, 0x7a, 0x32, 0xae, 0xe7, 0x44, 0xb4, 0xd5,
	0x89, 0xbd, 0x1d, 0x66, 0xc6, 0x64, 0xe6, 0x56, 0xdd, 0x7d, 0xd2, 0xfd, 0x02, 0x37, 0x6b, 0xba,
	0x22, 0xcf, 0x29, 0x93, 0x5b, 0xfc, 0x19, 0xda, 0x12, 0x79, 0x5e, 0x53, 0x45, 0x44, 0x75, 0x06,
	0x47, 0xd2, 0x71, 0x51, 0x14, 0x52, 0x55, 0xa2, 0xe9, 0xbf, 0x68, 0x37, 0x85, 0xd9, 0x28, 0x34,
	0x6a, 0x0a, 0x6d, 0x89, 0xbb, 0x0d, 0x6c, 0xb8, 0x5e, 0x2d, 0xc9, 0x35, 0x9c, 0xae, 0xc2, 0x33,
	0x5d, 0x2a, 0x5f, 0x6f, 0xc9, 0xc6, 0x52, 0x0d, 0xb6, 0x4f, 0xe9, 0xa0, 0x31, 0x4d, 0xab, 0xe2,
	0x99, 0xd7, 0x12, 0xec, 0x90, 0xd0, 0x6e, 0x47, 0x47, 0xd4, 0xd1, 0x07, 0xb8, 0x5e, 0xdf, 0x66,
	0xf9, 0x88, 0xa5, 0x21, 0x47, 0x2e, 0x61, 0xd4, 0x7e, 0x09, 0xa7, 0x55, 0x7d, 0x00, 0xb3, 0x5d,
	0x19, 0x34, 0xbe, 0xfa, 0x6b, 0xd5, 0x7f, 0xfa, 0x09, 0xde, 0xac, 0xc5, 0x37, 0x4d, 0xd3, 0x16,
	0x06, 0x9d, 0x60, 0x5b, 0x8d, 0x75, 0xd2, 0xbe, 0xbf, 0x9c, 0xb6, 0xcb, 0xed, 0x3e, 0xd9, 0x5e,
	0x73, 0x21, 0x85, 0x76, 0x75, 0x20, 0xf9, 0x9c, 0xbe, 0x87, 0xdb, 0x35, 0xfe, 0x42, 0xeb, 0x87,
	0x55, 0x4d, 0xa7, 0x70, 0x98, 0x2d, 0x84, 0xb4, 0xb2, 0x39, 0xfb, 0x6f, 0x2f, 0x84, 0x9b, 0x30,
	0xe3, 0x15, 0xd0, 0xa9, 0xa6, 0xb9, 0x4c, 0x3a, 0x36, 0x2a, 0x25, 0x7d, 0xdd, 0x23, 0xd2, 0x06,
	0x55, 0xab, 0x9b, 0x4e, 0x7b, 0xa7, 0xd1, 0xbe, 0x26, 0x63, 0x38, 0x78, 0x92, 0x4a, 0xb5, 0xfe,
	0xfc, 0xe9, 0xaf, 0x22, 0xfd, 0x87, 0x50, 0x55, 0x05, 0x25, 0x00, 0xf4, 0xdd, 0xac, 0x73, 0x38,
	0xff, 0x77, 0xfd, 0x62, 0x86, 0x1f, 0xc7, 0x33, 0x85, 0x7e, 0x16, 0x68, 0xc2, 0x8b, 0xdb, 0x70,
	0x2f, 0x6e, 0xdc, 0x0b, 0xbc, 0xd5, 0x96, 0xee, 0x37, 0x5b, 0xe7, 0x0c, 0x66, 0x52, 0x14, 0xb5,
	0xc4, 0x80, 0x24, 0xb6, 0x1e, 0x1f, 0x2d, 0xe4, 0x46, 0xdf, 0x6e, 0xa1, 0x9f, 0x56, 0x7d, 0x7f,
	0x7d, 0xe1, 0x95, 0x46, 0x5d, 0x01, 0xdd, 0x91, 0x24, 0x23, 0x18, 0x50, 0xd5, 0x95, 0xb1, 0xff,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x63, 0xf8, 0x9e, 0x4f, 0x1f, 0x06, 0x00, 0x00,
}