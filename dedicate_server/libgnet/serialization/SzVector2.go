// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package structure

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SzVector2 struct {
	_tab flatbuffers.Struct
}

func (rcv *SzVector2) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SzVector2) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *SzVector2) X() int32 {
	return rcv._tab.GetInt32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *SzVector2) MutateX(n int32) bool {
	return rcv._tab.MutateInt32(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *SzVector2) Y() int32 {
	return rcv._tab.GetInt32(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}
func (rcv *SzVector2) MutateY(n int32) bool {
	return rcv._tab.MutateInt32(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

func CreateSzVector2(builder *flatbuffers.Builder, x int32, y int32) flatbuffers.UOffsetT {
	builder.Prep(4, 8)
	builder.PrependInt32(y)
	builder.PrependInt32(x)
	return builder.Offset()
}
