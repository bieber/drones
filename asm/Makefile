# make all just builds the .go files, use go install to build the assembler
all: lex.yy.go y.go

lex.yy.go: lexer.l
	golex lexer.l
	gofmt -w lex.yy.go

y.go: parser.y
	go tool yacc parser.y
	gofmt -w y.go

clean:
	rm y.output y.go lex.yy.go

