all: FreeSansBold.ttf.go FreeSerifBold.ttf.go
	gofmt -w .

FreeSansBold.ttf.go: FreeSansBold.ttf
	go-bindata -f FreeSansBold -i FreeSansBold.ttf -m -p res

FreeSerifBold.ttf.go: FreeSerifBold.ttf
	go-bindata -f FreeSerifBold -i FreeSerifBold.ttf -m -p res
