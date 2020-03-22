# Introduction

The purpose of this repository is to provide an exercise for an entry-level programmer to learn [Control Flow](https://en.wikipedia.org/wiki/Control_flow) and [2D Computer Graphics](https://en.wikipedia.org/wiki/2D_computer_graphics) for a program that currently generates a [Julia Set](https://en.wikipedia.org/wiki/Julia_set) [Fractal](https://en.wikipedia.org/wiki/Fractal) based on a fixed number of [iteration](https://en.wikipedia.org/wiki/Iteration)s and a single [color space](https://en.wikipedia.org/wiki/Color_space).

The audience of this execise is people that have not programmed before but like [Math](https://en.wikipedia.org/wiki/Mathematics) more specifically [Geometry](https://en.wikipedia.org/wiki/Geometry), would like to dabble with [Computer Programming](https://en.wikipedia.org/wiki/Computer_programming) and to solve a problem that gives visual feedback as output.

# Fractals

In the most simpliest terms, Fractals appear the same at different depths or levels, as illustrated in successive magnifications of the [Mandelbrot set](https://en.wikipedia.org/wiki/Mandelbrot_set). Also, the "style" of this repeating detail depends on the region or section of the set being examined.

## [Mandelbrot set](https://en.wikipedia.org/wiki/Mandelbrot_set)

A mandelbrot set is created by treating real and imaginary parts of a [complex number](https://en.wikipedia.org/wiki/Complex_number) as image coordinates on a [complex plane](https://en.wikipedia.org/wiki/Complex_plane). In mathematics, the complex plane or z-plane is a geometric representation of the complex numbers established by the real axis and the perpendicular imaginary axis. Pixels may then be colored according to how soon the sequence crosses an arbitrarily chosen threshold, with a special color (usually black) used for the values of the complex number for which the sequence has not crossed the threshold after the predetermined number of iterations (this is necessary to clearly distinguish the Mandelbrot set image from the image of its complement).

Here is an image of the mandelbrot set:
![Mandelbrot set](https://upload.wikimedia.org/wikipedia/commons/thumb/2/21/Mandel_zoom_00_mandelbrot_set.jpg/1200px-Mandel_zoom_00_mandelbrot_set.jpg)

## Mandelbrot and Julia Sets

The Mandelbrot Set can be considered a map of all [Julia set](https://en.wikipedia.org/wiki/Julia_set)s because it uses a different complex number at each location, as if transforming from one Julia Set to another across space.

Here is a "map" of six Julia Sets and their corresponding locations in the Mandelbrot set from [https://www.karlsims.com/julia.html](https://www.karlsims.com/julia.html):
![Mandelbrot and Julia Sets](https://www.karlsims.com/julias-and-mandelbrot.jpg)

## Julia sets

Julia Sets consist of values that contain arbitrarily small perturbations to cause drastic changes in the sequence of complex numbers generated from functions. The behavior of Julia set functions is considered to be "chaotic".

Here is a picture of a Julia Set from imgur.com:

![Julia set](https://i.imgur.com/YrMem.png)

## This Program
This fractal program was adapted from [this repo](https://github.com/TheTeaCat/romanesgo).
This program runs 512 times to draw a Julia Set fractal in [grayscale](https://en.wikipedia.org/wiki/Grayscale). 

To run this program do the following:

0. [Install Golang](https://golang.org/doc/install)
1. Setup your `GOPATH`. On a mac, enter `PATH=$PATH:$GOPATH/bin` into your $HOME/.bashrc
2. Double check that you can run go programs by typing: `go version` from the command line.
   If you don't see the version, ensure that you have edited and sourced the .bashrc file into your current terminal:
   ```source $HOME/.bashrc```
3. Run the program
`go run main.go`
4. Review the resulting `julia.png`. It's really cool, right!!

## Your Exercise

You have two assignments:

1: Make the number of iterations dynamic such that the program can be run with any number of iterations. 
   For example: 5, 25, 50, 75, 100, 300, 500

2: In a sentance or two, describe the difference between the images generated at each scale of iteration.

## Bonus Exercises

B1: Provide the ability to change the color space from grayscle to color.

B2: In a sentance or two, describe the difference between the grayscale and color code space as its applied to in the fratal.