/*
Copyright 2020 The pdfcpu Authors.

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

package test

import (
	"path/filepath"
	"testing"

	"github.com/vharish836/pdfcpu/pkg/api"
	"github.com/vharish836/pdfcpu/pkg/cli"
	"github.com/vharish836/pdfcpu/pkg/pdfcpu"
)

func testGrid(t *testing.T, msg string, inFiles []string, outFile string, selectedPages []string, desc string, rows, cols int, isImg bool) {
	t.Helper()

	var (
		nup *pdfcpu.NUp
		err error
	)

	if isImg {
		if nup, err = api.ImageGrid(rows, cols, desc); err != nil {
			t.Fatalf("%s %s: %v\n", msg, outFile, err)
		}
	} else {
		if nup, err = api.PDFGrid(rows, cols, desc); err != nil {
			t.Fatalf("%s %s: %v\n", msg, outFile, err)
		}
	}

	cmd := cli.NUpCommand(inFiles, outFile, selectedPages, nup, nil)
	if _, err := cli.Process(cmd); err != nil {
		t.Fatalf("%s %s: %v\n", msg, outFile, err)
	}

	if err := validateFile(t, outFile, nil); err != nil {
		t.Fatalf("%s: %v\n", msg, err)
	}
}

func TestGridCommand(t *testing.T) {
	for _, tt := range []struct {
		msg           string
		inFiles       []string
		outFile       string
		selectedPages []string
		desc          string
		rows, cols    int
		isImg         bool
	}{
		{"TestGridFromPDF",
			[]string{filepath.Join(inDir, "Acroforms2.pdf")},
			filepath.Join(outDir, "testGridFromPDF.pdf"),
			nil, "f:LegalL", 1, 3, false},

		{"TestGridFromImages",
			[]string{
				filepath.Join(resDir, "pdfchip3.png"),
				filepath.Join(resDir, "demo.png"),
				filepath.Join(resDir, "snow.jpg"),
			},
			filepath.Join(outDir, "testGridFromImages.pdf"),
			nil, "d:500 500, m:20, b:off", 1, 3, true},
	} {
		testGrid(t, tt.msg, tt.inFiles, tt.outFile, tt.selectedPages, tt.desc, tt.rows, tt.cols, tt.isImg)
	}
}
