package utils_test

import (
	"testing"

	"github.com/anton7r/pgxe/internal/utils"
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

	prepped := utils.PrepReflect(&testStruct{Field2: 123})

	_, err := utils.GetNamedField(prepped, "Field22")

	if err != nil {
		if err.Error() != "field 'Field22' not found" {
			t.Error("Function errored:" + err.Error())
		}
	}
}

func TestIntConversion(t *testing.T) {
	prepped := utils.PrepReflect(&testStruct{Field2: 123})

	ret, err := utils.GetNamedField(prepped, "Field2")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "123" {
		t.Error("WRONG")
	}
}

func TestInt8Conversion(t *testing.T) {
	prep := utils.PrepReflect(&testStruct{Field4: 123})

	ret, err := utils.GetNamedField(prep, "Field4")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "123" {
		t.Error("WRONG")
	}
}

func TestInt16Conversion(t *testing.T) {
	pr := utils.PrepReflect(&testStruct{Field5: 123})

	ret, err := utils.GetNamedField(pr, "Field5")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "123" {
		t.Error("WRONG")
	}
}

func TestInt32Conversion(t *testing.T) {
	pr := utils.PrepReflect(&testStruct{Field6: 123})

	ret, err := utils.GetNamedField(pr, "Field6")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "123" {
		t.Error("WRONG")
	}
}

func TestInt64Conversion(t *testing.T) {
	pr := utils.PrepReflect(&testStruct{Field7: 123})

	ret, err := utils.GetNamedField(pr, "Field7")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "123" {
		t.Error("WRONG")
	}
}

func TestUintntConversion(t *testing.T) {
	pr := utils.PrepReflect(&testStruct{Field12: 123})

	ret, err := utils.GetNamedField(pr, "Field12")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "123" {
		t.Error("WRONG")
	}
}

func TestUint8Conversion(t *testing.T) {
	pr := utils.PrepReflect(&testStruct{Field14: 123})

	ret, err := utils.GetNamedField(pr, "Field14")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "123" {
		t.Error("WRONG")
	}
}

func TestUint16Conversion(t *testing.T) {
	pr := utils.PrepReflect(&testStruct{Field15: 123})

	ret, err := utils.GetNamedField(pr, "Field15")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "123" {
		t.Error("WRONG")
	}
}

func TestUint32Conversion(t *testing.T) {
	pr := utils.PrepReflect(&testStruct{Field16: 123})

	ret, err := utils.GetNamedField(pr, "Field16")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "123" {
		t.Error("WRONG")
	}
}

func TestUint64Conversion(t *testing.T) {
	pr := utils.PrepReflect(&testStruct{Field17: 123})

	ret, err := utils.GetNamedField(pr, "Field17")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "123" {
		t.Error("WRONG")
	}
}

func TestStringField(t *testing.T) {
	pr := utils.PrepReflect(&testStruct{Field: "123"})

	ret, err := utils.GetNamedField(pr, "Field")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "'123'" {
		t.Error("THE OUTPUT STRING IS WRONG: " + ret)
	}
}

func TestBoolFalseField(t *testing.T) {
	pr := utils.PrepReflect(&testStruct{Field3: false})

	ret, err := utils.GetNamedField(pr, "Field3")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "FALSE" {
		t.Error("WRONG")
	}
}

func TestBoolTrueField(t *testing.T) {
	pr := utils.PrepReflect(&testStruct{Field3: true})

	ret, err := utils.GetNamedField(pr, "Field3")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "TRUE" {
		t.Error("WRONG")
	}
}

func TestFloat32(t *testing.T) {
	pr := utils.PrepReflect(&testStruct{Field8: 1.98})

	ret, err := utils.GetNamedField(pr, "Field8")

	if err != nil {
		t.Error("Function errored:" + err.Error())
		return
	}

	if ret != "1.980000" {
		t.Error("WRONG expected 1.980000 but got " + ret)
	}
}

func TestFloat64(t *testing.T) {
	prep := utils.PrepReflect(&testStruct{Field9: 1.98})

	ret, err := utils.GetNamedField(prep, "Field9")

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

	prep := utils.PrepReflect(&testStruct{FieldUnsupported: &c})

	_, err := utils.GetNamedField(prep, "FieldUnsupported")

	if err == nil {
		t.Error("pointer was not errored")
		return
	}
}

func TestUnsupportedType2(t *testing.T) {
	var c []int = []int{123, 123}

	prepped := utils.PrepReflect(&testStruct{FieldUnsupported2: c})

	_, err := utils.GetNamedField(prepped, "FieldUnsupported2")

	if err == nil {
		t.Error("pointer was not errored")
		return
	}
}

func TestWrongType(t *testing.T) {
	str := "string"
	prep := utils.PrepReflect(&str)
	_, err := utils.GetNamedField(prep, "string")
	if err == nil {
		t.Error("Did not error with wrong type")
	}
}

func BenchmarkGetNamedField(b *testing.B) {
	ts := utils.PrepReflect(&testStruct{Field: "Field"})

	for i := 0; i < b.N; i++ {
		_, err := utils.GetNamedField(ts, "Field")
		if err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkGetNamedField_2(b *testing.B) {
	ts := utils.PrepReflect(&testStruct2{Field: "Field"})

	for i := 0; i < b.N; i++ {
		_, err := utils.GetNamedField(ts, "Field")
		if err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkGetNamedField_3(b *testing.B) {
	ts := utils.PrepReflect(&testStruct{Field: "Field"})

	for i := 0; i < b.N; i++ {
		_, err := utils.GetNamedField(ts, "Field")
		if err != nil {
			b.Error(err.Error())
		}

		_, err = utils.GetNamedField(ts, "Field2")
		if err != nil {
			b.Error(err.Error())
		}

		_, err = utils.GetNamedField(ts, "Field3")
		if err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkGetNamedField_4(b *testing.B) {
	ts := utils.PrepReflect(&testStruct{Field: "Field"})

	for i := 0; i < b.N; i++ {
		for x := 0; x < 5; x++ {
			_, err := utils.GetNamedField(ts, "Field")
			if err != nil {
				b.Error(err.Error())
			}

			_, err = utils.GetNamedField(ts, "Field2")
			if err != nil {
				b.Error(err.Error())
			}

			_, err = utils.GetNamedField(ts, "Field3")
			if err != nil {
				b.Error(err.Error())
			}
		}
	}
}