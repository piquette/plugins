// Code generated by "stringer -type=EnumElementType datatypes.go byteconversions.go"; DO NOT EDIT

package io

import "fmt"

const _EnumElementType_name = "FLOAT32INT32FLOAT64INT64EPOCHBYTEBOOLNONESTRINGINT16UINT8UINT16UINT32UINT64"

var _EnumElementType_index = [...]uint8{0, 7, 12, 19, 24, 29, 33, 37, 41, 47, 52, 57, 63, 69, 75}

func (i EnumElementType) String() string {
	if i >= EnumElementType(len(_EnumElementType_index)-1) {
		return fmt.Sprintf("EnumElementType(%d)", i)
	}
	return _EnumElementType_name[_EnumElementType_index[i]:_EnumElementType_index[i+1]]
}
