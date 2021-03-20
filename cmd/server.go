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
	"io"
	"log"
	"net/http"
	"strconv"

	_ "embed"

	"github.com/spf13/cobra"

	"github.com/kjinho/pdftool/src/utils"
)

//go:embed assets/index.html
var indexFile string

var serverPort int

const MAX_UPLOAD_SIZE = 1024 * 1024 * 50 // 50MB

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprint(w, indexFile)
	//http.ServeFile(w, r, "assets/index.html")
}

func batesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Bates stamp processing.")
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "The uploaded file is too big. Please choose a file that's less than 50MB in size.", http.StatusBadRequest)
		return
	}

	// The argument to FormFile must match the name attribute
	// of the file input on the frontend
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filetype := http.DetectContentType(buff)
	if filetype != "application/pdf" {
		http.Error(w, "The provided file format is not allowed. Please upload a PDF.", http.StatusBadRequest)
		return
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("File: %s", file)
	prefix := r.FormValue("prefix")
	log.Printf("Prefix: %s", prefix)
	widthString := r.FormValue("width")
	width, err := strconv.ParseInt(widthString, 10, 0)
	if err != nil {
		width = 10
	}
	log.Printf("Width: %d", width)
	divider := r.FormValue("divider")
	log.Printf("Divider: %s", divider)
	startnoString := r.FormValue("startno")
	startno, err := strconv.ParseInt(startnoString, 10, 0)
	if err != nil {
		startno = 1
	}
	log.Printf("Start Number: %d", startno)

	w.Header().Add("Content-Type", "application/pdf")

	utils.BatesStampRS(file, w, utils.GenerateFmtString(prefix, divider, int(width)), startno)
}

func draftHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Draft stamp processing.")
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "The uploaded file is too big. Please choose a file that's less than 50MB in size.", http.StatusBadRequest)
		return
	}

	// The argument to FormFile must match the name attribute
	// of the file input on the frontend
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filetype := http.DetectContentType(buff)
	if filetype != "application/pdf" {
		http.Error(w, "The provided file format is not allowed. Please upload a PDF.", http.StatusBadRequest)
		return
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("File: %s", file)

	w.Header().Add("Content-Type", "application/pdf")

	utils.DraftStampRS(file, w)
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "an HTTP service to process PDF files",
	Long: `
server provides an HTTP service to process PDF files through
an HTML interface. Use a web browser to access the service. The
service processes the files in memory without saving anything
to disk, so there is a maximum file size of ` + fmt.Sprintf("%d", MAX_UPLOAD_SIZE) + ` bytes.
`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Starting server on http://%s:%d.\nPress ctrl-c to quit.\n", "localhost", serverPort)
		mux := http.NewServeMux()
		mux.HandleFunc("/", indexHandler)
		mux.HandleFunc("/bates", batesHandler)
		mux.HandleFunc("/draft", draftHandler)

		if err := http.ListenAndServe(fmt.Sprintf(":%d", serverPort), mux); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serverCmd.Flags().IntVarP(&serverPort, "port", "p", 8080, "port for access")
}
