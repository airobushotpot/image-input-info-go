// Package imageinfo reads basic metadata from local PNG, JPEG, and GIF files.
// It does not upload files or contact any web service.
package imageinfo

import (
    "fmt"
    "image"
    _ "image/gif"
    _ "image/jpeg"
    _ "image/png"
    "os"
)

// Info describes encoded image dimensions, format, and file size in bytes.
// Width and Height do not account for EXIF orientation.
type Info struct {
    Width int
    Height int
    Format string
    Bytes int64
}

// InspectFile reads image configuration from a regular local file.
// A readable header does not guarantee that the full image is valid.
func InspectFile(path string) (Info, error) {
    f, err := os.Open(path)
    if err != nil { return Info{}, fmt.Errorf("open image: %w", err) }
    defer f.Close()
    stat, err := f.Stat()
    if err != nil { return Info{}, fmt.Errorf("stat image: %w", err) }
    if !stat.Mode().IsRegular() { return Info{}, fmt.Errorf("image is not a regular file") }
    cfg, format, err := image.DecodeConfig(f)
    if err != nil { return Info{}, fmt.Errorf("read image header: %w", err) }
    return Info{Width: cfg.Width, Height: cfg.Height, Format: format, Bytes: stat.Size()}, nil
}
