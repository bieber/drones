/*
 * This file is part of drones.
 *
 * drones is free software: you can redistribute it and/or modify it
 * under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * drones is distributed in the hope that it will be useful,
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with drones.  If not, see <http://www.gnu.org/licenses/>.
 */
package asm

import (
	"bufio"
	"fmt"
	"strconv"
)

type lexer struct {
	src     *bufio.Reader
	buf     []byte
	current byte
	empty   bool
	line    int
}

func newLexer(in *bufio.Reader) *lexer {
	l := &lexer{src: in, line: 1}
	if b, err := in.ReadByte(); err == nil {
		l.current = b
	}
	return l
}

func (l *lexer) getc() byte {
	if l.current != 0 {
		l.buf = append(l.buf, l.current)
		if l.current == '\n' {
			l.line++
		}
	}
	l.current = 0
	if b, err := l.src.ReadByte(); err == nil {
		l.current = b
	}
	return l.current
}

func (l *lexer) Error(e string) {
	errorCodes = append(errorCodes, fmt.Sprintf("Line %d: %s", l.line, e))
}

func (l *lexer) Lex(lval *yySymType) int {
	c := l.current
	if l.empty {
		c, l.empty = l.getc(), false
	}

yystate0:

	l.buf = l.buf[:0]

	goto yystart1

	goto yystate1 // silence unused label error
yystate1:
	c = l.getc()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '%':
		goto yystate4
	case c == '0':
		goto yystate8
	case c == '\n':
		goto yystate3
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate2
	case c >= '1' && c <= '9':
		goto yystate9
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate2:
	c = l.getc()
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate2
	}

yystate3:
	c = l.getc()
	goto yyrule2

yystate4:
	c = l.getc()
	switch {
	default:
		goto yyabort
	case c == 'o':
		goto yystate5
	}

yystate5:
	c = l.getc()
	switch {
	default:
		goto yyabort
	case c == 'r':
		goto yystate6
	}

yystate6:
	c = l.getc()
	switch {
	default:
		goto yyabort
	case c == 'g':
		goto yystate7
	}

yystate7:
	c = l.getc()
	goto yyrule3

yystate8:
	c = l.getc()
	switch {
	default:
		goto yyrule4
	case c == 'x':
		goto yystate10
	case c >= '0' && c <= '9':
		goto yystate9
	}

yystate9:
	c = l.getc()
	switch {
	default:
		goto yyrule4
	case c >= '0' && c <= '9':
		goto yystate9
	}

yystate10:
	c = l.getc()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate11
	}

yystate11:
	c = l.getc()
	switch {
	default:
		goto yyrule4
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate11
	}

yystate12:
	c = l.getc()
	switch {
	default:
		goto yyrule6
	case c == ':':
		goto yystate13
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate13:
	c = l.getc()
	goto yyrule5

yyrule1: // [ \t\r]+
	{

		/* ignore */
		goto yystate0
	}
yyrule2: // \n
	{

		return NEWLINE
	}
yyrule3: // %org
	{

		return ORG
	}
yyrule4: // {DECIMAL}|{HEX}
	{

		num, err := strconv.ParseInt(string(l.buf), 0, 16)
		if err == nil {
			lval.num = uint16(num)
			return NUM
		} else {
			return ILLEGAL
		}
		goto yystate0
	}
yyrule5: // {ID}:
	{

		lval.str = string(l.buf)
		return LABEL

	}
yyrule6: // {ID}
	{

		lval.str = string(l.buf)
		return IDENT
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	l.empty = true
	return int(c)
}
