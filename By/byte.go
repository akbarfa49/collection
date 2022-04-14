package By

import (
	"reflect"
	"strconv"
	"unsafe"
)

//convert Byte To String Unsafe way. dont change the byte if string still survive
func ToSUnsafe(b []byte) (s string) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	sh.Len = bh.Len
	sh.Data = bh.Data
	return
}

//ToS copy to string
func ToS(b []byte) string {
	return string(b)
}

//string number compatible
func ToInt64(b []byte) (i int64) {
	//strconv copy the value so safe to change byte after this
	i, _ = strconv.ParseInt(ToSUnsafe(b), 10, 64)
	return
}

func ToFloat64(b []byte) (i float64) {
	//strconv copy the value so safe to change byte after this
	i, _ = strconv.ParseFloat(ToSUnsafe(b), 64)
	return
}

//string number compatible
func ToU(b []byte) (u uint64) {
	//strconv copy the value so safe to change byte after this
	u, _ = strconv.ParseUint(ToSUnsafe(b), 10, 64)
	return
}
