package realesrgan

/*
#cgo LDFLAGS: -L. -linfer
#include "infer.h"

#include <stdlib.h>
*/
import "C"
import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"unsafe"
)

type ImageFormat string

const (
	ImageFormatPNG  ImageFormat = "png"
	ImageFormatJPEG ImageFormat = "jpg"
)

func (img ImageFormat) string() string {
	return string(img)
}

type RealESRGANParams struct {
	GPUID      int
	TTA        int
	Scale      uint
	TileSize   uint
	Prepadding uint
	ModelPath  string
	ParamPath  string
}

func RealESRGAN(params RealESRGANParams, inputPath string, outputPath string, imageFormat ImageFormat) error {
	// Sample input data and format
	file, err := os.Open(inputPath)
	if err != nil {
		logrus.Errorf("Error opening file: %v", err)
		return err
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		logrus.Errorf("error occurred: %v\n", err)
		return err
	}

	// Convert Go slice to C array
	var (
		InputImagePtr = (*C.uchar)(unsafe.Pointer(&b[0]))
		InputImageLen = C.size_t(len(b))

		GPUID      = C.int(params.GPUID)
		TTA        = C.int(params.TTA)
		Scale      = C.int(params.Scale)
		TileSize   = C.int(params.TileSize)
		Prepadding = C.int(params.Prepadding)

		ModelPath = C.CString(params.ModelPath)
		ParamPath = C.CString(params.ParamPath)

		InputImageFormat = C.CString(imageFormat.string())
	)

	// Create RealESRGAN instance
	realesrgan := C.create_realesrgan_instance(GPUID, TTA, Scale, TileSize, Prepadding, ModelPath, ParamPath)
	buffer := C.process_image(InputImagePtr, InputImageLen, InputImageFormat, realesrgan, 4)

	// Access the returned buffer data
	data := C.GoBytes(unsafe.Pointer(buffer.data), C.int(buffer.size))

	// Write the output to a file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		logrus.Errorf("Error creating output file: %v", err)
		return err
	}
	defer outputFile.Close()

	outputBuffer := bytes.NewBuffer(data)
	_, err = io.Copy(outputFile, outputBuffer)
	if err != nil {
		logrus.Errorf("Error saving output file: %v", err)
		return err
	}

	// Free up all memory
	defer C.free(unsafe.Pointer(ModelPath))
	defer C.free(unsafe.Pointer(ParamPath))
	defer C.free(unsafe.Pointer(InputImageFormat))
	defer C.delete_realesrgan(realesrgan) // Ensure resources are freed
	defer C.free_buffer(buffer)

	return nil
}

func RealESRGANInMemory(params RealESRGANParams, inputImage []byte, imageFormat ImageFormat) ([]byte, error) {
	// Convert Go slice to C array
	var (
		InputImagePtr = (*C.uchar)(unsafe.Pointer(&inputImage[0]))
		InputImageLen = C.size_t(len(inputImage))

		GPUID      = C.int(params.GPUID)
		TTA        = C.int(params.TTA)
		Scale      = C.int(params.Scale)
		TileSize   = C.int(params.TileSize)
		Prepadding = C.int(params.Prepadding)

		ModelPath = C.CString(params.ModelPath)
		ParamPath = C.CString(params.ParamPath)

		InputImageFormat = C.CString(imageFormat.string())
	)

	// Create RealESRGAN instance
	realesrgan := C.create_realesrgan_instance(GPUID, TTA, Scale, TileSize, Prepadding, ModelPath, ParamPath)
	buffer := C.process_image(InputImagePtr, InputImageLen, InputImageFormat, realesrgan, 4)

	// Access the returned buffer data
	data := C.GoBytes(unsafe.Pointer(buffer.data), C.int(buffer.size))

	// Free up all memory
	defer C.free(unsafe.Pointer(ModelPath))
	defer C.free(unsafe.Pointer(ParamPath))
	defer C.free(unsafe.Pointer(InputImageFormat))
	defer C.delete_realesrgan(realesrgan) // Ensure resources are freed
	defer C.free_buffer(buffer)

	return data, nil
}
