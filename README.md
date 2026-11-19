# Image Input Info for Go

A small Go library for reading the encoded dimensions, format, and file size of
local JPEG, PNG, and GIF images. It can help organize images before a separate
image-to-video workflow, such as [Image to Video AI](https://imagetovideoai.tools/).

This library does not upload images, generate video, or call a website API.

## Usage

```go
package main

import (
    "fmt"
    imageinfo "github.com/airobushotpot/image-input-info-go"
)

func main() {
    info, err := imageinfo.InspectFile("image.png")
    if err != nil {
        panic(err)
    }
    fmt.Printf("%dx%d %s %d bytes\n", info.Width, info.Height, info.Format, info.Bytes)
}
```

The result uses encoded dimensions and does not apply EXIF orientation. Header
inspection does not validate all pixels or guarantee that a file is safe or
compatible with any particular service. The package requires Go 1.22 or later.
Installation and functionality have not been independently tested.
