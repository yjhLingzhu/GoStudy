package main

import (
	"bufio"
	"fmt"
	"image"

	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	// datasource := "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X7A3wR/aS8Jt4d8eeB7S9iBJVXj6H8K/J7/go1/wa8+EvEfgTU/iF+zdp6abqdjAWh0a1t8/aWGT949K/Bv43/s+/E/9n3xbceDfiX4an0+8tpCjpKveuJooooAJOBXt37J/7A37QX7Xevw6b8NfBd7NaO+2XUFti8UZ9yDX7F/8E5/+DUu10vxHpnj39qjxBb6lYB0lSw095IZFYcjdnII9q/bT4G/sufBv4BeG7Twp8M/AlhZW9lGEjmWzjEhA7s4UEmvSF0+DdvdQTjoOAKnACjCgAegoooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooqrqOlW2oWslrLGCsikMuOtfm7/AMFg/wDgip8Mf2yPBNx4j8LaHDZ67ZW7yxzRJsEhCnghQCT9a/l8/ao/Zl8f/sv/ABS1H4f+N9DntGtblkiaaMrvA9M15lRRX3P/AMEW/wDgkr4x/wCCiPx8j0zxTp97p3hXTQJdRvjCU3cjCrvADZ+tf1Gfsm/8E6/gF+yP4Js/A3w08MQRLbwBZbpbYK8zf3mwete+6dodnp0CwRpkIMKO1XAABgCiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiioL+yivYGikQEMpBB71+UH/Ber/gjx4f/aq+GNz43+HOhRw+I7ISTpNEgUzYXOGb8K/l+8a+C/EPgDxJd+FfE2nS213ZztFLHLGVOVOD1ArKor+1X/gm3+zz8HPg78FYdF+GPgq30mSK5C3IjHLMMjk19SxRiJAo/E+tOoooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooorM8TaDaa7pkthdxKySoVYEeor+eP/g43/4I72vg6S9/aU+Enh/bGSX1KK3h4BJ5Yn3r8NZ4Jbad7eZCroxVlPYimV/d58EtB07QNAFtp8W1WlDN9a9AooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooIBGCK8y/ab+Bnhj44/CzVvAXiXS4ri3vbRkIkjBJyPcV/H9/wVy/YZ8S/sXftPavoM+lyRaTf3LzafIUO3aWPGTXyjX95fwsubS60C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oooooooooooooooooooooooooooooooooooooops/wDqH/3D/KqEJTcwKsef4TRpVyt7BHcoBhsHitGvlX/gshrr+H/+CdXxK1KN8OugSqMn+9hf61/PB/wa920l9/wU3iuG5KQlmx7u1f1gaYAse0dlFWaKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKbP/qH/ANw/yrPiKITudhn0FJolkdPs4bUvu24Bb1rTr4y/4LqG4/4dqfEZLcn/AJBXzY9N61+FX/BpvpFvf/8ABRDWLmSDc1tpalD/AHcu/wDhX9TVgpCFvoKsUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUU2fmFx/sn+VZpI3EGXbz1qxbgh1B/vCrlfG//AAXHiaX/AIJufEaMH72kdv8AfWvx7/4M9vhXb6h+0f4/+JzId1lElnHu/wCBN/7NX9I9mCIsnueKlooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooops/+pf8A3T/KqMSKzksyj61LGD5qnPcVbr5C/wCC1SLN/wAE9PH8EijDaQ2fzFfnF/wZ6/D1dN8E/EPxjLGc3niIiJivUKqiv3Wtf+Pdfx/nUlFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFNn/ANQ/+4f5VSgGWOFBPvT4HDupB/iFXK+P/wDgtdN5P/BPrx6QOf7Hf+Yr5d/4NX/h7H4d/YxbxMsG06lqVxKWx1+cr/Sv1ggAWFQB/DTqKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKbP/qX/wB0/wAqooqZPmflUsTBpRtXA3DFW6+LP+C7esx6D/wTw8b3DOMy2McKg9y8yL/WoP8Aghb8MrP4c/sA+DIYYVV7uwa4faP70j/4V9sxgrGqkchRmlooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooops/ELn/AGT/ACrNkW4kb9wR75qzHjzRgcbhVpjtUt6CvzO/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"

	// b_img, _ := base64.StdEncoding.DecodeString(datasource)                  //转成图片文件并把文件写入到buffer
	// err2 := ioutil.WriteFile("C:/Users/Ykim H/Desktop/1/1.png", b_img, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）  蒙板存为png后使用image打开会报错unknown format
	// if err2 != nil {
	// 	fmt.Println(123)
	// }

	// 对比像素点，这种方式效果不好
	f, err := os.Open("C:/Users/Ykim H/Desktop/1/1.jpg")
	f2, _ := os.Open("D:/Data/zhiyuzhou/portrait_matting/20211216/apuhsrnGOIUQQiSe.png")
	defer f.Close()
	defer f2.Close()
	if err != nil {
		fmt.Println(123)
		panic(err)
	}
	img, err := jpeg.Decode(f)
	img2, _, _ := image.Decode(f2)
	if err != nil {
		fmt.Println(456)
		panic(err)
	}
	fmt.Println(img.At(0, 2))
	fmt.Println(img2.At(0, 2))

	img1 := image.NewRGBA(image.Rect(0, 0, 630, 1120))
	for x := 0; x < img.Bounds().Dx(); x++ { // 换背景
		for y := 0; y < img.Bounds().Dy(); y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			if r == 0 && g == 0 && b == 0 {
				// fmt.Println("进来了")
				img1.Set(x, y, color.RGBA{0, 191, 255, 255})
			} else {
				img1.Set(x, y, img2.At(x, y))
			}
		}
	}

	outFile, err := os.Create("C:/Users/Ykim H/Desktop/1/2.jpg") // 创建一张空的图片文件
	defer outFile.Close()
	if err != nil {
		panic(err)
	}
	b := bufio.NewWriter(outFile) // NewWriter返回一个新的Writer，其缓冲区具有默认大小
	err = png.Encode(b, img1)     // 将绘制完的图片img重新编码并输入到Writer b中
	if err != nil {
		panic(err)
	}
	err = b.Flush() // 将所有缓冲数据写入底层io.Writer
	if err != nil {
		panic(err)
	}
}