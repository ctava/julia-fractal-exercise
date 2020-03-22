package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"sync"
	"time"
)

type complexNumber struct {
	real      float64
	imaginary float64
}

func (c complexNumber) abs() float64 {
	return math.Sqrt((c.real * c.real) + (c.imaginary * c.imaginary))
}

func (c complexNumber) exp() (c1 complexNumber) {
	c1.real = math.Exp(c.real) * math.Cos(c.imaginary)
	c1.imaginary = math.Exp(c.real) * math.Sin(c.imaginary)
	return c1
}

func (c complexNumber) mul(c1 complexNumber) (c2 complexNumber) {
	c2.real = (c.real * c1.real) - (c.imaginary * c1.imaginary)
	c2.imaginary = (c.real * c1.imaginary) + (c.imaginary * c1.real)
	return c2
}

func (c complexNumber) div(c1 complexNumber) (c2 complexNumber) {
	c2.real = ((c.real * c1.real) + (c.imaginary * c1.imaginary)) / ((c1.real * c1.real) + (c1.imaginary * c1.imaginary))
	c2.imaginary = ((c.imaginary * c1.real) - (c.real * c1.imaginary)) / ((c1.real * c1.real) + (c1.imaginary * c1.imaginary))
	return c2
}

func (c complexNumber) add(c1 complexNumber) (c2 complexNumber) {
	c2.real = c.real + c1.real
	c2.imaginary = c.imaginary + c1.imaginary
	return c2
}

type fractal struct {
	fractalImg      *image.NRGBA
	cReal           float64
	cImaginary      float64
	imgXCenter      float64
	imgYCenter      float64
	zoom            float64
	scale           float64
	width           int
	height          int
	routines        int
	maxIterations   int
	numberOfSamples int
}

func getNewFractal(imageConstants []int, fractalProcessingConstants []int, imageProcessingConstants []float64) fractal {

	newFractal := fractal{
		width:           imageConstants[0],
		height:          imageConstants[1],
		routines:        fractalProcessingConstants[0],
		maxIterations:   fractalProcessingConstants[1],
		numberOfSamples: fractalProcessingConstants[2],
		cReal:           imageProcessingConstants[0],
		cImaginary:      imageProcessingConstants[1],
		imgXCenter:      imageProcessingConstants[2],
		imgYCenter:      imageProcessingConstants[3],
		zoom:            imageProcessingConstants[4],
		fractalImg:      image.NewNRGBA(image.Rect(0, 0, imageConstants[0], imageConstants[1])),
	}

	//Pick the smaller of the two dimensions (width and height)
	//and use that length in pixels as the scale for both axis
	if newFractal.width < newFractal.height {
		newFractal.scale = float64(newFractal.width)
	} else {
		newFractal.scale = float64(newFractal.height)
	}

	return newFractal
}

func (f fractal) simpleGreyscale(iterations int) (R, G, B, A float64) {
	max := 255.0
	color := float64(max*float64(iterations)) / float64(f.maxIterations)
	return color, color, color, max
}

func (f fractal) pixelToCoordinate(xPixel, yPixel float64) (xCoordinate, yCoordinate float64) {
	xCoordinate = ((xPixel - (float64(f.width) / 2)) * ((2 / f.scale) / f.zoom)) + f.imgXCenter
	yCoordinate = ((yPixel - (float64(f.height) / 2)) * ((2 / f.scale) / f.zoom)) + f.imgYCenter
	return xCoordinate, yCoordinate
}

func (f fractal) renderFractal(wg *sync.WaitGroup, routineNumber int, samples int, pointFunc func(float64, float64, int) (R, G, B, A float64)) {

	offsets := make([]float64, samples)

	for sample := 0; sample < samples; sample++ {
		offsets[sample] = (1 + float64(2*sample) - float64(samples)) / float64(2*(samples))
	}
	samplesSquared := float64(samples * samples)

	routines := f.routines
	size := f.width * f.height
	for i := routineNumber; i < size; i = i + routines {
		xPixel := i % f.width
		yPixel := i / f.width

		R, G, B, A := 0.0, 0.0, 0.0, 0.0

		for xSample := 0; xSample < samples; xSample++ {
			for ySample := 0; ySample < samples; ySample++ {
				xCoordinate, yCoordinate := f.pixelToCoordinate(float64(xPixel)+offsets[xSample], float64(yPixel)+offsets[ySample])

				r, g, b, a := pointFunc(xCoordinate, yCoordinate, f.maxIterations)

				R, G, B, A = R+(r/samplesSquared),
					G+(g/samplesSquared),
					B+(b/samplesSquared),
					A+(a/samplesSquared)
			}
		}
		f.fractalImg.Set(xPixel, yPixel, color.RGBA{uint8(R), uint8(G), uint8(B), uint8(A)})
	}

	wg.Done()
}

func (f fractal) julia() func(float64, float64, int) (R, G, B, A float64) {

	getJuliaPointFunc := func(xCoordinate, yCoordinate float64, maxIterations int) (R, G, B, A float64) {
		c := complexNumber{f.cReal, f.cImaginary}
		z := complexNumber{xCoordinate, yCoordinate}

		iterations := 0
		for iterations = 0; z.mul(z).add(c).abs() <= 2 && iterations < f.maxIterations; iterations++ {
			z = z.mul(z).add(c)
		}

		return f.simpleGreyscale(iterations)
	}
	return getJuliaPointFunc
}

func (f fractal) generate(pointFunc func(float64, float64, int) (R, G, B, A float64)) {
	var wg sync.WaitGroup
	wg.Add(f.routines)

	for routine := 0; routine < f.routines; routine++ {
		go f.renderFractal(&wg, routine, f.numberOfSamples, pointFunc)
	}

	wg.Wait()
}

func main() {

	imgWidth := 2600
	imgHeight := 2000
	imageConstants := []int{imgWidth, imgHeight}

	goRoutines := 8
	maxIterations := 512
	noOfSamples := 2
	fractalProcessingConstants := []int{goRoutines, maxIterations, noOfSamples}

	//1
	//cReal := -0.79
	//cImaginary := 0.15

	//2
	//cReal := -0.162
	//cImaginary := 1.04

	//3
	//cReal := 0.3
	//cImaginary := -0.01

	//4
	//cReal := -1.476
	//cImaginary := 0.0

	//5
	//cReal := -0.12
	//cImaginary := -0.77

	//6
	cReal := 0.28
	cImaginary := 0.008

	imgXCenter := 0.0
	imgYCenter := -0.0
	zoom := 1.5
	imageProcessingConstants := []float64{cReal, cImaginary, imgXCenter, imgYCenter, zoom}
	f := getNewFractal(imageConstants, fractalProcessingConstants, imageProcessingConstants)

	startTime := time.Now()
	f.generate(f.julia())
	duration := time.Since(startTime)
	fmt.Println("\nTime taken:", duration)

	outputFileName := "julia.png"
	newFile, _ := os.Create(outputFileName)
	png.Encode(newFile, f.fractalImg)
}
