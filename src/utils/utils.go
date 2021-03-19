/*
Package utils provides utilities usable in the other functions.

Copyright © 2021 Jin-Ho King <j@kingesq.us>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package utils

import (
	"fmt"
	"io"
	"log"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

// BatesStampRS adds a bates stamp to each page of rs and writes to w
func BatesStampRS(rs io.ReadSeeker, w io.Writer, fmtString string, startno int64) error {
	_, err := rs.Seek(0, io.SeekStart)
	if err != nil {
		log.Printf("Error seeking to beginning of file\n%s", err)
		return err
	}

	pageCount, err := api.PageCount(rs, nil)
	if err != nil {
		return err
	}
	m := map[int]*pdfcpu.Watermark{}

	for i := 0; i < pageCount; i++ {
		text := fmt.Sprintf(fmtString, startno+int64(i))
		fontName := "Helvetica"
		points := 12
		pos := "br"
		rot := 0
		ma := "2"
		fillc := "#000000"
		offset := "-5 5"
		scale := "1 abs"
		bgcolor := "#ffffff"
		opacity := "1"
		border := "1 #000000"
		desc := fmt.Sprintf(
			"font:%s, points:%d, scale:%s, pos:%s, rot:%d, ma:%s, fillc:%s, offset:%s, bgcolor:%s, op:%s, border:%s",
			fontName,
			points,
			scale,
			pos,
			rot,
			ma,
			fillc,
			offset,
			bgcolor,
			opacity,
			border,
		)

		wm, err := api.TextWatermark(text, desc, true, false, pdfcpu.POINTS)
		if err != nil {
			return err
		}
		m[i+1] = wm // PDF page numbering starts at 1
	}

	_, err = rs.Seek(0, io.SeekStart)
	if err != nil {
		log.Printf("Error seeking to beginning of file\n%s", err)
		return err
	}

	if err := api.AddWatermarksMap(rs, w, m, nil); err != nil {
		return err
	}
	return nil
}

// BatesStamp adds a bates stamp to each page of inFile and writes to outFile
func BatesStamp(inFile string, outFile string, fmtString string, startno int64) error {

	pageCount, err := api.PageCountFile(inFile)
	if err != nil {
		return err
	}

	m := map[int]*pdfcpu.Watermark{}

	for i := 0; i < pageCount; i++ {
		text := fmt.Sprintf(fmtString, startno+int64(i))
		fontName := "Helvetica"
		points := 12
		pos := "br"
		rot := 0
		ma := "2"
		fillc := "#000000"
		offset := "-5 5"
		scale := "1 abs"
		bgcolor := "#ffffff"
		opacity := "1"
		border := "1 #000000"
		desc := fmt.Sprintf(
			"font:%s, points:%d, scale:%s, pos:%s, rot:%d, ma:%s, fillc:%s, offset:%s, bgcolor:%s, op:%s, border:%s",
			fontName,
			points,
			scale,
			pos,
			rot,
			ma,
			fillc,
			offset,
			bgcolor,
			opacity,
			border,
		)

		wm, err := api.TextWatermark(text, desc, true, false, pdfcpu.POINTS)
		if err != nil {
			return err
		}
		m[i+1] = wm // PDF page numbering starts at 1
	}

	if err := api.AddWatermarksMapFile(inFile, outFile, m, nil); err != nil {
		return err
	}
	return nil
}

func GenerateFmtString(prefix string, separator string, padding int) string {
	return fmt.Sprintf("%s%s%s0%dd", prefix, separator, "%", padding)
}
