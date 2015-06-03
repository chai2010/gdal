// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Auto Generated By 'go generate', DONOT EDIT!!!

package gdal

import (
	"image"
	"image/color"
	"reflect"
)

type imageRGBA128i struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

// newImageRGBA128i returns a new imageRGBA128i with the given bounds.
func newImageRGBA128i(r image.Rectangle) *imageRGBA128i {
	return new(imageRGBA128i).Init(make([]uint8, 16*r.Dx()*r.Dy()), 16*r.Dx(), r)
}

func (p *imageRGBA128i) Init(pix []uint8, stride int, rect image.Rectangle) *imageRGBA128i {
	*p = imageRGBA128i{
		M: struct {
			Pix    []uint8
			Stride int
			Rect   image.Rectangle
		}{
			Pix:    pix,
			Stride: stride,
			Rect:   rect,
		},
	}
	return p
}

func (p *imageRGBA128i) Pix() []byte           { return p.M.Pix }
func (p *imageRGBA128i) Stride() int           { return p.M.Stride }
func (p *imageRGBA128i) Rect() image.Rectangle { return p.M.Rect }
func (p *imageRGBA128i) Channels() int         { return 4 }
func (p *imageRGBA128i) Depth() reflect.Kind   { return reflect.Int32 }

func (p *imageRGBA128i) ColorModel() color.Model { return colorRGBA128iModel }

func (p *imageRGBA128i) Bounds() image.Rectangle { return p.M.Rect }

func (p *imageRGBA128i) At(x, y int) color.Color {
	return p.RGBA128iAt(x, y)
}

func (p *imageRGBA128i) RGBA128iAt(x, y int) colorRGBA128i {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorRGBA128i{}
	}
	i := p.PixOffset(x, y)
	return pRGBA128iAt(p.M.Pix[i:])
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *imageRGBA128i) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*16
}

func (p *imageRGBA128i) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := colorRGBA128iModel.Convert(c).(colorRGBA128i)
	pSetRGBA128i(p.M.Pix[i:], c1)
	return
}

func (p *imageRGBA128i) SetRGBA128i(x, y int, c colorRGBA128i) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	pSetRGBA128i(p.M.Pix[i:], c)
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *imageRGBA128i) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &imageRGBA128i{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(imageRGBA128i).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *imageRGBA128i) Opaque() bool {
	if p.M.Rect.Empty() {
		return true
	}
	for y := p.M.Rect.Min.Y; y < p.M.Rect.Max.Y; y++ {
		for x := p.M.Rect.Min.X; x < p.M.Rect.Max.X; x++ {
			if _, _, _, a := p.At(x, y).RGBA(); a == 0xFFFF {
				return false
			}
		}
	}
	return true
}
