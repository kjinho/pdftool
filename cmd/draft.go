/*
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
package cmd

import (
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/spf13/cobra"
)

// draftCmd represents the draft command
var draftCmd = &cobra.Command{
	Use:   "draft inFile outFile",
	Short: "Add a `DRAFT` watermark",
	Long: `
draft adds a "DRAFT" watermark to each page of the PDF.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat(args[0])
		if err != nil {
			log.Fatalf("inFile `%s` does not exist", args[0])
		}
		_, err = api.PageCountFile(args[0])
		if err != nil {
			log.Fatalf("error with inFile `%s`: %s", args[0], err)
		}
		_, err = os.Stat(args[1])
		if !Overwrite && err == nil {
			log.Fatalf("outFile `%s` already exists. To overwrite, use --force", args[1])
		}
		pages, err := api.ParsePageSelection("")
		if err != nil {
			log.Fatalf("error: %s", err)
		}
		err = api.AddTextWatermarksFile(
			args[0],
			args[1],
			pages,
			true,
			"DRAFT",
			"points:48, scale:1, op:0.2",
			pdfcpu.NewDefaultConfiguration())
		if err != nil {
			log.Fatalf("error: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(draftCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// draftCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// draftCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	draftCmd.Flags().BoolVarP(&Overwrite, "force", "f", false, "overwrite the output file (default: error on existing output file)")
}