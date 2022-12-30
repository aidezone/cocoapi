# coco
coco dataset api for go

# Author
forked from dereklstinson/coco
modify by gaoyuan1

# testcase
```bash
cd GolangAPI
go test -v --count=1 .
```

# How to use
```bash
export CGO_ENABLED=1

go get -u github.com/aidezone/cocoapi/GolangAPI

go mod tidy

```

```golang
// main.go

package main

import (
    coco "github.com/aidezone/cocoapi/GolangAPI"
)

func main() {
    
    dataset, err = ioutil.ReadFile("../anno/stuff_val2017.json")
    if err !=nil{
        fmt.Println("err:", err)
        return
    }

    cocoApi := NewCocoApi(dataset);
    cocoApi.LoadAnns(cocoApi.GetAnnIds(cocoApi.GetImgIds(nil), nil, nil, 3))
}

```