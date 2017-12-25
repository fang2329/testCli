// Code generated by protoc-gen-go. DO NOT EDIT.
// source: texas_logic_msg.proto

package net

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 游戏状态
type MsgTexasGameInfoFreeRep struct {
	CellMinscore     *int64 `protobuf:"varint,1,opt,name=cell_minscore" json:"cell_minscore,omitempty"`
	CellMaxscore     *int64 `protobuf:"varint,2,opt,name=cell_maxscore" json:"cell_maxscore,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgTexasGameInfoFreeRep) Reset()                    { *m = MsgTexasGameInfoFreeRep{} }
func (m *MsgTexasGameInfoFreeRep) String() string            { return proto.CompactTextString(m) }
func (*MsgTexasGameInfoFreeRep) ProtoMessage()               {}
func (*MsgTexasGameInfoFreeRep) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *MsgTexasGameInfoFreeRep) GetCellMinscore() int64 {
	if m != nil && m.CellMinscore != nil {
		return *m.CellMinscore
	}
	return 0
}

func (m *MsgTexasGameInfoFreeRep) GetCellMaxscore() int64 {
	if m != nil && m.CellMaxscore != nil {
		return *m.CellMaxscore
	}
	return 0
}

type MsgTexasGameInfoPlayRep struct {
	// 加注信息
	CellScore     *int64  `protobuf:"varint,1,opt,name=cell_score" json:"cell_score,omitempty"`
	TurnMaxScore  *int64  `protobuf:"varint,2,opt,name=turn_max_score" json:"turn_max_score,omitempty"`
	TurnLessScore *int64  `protobuf:"varint,3,opt,name=turn_less_score" json:"turn_less_score,omitempty"`
	CellMaxScore  *int64  `protobuf:"varint,4,opt,name=cell_max_score" json:"cell_max_score,omitempty"`
	AddLessScore  *int64  `protobuf:"varint,5,opt,name=add_less_score" json:"add_less_score,omitempty"`
	TableScores   []int64 `protobuf:"varint,6,rep,name=table_scores" json:"table_scores,omitempty"`
	TotalScores   []int64 `protobuf:"varint,7,rep,name=total_scores" json:"total_scores,omitempty"`
	CenterScore   *int64  `protobuf:"varint,8,opt,name=center_score" json:"center_score,omitempty"`
	// 状态信息
	DUser        *uint32  `protobuf:"varint,9,opt,name=d_user" json:"d_user,omitempty"`
	CurrentUser  *uint32  `protobuf:"varint,10,opt,name=current_user" json:"current_user,omitempty"`
	PlayStatus   []uint32 `protobuf:"varint,11,rep,name=play_status" json:"play_status,omitempty"`
	BalanceCount *uint32  `protobuf:"varint,12,opt,name=balance_count" json:"balance_count,omitempty"`
	// 扑克信息
	CenterCards      []uint32 `protobuf:"varint,13,rep,name=center_cards" json:"center_cards,omitempty"`
	HandCards        []uint32 `protobuf:"varint,14,rep,name=hand_cards" json:"hand_cards,omitempty"`
	OperTime         *uint32  `protobuf:"varint,15,opt,name=oper_time" json:"oper_time,omitempty"`
	MinChipinUser    *uint32  `protobuf:"varint,16,opt,name=min_chipin_user" json:"min_chipin_user,omitempty"`
	MaxChipinUser    *uint32  `protobuf:"varint,17,opt,name=max_chipin_user" json:"max_chipin_user,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *MsgTexasGameInfoPlayRep) Reset()                    { *m = MsgTexasGameInfoPlayRep{} }
func (m *MsgTexasGameInfoPlayRep) String() string            { return proto.CompactTextString(m) }
func (*MsgTexasGameInfoPlayRep) ProtoMessage()               {}
func (*MsgTexasGameInfoPlayRep) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *MsgTexasGameInfoPlayRep) GetCellScore() int64 {
	if m != nil && m.CellScore != nil {
		return *m.CellScore
	}
	return 0
}

func (m *MsgTexasGameInfoPlayRep) GetTurnMaxScore() int64 {
	if m != nil && m.TurnMaxScore != nil {
		return *m.TurnMaxScore
	}
	return 0
}

func (m *MsgTexasGameInfoPlayRep) GetTurnLessScore() int64 {
	if m != nil && m.TurnLessScore != nil {
		return *m.TurnLessScore
	}
	return 0
}

func (m *MsgTexasGameInfoPlayRep) GetCellMaxScore() int64 {
	if m != nil && m.CellMaxScore != nil {
		return *m.CellMaxScore
	}
	return 0
}

func (m *MsgTexasGameInfoPlayRep) GetAddLessScore() int64 {
	if m != nil && m.AddLessScore != nil {
		return *m.AddLessScore
	}
	return 0
}

func (m *MsgTexasGameInfoPlayRep) GetTableScores() []int64 {
	if m != nil {
		return m.TableScores
	}
	return nil
}

func (m *MsgTexasGameInfoPlayRep) GetTotalScores() []int64 {
	if m != nil {
		return m.TotalScores
	}
	return nil
}

func (m *MsgTexasGameInfoPlayRep) GetCenterScore() int64 {
	if m != nil && m.CenterScore != nil {
		return *m.CenterScore
	}
	return 0
}

func (m *MsgTexasGameInfoPlayRep) GetDUser() uint32 {
	if m != nil && m.DUser != nil {
		return *m.DUser
	}
	return 0
}

func (m *MsgTexasGameInfoPlayRep) GetCurrentUser() uint32 {
	if m != nil && m.CurrentUser != nil {
		return *m.CurrentUser
	}
	return 0
}

func (m *MsgTexasGameInfoPlayRep) GetPlayStatus() []uint32 {
	if m != nil {
		return m.PlayStatus
	}
	return nil
}

func (m *MsgTexasGameInfoPlayRep) GetBalanceCount() uint32 {
	if m != nil && m.BalanceCount != nil {
		return *m.BalanceCount
	}
	return 0
}

func (m *MsgTexasGameInfoPlayRep) GetCenterCards() []uint32 {
	if m != nil {
		return m.CenterCards
	}
	return nil
}

func (m *MsgTexasGameInfoPlayRep) GetHandCards() []uint32 {
	if m != nil {
		return m.HandCards
	}
	return nil
}

func (m *MsgTexasGameInfoPlayRep) GetOperTime() uint32 {
	if m != nil && m.OperTime != nil {
		return *m.OperTime
	}
	return 0
}

func (m *MsgTexasGameInfoPlayRep) GetMinChipinUser() uint32 {
	if m != nil && m.MinChipinUser != nil {
		return *m.MinChipinUser
	}
	return 0
}

func (m *MsgTexasGameInfoPlayRep) GetMaxChipinUser() uint32 {
	if m != nil && m.MaxChipinUser != nil {
		return *m.MaxChipinUser
	}
	return 0
}

// 游戏开始
type MsgTexasStartRep struct {
	CurrentUser      *uint32   `protobuf:"varint,1,opt,name=current_user" json:"current_user,omitempty"`
	DUser            *uint32   `protobuf:"varint,2,opt,name=d_user" json:"d_user,omitempty"`
	MaxChipinUser    *uint32   `protobuf:"varint,3,opt,name=max_chipin_user" json:"max_chipin_user,omitempty"`
	CellScore        *int64    `protobuf:"varint,4,opt,name=cell_score" json:"cell_score,omitempty"`
	TurnMaxScore     *int64    `protobuf:"varint,5,opt,name=turn_max_score" json:"turn_max_score,omitempty"`
	TurnLessScore    *int64    `protobuf:"varint,6,opt,name=turn_less_score" json:"turn_less_score,omitempty"`
	AddLessScore     *int64    `protobuf:"varint,7,opt,name=add_less_score" json:"add_less_score,omitempty"`
	CardDatas        *MsgCards `protobuf:"bytes,8,opt,name=card_datas" json:"card_datas,omitempty"`
	MinChipinUser    *uint32   `protobuf:"varint,9,opt,name=min_chipin_user" json:"min_chipin_user,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *MsgTexasStartRep) Reset()                    { *m = MsgTexasStartRep{} }
func (m *MsgTexasStartRep) String() string            { return proto.CompactTextString(m) }
func (*MsgTexasStartRep) ProtoMessage()               {}
func (*MsgTexasStartRep) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

func (m *MsgTexasStartRep) GetCurrentUser() uint32 {
	if m != nil && m.CurrentUser != nil {
		return *m.CurrentUser
	}
	return 0
}

func (m *MsgTexasStartRep) GetDUser() uint32 {
	if m != nil && m.DUser != nil {
		return *m.DUser
	}
	return 0
}

func (m *MsgTexasStartRep) GetMaxChipinUser() uint32 {
	if m != nil && m.MaxChipinUser != nil {
		return *m.MaxChipinUser
	}
	return 0
}

func (m *MsgTexasStartRep) GetCellScore() int64 {
	if m != nil && m.CellScore != nil {
		return *m.CellScore
	}
	return 0
}

func (m *MsgTexasStartRep) GetTurnMaxScore() int64 {
	if m != nil && m.TurnMaxScore != nil {
		return *m.TurnMaxScore
	}
	return 0
}

func (m *MsgTexasStartRep) GetTurnLessScore() int64 {
	if m != nil && m.TurnLessScore != nil {
		return *m.TurnLessScore
	}
	return 0
}

func (m *MsgTexasStartRep) GetAddLessScore() int64 {
	if m != nil && m.AddLessScore != nil {
		return *m.AddLessScore
	}
	return 0
}

func (m *MsgTexasStartRep) GetCardDatas() *MsgCards {
	if m != nil {
		return m.CardDatas
	}
	return nil
}

func (m *MsgTexasStartRep) GetMinChipinUser() uint32 {
	if m != nil && m.MinChipinUser != nil {
		return *m.MinChipinUser
	}
	return 0
}

// 用户下注
type MsgTexasAddscoreReq struct {
	AddScore         *int64 `protobuf:"varint,1,opt,name=add_score" json:"add_score,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgTexasAddscoreReq) Reset()                    { *m = MsgTexasAddscoreReq{} }
func (m *MsgTexasAddscoreReq) String() string            { return proto.CompactTextString(m) }
func (*MsgTexasAddscoreReq) ProtoMessage()               {}
func (*MsgTexasAddscoreReq) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

func (m *MsgTexasAddscoreReq) GetAddScore() int64 {
	if m != nil && m.AddScore != nil {
		return *m.AddScore
	}
	return 0
}

type MsgTexasAddscoreRep struct {
	CurrentUser      *uint32  `protobuf:"varint,1,opt,name=current_user" json:"current_user,omitempty"`
	AddscoreUser     *uint32  `protobuf:"varint,2,opt,name=addscore_user" json:"addscore_user,omitempty"`
	AddscoreCount    *int64   `protobuf:"varint,3,opt,name=addscore_count" json:"addscore_count,omitempty"`
	TurnLessScore    *int64   `protobuf:"varint,4,opt,name=turn_less_score" json:"turn_less_score,omitempty"`
	TurnMaxScore     *int64   `protobuf:"varint,5,opt,name=turn_max_score" json:"turn_max_score,omitempty"`
	AddLessScore     *int64   `protobuf:"varint,6,opt,name=add_less_score" json:"add_less_score,omitempty"`
	Showhands        []uint32 `protobuf:"varint,7,rep,name=showhands" json:"showhands,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *MsgTexasAddscoreRep) Reset()                    { *m = MsgTexasAddscoreRep{} }
func (m *MsgTexasAddscoreRep) String() string            { return proto.CompactTextString(m) }
func (*MsgTexasAddscoreRep) ProtoMessage()               {}
func (*MsgTexasAddscoreRep) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

func (m *MsgTexasAddscoreRep) GetCurrentUser() uint32 {
	if m != nil && m.CurrentUser != nil {
		return *m.CurrentUser
	}
	return 0
}

func (m *MsgTexasAddscoreRep) GetAddscoreUser() uint32 {
	if m != nil && m.AddscoreUser != nil {
		return *m.AddscoreUser
	}
	return 0
}

func (m *MsgTexasAddscoreRep) GetAddscoreCount() int64 {
	if m != nil && m.AddscoreCount != nil {
		return *m.AddscoreCount
	}
	return 0
}

func (m *MsgTexasAddscoreRep) GetTurnLessScore() int64 {
	if m != nil && m.TurnLessScore != nil {
		return *m.TurnLessScore
	}
	return 0
}

func (m *MsgTexasAddscoreRep) GetTurnMaxScore() int64 {
	if m != nil && m.TurnMaxScore != nil {
		return *m.TurnMaxScore
	}
	return 0
}

func (m *MsgTexasAddscoreRep) GetAddLessScore() int64 {
	if m != nil && m.AddLessScore != nil {
		return *m.AddLessScore
	}
	return 0
}

func (m *MsgTexasAddscoreRep) GetShowhands() []uint32 {
	if m != nil {
		return m.Showhands
	}
	return nil
}

// 用户放弃
type MsgTexasGiveupReq struct {
	GiveupReason     *uint32 `protobuf:"varint,1,opt,name=giveup_reason" json:"giveup_reason,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgTexasGiveupReq) Reset()                    { *m = MsgTexasGiveupReq{} }
func (m *MsgTexasGiveupReq) String() string            { return proto.CompactTextString(m) }
func (*MsgTexasGiveupReq) ProtoMessage()               {}
func (*MsgTexasGiveupReq) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{5} }

func (m *MsgTexasGiveupReq) GetGiveupReason() uint32 {
	if m != nil && m.GiveupReason != nil {
		return *m.GiveupReason
	}
	return 0
}

type MsgTexasGiveupRep struct {
	GiveupUser       *uint32 `protobuf:"varint,1,opt,name=giveup_user" json:"giveup_user,omitempty"`
	LostScore        *int64  `protobuf:"varint,2,opt,name=lost_score" json:"lost_score,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgTexasGiveupRep) Reset()                    { *m = MsgTexasGiveupRep{} }
func (m *MsgTexasGiveupRep) String() string            { return proto.CompactTextString(m) }
func (*MsgTexasGiveupRep) ProtoMessage()               {}
func (*MsgTexasGiveupRep) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{6} }

func (m *MsgTexasGiveupRep) GetGiveupUser() uint32 {
	if m != nil && m.GiveupUser != nil {
		return *m.GiveupUser
	}
	return 0
}

func (m *MsgTexasGiveupRep) GetLostScore() int64 {
	if m != nil && m.LostScore != nil {
		return *m.LostScore
	}
	return 0
}

// 发牌
type MsgTexasSendcardRep struct {
	PublicCard       *uint32  `protobuf:"varint,1,opt,name=public_card" json:"public_card,omitempty"`
	CurrentUser      *uint32  `protobuf:"varint,2,opt,name=current_user" json:"current_user,omitempty"`
	SendCardCount    *uint32  `protobuf:"varint,3,opt,name=send_card_count" json:"send_card_count,omitempty"`
	CenterCards      []uint32 `protobuf:"varint,4,rep,name=center_cards" json:"center_cards,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *MsgTexasSendcardRep) Reset()                    { *m = MsgTexasSendcardRep{} }
func (m *MsgTexasSendcardRep) String() string            { return proto.CompactTextString(m) }
func (*MsgTexasSendcardRep) ProtoMessage()               {}
func (*MsgTexasSendcardRep) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{7} }

func (m *MsgTexasSendcardRep) GetPublicCard() uint32 {
	if m != nil && m.PublicCard != nil {
		return *m.PublicCard
	}
	return 0
}

func (m *MsgTexasSendcardRep) GetCurrentUser() uint32 {
	if m != nil && m.CurrentUser != nil {
		return *m.CurrentUser
	}
	return 0
}

func (m *MsgTexasSendcardRep) GetSendCardCount() uint32 {
	if m != nil && m.SendCardCount != nil {
		return *m.SendCardCount
	}
	return 0
}

func (m *MsgTexasSendcardRep) GetCenterCards() []uint32 {
	if m != nil {
		return m.CenterCards
	}
	return nil
}

// 游戏结束
type TexasWinChair struct {
	WinChair         []uint32 `protobuf:"varint,1,rep,name=win_chair" json:"win_chair,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TexasWinChair) Reset()                    { *m = TexasWinChair{} }
func (m *TexasWinChair) String() string            { return proto.CompactTextString(m) }
func (*TexasWinChair) ProtoMessage()               {}
func (*TexasWinChair) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{8} }

func (m *TexasWinChair) GetWinChair() []uint32 {
	if m != nil {
		return m.WinChair
	}
	return nil
}

type MsgTexasGameEndRep struct {
	TotalEnd           *uint32          `protobuf:"varint,1,opt,name=total_end" json:"total_end,omitempty"`
	GameScores         []int64          `protobuf:"varint,2,rep,name=game_scores" json:"game_scores,omitempty"`
	CardDatas          []*MsgCards      `protobuf:"bytes,3,rep,name=card_datas" json:"card_datas,omitempty"`
	LastCenterCardData []*MsgCards      `protobuf:"bytes,4,rep,name=last_center_cardData" json:"last_center_cardData,omitempty"`
	WinChairs          []*TexasWinChair `protobuf:"bytes,5,rep,name=win_chairs" json:"win_chairs,omitempty"`
	XXX_unrecognized   []byte           `json:"-"`
}

func (m *MsgTexasGameEndRep) Reset()                    { *m = MsgTexasGameEndRep{} }
func (m *MsgTexasGameEndRep) String() string            { return proto.CompactTextString(m) }
func (*MsgTexasGameEndRep) ProtoMessage()               {}
func (*MsgTexasGameEndRep) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{9} }

func (m *MsgTexasGameEndRep) GetTotalEnd() uint32 {
	if m != nil && m.TotalEnd != nil {
		return *m.TotalEnd
	}
	return 0
}

func (m *MsgTexasGameEndRep) GetGameScores() []int64 {
	if m != nil {
		return m.GameScores
	}
	return nil
}

func (m *MsgTexasGameEndRep) GetCardDatas() []*MsgCards {
	if m != nil {
		return m.CardDatas
	}
	return nil
}

func (m *MsgTexasGameEndRep) GetLastCenterCardData() []*MsgCards {
	if m != nil {
		return m.LastCenterCardData
	}
	return nil
}

func (m *MsgTexasGameEndRep) GetWinChairs() []*TexasWinChair {
	if m != nil {
		return m.WinChairs
	}
	return nil
}

// buyin
type MsgTexasBuyinReq struct {
	Score            *int64 `protobuf:"varint,1,opt,name=score" json:"score,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgTexasBuyinReq) Reset()                    { *m = MsgTexasBuyinReq{} }
func (m *MsgTexasBuyinReq) String() string            { return proto.CompactTextString(m) }
func (*MsgTexasBuyinReq) ProtoMessage()               {}
func (*MsgTexasBuyinReq) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{10} }

func (m *MsgTexasBuyinReq) GetScore() int64 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

type MsgTexasBuyinRep struct {
	Score            *int64  `protobuf:"varint,1,opt,name=score" json:"score,omitempty"`
	Result           *uint32 `protobuf:"varint,2,opt,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgTexasBuyinRep) Reset()                    { *m = MsgTexasBuyinRep{} }
func (m *MsgTexasBuyinRep) String() string            { return proto.CompactTextString(m) }
func (*MsgTexasBuyinRep) ProtoMessage()               {}
func (*MsgTexasBuyinRep) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{11} }

func (m *MsgTexasBuyinRep) GetScore() int64 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *MsgTexasBuyinRep) GetResult() uint32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

// 亮牌
type MsgTexasShowCardReq struct {
	Cards            []uint32 `protobuf:"varint,1,rep,name=cards" json:"cards,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *MsgTexasShowCardReq) Reset()                    { *m = MsgTexasShowCardReq{} }
func (m *MsgTexasShowCardReq) String() string            { return proto.CompactTextString(m) }
func (*MsgTexasShowCardReq) ProtoMessage()               {}
func (*MsgTexasShowCardReq) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{12} }

func (m *MsgTexasShowCardReq) GetCards() []uint32 {
	if m != nil {
		return m.Cards
	}
	return nil
}

type MsgTexasShowCardRep struct {
	ShowChairid      *uint32  `protobuf:"varint,1,opt,name=show_chairid" json:"show_chairid,omitempty"`
	Cards            []uint32 `protobuf:"varint,2,rep,name=cards" json:"cards,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *MsgTexasShowCardRep) Reset()                    { *m = MsgTexasShowCardRep{} }
func (m *MsgTexasShowCardRep) String() string            { return proto.CompactTextString(m) }
func (*MsgTexasShowCardRep) ProtoMessage()               {}
func (*MsgTexasShowCardRep) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{13} }

func (m *MsgTexasShowCardRep) GetShowChairid() uint32 {
	if m != nil && m.ShowChairid != nil {
		return *m.ShowChairid
	}
	return 0
}

func (m *MsgTexasShowCardRep) GetCards() []uint32 {
	if m != nil {
		return m.Cards
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgTexasGameInfoFreeRep)(nil), "net.msg_texas_game_info_free_rep")
	proto.RegisterType((*MsgTexasGameInfoPlayRep)(nil), "net.msg_texas_game_info_play_rep")
	proto.RegisterType((*MsgTexasStartRep)(nil), "net.msg_texas_start_rep")
	proto.RegisterType((*MsgTexasAddscoreReq)(nil), "net.msg_texas_addscore_req")
	proto.RegisterType((*MsgTexasAddscoreRep)(nil), "net.msg_texas_addscore_rep")
	proto.RegisterType((*MsgTexasGiveupReq)(nil), "net.msg_texas_giveup_req")
	proto.RegisterType((*MsgTexasGiveupRep)(nil), "net.msg_texas_giveup_rep")
	proto.RegisterType((*MsgTexasSendcardRep)(nil), "net.msg_texas_sendcard_rep")
	proto.RegisterType((*TexasWinChair)(nil), "net.texas_win_chair")
	proto.RegisterType((*MsgTexasGameEndRep)(nil), "net.msg_texas_game_end_rep")
	proto.RegisterType((*MsgTexasBuyinReq)(nil), "net.msg_texas_buyin_req")
	proto.RegisterType((*MsgTexasBuyinRep)(nil), "net.msg_texas_buyin_rep")
	proto.RegisterType((*MsgTexasShowCardReq)(nil), "net.msg_texas_show_card_req")
	proto.RegisterType((*MsgTexasShowCardRep)(nil), "net.msg_texas_show_card_rep")
}

func init() { proto.RegisterFile("texas_logic_msg.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 660 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x54, 0x41, 0x6e, 0xdb, 0x3a,
	0x10, 0x85, 0xa2, 0x38, 0xf9, 0x9e, 0x44, 0x76, 0xac, 0xd8, 0x89, 0x10, 0xfc, 0x85, 0x21, 0x64,
	0x61, 0xa0, 0x6d, 0x16, 0x45, 0xd7, 0xed, 0xa6, 0xcb, 0xde, 0x81, 0xa0, 0xa5, 0x49, 0x2c, 0x80,
	0x26, 0x55, 0x92, 0x6a, 0x92, 0x8b, 0xf4, 0x2e, 0x05, 0x7a, 0xa6, 0x9e, 0xa1, 0xe0, 0x50, 0x92,
	0x69, 0xc7, 0xea, 0x52, 0x8f, 0x8f, 0x8f, 0x33, 0xef, 0xcd, 0x08, 0x16, 0x16, 0x5f, 0xb8, 0x61,
	0x42, 0x3d, 0x55, 0x05, 0xdb, 0x9a, 0xa7, 0x87, 0x5a, 0x2b, 0xab, 0xd2, 0x58, 0xa2, 0xbd, 0x5b,
	0xac, 0xb9, 0x41, 0x07, 0x32, 0x63, 0x75, 0x53, 0x58, 0x7f, 0x76, 0x77, 0x85, 0x5a, 0x2b, 0xcd,
	0x0a, 0x55, 0xa2, 0x47, 0xf2, 0x6f, 0xf0, 0xbf, 0x63, 0x79, 0xa9, 0x27, 0xbe, 0x45, 0x56, 0xc9,
	0x47, 0xc5, 0x1e, 0x35, 0x22, 0xd3, 0x58, 0xa7, 0x0b, 0x48, 0x0a, 0x14, 0x82, 0x6d, 0x2b, 0x69,
	0x0a, 0xa5, 0x31, 0x8b, 0x96, 0xd1, 0x2a, 0xde, 0xc1, 0xfc, 0xc5, 0xc3, 0x27, 0x0e, 0xce, 0x7f,
	0xc6, 0xc7, 0xe5, 0x6a, 0xc1, 0x5f, 0x49, 0x2e, 0x05, 0xa0, 0x7b, 0xa1, 0xd6, 0x0d, 0x4c, 0x6c,
	0xa3, 0xa5, 0xd3, 0x62, 0x81, 0x58, 0x7a, 0x0b, 0x53, 0xc2, 0x05, 0x1a, 0xd3, 0x1e, 0xc4, 0xdd,
	0x85, 0xee, 0xf1, 0x16, 0x3f, 0xed, 0x70, 0x5e, 0x96, 0x21, 0x7f, 0x44, 0xf8, 0x1c, 0x2e, 0x2d,
	0x5f, 0x0b, 0xf4, 0xa0, 0xc9, 0xce, 0x96, 0x71, 0x8b, 0x2a, 0xcb, 0x45, 0x87, 0x9e, 0x77, 0x68,
	0x81, 0xd2, 0xa2, 0x6e, 0x15, 0xfe, 0x23, 0x85, 0x09, 0x9c, 0x95, 0xac, 0x31, 0xa8, 0xb3, 0xf1,
	0x32, 0x5a, 0x25, 0xc4, 0x6a, 0xb4, 0x46, 0x69, 0x3d, 0x0a, 0x84, 0x5e, 0xc3, 0x05, 0x35, 0x6a,
	0x2c, 0xb7, 0x8d, 0xc9, 0x2e, 0x96, 0xf1, 0x2a, 0x71, 0x4e, 0xad, 0xb9, 0xe0, 0xb2, 0x40, 0x56,
	0xa8, 0x46, 0xda, 0xec, 0xb2, 0x57, 0xf0, 0xef, 0x14, 0x5c, 0x97, 0x26, 0x4b, 0x88, 0x9c, 0x02,
	0x6c, 0xb8, 0x2c, 0x5b, 0x6c, 0x42, 0xd8, 0x0c, 0xc6, 0xaa, 0x46, 0xcd, 0x6c, 0xb5, 0xc5, 0x6c,
	0x4a, 0x97, 0x6f, 0x61, 0xba, 0xad, 0x24, 0x2b, 0x36, 0x55, 0x5d, 0x49, 0x5f, 0xc1, 0x55, 0x7f,
	0xc0, 0x5f, 0xf6, 0x0e, 0x66, 0xee, 0x20, 0xff, 0x13, 0xc1, 0xf5, 0x2e, 0x18, 0x63, 0xb9, 0xb6,
	0x94, 0xc7, 0x61, 0x23, 0x11, 0xc9, 0xec, 0xda, 0x3d, 0x19, 0x92, 0x8d, 0xe9, 0x60, 0x3f, 0xce,
	0xd3, 0x81, 0x38, 0x47, 0x43, 0x71, 0x9e, 0x0d, 0xc4, 0x76, 0x4e, 0x78, 0x0e, 0xe0, 0x7c, 0x60,
	0x25, 0xb7, 0xdc, 0x50, 0x10, 0x17, 0x1f, 0x27, 0x0f, 0x12, 0xed, 0x83, 0xeb, 0x84, 0x2c, 0x3a,
	0xe6, 0x04, 0x25, 0x94, 0xbf, 0x83, 0x9b, 0x5d, 0xbf, 0xbc, 0x2c, 0x49, 0x98, 0x69, 0xfc, 0xee,
	0xfc, 0x74, 0xcf, 0x05, 0x13, 0x98, 0xff, 0x8a, 0x06, 0xd8, 0x43, 0x06, 0x2d, 0x20, 0xe9, 0x59,
	0x81, 0x4f, 0xbe, 0x13, 0x0f, 0xfb, 0xb0, 0xe3, 0xa1, 0xd6, 0xff, 0xed, 0xd5, 0x5b, 0x4b, 0xbc,
	0x55, 0x33, 0x18, 0x9b, 0x8d, 0x7a, 0x76, 0x33, 0xe2, 0x07, 0x36, 0xc9, 0x3f, 0xc0, 0x3c, 0xd8,
	0xb8, 0xea, 0x07, 0x36, 0x35, 0xb5, 0xb9, 0x80, 0xa4, 0xff, 0xe2, 0x46, 0x49, 0x5f, 0x79, 0xfe,
	0xe5, 0x28, 0xbd, 0x76, 0xb3, 0xdb, 0x7e, 0x05, 0x6d, 0xa6, 0x00, 0x42, 0x19, 0x1b, 0x6e, 0x65,
	0xae, 0x43, 0xab, 0x0c, 0xca, 0x92, 0x32, 0x6a, 0x25, 0xea, 0x66, 0x2d, 0xaa, 0x82, 0xb2, 0x69,
	0x25, 0x0e, 0xfd, 0xeb, 0x07, 0xca, 0x5d, 0x25, 0x62, 0xe0, 0xd4, 0xdb, 0xb5, 0x38, 0xa5, 0x1e,
	0xef, 0x61, 0xea, 0xdf, 0x7b, 0xa6, 0xac, 0x79, 0xa5, 0x9d, 0x13, 0xfd, 0x47, 0x16, 0x11, 0xeb,
	0xf7, 0x5e, 0x8a, 0xf4, 0xf3, 0x71, 0x6f, 0xb8, 0xd2, 0x66, 0x30, 0xf6, 0xbb, 0x8e, 0xb2, 0x2b,
	0xcc, 0x35, 0xec, 0x28, 0xed, 0xf6, 0x9f, 0xd0, 0xf6, 0xef, 0x8f, 0x5c, 0xbc, 0x8c, 0x8f, 0x8c,
	0xdc, 0x7b, 0x98, 0x0b, 0x6e, 0x2c, 0x0b, 0xea, 0xfc, 0xca, 0x2d, 0xa7, 0x52, 0xdf, 0xb2, 0x57,
	0x00, 0x7d, 0x9d, 0x26, 0x1b, 0x11, 0x67, 0x4e, 0x9c, 0x83, 0x8e, 0xf2, 0xfb, 0x70, 0x43, 0xd7,
	0xcd, 0x6b, 0x25, 0x29, 0xc7, 0x04, 0x46, 0xe1, 0xa8, 0x7e, 0x3a, 0xc6, 0xaa, 0x0f, 0x58, 0x6e,
	0x81, 0x35, 0x9a, 0x46, 0x58, 0xef, 0x77, 0xbe, 0x82, 0xdb, 0x20, 0xb4, 0x8d, 0x7a, 0x66, 0x6d,
	0x6a, 0xa4, 0xef, 0xad, 0xf6, 0x26, 0x7e, 0x1e, 0x62, 0xd2, 0x2a, 0x78, 0xc0, 0x95, 0x5b, 0x75,
	0x3e, 0xf6, 0xf7, 0x9d, 0x83, 0xc9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0x29, 0xf7, 0xde,
	0x95, 0x06, 0x00, 0x00,
}
