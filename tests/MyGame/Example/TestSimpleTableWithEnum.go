// automatically generated by the FlatBuffers compiler, do not modify

package Example

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type TestSimpleTableWithEnum struct {
	_tab flatbuffers.Table
}

func GetRootAsTestSimpleTableWithEnum(buf []byte, offset flatbuffers.UOffsetT) *TestSimpleTableWithEnum {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TestSimpleTableWithEnum{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *TestSimpleTableWithEnum) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TestSimpleTableWithEnum) Color() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 2
}

func (rcv *TestSimpleTableWithEnum) MutateColor(n int8) bool {
	return rcv._tab.MutateInt8Slot(4, n)
}

func TestSimpleTableWithEnumStart(builder *flatbuffers.Builder) { builder.StartObject(1) }
func TestSimpleTableWithEnumAddColor(builder *flatbuffers.Builder, color int8) { builder.PrependInt8Slot(0, color, 2) }
func TestSimpleTableWithEnumEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
