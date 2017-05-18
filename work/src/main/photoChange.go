package main

import (
	"os"
	"image/jpeg"
	"image"
	"image/color"
	"image/draw"
      "github.com/disintegration/imaging"

	"log"
)
/*模糊图*/
func quality(name string)  {

	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	src,err:=imaging.Decode(file)
	blurimage:=imaging.Blur(src,2)
	err = imaging.Save(blurimage, "_模糊"+name)
	if err != nil {
		log.Fatalf("Save failed: %v", err)
	}
}
func gray(name string)  {

	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	src,err:=imaging.Decode(file)
	grayimage:=imaging.Grayscale(src)
	err = imaging.Save(grayimage, "_灰调"+name)
	if err != nil {
		log.Fatalf("Save failed: %v", err)
	}
}
func fliph(name string)  {

	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	src,err:=imaging.Decode(file)
	fliimage:=imaging.FlipH(src)
	err = imaging.Save(fliimage, "_镜像"+name)
	if err != nil {
		log.Fatalf("Save failed: %v", err)
	}
}
/*镜像合成图*/
	func mirro(name string) {

		file, err := os.Open(name)
		if err != nil {
			panic(err)
		}

		src,err:=imaging.Decode(file)
		src = imaging.Resize(src, 512, 256, imaging.Lanczos)
		rotatesrc:=imaging.Rotate180(src)
		rotatesrc = imaging.Resize(rotatesrc, 512, 256, imaging.Lanczos)
		//把两张颠倒的图片合成一张
		dst := imaging.New(512, 512, color.NRGBA{0, 0, 0, 0})
		dst = imaging.Paste(dst, src, image.Pt(0, 0))
		dst = imaging.Paste(dst, rotatesrc, image.Pt(0, 256))

		// 把黏贴的两张图片存起来
		err = imaging.Save(dst, "_对称"+name)
		if err != nil {
			log.Fatalf("Save failed: %v", err)
		}
	}
/*缩略图*/
func small(name string) {

	f1, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	defer f1.Close()
	f3, err := os.Create("_缩小" + name)
	if err != nil {
		panic(err)
	}
	defer f3.Close()
	m1, err := imaging.Decode(f1)
	if err != nil {
		panic(err)
	}
	bounds:=m1.Bounds()
	m := image.NewRGBA(bounds)
	white:=color.RGBA{255,255,255,255}
	draw.Draw(m,bounds,&image.Uniform{white},image.ZP,draw.Src)
	draw.Draw(m, bounds, m1, image.ZP, draw.Src)
	dstImagin128:=imaging.Resize(m,128,128,imaging.Lanczos)
	err = jpeg.Encode(f3, dstImagin128, &jpeg.Options{90})
	if err != nil {
		panic(err)
	}

}
/*四合一*/
func change(name string) {

	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	src,err:=imaging.Decode(file)
	src = imaging.CropAnchor(src, 350, 350, imaging.Center)
	// 设置图片大小
	src = imaging.Resize(src, 256, 0, imaging.Lanczos)
	// 模糊图片
	img1 := imaging.Blur(src, 2)
	// 锐化图片
	img2 := imaging.Grayscale(src)
	img2 = imaging.AdjustContrast(img2, 20)
	img2 = imaging.Sharpen(img2, 2)

	// 把图片内所有颜色颠倒
	img3 := imaging.Invert(src)

	// 卷积图像
	img4 := imaging.Convolve3x3(
		src,
		[9]float64{
			-1, -1, 0,
			-1, 1, 1,
			0, 1, 1,
		},
		nil,
	)

	// 把四张图片黏在一起
	add := imaging.New(512, 512, color.NRGBA{0, 0, 0, 0})
	add = imaging.Paste(add, img1, image.Pt(0, 0))
	add = imaging.Paste(add, img2, image.Pt(0, 256))
	add = imaging.Paste(add, img3, image.Pt(256, 0))
	add = imaging.Paste(add, img4, image.Pt(256, 256))


	err = imaging.Save(add, "_四合一"+name)
	if err != nil {
		log.Fatalf("Save failed: %v", err)
	}


}
/*卷积处理*/
func convolve(name string) {

	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	src, err := imaging.Decode(file)
	convolvesrc:=imaging.Convolve3x3(src,[9]float64{
		-1, -1, 0,
		-1, 1, 1,
		0, 1, 1,
	},
		nil,
	)
	err = imaging.Save(convolvesrc, "_卷积"+name)
	if err != nil {
		log.Fatalf("Save failed: %v", err)
	}
}
/*锐化*/
func sharpen(name string){
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	src, err := imaging.Decode(file)
	sharpensrc:=imaging.Sharpen(src,4)

	err = imaging.Save(sharpensrc, "_锐化"+name)
	if err != nil {
		log.Fatalf("Save failed: %v", err)
	}
}
/*颜色转换*/
func invert(name string) {

	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	src, err := imaging.Decode(file)
	invertsrc:=imaging.Invert(src)
	err = imaging.Save(invertsrc, "_颜色颠倒"+name)
	if err != nil {
		log.Fatalf("Save failed: %v", err)
	}
}
/*未用，创建文件夹的*/
func makefile()  {
	os.MkdirAll("./image", os.ModePerm)
}
