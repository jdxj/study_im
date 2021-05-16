package protobuf

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/valyala/bytebufferpool"
)

var (
	ErrIDAlreadyExists      = errors.New("id already exists")
	ErrMsgAlreadyExists     = errors.New("msg already exists")
	ErrProtobufDataTooShort = errors.New("protobuf data too short")
)

func NewProcessor() *Processor {
	p := &Processor{
		littleEndian: false,
		bufferPool:   &bytebufferpool.Pool{},
		idMsg:        make(map[uint16]reflect.Type),
		msgID:        make(map[reflect.Type]uint16),
	}
	return p
}

// id | protobuf

type Processor struct {
	littleEndian bool
	bufferPool   *bytebufferpool.Pool

	idMsg map[uint16]reflect.Type
	msgID map[reflect.Type]uint16
}

func (p *Processor) SetByteOrder(littleEndian bool) {
	p.littleEndian = littleEndian
}

func (p *Processor) Register(id uint16, msg proto.Message) {
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

func (p *Processor) Unmarshal(data []byte) (uint16, interface{}, error) {
	if len(data) < 2 {
		return 0, nil, ErrProtobufDataTooShort
	}

	var id uint16
	if p.littleEndian {
		id = binary.LittleEndian.Uint16(data)
	} else {
		id = binary.BigEndian.Uint16(data)
	}

	msgType, ok := p.idMsg[id]
	if !ok {
		return 0, nil, fmt.Errorf("message id %d not registered", id)
	}

	msg := reflect.New(msgType.Elem()).Interface()
	return id, msg, proto.Unmarshal(data[2:], msg.(proto.Message))
}

// id | msg

func (p *Processor) Marshal(msg interface{}) ([]byte, error) {
	msgType := reflect.TypeOf(msg)
	id, ok := p.msgID[msgType]
	if !ok {
		return nil, fmt.Errorf("message %s not registered", msgType)
	}

	idBuf := make([]byte, 2)
	if p.littleEndian {
		binary.LittleEndian.PutUint16(idBuf, id)
	} else {
		binary.BigEndian.PutUint16(idBuf, id)
	}

	content, err := proto.Marshal(msg.(proto.Message))
	if err != nil {
		return nil, err
	}

	buffer := p.bufferPool.Get()
	defer p.bufferPool.Put(buffer)

	_, err = buffer.Write(idBuf)
	if err != nil {
		return nil, err
	}
	_, err = buffer.Write(content)
	return buffer.Bytes(), err
}
