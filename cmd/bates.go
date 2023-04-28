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
	"fmt"
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/spf13/cobra"

	"github.com/kjinho/pdftool/src/utils"
)

var prefix string
var separator string
var buffer int
var startNo int64

// batesCmd represents the bates command
var batesCmd = &cobra.Command{
	Use:   "bates inFile1 ...",
	Short: "Bates stamp PDF files",
	Long: `
bates provides a tool to Bates stamp PDFs. Mandatory arguments are the 
list of inFiles (input PDFs) that should be Bates stamped.

Optional flags are available to modify the content of the Bates stamp 
(e.g., prefix, separator, number width, starting number).

For example, given a 10-page PDF infile.pdf,

  $ pdftool bates infile.pdf -p ABCD -s _ -n 101

takes infile.pdf and writes a new PDF with bates numbers starting with
ABCD_0000000101 on the first page of the PDF. The output filename will
be infile-ABCD_0000000101-ABCD_0000000110.pdf.
  `,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		nargs := len(args)
		xstartNo := startNo
		for i := 0; i < nargs; i++ {
			_, err := os.Stat(args[i])
			if err != nil {
				log.Fatalf("inFile `%s` does not exist", args[i])
			}
			pageCount, err := api.PageCountFile(args[i])
			if err != nil {
				log.Fatalf("error with inFile `%s`: %s", args[i], err)
			}
			fmtString := utils.GenerateFmtString(prefix, separator, buffer)
			startBates := fmt.Sprintf(fmtString, xstartNo)
			stopBates := fmt.Sprintf(fmtString, xstartNo+int64(pageCount)-1)

			newFilename := generateNewFilename(args[i], "-"+startBates+"-"+stopBates)
			_, err = os.Stat(newFilename)
			if !Overwrite && err == nil {
				log.Fatalf("outFile `%s` already exists. To overwrite, use --force", newFilename)
			}
			log.Printf(
				"Performing bates numbering\nInput:\t%s\nOutput:\t%s\nStart:\t%s\nStop:\t%s\n",
				args[i],
				newFilename,
				startBates,
				stopBates,
			)

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
			utils.BatesStampRS(fIn, fOut, fmtString, xstartNo)

			xstartNo += int64(pageCount)
		}
	},
}

func init() {
	rootCmd.AddCommand(batesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// batesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// batesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	batesCmd.Flags().StringVarP(&prefix, "prefix", "p", "Bates", "bates numbering prefix")
	batesCmd.Flags().StringVarP(&separator, "separator", "s", "-", "separator")
	batesCmd.Flags().IntVarP(&buffer, "width", "w", 8, "number of characters for number")
	batesCmd.Flags().Int64VarP(&startNo, "number", "n", 1, "number to start on")
	batesCmd.Flags().BoolVarP(&Overwrite, "force", "f", false, "overwrite the output file (default: error on existing output file)")
}
