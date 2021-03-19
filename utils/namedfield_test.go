package utils_test

import (
	"github.com/anton7r/pgxe/utils"
	"testing"
)

type testStruct struct {
	Field             string
	Field2            int
	Field3            bool
	Field4            int8
	Field5            int16
	Field6            int32
	Field7            int64
	Field8            float32
	Field9            float64
	Field12           uint
	Field14           uint8
	Field15           uint16
	Field16           uint32
	Field17           uint64
	FieldUnsupported  *int
	FieldUnsupported2 []int
}

type testStruct2 struct {
	Field string
}

func TestMissingField(t *testing.T) {

	_, err := utils.GetNamedField(&testStruct{Field2: 123}, "Field22")

	if err != nil {
		if err.Error() != "field 'field22' not found" {
			t.Error("Function errored:" + err.Error())
		}
	}
}

func TestIntConversion(t *testing.T) {
	ret, err := utils.GetNamedField(&testStruct{Field2: 123}, "Field2")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "123" {
		t.Error("WRONG")
	}
}

func TestInt8Conversion(t *testing.T) {
	ret, err := utils.GetNamedField(&testStruct{Field4: 123}, "Field4")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "123" {
		t.Error("WRONG")
	}
}

func TestInt16Conversion(t *testing.T) {
	ret, err := utils.GetNamedField(&testStruct{Field5: 123}, "Field5")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "123" {
		t.Error("WRONG")
	}
}

func TestInt32Conversion(t *testing.T) {
	ret, err := utils.GetNamedField(&testStruct{Field6: 123}, "Field6")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "123" {
		t.Error("WRONG")
	}
}

func TestInt64Conversion(t *testing.T) {
	ret, err := utils.GetNamedField(&testStruct{Field7: 123}, "Field7")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "123" {
		t.Error("WRONG")
	}
}

func TestUintntConversion(t *testing.T) {
	ret, err := utils.GetNamedField(&testStruct{Field12: 123}, "Field12")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "123" {
		t.Error("WRONG")
	}
}

func TestUint8Conversion(t *testing.T) {
	ret, err := utils.GetNamedField(&testStruct{Field14: 123}, "Field14")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "123" {
		t.Error("WRONG")
	}
}

func TestUint16Conversion(t *testing.T) {
	ret, err := utils.GetNamedField(&testStruct{Field15: 123}, "Field15")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "123" {
		t.Error("WRONG")
	}
}

func TestUint32Conversion(t *testing.T) {
	ret, err := utils.GetNamedField(&testStruct{Field16: 123}, "Field16")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "123" {
		t.Error("WRONG")
	}
}

func TestUint64Conversion(t *testing.T) {
	ret, err := utils.GetNamedField(&testStruct{Field17: 123}, "Field17")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "123" {
		t.Error("WRONG")
	}
}

func TestStringField(t *testing.T) {
	ret, err := utils.GetNamedField(&testStruct{Field: "123"}, "Field")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "'123'" {
		t.Error("THE OUTPUT STRING IS WRONG: " + ret)
	}
}

func TestBoolFalseField(t *testing.T) {
	ret, err := utils.GetNamedField(&testStruct{Field3: false}, "Field3")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "FALSE" {
		t.Error("WRONG")
	}
}

func TestBoolTrueField(t *testing.T) {
	ret, err := utils.GetNamedField(&testStruct{Field3: true}, "Field3")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "TRUE" {
		t.Error("WRONG")
	}
}

func TestFloat32(t *testing.T) {
	ret, err := utils.GetNamedField(&testStruct{Field8: 1.98}, "Field8")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "1.980000" {
		t.Error("WRONG expected 1.980000 but got " + ret)
	}
}

func TestFloat64(t *testing.T) {
	ret, err := utils.GetNamedField(&testStruct{Field9: 1.98}, "Field9")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "1.980000" {
		t.Error("WRONG expected 1.980000 but got " + ret)
	}
}

func TestUnsupportedType(t *testing.T) {
	c := 103

	_, err := utils.GetNamedField(&testStruct{FieldUnsupported: &c}, "FieldUnsupported")

	if err == nil {
		t.Error("pointer was not errored")
		return
	}
}

func TestUnsupportedType2(t *testing.T) {
	var c []int = []int{123, 123}

	_, err := utils.GetNamedField(&testStruct{FieldUnsupported2: c}, "FieldUnsupported2")

	if err == nil {
		t.Error("pointer was not errored")
		return
	}
}

func BenchmarkGetNamedField(b *testing.B) {
	ts := &testStruct{Field:"Field"}

	for i := 0; i < b.N; i++ {
		_, err := utils.GetNamedField(ts, "Field")
		if err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkGetNamedField2(b *testing.B) {
	ts := &testStruct2{Field:"Field"}

	for i := 0; i < b.N; i++ {
		_, err := utils.GetNamedField(ts, "Field")
		if err != nil {
			b.Error(err.Error())
		}
	}
}