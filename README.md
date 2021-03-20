# pdftool

```pdftool``` provides tools to process PDFs. Many of these tools
are helpful for lawyers that work with PDFs.

## Installation

### From Source

    go install github.com/kjinho/pdftool@latest

## Usage

    pdftool [command]

Available Commands:

    bates       Bates stamp PDF files
    draft       Add a `DRAFT` watermark
    help        Help about any command
    server      an HTTP service to process PDF files

Flags:

        --config string   config file (default is $HOME/.pdftool.yaml)
    -h, --help            help for pdftool

Use ```pdftool [command] --help``` for more information about a command.
