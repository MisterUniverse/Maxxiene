package fileReader

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	f := "../nasa_test.json"
	_, err := ReadFile(f)
	if err != nil {
		t.Errorf("ReadFile(%v) = %T; want ", f, err)
	}

}

func TestConvert(t *testing.T) {
	f := "../nasa_test.json"
	buff, er := ReadFile(f)
	if er != nil {
		t.Errorf("faild to read file")
	}
	val, err := ConvertFileBufferToStr(buff)
	if err != nil {
		t.Errorf("ConvertFileBufferToStr(File(f)) falied")
	}
	fmt.Printf("%v", val)
}
