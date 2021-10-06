// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package structure

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SzRectangle struct {
	_tab flatbuffers.Struct
}

func (rcv *SzRectangle) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SzRectangle) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *SzRectangle) TopLeft(obj *SzVector2) *SzVector2 {
	if obj == nil {
		obj = new(SzVector2)
	}
	obj.Init(rcv._tab.Bytes, rcv._tab.Pos+0)
	return obj
}
func (rcv *SzRectangle) BotRight(obj *SzVector2) *SzVector2 {
	if obj == nil {
		obj = new(SzVector2)
	}
	obj.Init(rcv._tab.Bytes, rcv._tab.Pos+8)
	return obj
}

func CreateSzRectangle(builder *flatbuffers.Builder, top_left_x int32, top_left_y int32, bot_right_x int32, bot_right_y int32) flatbuffers.UOffsetT {
	builder.Prep(4, 16)
	builder.Prep(4, 8)
	builder.PrependInt32(bot_right_y)
	builder.PrependInt32(bot_right_x)
	builder.Prep(4, 8)
	builder.PrependInt32(top_left_y)
	builder.PrependInt32(top_left_x)
	return builder.Offset()
}