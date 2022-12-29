package coco

import (
	"encoding/json"
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

type CocoData struct {
	Info        Information    `json:"info,omitempty"`
	Images      []Image        `json:"images,omitempty"`
	Annotations []Annotation `json:"annotations,omitempty"`
	Licenses    []License      `json:"licenses,omitempty"`

	// ImageCaption does not have this property
	Categories  []Categories `json:"categories,omitempty"`
}

//Annotation is the object detection annotation
type Annotation struct {
	ImageID      int        `json:"image_id,omitempty"`

	// PanopticSegmentation does not have this property
	ID           int        `json:"id,omitempty"`

	// PanopticSegmentation and ImageCaption does not have this property
	CategoryID   int        `json:"category_id,omitempty"`
	Segmentation Segment    `json:"segmentation,omitempty"`
	Area         float32    `json:"area,omitempty"`
	Bbox         [4]float32 `json:"bbox,omitempty"`

	// PanopticSegmentation and StuffSegmentation and ImageCaption does not have this property
	Iscrowd      byte       `json:"iscrowd,omitempty"`

	// ImageCaption own property
	Caption      string     `json:"caption,omitempty"`

	// KeypointDetection own property
	Keypoints    []float32  `json:"keypoints,omitempty"`
	NumKeypoints int        `json:"num_keypoints,omitempty"`

	// PanopticSegmentation own property
	FileName     string          `json:"file_name,omitempty"`
	SegmentsInfo []PSSegmentInfo `json:"segments_info,omitempty"`
}

//Edge desribes a 2 point edge Probably [x,y] I haven't tested it yet
type Edge [2]uint32

//Categories is the object detection categories.
type Categories struct {
	ID            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Supercategory string `json:"supercategory,omitempty"`

	// KeypointDetection own property
	Keypoints     []string `json:"keypoints,omitempty"`
	Skeleton      []Edge   `json:"skeleton,omitempty"`

	// PanopticSegmentation own property
	Isthing       byte      `json:"isthing,omitempty"`
	Color         [3]uint32 `json:"color,omitempty"`
}

//PSSegmentInfo contains segment info for the annotation
type PSSegmentInfo struct {
	ID         int        `json:"id,omitempty"`
	CategoryID int        `json:"category_id,omitempty"`
	Area       int        `json:"area,omitempty"`
	Bbox       [4]float32 `json:"bbox,omitempty"`
	Iscrowd    byte       `json:"iscrowd,omitempty"`
}

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



