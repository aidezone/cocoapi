package models

import (
	"encoding/json"
	// "fmt"
	"errors"
)

//Information is basic image and Coco information and is shared between all the data formats
type Information struct {
	Year        int    `json:"year,omitempty"`
	Version     string `json:"version,omitempty"`
	Description string `json:"description,omitempty"`
	Contributor string `json:"contributor,omitempty"`
	URL         string `json:"url,omitempty"`
	DateCreated string `json:"date_created,omitempty"`
}

//Image is the image information and is shared between all the dataformats
type Image struct {
	ID           int    `json:"id,omitempty"`
	Width        int    `json:"width,omitempty"`
	Height       int    `json:"height,omitempty"`
	FileName     string `json:"file_name,omitempty"`
	License      int    `json:"license,omitempty"`
	FlickrURL    string `json:"flickr_url,omitempty"`
	CocoURL      string `json:"coco_url,omitempty"`
	DateCaptured string `json:"date_captured,omitempty"`
}

//License is the license information and is shared between all the formats
type License struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type DataFormats interface {
	ObjectDetection | StuffSegmentation | PanopticSegmentation
}


// type Segment interface {
// 	SegmentationRLE | SegmentationRLEUncompressed | SegmentationPolygon
// }

// /*

// These are helpers to work with the shared json catagory.


// */

// //RLEgo is a struct that will take the info from json to an RLE format
// type RLEgo struct {
// 	Counts []uint32
// 	Size   []uint32
// }

// //Polygon from what ive seen looks to be in the form of [1][x0, x1,x2,x3 . . .] but it could be more on the first part.
// type Polygon [][]float32

// //SegmentationHelper is used for segmentation
// type SegmentationHelper struct {
// 	Poly Polygon
// 	Rle  RLEgo
// }


//Segment interface a placeholder for Segmentation data structures
//It can either be an RLE or a Polygon
type Segment struct {
	SegmentationHelper
}

//SegmentationHelper is used for segmentation
type SegmentationHelper interface {
	SegmentationType() string
}

type SegmentationRLE struct {
	Counts string `json:"counts,omitempty"`
	Size   [2]uint32 `json:"size,omitempty"`
}

type SegmentationRLEUncompressed struct {
	Counts []float32 `json:"counts,omitempty"`
	Size   [2]uint32 `json:"size,omitempty"`
}

type SegmentationPolygon [][]float32

func (s *SegmentationRLE) SegmentationType() string {
	return "RLE"
}

func (s *SegmentationRLEUncompressed) SegmentationType() string {
	return "RLEUncompressed"
}

func (s *SegmentationPolygon) SegmentationType() string {
	return "Polygon"
}

func (s *Segment) UnmarshalJSON(jBytes []byte) error {
	seg, err := decodeToSegmentation(jBytes)
	s.SegmentationHelper = seg
	if err != nil {
		return err
	}
	return nil
}

func decodeToSegmentation(jBytes []byte) (SegmentationHelper, error) {
	// fmt.Printf("parsing nested json %s \n", string(jBytes))

	var err error;
	segRleUnc := &SegmentationRLEUncompressed{}
	err = json.Unmarshal(jBytes, segRleUnc)
	if err == nil {
		return segRleUnc, nil
	}

	segRle := &SegmentationRLE{}
	err = json.Unmarshal(jBytes, segRle)
	if err == nil {
		return segRle, nil
	}

	segPolygon := &SegmentationPolygon{}
	err = json.Unmarshal(jBytes, segPolygon)
	if err == nil {
		return segPolygon, nil
	}

	return nil, errors.New("decode segmentation error")
}



