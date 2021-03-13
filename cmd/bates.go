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

var overwrite bool
var prefix string
var separator string
var buffer int
var startNo int64

// batesCmd represents the bates command
var batesCmd = &cobra.Command{
	Use:   "bates inFile outFile",
	Short: "Bates stamp PDF files",
	Long: `
bates provides a tool to Bates stamp PDFs. Mandatory arguments are the 
inFile (input PDF) and the outFile (output PDF).

Optional flags are available to modify the content of the Bates stamp 
(e.g., prefix, separator, number width, starting number).

For example,

  $ pdftool bates infile.pdf outfile.pdf -p ABCD -s _ -n 101

takes infile.pdf and writes a new outfile.pdf with bates numbers
starting with ABCD_0000000101 on the first page of the PDF.
  `,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// i, err := strconv.ParseInt(args[2], 10, 64)
		// if err != nil {
		// 	log.Fatalf("argument `%s` is not an integer\nError: %s", args[2], err)
		// } else if i < 0 {
		// 	log.Fatalf("starting number `%d` is not a positive integer", i)
		// }
		_, err := os.Stat(args[0])
		if err != nil {
			log.Fatalf("inFile `%s` does not exist", args[0])
		}
		pageCount, err := api.PageCountFile(args[0])
		if err != nil {
			log.Fatalf("error with inFile `%s`: %s", args[0], err)
		}
		_, err = os.Stat(args[1])
		if !overwrite && err == nil {
			log.Fatalf("outFile `%s` already exists. To overwrite, use --force", args[1])
		}
		fmtString := utils.GenerateFmtString(prefix, separator, buffer)
		startBates := fmt.Sprintf(fmtString, startNo)
		stopBates := fmt.Sprintf(fmtString, startNo+int64(pageCount)-1)
		fmt.Printf(
			"Performing bates numbering\nInput:\t%s\nOutput:\t%s\nStart:\t%s\nStop:\t%s\n",
			args[0],
			args[1],
			startBates,
			stopBates,
		)
		utils.BatesStamp(args[0], args[1], fmtString, startNo)
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
	batesCmd.Flags().BoolVarP(&overwrite, "force", "f", false, "overwrite the output file (default: error on existing output file)")
	batesCmd.Flags().StringVarP(&prefix, "prefix", "p", "Bates", "bates numbering prefix")
	batesCmd.Flags().StringVarP(&separator, "separator", "s", "-", "separator")
	batesCmd.Flags().IntVarP(&buffer, "width", "w", 10, "number of characters for number")
	batesCmd.Flags().Int64VarP(&startNo, "number", "n", 1, "number to start on")
}
