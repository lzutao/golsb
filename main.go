package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const byte_size = 8

func init() {
	// disable all log
	log.SetOutput(ioutil.Discard)
}

func usage() {
	fmt.Printf("usage: %s [-h] [--encode | --decode] [--message=message] <image.jpg>\n\n", os.Args[0])
	fmt.Printf("Steganography image file\n\n")
	fmt.Println("optional arguments:")
	flag.PrintDefaults()
	fmt.Println("\n[+] Written by lzutao")
}

func encodeMessage(img image.Image, msg string, outfile string) {
	log.Fatal("Not implemented!")
}

func decodeMessage(img image.Image) string {
	var (
		bounds = img.Bounds()
		width_min, width_max = bounds.Min.X, bounds.Max.X
		height_min, height_max = bounds.Min.Y, bounds.Max.Y
	)

	var out []byte
	for y := height_min; y < height_max; y++ {
		for x := width_min; x < width_max; x++ {
			pix := img.At(x, y)
			r, g, b, _ := pix.RGBA()
			r = r & 1
			g = g & 1
			b = b & 1
			out = append(out, byte(r), byte(g), byte(b))
		}
	}

	var rs []byte
	var out_len = len(out)
	for i := 0; i < out_len; i += byte_size {
		var c byte = 0;
		for j := 0; j < byte_size; j++ {
			c <<= 1
			c |= out[i + j]
		}
		rs = append(rs, c)
	}
	return string(rs[:])
}

func main() {
	var (
		isEncode bool
		isDecode bool
		message string
	)

	flag.BoolVar(&isEncode, "encode", false, "encode a message into an image")
	flag.BoolVar(&isDecode, "decode", false, "decode a hidden message from an image")
	flag.StringVar(&message, "message", "", "message to be encoded to the image")

	flag.Usage = usage
	flag.Parse()

	var args = flag.Args()
	if len(args) == 0 {
		usage()
		os.Exit(1)
	}
	var imgName = args[0]

	file, err := os.Open(imgName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var img image.Image
	var ext = filepath.Ext(imgName)
	switch ext {
	case ".png":
		img, err = png.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
	case ".jpg":
		img, err = jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Please use image filename with .jpg or .png extension")
	}

	if isEncode {
		encodeMessage(img, message, "a" + ext)
	} else if isDecode {
		s := decodeMessage(img)
		fmt.Println(s)
	} else {
		fmt.Println("Choose --encode or --decode to continue.")
		flag.PrintDefaults()
	}
}
