package ioutils_test

import (
	"bytes"
	"mime/multipart"
	"testing"

	"playground.com/server/pkg/ioutils"
)

type test struct {
	F1 string `json:"f1"`
	F2 string
	f3 string
	F4 []int `json:"f4"`
	F5 struct {
		Nested int `json:"nested"`
	} `json:"f5"`
}

func TestReadFormToStruct(t *testing.T) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	v1, v2, v3 := "test_1", "test_2", "test_3"
	_ = w.WriteField("f1", v1)
	_ = w.WriteField("F2", v2)
	_ = w.WriteField("f3", v3)
	_ = w.WriteField("f4", "1")
	_ = w.WriteField("f4", "2")
	_ = w.WriteField("f5[nested]", "5")
	w.Close()
	r := multipart.NewReader(&b, w.Boundary())
	form, _ := r.ReadForm(10000)

	var target test

	err := ioutils.ReadFormToStruct(form, &target)
	if err != nil {
		t.Fatal(err)
	}

	if target.F1 != v1 {
		t.Errorf("want: %s, got: %s", v1, target.F1)
	}

	if target.F2 != v2 {
		t.Errorf("failed to write to field by name")
	}

	if target.f3 != "" {
		t.Errorf("write to unexported field")
	}

	if len(target.F4) != 2 || target.F4[0] != 1 || target.F4[1] != 2 {
		t.Errorf("error while writing nested slice")
	}
}
