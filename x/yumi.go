// favicon encoder 
// 32x32 png and json base64 data uri formats
package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"
)

var (
	Png      string
	Json     string
	PngFile  *os.File
	JsonFile *os.File
	Pixels   *image.RGBA
)

type Favio struct {
	DataUri string
}

func Filepath(p, s string) string {
	now := time.Now().UnixNano()
	r := fmt.Sprintf("%s%d%s", p, now, s)
	return r
}

func Path() {
	Png = Filepath("favicon_", ".png")
	Json = Filepath("dataUri_", ".json")
}

func Open() {
	var err error
	PngFile, err = os.Create(Png)
	if err != nil {
		fmt.Println(err)
	}
	JsonFile, err = os.Create(Json)
	if err != nil {
		fmt.Println(err)
	}
}

func Image() {
	rect := image.Rect(0, 0, 32, 32)
	Pixels = image.NewRGBA(rect)
}

func Paint() {
    // okaq xonix
    c0 := color.RGBA{27,27,27,255}
    s0 := Pixels.Bounds().Dx() * Pixels.Bounds().Dy()
	for i := 0; i < s0; i++ {
		x0 := i % Pixels.Bounds().Dx()
		y0 := i / Pixels.Bounds().Dy()
		Pixels.SetRGBA(x0, y0, c0)
	}
	fmt.Println(Pixels.At(0, 0))
}

func Write() {
	err := png.Encode(PngFile, Pixels)
	if err != nil {
		fmt.Println(err)
	}
	// following line encodes pixel data only, not full png file
	fi0, err := PngFile.Stat()
	if err != nil {
		fmt.Println(err)
	}
	d1 := make([]byte, fi0.Size())
	_, err = PngFile.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
	}
	_, err = PngFile.Read(d1)
	if err != nil {
		fmt.Println(err)
	}
	s0 := base64.StdEncoding.EncodeToString(d1)
	s1 := "data:image/png;base64,"
	s2 := fmt.Sprintf("%s%s", s1, s0)
	d0 := Favio{DataUri: s2}
	b0, err := json.Marshal(d0)
	if err != nil {
		fmt.Println(err)
	}
	_, err = JsonFile.Write(b0)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	Path()
	Open()
	defer PngFile.Close()
	defer JsonFile.Close()
	Image()
	Paint()
	Write()
}

// yumi favicon encoder
// AQ <aq@okaq.com>
// v day 2021



