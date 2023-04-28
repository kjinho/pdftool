VERSION := "v0.1.3"
OUTPUT := pdftool
INSTALL_DIR := ~/.local/bin

OTHER_FILES := Makefile
MOD_FILES := go.mod go.sum
SRC_FILES := src/utils/utils.go
CMD_FILES := cmd/bates.go cmd/copy.go cmd/draft.go \
cmd/root.go cmd/server.go cmd/utils.go cmd/version.go \
cmd/assets/index.html cmd/assets/normalize.css \
cmd/assets/skeleton.css

.PHONY: clean all update

all: $(OUTPUT)

$(OUTPUT): main.go $(OTHER_FILES) $(MOD_FILES) $(SRC_FILES) $(CMD_FILES)
	go build -o $@ -ldflags="-X 'github.com/kjinho/pdftool/cmd.VersionNumber=$(VERSION)'"

clean:
	rm -rf $(OUTPUT)

update:
	go get -u ./...
	go mod tidy

install: $(OUTPUT)
	install -m 755 $(OUTPUT) $(INSTALL_DIR)
