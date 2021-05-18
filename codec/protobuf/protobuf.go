package protobuf

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/valyala/bytebufferpool"
)

var (
	ErrIDAlreadyExists      = errors.New("id already exists")
	ErrMsgAlreadyExists     = errors.New("msg already exists")
	ErrProtobufDataTooShort = errors.New("protobuf data too short")
	ErrIncompatibleVersion  = errors.New("incompatible version")
)

// 应用层协议
//type RawData struct {
//	version   uint32
//	cmd       uint32
//	seq       uint32
//	timestamp uint32
//	bodyLen   uint32
//	body      []byte
//}

const (
	CurVersion = 1
)

type RawMsg struct {
	Version   uint32
	Cmd       uint32
	Seq       uint32
	Timestamp uint32
	Msg       interface{}
}

func NewProcessor() *Processor {
	p := &Processor{
		byteOrder:  binary.BigEndian,
		bufferPool: &bytebufferpool.Pool{},
		idMsg:      make(map[uint32]reflect.Type),
		msgID:      make(map[reflect.Type]uint32),
	}
	return p
}

type Processor struct {
	byteOrder  binary.ByteOrder
	bufferPool *bytebufferpool.Pool

	idMsg map[uint32]reflect.Type
	msgID map[reflect.Type]uint32
}

func (p *Processor) SetByteOrder(byteOrder binary.ByteOrder) {
	p.byteOrder = byteOrder
}

func (p *Processor) Register(id uint32, msg proto.Message) {
	if _, ok := p.idMsg[id]; ok {
		log.Fatalln(ErrIDAlreadyExists)
	}

	msgType := reflect.TypeOf(msg)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		log.Fatalln("protobuf message pointer required")
	}
	if _, ok := p.msgID[msgType]; ok {
		log.Fatalln(ErrMsgAlreadyExists)
	}

	p.idMsg[id] = msgType
	p.msgID[msgType] = id
}

func (p *Processor) Unmarshal(data []byte) (*RawMsg, error) {
	if len(data) < 20 {
		return nil, ErrProtobufDataTooShort
	}

	byteOrder := p.byteOrder
	version := byteOrder.Uint32(data[0:4])
	cmd := byteOrder.Uint32(data[4:8])
	seq := byteOrder.Uint32(data[8:12])
	timestamp := byteOrder.Uint32(data[12:16])
	bodyLen := byteOrder.Uint32(data[16:20])

	if version != CurVersion {
		return nil, ErrIncompatibleVersion
	}

	msgType, ok := p.idMsg[cmd]
	if !ok {
		return nil, fmt.Errorf("message id %d not registered", cmd)
	}

	if len(data[20:]) != int(bodyLen) {
		return nil, fmt.Errorf("invalid message length")
	}

	msg := reflect.New(msgType.Elem()).Interface()
	err := proto.Unmarshal(data[20:], msg.(proto.Message))
	if err != nil {
		return nil, err
	}

	rawMsg := &RawMsg{
		Version:   version,
		Cmd:       cmd,
		Seq:       seq,
		Timestamp: timestamp,
		Msg:       msg,
	}
	return rawMsg, nil
}

func (p *Processor) Marshal(seq uint32, msg interface{}) ([]byte, error) {
	msgType := reflect.TypeOf(msg)
	cmd, ok := p.msgID[msgType]
	if !ok {
		return nil, fmt.Errorf("message %s not registered", msgType)
	}

	buffer := p.bufferPool.Get()
	defer p.bufferPool.Put(buffer)

	content, err := proto.Marshal(msg.(proto.Message))
	if err != nil {
		return nil, err
	}

	fieldBuf := make([]byte, 20)
	byteOrder := p.byteOrder
	byteOrder.PutUint32(fieldBuf[0:4], CurVersion)
	byteOrder.PutUint32(fieldBuf[4:8], cmd)
	byteOrder.PutUint32(fieldBuf[8:12], seq)
	byteOrder.PutUint32(fieldBuf[12:16], uint32(time.Now().Unix()))
	byteOrder.PutUint32(fieldBuf[16:20], uint32(len(content)))

	_, err1 := buffer.Write(fieldBuf)
	_, err2 := buffer.Write(content)
	if err1 != nil || err2 != nil {
		return nil, fmt.Errorf("write data, err1: %s, err2: %s", err1, err2)
	}
	return buffer.Bytes(), nil
}
