package libmqtt

import "bytes"

type Packet interface {
	Type() CtrlType

	// Bytes dump a mqtt Packet object to mqtt bytes into provided buffer
	Bytes(*bytes.Buffer) error
}

type Topic struct {
	Name string
	Qos  QosLevel
}

const (
	maxMsgSize = 0xffffff7f
)

type CtrlType = byte

const (
	CtrlConn      CtrlType = iota + 1
	CtrlConnAck
	CtrlPublish
	CtrlPubAck
	CtrlPubRecv
	CtrlPubRel
	CtrlPubComp
	CtrlSubscribe
	CtrlSubAck
	CtrlUnSub
	CtrlUnSubAck
	CtrlPingReq
	CtrlPingResp
	CtrlDisConn
)

type QosLevel = byte

type ProtocolLevel = byte

const (
	V31  ProtocolLevel = iota + 3
	V311
)

const (
	Qos0 QosLevel = iota
	Qos1
	Qos2
)

var (
	mqtt = []byte{'M', 'Q', 'T', 'T'}
)

type ConnAckCode = byte

const (
	ConnAccepted          ConnAckCode = iota
	ConnBadProtocol
	ConnIdRejected
	ConnServerUnavailable
	ConnBadIdentity
	ConnAuthFail
	ConnNetErr            ConnAckCode = 0xf0
	ConnBadPacket         ConnAckCode = 0xf1
	ConnDialErr           ConnAckCode = 0xf2
)

type SubAckCode = byte

const (
	SubOkMaxQos0 SubAckCode = iota
	SubOkMaxQos1
	SubOkMaxQos2
	SubFail      = 0x80
)

var (
	disConnBytes  = []byte{CtrlDisConn << 4, 0x00}
	pingReqBytes  = []byte{CtrlPingReq << 4, 0x00}
	pingRespBytes = []byte{CtrlPingResp << 4, 0x00}
)