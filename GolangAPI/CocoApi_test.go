package coco

import (
	"testing"
	"fmt"
	"io/ioutil"
	// "coco/models"
	"encoding/json"
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

func Test_EncodeMaskToSegment(t *testing.T) {
	size := [2]uint32{5, 6} // segmentation.size
	originMask := []byte{0,0,0,0,0,1,1,1,1,1,1,0,0,0,0,1,1,0,1,1,0,0,0,0,0,1,1,0,1,1}
	seg := EncodeMaskToSegment(originMask, size)
    fmt.Println("seg: ", seg)
}

func Test_DecodeSegmentToMask(t *testing.T) {
	size := [2]uint32{5, 6} // segmentation.size
	originMask := []byte{0,0,0,0,0,1,1,1,1,1,1,0,0,0,0,1,1,0,1,1,0,0,0,0,0,1,1,0,1,1}
	seg := EncodeMaskToSegment(originMask, size)
    mask := DecodeSegmentToMask(seg)
    fmt.Println("mask: ", mask)
}

func Test_compressRLE(t *testing.T) {
	cnts := []uint32{70375,8,415,12,411,15,409,17,407,19,405,20,405,21,404,21,404,21,402,23,400,26,398,29,395,31,393,33,392,34,390,41,384,43,382,44,381,45,380,46,379,46,379,47,378,47,378,47,378,47,378,48,377,47,378,49,375,50,375,50,375,50,375,50,375,46,379,43,382,41,385,38,388,18,2,8,1,7,391,4,6,4,6,4,6,3,69652,6,418,8,416,10,414,12,413,12,413,12,413,13,412,13,412,13,412,13,412,13,412,13,412,13,412,12,413,12,414,10,416,8,419,4,844,7,417,9,415,11,413,14,411,15,410,16,409,17,408,17,408,17,408,19,407,19,407,19,407,19,407,18,409,17,411,14,412,13,412,13,412,14,411,15,6,4,400,15,4,6,400,15,3,8,400,14,2,9,401,13,1,10,402,12,1,10,402,23,403,10,1,11,404,8,2,22,395,4,4,39,386,41,384,41,383,42,383,42,383,43,381,44,381,44,381,44,381,44,381,45,380,45,380,45,380,45,380,46,379,46,380,45,381,44,381,45,381,44,382,43,383,42,384,41,386,40,406,18,408,18,407,18,407,18,407,19,407,18,407,18,407,19,407,18,408,17,408,17,409,17,410,15,412,13,372,6,37,9,372,8,416,10,414,12,413,12,413,13,412,14,411,15,410,15,410,15,410,15,410,15,410,15,410,15,410,15,409,21,404,22,402,24,401,25,400,25,400,25,400,25,400,25,390,35,389,36,388,36,388,36,387,37,387,38,386,39,385,39,386,38,387,36,389,37,388,38,386,39,386,39,386,39,386,39,386,39,386,39,386,39,386,39,386,39,387,38,388,23,2,12,389,18,6,12,391,12,10,12,80,3,330,12,79,16,318,12,79,28,306,11,80,28,306,11,80,28,306,11,80,29,305,11,80,29,305,11,80,29,306,10,81,29,305,10,81,29,306,9,81,29,307,7,82,30,308,4,83,30,396,29,397,29,397,28,397,28,398,28,397,28,397,29,396,29,396,29,396,30,395,30,395,30,395,30,395,30,395,30,396,29,396,29,397,28,398,27,399,26,400,24,402,23,402,23,402,23,403,21,405,20,406,19,407,18,408,16,410,14,412,12,415,8,31756,20,403,27,397,30,394,32,392,34,391,35,389,36,389,36,389,36,389,36,389,37,388,37,388,37,388,37,388,37,389,36,390,36,389,36,390,35,391,34,391,34,391,34,391,35,149}
	cntsRle := compressRLE(cnts,425,640)
	cntsRleDecodedMask := cntsRle.Decode()
	fmt.Println("cntsRle.Decode() => cntsRleDecodedMask", 425*640, len(cntsRleDecodedMask))

}

func Test_EncodeRLEToSegment(t *testing.T) {
	cnts := []uint32{70375,8,415,12,411,15,409,17,407,19,405,20,405,21,404,21,404,21,402,23,400,26,398,29,395,31,393,33,392,34,390,41,384,43,382,44,381,45,380,46,379,46,379,47,378,47,378,47,378,47,378,48,377,47,378,49,375,50,375,50,375,50,375,50,375,46,379,43,382,41,385,38,388,18,2,8,1,7,391,4,6,4,6,4,6,3,69652,6,418,8,416,10,414,12,413,12,413,12,413,13,412,13,412,13,412,13,412,13,412,13,412,13,412,12,413,12,414,10,416,8,419,4,844,7,417,9,415,11,413,14,411,15,410,16,409,17,408,17,408,17,408,19,407,19,407,19,407,19,407,18,409,17,411,14,412,13,412,13,412,14,411,15,6,4,400,15,4,6,400,15,3,8,400,14,2,9,401,13,1,10,402,12,1,10,402,23,403,10,1,11,404,8,2,22,395,4,4,39,386,41,384,41,383,42,383,42,383,43,381,44,381,44,381,44,381,44,381,45,380,45,380,45,380,45,380,46,379,46,380,45,381,44,381,45,381,44,382,43,383,42,384,41,386,40,406,18,408,18,407,18,407,18,407,19,407,18,407,18,407,19,407,18,408,17,408,17,409,17,410,15,412,13,372,6,37,9,372,8,416,10,414,12,413,12,413,13,412,14,411,15,410,15,410,15,410,15,410,15,410,15,410,15,410,15,409,21,404,22,402,24,401,25,400,25,400,25,400,25,400,25,390,35,389,36,388,36,388,36,387,37,387,38,386,39,385,39,386,38,387,36,389,37,388,38,386,39,386,39,386,39,386,39,386,39,386,39,386,39,386,39,386,39,387,38,388,23,2,12,389,18,6,12,391,12,10,12,80,3,330,12,79,16,318,12,79,28,306,11,80,28,306,11,80,28,306,11,80,29,305,11,80,29,305,11,80,29,306,10,81,29,305,10,81,29,306,9,81,29,307,7,82,30,308,4,83,30,396,29,397,29,397,28,397,28,398,28,397,28,397,29,396,29,396,29,396,30,395,30,395,30,395,30,395,30,395,30,396,29,396,29,397,28,398,27,399,26,400,24,402,23,402,23,402,23,403,21,405,20,406,19,407,18,408,16,410,14,412,12,415,8,31756,20,403,27,397,30,394,32,392,34,391,35,389,36,389,36,389,36,389,36,389,37,388,37,388,37,388,37,388,37,389,36,390,36,389,36,390,35,391,34,391,34,391,34,391,35,149}
	segRLEUnc := &SegmentationRLEUncompressed{
		Counts: cnts,
		Size: [2]uint32{425,640},
	}
	seg := EncodeRLEToSegment(segRLEUnc)
    fmt.Println("seg: ", seg)
}

func Test_encodeRLE(t *testing.T) {
	size := [2]uint32{5, 6} // segmentation.size
	originMask := []byte{0,0,0,0,0,1,1,1,1,1,1,0,0,0,0,1,1,0,1,1,0,0,0,0,0,1,1,0,1,1}
	fmt.Println("originMask", originMask)

	rle := encodeRLE(originMask, size[0], size[1], 1)
	fmt.Println("encodeRLE => rle", rle)

	decodedMask := rle.Decode()
	fmt.Println("rle.Decode() => decodedMask", decodedMask)
}

func Test_GetAnnIds(t *testing.T) {
	// without filter
	var anns []int

	anns = datasetMetaObj.GetAnnIds(nil, nil, nil, 0)
    fmt.Println("anns len: ", len(anns))

    // filter by imgIds
	anns = datasetMetaObj.GetAnnIds([]int{397133, 87038, 6818}, nil, nil, 0)
    fmt.Println("anns len: ", len(anns), anns)

    anns = datasetMetaObj.GetAnnIds([]int{397133, 87038, 6818}, []int{112, 123}, nil, 0)
    fmt.Println("anns len: ", len(anns), anns)

    anns = datasetMetaObj.GetAnnIds([]int{397133, 87038, 6818}, []int{112, 123}, []int{17008, 17010}, 0)
    fmt.Println("anns len: ", len(anns), anns)
}

func Test_GetCatIds(t *testing.T) {
	// without filter
	var resultIds []int

	var names = []string{
		"banner",
		"branch",
		"cabinet",
		"ceiling-other",
	}

	var superNames = []string{
		"building",
		"furniture-stuff",
		"ceiling",
	}

	resultIds = datasetMetaObj.GetCatIds(nil, nil)
    fmt.Println("GetCatIds resultIds len: ", len(resultIds))

    // filter by Name
	resultIds = datasetMetaObj.GetCatIds(names, nil)
    fmt.Println("GetCatIds resultIds len: ", len(resultIds), resultIds)

    // filter by SupName
	resultIds = datasetMetaObj.GetCatIds(nil, superNames)
    fmt.Println("GetCatIds resultIds len: ", len(resultIds), resultIds)

	// filter by Both
	resultIds = datasetMetaObj.GetCatIds(names, superNames)
    fmt.Println("GetCatIds resultIds len: ", len(resultIds), resultIds)
}

func Test_GetImgIds(t *testing.T) {
	// without filter
	var resultIds []int
	resultIds = datasetMetaObj.GetImgIds(nil)
    fmt.Println("GetImgIds resultIds len: ", len(resultIds))

    // filter by catids
	resultIds = datasetMetaObj.GetImgIds([]int{98, 102})
    fmt.Println("GetImgIds resultIds len: ", len(resultIds), resultIds)
}

func Test_LoadAnns(t *testing.T) {
	ids := datasetMetaObj.GetAnnIds([]int{397133, 87038, 6818}, []int{112, 123}, nil, 0)
    fmt.Println("LoadAnns ids len: ", len(ids), ids)

    // filter by catids
	results := datasetMetaObj.LoadAnns(ids)
    fmt.Println("LoadAnns result: ", len(results), results)
}

func Test_LoadCats(t *testing.T) {

	var names = []string{
		"banner",
		"branch",
		"cabinet",
		"ceiling-other",
	}

	var superNames = []string{
		"building",
		"furniture-stuff",
		"ceiling",
	}

	// filter by Both
	ids := datasetMetaObj.GetCatIds(names, superNames)
    fmt.Println("LoadCats ids len: ", len(ids), ids)

    // filter by catids
	results := datasetMetaObj.LoadCats(ids)
    fmt.Println("LoadCats result: ", len(results), results)
}

func Test_LoadImgs(t *testing.T) {
	ids := []int{397133, 87038, 6818}
    fmt.Println("LoadImgs ids len: ", len(ids), ids)

    // filter by catids
	results := datasetMetaObj.LoadImgs(ids)
    fmt.Println("LoadImgs result: ", len(results), results)
}

func Test_decoderExample(t *testing.T) {
	var err error
	

	var labelInfo CocoData
	err = json.Unmarshal(datasetMeta, &labelInfo)
	if err !=nil{
		fmt.Println("json.unmarshal failed,err:",err)
		return
	}
	fmt.Println(len(labelInfo.Annotations))

	for _, info1 := range labelInfo.Annotations{


		seg := info1.Segmentation.SegmentationHelper

		stype := seg.SegmentationType()

		switch (stype) {
		case "Polygon":
			segmentation := seg.(*SegmentationPolygon)
			fmt.Println("Type is Polygon. Create PolygonTool labeled Result", segmentation)
		default:
			segmentation := DecodeSegmentToMask(seg)
			fmt.Println("Type is RLE or RLEUncompressed. Create PolygonTool labeled Result", len(segmentation))
		}

		break;

	}

}


// func decoderExample(info []byte){
// 	var err error
	

// 	var labelInfo CocoData
// 	err = json.Unmarshal(info,&labelInfo)
// 	if err !=nil{
// 		fmt.Println("json.unmarshal failed,err:",err)
// 		return
// 	}
// 	fmt.Println(len(labelInfo.Annotations))

// 	size := [2]uint32{5, 6} // segmentation.size
// 	originMask := []byte{0,0,0,0,0,1,1,1,1,1,1,0,0,0,0,1,1,0,1,1,0,0,0,0,0,1,1,0,1,1}
// 	fmt.Println("originMask", originMask)

// 	rle := encodeRLE(originMask, size[0], size[1], 1)
// 	fmt.Println("encodeRLE => rle", rle)

// 	decodedMask := rle.Decode()
// 	fmt.Println("rle.Decode() => decodedMask", decodedMask)

// 	// countsString := C.GoString((* C.char)(rle.ToChar().Cc)) // segmentation.counts
// 	// fmt.Println("rle.ToChar().Cc => countsString", countsString)

// 	// countsByte := []byte(countsString)
// 	// cocoChar := Char{C.CBytes(countsByte)}
// 	// fmt.Println("countsString create Char => cocoChar", cocoChar)

// 	// cocoCharRle := cocoChar.ToRLE(size[0], size[1])
// 	// fmt.Println("cocoChar.ToRLE => cocoCharRle", cocoCharRle)

// 	// finalMask := cocoCharRle.Decode()
// 	// fmt.Println("cocoCharRle.Decode() => finalMask", finalMask)

// 	cnts := []uint32{70375,8,415,12,411,15,409,17,407,19,405,20,405,21,404,21,404,21,402,23,400,26,398,29,395,31,393,33,392,34,390,41,384,43,382,44,381,45,380,46,379,46,379,47,378,47,378,47,378,47,378,48,377,47,378,49,375,50,375,50,375,50,375,50,375,46,379,43,382,41,385,38,388,18,2,8,1,7,391,4,6,4,6,4,6,3,69652,6,418,8,416,10,414,12,413,12,413,12,413,13,412,13,412,13,412,13,412,13,412,13,412,13,412,12,413,12,414,10,416,8,419,4,844,7,417,9,415,11,413,14,411,15,410,16,409,17,408,17,408,17,408,19,407,19,407,19,407,19,407,18,409,17,411,14,412,13,412,13,412,14,411,15,6,4,400,15,4,6,400,15,3,8,400,14,2,9,401,13,1,10,402,12,1,10,402,23,403,10,1,11,404,8,2,22,395,4,4,39,386,41,384,41,383,42,383,42,383,43,381,44,381,44,381,44,381,44,381,45,380,45,380,45,380,45,380,46,379,46,380,45,381,44,381,45,381,44,382,43,383,42,384,41,386,40,406,18,408,18,407,18,407,18,407,19,407,18,407,18,407,19,407,18,408,17,408,17,409,17,410,15,412,13,372,6,37,9,372,8,416,10,414,12,413,12,413,13,412,14,411,15,410,15,410,15,410,15,410,15,410,15,410,15,410,15,409,21,404,22,402,24,401,25,400,25,400,25,400,25,400,25,390,35,389,36,388,36,388,36,387,37,387,38,386,39,385,39,386,38,387,36,389,37,388,38,386,39,386,39,386,39,386,39,386,39,386,39,386,39,386,39,386,39,387,38,388,23,2,12,389,18,6,12,391,12,10,12,80,3,330,12,79,16,318,12,79,28,306,11,80,28,306,11,80,28,306,11,80,29,305,11,80,29,305,11,80,29,306,10,81,29,305,10,81,29,306,9,81,29,307,7,82,30,308,4,83,30,396,29,397,29,397,28,397,28,398,28,397,28,397,29,396,29,396,29,396,30,395,30,395,30,395,30,395,30,395,30,396,29,396,29,397,28,398,27,399,26,400,24,402,23,402,23,402,23,403,21,405,20,406,19,407,18,408,16,410,14,412,12,415,8,31756,20,403,27,397,30,394,32,392,34,391,35,389,36,389,36,389,36,389,36,389,37,388,37,388,37,388,37,388,37,389,36,390,36,389,36,390,35,391,34,391,34,391,34,391,35,149}
// 	cntsRle := CompressRLE(cnts,425,640)
// 	// cntsString := C.GoString((* C.char)(cntsRle.ToChar().Cc)) // segmentation.counts
// 	// fmt.Println("cntsRle.ToChar().Cc => cntsString", cntsString)

// 	cntsRleDecodedMask := cntsRle.Decode()
// 	fmt.Println("cntsRle.Decode() => cntsRleDecodedMask", 425*640, len(cntsRleDecodedMask))
// 	// return;


// 	for _, info1 := range labelInfo.Annotations{


// 		seg := info1.Segmentation.SegmentationHelper

// 		stype := seg.SegmentationType()

// 		switch (stype) {
// 		case "Polygon":
// 			segmentation := seg.(*SegmentationPolygon)
// 			fmt.Println("Type is Polygon. Create PolygonTool labeled Result", segmentation)
// 		default:
// 			segmentation := seg.(*SegmentationRLE)
// 			r := []byte(segmentation.Counts)
// 			c := Char{C.CBytes(r)}
// 			rle := c.ToRLE(segmentation.Size[0], segmentation.Size[1])
// 			mask := rle.Decode()
// 			fmt.Println("Type is RLE. Create MaskTool labeled Result", len(mask))
// 		case "RLEUncompressed":
// 			segmentation := seg.(*SegmentationRLEUncompressed)
// 			// fake segmentation.Counts
// 			info1_Segmentation_Counts := cnts
// 			cntsRle := CompressRLE(info1_Segmentation_Counts, segmentation.Size[0], segmentation.Size[1])
// 			mask := cntsRle.Decode()
// 			fmt.Println("Type is RLEUncompressed. Create MaskTool labeled Result", len(mask))
// 		}

// 		break;

// 	}

// }
