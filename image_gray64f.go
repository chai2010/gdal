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

type imageGray64f struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

// newImageGray64f returns a new imageGray64f with the given bounds.
func newImageGray64f(r image.Rectangle) *imageGray64f {
	return new(imageGray64f).Init(make([]uint8, 8*r.Dx()*r.Dy()), 8*r.Dx(), r)
}

func (p *imageGray64f) Init(pix []uint8, stride int, rect image.Rectangle) *imageGray64f {
	*p = imageGray64f{
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

func (p *imageGray64f) Pix() []byte           { return p.M.Pix }
func (p *imageGray64f) Stride() int           { return p.M.Stride }
func (p *imageGray64f) Rect() image.Rectangle { return p.M.Rect }
func (p *imageGray64f) Channels() int         { return 1 }
func (p *imageGray64f) Depth() reflect.Kind   { return reflect.Float64 }

func (p *imageGray64f) ColorModel() color.Model { return colorGray64fModel }

func (p *imageGray64f) Bounds() image.Rectangle { return p.M.Rect }

func (p *imageGray64f) At(x, y int) color.Color {
	return p.Gray64fAt(x, y)
}

func (p *imageGray64f) Gray64fAt(x, y int) colorGray64f {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorGray64f{}
	}
	i := p.PixOffset(x, y)
	return pGray64fAt(p.M.Pix[i:])
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *imageGray64f) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*8
}

func (p *imageGray64f) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := colorGray64fModel.Convert(c).(colorGray64f)
	pSetGray64f(p.M.Pix[i:], c1)
	return
}

func (p *imageGray64f) SetGray64f(x, y int, c colorGray64f) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	pSetGray64f(p.M.Pix[i:], c)
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *imageGray64f) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &imageGray64f{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(imageGray64f).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *imageGray64f) Opaque() bool {
	return true
}
