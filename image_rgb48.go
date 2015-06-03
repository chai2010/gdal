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

type imageRGB48 struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

// newImageRGB48 returns a new imageRGB48 with the given bounds.
func newImageRGB48(r image.Rectangle) *imageRGB48 {
	return new(imageRGB48).Init(make([]uint8, 6*r.Dx()*r.Dy()), 6*r.Dx(), r)
}

func (p *imageRGB48) Init(pix []uint8, stride int, rect image.Rectangle) *imageRGB48 {
	*p = imageRGB48{
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

func (p *imageRGB48) Pix() []byte           { return p.M.Pix }
func (p *imageRGB48) Stride() int           { return p.M.Stride }
func (p *imageRGB48) Rect() image.Rectangle { return p.M.Rect }
func (p *imageRGB48) Channels() int         { return 3 }
func (p *imageRGB48) Depth() reflect.Kind   { return reflect.Uint16 }

func (p *imageRGB48) ColorModel() color.Model { return colorRGB48Model }

func (p *imageRGB48) Bounds() image.Rectangle { return p.M.Rect }

func (p *imageRGB48) At(x, y int) color.Color {
	return p.RGB48At(x, y)
}

func (p *imageRGB48) RGB48At(x, y int) colorRGB48 {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorRGB48{}
	}
	i := p.PixOffset(x, y)
	return pRGB48At(p.M.Pix[i:])
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *imageRGB48) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*6
}

func (p *imageRGB48) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := colorRGB48Model.Convert(c).(colorRGB48)
	pSetRGB48(p.M.Pix[i:], c1)
	return
}

func (p *imageRGB48) SetRGB48(x, y int, c colorRGB48) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	pSetRGB48(p.M.Pix[i:], c)
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *imageRGB48) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &imageRGB48{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(imageRGB48).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *imageRGB48) Opaque() bool {
	return true
}
