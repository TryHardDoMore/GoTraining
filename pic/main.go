package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"reflect"
	"os"
)

func decodePng(path string) (image.Image, string, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	return image.Decode(file)
}

func getPixelValue(img image.Image, x, y int) color.Color {
	return img.At(x, y)
}

func getImageMap(img image.Image) [][]color.Color {
	// Getting shape of image
	rect := img.Bounds()
	// Getting cordinates
	dx, dy := rect.Dx(), rect.Dy()
	imgColor := make([][]color.Color, dy)
	// Img to 2 dimmentional array
	for y := 0; y < dy; y++ {
		imgColor[y] = make([]color.Color, dx)
		for x := 0; x < dx; x++ {
			imgColor[y][x] = img.At(x, y)
		}
	}
	return imgColor
}

func calcCapacity(img image.Image) int {
	rect := img.Bounds()
	dx, dy := rect.Max.X, rect.Max.Y
	cap := (dx*dy - 2*dx)/64
	return cap
}


func updateImage(img image.Image, pattern []byte) image.Image{
	imag := image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X, img.Bounds().Max.Y))
	for index := 0; index < img.Bounds().Max.X; index++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			clr := img.At(index,y)
			clr.RGBA()
			r,g,b,c := 	reflect.ValueOf(clr).Field(0).Interface().(uint8),
						reflect.ValueOf(clr).Field(1).Interface().(uint8), 
						reflect.ValueOf(clr).Field(2).Interface().(uint8),
						reflect.ValueOf(clr).Field(3).Interface().(uint8)
			b  &= 0xF0
			imag.Set(index,y, color.RGBA{r,g,b,c})
			// Lb to 0  b  &= 0xFE
			// LB to 1  b |= 0x01
		}
	}
	return imag
}

func main() {
	// var message = flag.String("text","RandomText","Message to be enrypted")
	flag.Parse()
	img, t, err := decodePng("Lenna.png")
	if err != nil{
		panic(err)
	}
	if t != "png" {
		panic("Wrong format of pic")
	}
	var pattern []byte
	imag := updateImage(img, pattern)
	

	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, imag)

}
