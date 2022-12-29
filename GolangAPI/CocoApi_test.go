package coco

import (
	"testing"
	"fmt"
	"io/ioutil"
	// "coco/models"
	// "encoding/json"
)


var datasetMeta []byte
var datasetMetaObj *CocoApi
var err error

func TestMain(m *testing.M) {


    fmt.Println("begin")
    datasetMeta, err = ioutil.ReadFile("../anno/stuff_val2017.json")
	if err !=nil{
		fmt.Println("err:", err)
		return
	}

	// meta := &models.ObjectDetection{}
	// err = json.Unmarshal(datasetMeta, meta)
	// if err != nil {
	// 	fmt.Println("json.unmarshal failed,err:",err)
	// 	return
	// }
	datasetMetaObj = NewCocoApi(datasetMeta);
    m.Run()
    fmt.Println("end")
}

func Test_GetAnnIds(t *testing.T) {
	anns, _ := datasetMetaObj.GetAnnIds(nil, nil, nil, 0)
    fmt.Println("anns len: ", len(anns))
}


func Test_decoderExample(t *testing.T) {
	decoderExample(datasetMeta)
}
