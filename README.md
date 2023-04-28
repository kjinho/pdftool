# pdftool

```pdftool``` provides tools to process PDFs. Many of these tools
are helpful for lawyers that work with PDFs.

## Installation

### From Source

    go install github.com/kjinho/pdftool@latest

Alternatively, you may use the provided Makefile to compile and install
pdftool. Remember to set the variables at the top of the Makefile
appropriately first (default is to install to `~/.local/bin`).

    git clone https://github.com/kjinho/pdftool
    cd pdftool
    make install 

## Usage

    pdftool [command]

Available Commands:

    bates       Bates stamp PDF files
    completion  Generate the autocompletion script for the specified shell
    copy        Add a `COPY` watermark
    draft       Add a `DRAFT` watermark
    help        Help about any command
    server      an HTTP service to process PDF files

Flags:

        --config string   config file (default is $HOME/.pdftool.yaml)
    -h, --help            help for pdftool

Use ```pdftool [command] --help``` for more information about a command.
