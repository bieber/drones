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

%{
package asm

import (
	"bufio"
	"fmt"
	"strconv"
)

type lexer struct {
	src 	*bufio.Reader
	buf		[]byte
	current	byte
	empty 	bool
	line	int
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
%}

%yyc c
%yyn c = l.getc()

DECIMAL	[0-9]+
HEX		0x[0-9a-fA-F]+
ID		[a-zA-Z][a-zA-Z0-9_]*

%%
	l.buf = l.buf[:0]

[ \t\r]+
	/* ignore */
\n
	return NEWLINE

%org
	return ORG

{DECIMAL}|{HEX}
	num, err := strconv.ParseInt(string(l.buf), 0, 16)
	if err == nil {
		lval.num = uint16(num)
		return NUM
	} else {
		return ILLEGAL
	}

{ID}:
	lval.str = string(l.buf)
	return LABEL
	
{ID}
	lval.str = string(l.buf)
	return IDENT
%%
	l.empty = true
	return int(c)
}

