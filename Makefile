GO = go
ARGS = build
BUILD_DIR = .
BINARY = ecm-distro-tools-ui

# disable warnings
.EXPORT_ALL_VARIABLES:
CGO_CFLAGS=$(shell go env CGO_CFLAGS) -w

$(BINARY): main.go
	$(GO) $(ARGS) $(BUILD_DIR)

clean:
	rm -f $(BINARY)
