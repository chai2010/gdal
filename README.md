- *赞助 BTC: 1Cbd6oGAUUyBi7X7MaR4np4nTmQZXVgkCW*
- *赞助 ETH: 0x623A3C3a72186A6336C79b18Ac1eD36e1c71A8a6*
- *Go语言付费QQ群: 1055927514*

----

# Go bindings for GDAL

[![Build Status](https://travis-ci.org/chai2010/gdal.svg)](https://travis-ci.org/chai2010/gdal)
[![GoDoc](https://godoc.org/github.com/chai2010/gdal?status.svg)](https://godoc.org/github.com/chai2010/gdal)

**Notes: Need Go1.5+!**

## Install

Install `GCC` or `MinGW` (http://tdm-gcc.tdragon.net/download) at first,
and then run these commands:

1. `go get github.com/chai2010/gdal`
2. `go run hello.go`

For windows-64bit user:

1. run `go generate github.com/chai2010/gdal` copy the `gdal-cgo-win64.dll` to `$(GOPATH)/bin`

For windows-32bit user:

1. Please build `gdal-cgo-win32.dll` and `./build-windows/lib/libgdal-cgo-386.a` first!

## Example

```Go
// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"

	"github.com/chai2010/gdal"
)

func main() {
	fmt.Printf("GDAL %d.%d.%d\n", gdal.MajorVersion, gdal.MinorVersion, gdal.RevVersion)

	// load data
	m, err := gdal.Load("./testdata/lena512color.png")
	if err != nil {
		log.Fatal("gdal.Load:", err)
	}

	// save bmp
	err = gdal.Save("output.bmp", m, nil)
	if err != nil {
		log.Fatal("gdal.Save:", err)
	}

	// save tiff
	err = gdal.Save("output.tiff", m, nil)
	if err != nil {
		log.Fatal("gdal.Save:", err)
	}

	// save jpeg-tiff data
	err = gdal.Save("output.jpeg.tiff", m, &gdal.Options{
		DriverName: "GTiff",
		ExtOptions: map[string]string{
			"COMPRESS":     "JPEG",
			"JPEG_QUALITY": "75",
		},
	})
	if err != nil {
		log.Fatal("gdal.Save:", err)
	}

	fmt.Println("Done.")
}
```

BUGS
====

Report bugs to <chaishushan@gmail.com>.

Thanks!
