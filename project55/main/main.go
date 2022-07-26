package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	f, err := os.Open("D:/Data/zhiyuzhou/portrait_resize/20211220/scalratio.png")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	TheImg, fm, err := image.Decode(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("fm:", fm)
	img := TheImg.(*image.RGBA)
	subImage := img.SubImage(image.Rect(5, 8, 5+295, 8+413)).(*image.RGBA)              // 裁剪
	outFile, _ := os.Create("D:/Data/zhiyuzhou/portrait_matting/20211216/gopher12.jpg") // 创建一张空的图片文件
	jpeg.Encode(outFile, subImage, &jpeg.Options{Quality: 100})
}
