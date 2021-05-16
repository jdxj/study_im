package frame

import (
	"encoding/binary"

	"github.com/panjf2000/gnet"
)

func NewLengthFieldBasedFrameCodec() *gnet.LengthFieldBasedFrameCodec {
	ec := gnet.EncoderConfig{
		ByteOrder:                       binary.BigEndian,
		LengthFieldLength:               4,
		LengthAdjustment:                0,
		LengthIncludesLengthFieldLength: false,
	}
	dc := gnet.DecoderConfig{
		ByteOrder:           binary.BigEndian,
		LengthFieldOffset:   0,
		LengthFieldLength:   4,
		LengthAdjustment:    0,
		InitialBytesToStrip: 4,
	}
	return gnet.NewLengthFieldBasedFrameCodec(ec, dc)
}
