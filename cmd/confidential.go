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
	"github.com/spf13/cobra"

	"github.com/kjinho/pdftool/src/utils"
)

var confidentialFilenameSuffix string

// draftCmd represents the draft command
var confidentialCmd = &cobra.Command{
	Use:   "confidential inFile1 ...",
	Short: "Add a `CONFIDENTIAL` watermark",
	Long: `
confidential adds a "CONFIDENTIAL" watermark to each page of the PDF.

By default, the output filename is given the suffix "-CONFIDENTIAL"`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		nargs := len(args)

		for i := 0; i < nargs; i++ {
			_, err := os.Stat(args[i])
			if err != nil {
				log.Fatalf("inFile `%s` does not exist", args[i])
			}
			_, err = api.PageCountFile(args[i])
			if err != nil {
				log.Fatalf("error with inFile `%s`: %s", args[i], err)
			}

			newFilename := generateNewFilename(args[i], confidentialFilenameSuffix)

			_, err = os.Stat(newFilename)
			if !Overwrite && err == nil {
				log.Fatalf("outFile `%s` already exists. To overwrite, use --force", newFilename)
			}

			fOut, err := os.Create(newFilename)
			if err != nil {
				log.Fatalf("Error creating file `%s`\n%s\n", newFilename, err)
			}
			defer fOut.Close()
			fIn, err := os.Open(args[i])
			if err != nil {
				log.Fatalf("Error opening file `%s`\n%s\n", args[i], err)
			}
			defer fIn.Close()
			err = utils.ConfidentialStampRS(fIn, fOut)
			if err != nil {
				log.Fatalf("Error stamping `%s`", args[i])
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(confidentialCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// draftCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// draftCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	confidentialCmd.Flags().BoolVarP(&Overwrite, "force", "f", false, "overwrite the output file (default: error on existing output file)")
	confidentialCmd.Flags().StringVar(&confidentialFilenameSuffix, "suffix", "-CONFIDENTIAL", "output filename suffix")
}
