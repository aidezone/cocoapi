package coco

// #include <stdlib.h>
// #include <stdio.h>
import "C"

import (
	"encoding/json"
	"unsafe"
)

// The following API functions are defined:
//  CocoApi	- CocoApi api class that loads COCO annotation file and prepare data structures.
//  DecodeSegmentToMask - Decode binary mask M encoded via run-length encoding.
//  EncodeMaskToSegment - Encode binary mask M using run-length encoding.
//  EncodeRLEToSegment  - Encode binary mask M using run-length encoding.
//  GetAnnIds  - Get ann ids that satisfy given filter conditions.
//  GetCatIds  - Get cat ids that satisfy given filter conditions.
//  GetImgIds  - Get img ids that satisfy given filter conditions.
//  LoadAnns   - Load anns with the specified ids.
//  LoadCats   - Load cats with the specified ids.
//  LoadImgs   - Load imgs with the specified ids.
//  ShowAnns   - Display the specified annotations.
// Throughout the API "ann"=annotation, "cat"=category, and "img"=image.
// Help on each functions can be accessed by: "help COCO>function".

type CocoApi struct {
	datasetMeta CocoData
	imgMap map[int]Image
	annMap map[int]Annotation
	catMap map[int]Categories
	// catNameMap map[int]Categories
	imgToAnnMap map[int][]int
	// annToImgMap map[int][]int
	// catToAnnMap map[int][]int
	// annToCatMap map[int][]int
	// imgToCatMap map[int][]int
	catToImgMap map[int][]int
}

func NewCocoApi(datasetMeta []byte) (cocoApi *CocoApi, err error) {
	cocoApi = &CocoApi{
		imgMap: make(map[int]Image),
		annMap: make(map[int]Annotation),
		catMap: make(map[int]Categories),
		// catNameMap: make(map[string]Categories),
		imgToAnnMap: make(map[int][]int),
		// catToAnnMap: make(map[int][]int),
		// imgToCatMap: make(map[int][]int),
		catToImgMap: make(map[int][]int),
	}
	err = cocoApi.init(datasetMeta)
	return
}

func DecodeSegmentToMask(segmentation SegmentationHelper) (mask []byte) {
	stype := segmentation.SegmentationType()
	var segment *SegmentationRLE;
	switch (stype) {
	case "RLE":
		segment = segmentation.(*SegmentationRLE)
		
	case "RLEUncompressed":
		segmentTmp := segmentation.(*SegmentationRLEUncompressed)
		segment = EncodeRLEToSegment(segmentTmp)
	}

	rleGoBytes := []byte(segment.Counts)
	// 正确用法
	charSegment := &Char{
		Cc: unsafe.Pointer(&rleGoBytes[0]),
	}
	// // 错误用法
	// relBytes := C.CBytes(rleGoBytes)
	// defer C.free(relBytes)
	// charSegment := &Char{
	// 	Cc: relBytes,
	// }
	// fmt.Println("rleGoBytes len: ", len(rleGoBytes))
	// // 旧版用法
	// rle := charSegment.ToRLE(segment.Size[0], segment.Size[1])
	// 新版增加长度参数，避免越界
	rle := charSegment.ToRLEWithByteLen(segment.Size[0], segment.Size[1], uint32(len(rleGoBytes)))
	mask = rle.Decode()
	return
}

func EncodeMaskToSegment(mask []byte, size [2]uint32) *SegmentationRLE {
	rle := encodeRLE(mask, size[0], size[1], 1)
	return rleToSegment(rle, size)
}

func EncodeRLEToSegment(segmentation *SegmentationRLEUncompressed) *SegmentationRLE {
	rle := compressRLE(segmentation.Counts, segmentation.Size[0], segmentation.Size[1])
	return rleToSegment(rle, segmentation.Size)
}

func rleToSegment(rle *RLE, size [2]uint32) *SegmentationRLE {
	charPointer := (* C.char)(rle.ToChar().Cc)
	countsString := C.GoString(charPointer) // segmentation.counts
	return &SegmentationRLE{
		Counts: countsString,
		Size: size,
	}
}

func (api *CocoApi) init(datasetMeta []byte) (err error) {
	err = json.Unmarshal(datasetMeta, &api.datasetMeta)
	if err != nil {
		// fmt.Println("json.unmarshal failed,err:",err)
		return
	}

	// createIndex
	imgs := api.datasetMeta.Images
	for i := 0; i < len(imgs); i++ {		
		api.imgMap[imgs[i].ID] = imgs[i]
	}

	cats := api.datasetMeta.Categories
	for i := 0; i < len(cats); i++ {
		// api.catNameMap[cats[i].Name] = cats[i]
		api.catMap[cats[i].ID] = cats[i]
	}

	anns := api.datasetMeta.Annotations
	for i := 0; i < len(anns); i++ {		
		api.annMap[anns[i].ID] = anns[i]
		api.imgToAnnMap[anns[i].ImageID] = append(api.imgToAnnMap[anns[i].ImageID], anns[i].ID)
		// api.catToAnnMap[anns[i].CategoryID] = append(api.catToAnnMap[anns[i].CategoryID], anns[i].ID)
		// api.imgToCatMap[anns[i].ImageID] = append(api.imgToCatMap[anns[i].ImageID], anns[i].CategoryID)
		api.catToImgMap[anns[i].CategoryID] = append(api.catToImgMap[anns[i].CategoryID], anns[i].ImageID)
	}
	return
}

func (api *CocoApi) GetLicense() ([]License) {
	return api.datasetMeta.Licenses
}

func (api *CocoApi) GetInfo() (Information) {
	return api.datasetMeta.Info
}

func (api *CocoApi) GetAnnIds(imgIds, catIds, areaRng []int, iscrowd byte) (ids []int) {
	var anns map[int]Annotation
	if len(imgIds) == 0 && len(catIds) == 0 && len(areaRng) == 0 {
		anns = api.annMap
	} else {
		if len(imgIds) != 0 {
			list := api.imgToAnnMap[imgIds[0]]
			for i := 1; i < len(imgIds); i++ {
				list = append(list, api.imgToAnnMap[imgIds[i]]...)
			}
			anns = make(map[int]Annotation)
			for i := 0; i < len(list); i++ {
				anns[list[i]] = api.annMap[list[i]]
			}
		} else {
			anns = api.annMap
		}
		if len(catIds) != 0 {
			catIdMap := make(map[int]int)
			for i := 0; i < len(catIds); i++ {
				catIdMap[catIds[i]] = catIds[i]
			}
			for k, v := range anns {
				if _, ok := catIdMap[v.CategoryID]; !ok {
					delete(anns, k)
				}
				if len(areaRng) == 0 {
					continue
				}
				if v.Area <= float32(areaRng[0]) || v.Area >= float32(areaRng[1]) {
					delete(anns, k)
				}
			}
		}
	}
	// fmt.Println(anns)
	for _, v := range anns {
		if iscrowd == 0 || iscrowd == 1 {
			if (v.Iscrowd != iscrowd) {
				continue
			}
		}
		ids = append(ids, v.ID)
	}
	return
}

func (api *CocoApi) GetCatIds(names, supCatNames []string) (ids []int) {
    
    nameMap := make(map[string]byte)
	for i := 0; i < len(names); i++ {
		nameMap[names[i]] = 1
	}
	superNameMap :=  make(map[string]byte)
	for i := 0; i < len(supCatNames); i++ {
		superNameMap[supCatNames[i]] = 1
	}

	for _, v := range api.catMap {
		if len(names) > 0 {
			if _, ok := nameMap[v.Name]; !ok {
				continue
			}
		}
		if len(supCatNames) > 0 {
			if _, ok := superNameMap[v.Supercategory]; !ok {
				continue
			}
		}
		ids = append(ids, v.ID)
	}

	return
}

func (api *CocoApi) GetImgIds(catIds []int) (ids []int) {

	if len(catIds) < 1 {
		ids = getImageMapKeys(api.imgMap)
		return
	}
	for i := 0; i < len(catIds); i++ {
		ids = append(ids, api.catToImgMap[catIds[i]]...)
	}
	ids = removeDuplicates(ids)
	return
}

func (api *CocoApi) LoadAnns(ids []int) (list []Annotation) {
	for _, v := range ids {
		list = append(list, api.annMap[v])
	}
	return
}

func (api *CocoApi) LoadCats(ids []int) (list []Categories) {
	for _, v := range ids {
		list = append(list, api.catMap[v])
	}
	return
}

func (api *CocoApi) LoadImgs(ids []int) (list []Image) {
	for _, v := range ids {
		list = append(list, api.imgMap[v])
	}
	return
}

func (api *CocoApi) ShowAnns(ids []int) ([]interface{}, error) {
	return nil, nil
}

func removeDuplicates(ids []int) (filteredIds []int) {
    if len(ids) == 0 {
        return
    }

    idMap := make(map[int]byte)
	for _, v:= range ids{
	    idMap[v] = 1
	}

	j := 0
	keys := make([]int, len(idMap))
	for k := range idMap {
		keys[j] = k
		j++
	}
	filteredIds = keys
	return
}

func getImageMapKeys(m map[int]Image) []int {
	j := 0
	keys := make([]int, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}
