package main

import "C"
import (
	"ASparkOfFire/realesrgan-golang.git/realesrgan"
	"github.com/sirupsen/logrus"
)

func main() {
	params := realesrgan.RealESRGANParams{
		GPUID:      1,
		TTA:        0,
		Scale:      4,
		TileSize:   32,
		Prepadding: 10,
		ModelPath:  "./models/remacri.bin",
		ParamPath:  "./models/remacri.param",
	}
	if err := realesrgan.RealESRGAN(params, "./test/input.jpg", "./test/output.jpg", realesrgan.ImageFormatJPEG); err != nil {
		logrus.Errorf("Error While Upscaling the image: %v", err)
	}
}
