package main

import (
	"bytes"
	"flag"
	"fmt"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/chai2010/webp"
)

var Input_src = flag.String("src", "", "-src='/tmp/a.png'")
var Input_mode = flag.String("mode", "lossless_rgba", "-mode=lossless_rgba|lossless_rgb|lossless_gray|rgba90|rgba80|rgba70|rgb90|rgb80|rgb70|gray90|gray80|gray70")
var Input_dest = flag.String("dest", "", "-dest='/tmp/a.webp'")
var Input_help = flag.String("help", "", "-src='/tmp/a.png' -dest='/tmp/a.webp'")
var Input_usage = flag.String("usage", "", "./image2webp -src='/tmp/a.png' -dest='/tmp/a.webp'")

func main() {
	flag.Parse()
	if _, find := os.Stat(*Input_src); find != nil {
		fmt.Println("File " + *Input_src + " not exists.")
		return
	}
	Img2Webp(*Input_src, *Input_dest)
}
func Img2Webp(src string, dest string) {
	if _, find := os.Stat(src); find == nil {
		lowSrc := strings.ToLower(src)
		if strings.HasSuffix(lowSrc, ".png") {
			err := Png2Webp(src, dest)
			if err != nil {
				log.Println(err)
			}
		}
		if strings.HasSuffix(lowSrc, ".jpeg") || strings.HasSuffix(lowSrc, ".jpg") {
			err := Jpeg2Webp(src, dest)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
func Png2Webp(src string, dest string) error {
	imgByte, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	pngImg, ok := png.Decode(bytes.NewReader(imgByte))
	if ok != nil {
		return ok
	}
	webpByte := []byte{}
	switch *Input_mode {
	case "lossless_rgba":
		webpByte, _ = webp.EncodeLosslessRGBA(pngImg)
	case "lossless_rgb":
		webpByte, _ = webp.EncodeLosslessRGB(pngImg)
	case "lossless_gray":
		webpByte, _ = webp.EncodeLosslessGray(pngImg)
	case "rgba90":
		webpByte, _ = webp.EncodeRGBA(pngImg, 90.0)
	case "rgba80":
		webpByte, _ = webp.EncodeRGBA(pngImg, 80.0)
	case "rgba70":
		webpByte, _ = webp.EncodeRGBA(pngImg, 70.0)
	case "rgb90":
		webpByte, _ = webp.EncodeRGB(pngImg, 90.0)
	case "rgb80":
		webpByte, _ = webp.EncodeRGB(pngImg, 80.0)
	case "rgb70":
		webpByte, _ = webp.EncodeRGB(pngImg, 70.0)
	case "gray90":
		webpByte, _ = webp.EncodeGray(pngImg, 90.0)
	case "gray80":
		webpByte, _ = webp.EncodeGray(pngImg, 80.0)
	case "gray70":
		webpByte, _ = webp.EncodeGray(pngImg, 70.0)
	default:
		webpByte, _ = webp.EncodeLosslessRGBA(pngImg)
	}
	fileInfo, _ := os.Stat(src)
	ioutil.WriteFile(dest, webpByte, fileInfo.Mode())
	return nil
}
func Jpeg2Webp(src string, dest string) error {
	imgByte, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	jpegImg, ok := jpeg.Decode(bytes.NewReader(imgByte))
	if ok != nil {
		return ok
	}
	webpByte := []byte{}
	switch *Input_mode {
	case "lossless_rgba":
		webpByte, _ = webp.EncodeLosslessRGBA(jpegImg)
	case "lossless_rgb":
		webpByte, _ = webp.EncodeLosslessRGB(jpegImg)
	case "lossless_gray":
		webpByte, _ = webp.EncodeLosslessGray(jpegImg)
	case "rgba90":
		webpByte, _ = webp.EncodeRGBA(jpegImg, 90.0)
	case "rgba80":
		webpByte, _ = webp.EncodeRGBA(jpegImg, 80.0)
	case "rgba70":
		webpByte, _ = webp.EncodeRGBA(jpegImg, 70.0)
	case "rgb90":
		webpByte, _ = webp.EncodeRGB(jpegImg, 90.0)
	case "rgb80":
		webpByte, _ = webp.EncodeRGB(jpegImg, 80.0)
	case "rgb70":
		webpByte, _ = webp.EncodeRGB(jpegImg, 70.0)
	case "gray90":
		webpByte, _ = webp.EncodeGray(jpegImg, 90.0)
	case "gray80":
		webpByte, _ = webp.EncodeGray(jpegImg, 80.0)
	case "gray70":
		webpByte, _ = webp.EncodeGray(jpegImg, 70.0)
	default:
		webpByte, _ = webp.EncodeLosslessRGBA(jpegImg)
	}
	fileInfo, _ := os.Stat(src)
	ioutil.WriteFile(dest, webpByte, fileInfo.Mode())
	return nil
}
