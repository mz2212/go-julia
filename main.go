package main

import (
  "github.com/fogleman/gg"
  "github.com/mz2212/go_julia/hsvrgb"
)

const width, height = 1920, 1080

func main() {
  canvas := gg.NewContext(width, height)
  maxIter := 300
  var constReal, constImag float64 = -0.7, 0.27015
  var zoom, moveX, moveY float64 = 1, 0, 0

  for x := 0; x <= width; x++ {
    for y := 0; y <= height; y++ {
      iter := crunch(constReal, constImag, maxIter, x, y, zoom, moveX, moveY)
      paint(x, y, iter, maxIter, canvas)
    }
  }

  canvas.SavePNG("test.png")
}

func paint(x, y, iter, maxIter int, c *gg.Context) {
  color := hsvrgb.Hsv2Rgb(hsvrgb.HsvColor{iter % 256, 255,
    255 * bool2int(iter < maxIter)})
  c.SetRGB255(color.R, color.G, color.B)
  c.SetPixel(x, y)
}

func crunch(constReal, constImag float64, maxIter, x, y int, zoom, moveX, moveY float64) int {
  var oldReal, oldImag float64
  var newReal, newImag float64 = (1.5 * (float64(x) - width / 2.0) / (0.5 * zoom * width) + moveX), ((float64(y) - height / 2.0) / (0.5 * zoom * height) + moveY)
  var i int
  for i = 0; i < maxIter; i++ {
    oldReal, oldImag = newReal, newImag

    newReal = oldReal * oldReal - oldImag * oldImag + constReal;
    newImag = 2 * oldReal * oldImag + constImag;
    if (newReal * newReal + newImag * newImag) > 4 {break}
  }
  return i
}

func bool2int(b bool) int {
  if b {
    return 1
  }
  return 0
}
