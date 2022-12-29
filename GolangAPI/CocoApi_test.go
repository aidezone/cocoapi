package coco

import (
	"testing"
	"fmt"
	"io/ioutil"
	// "coco/models"
)


var datasetMeta []byte
var err error

func TestMain(m *testing.M) {
    fmt.Println("begin")
    datasetMeta, err = ioutil.ReadFile("../anno/stuff_val2017.json")
	if err !=nil{
		fmt.Println("err:", err)
		return
	}
    m.Run()
    fmt.Println("end")
}


func Test_decoderExample(t *testing.T) {
	decoderExample(datasetMeta)
}
