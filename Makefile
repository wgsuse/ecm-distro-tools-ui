demo.exe: clean
	go build -o demo.exe -ldflags "-H windowsgui" demo.go attr_windows.go

clean:
	rm -f demo.exe
