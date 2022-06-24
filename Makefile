.PHONY: clean

# We are on windows
ifdef OS
   EXT = .exe
   LDFLAGS = -H windowsgui
endif

GO         = go
BINDIR    := bin
BINARY    := ecm-distro-tools-ui$(EXT)

# disable warnings
.EXPORT_ALL_VARIABLES:
CGO_CFLAGS=$(shell $(GO) env CGO_CFLAGS) -w

$(BINDIR)/$(BINARY): $(shell ls *.go)
	$(GO) build -o $(BINDIR)/$(BINARY) -ldflags "$(LDFLAGS)" .

clean:
	rm -f $(BINDIR)/$(BINARY)
