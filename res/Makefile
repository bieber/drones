all: FreeSansBold.ttf.go
	gofmt -w .

FreeSansBold.ttf.go: FreeSansBold.ttf
	go-bindata -f FreeSansBold -i FreeSansBold.ttf -m -p res
