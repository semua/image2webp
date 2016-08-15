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
	webpByte, trans := webp.EncodeLosslessRGBA(pngImg)
	if trans != nil {
		return trans
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
	webpByte, trans := webp.EncodeLosslessRGBA(jpegImg)
	if trans != nil {
		return trans
	}
	fileInfo, _ := os.Stat(src)
	ioutil.WriteFile(dest, webpByte, fileInfo.Mode())
	return nil
}
