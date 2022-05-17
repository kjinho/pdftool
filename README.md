# pdftool

```pdftool``` provides tools to process PDFs. Many of these tools
are helpful for lawyers that work with PDFs.

## Installation

### From Source

    go install github.com/kjinho/pdftool@latest

If you wish to set the version number in the binary to 'v0.1.2', for
example, build from source with the following linker flags:

    go build -ldflags="-X 'github.com/kjinho/pdftool/cmd.VersionNumber=v0.1.2'"

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
