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

type imageRGB96f struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

// newImageRGB96f returns a new imageRGB96f with the given bounds.
func newImageRGB96f(r image.Rectangle) *imageRGB96f {
	return new(imageRGB96f).Init(make([]uint8, 12*r.Dx()*r.Dy()), 12*r.Dx(), r)
}

func (p *imageRGB96f) Init(pix []uint8, stride int, rect image.Rectangle) *imageRGB96f {
	*p = imageRGB96f{
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

func (p *imageRGB96f) Pix() []byte           { return p.M.Pix }
func (p *imageRGB96f) Stride() int           { return p.M.Stride }
func (p *imageRGB96f) Rect() image.Rectangle { return p.M.Rect }
func (p *imageRGB96f) Channels() int         { return 3 }
func (p *imageRGB96f) Depth() reflect.Kind   { return reflect.Float32 }

func (p *imageRGB96f) ColorModel() color.Model { return colorRGB96fModel }

func (p *imageRGB96f) Bounds() image.Rectangle { return p.M.Rect }

func (p *imageRGB96f) At(x, y int) color.Color {
	return p.RGB96fAt(x, y)
}

func (p *imageRGB96f) RGB96fAt(x, y int) colorRGB96f {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorRGB96f{}
	}
	i := p.PixOffset(x, y)
	return pRGB96fAt(p.M.Pix[i:])
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *imageRGB96f) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*12
}

func (p *imageRGB96f) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := colorRGB96fModel.Convert(c).(colorRGB96f)
	pSetRGB96f(p.M.Pix[i:], c1)
	return
}

func (p *imageRGB96f) SetRGB96f(x, y int, c colorRGB96f) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	pSetRGB96f(p.M.Pix[i:], c)
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *imageRGB96f) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &imageRGB96f{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(imageRGB96f).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *imageRGB96f) Opaque() bool {
	return true
}
