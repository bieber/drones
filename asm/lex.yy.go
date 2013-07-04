/*
 * Copyright 2013, Robert Bieber
 *
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
	"strconv"
)

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
	case c == '#':
		goto yystate4
	case c == '%':
		goto yystate6
	case c == '-':
		goto yystate15
	case c == '0':
		goto yystate17
	case c == '\n':
		goto yystate3
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate2
	case c >= '1' && c <= '9':
		goto yystate16
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate20
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
	goto yyrule3

yystate4:
	c = l.getc()
	switch {
	default:
		goto yyabort
	case c == '\n':
		goto yystate5
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate4
	}

yystate5:
	c = l.getc()
	goto yyrule2

yystate6:
	c = l.getc()
	switch {
	default:
		goto yyabort
	case c == 'o':
		goto yystate7
	case c == 'w':
		goto yystate10
	}

yystate7:
	c = l.getc()
	switch {
	default:
		goto yyabort
	case c == 'r':
		goto yystate8
	}

yystate8:
	c = l.getc()
	switch {
	default:
		goto yyabort
	case c == 'g':
		goto yystate9
	}

yystate9:
	c = l.getc()
	goto yyrule4

yystate10:
	c = l.getc()
	switch {
	default:
		goto yyabort
	case c == 'o':
		goto yystate11
	}

yystate11:
	c = l.getc()
	switch {
	default:
		goto yyabort
	case c == 'r':
		goto yystate12
	}

yystate12:
	c = l.getc()
	switch {
	default:
		goto yyabort
	case c == 'd':
		goto yystate13
	}

yystate13:
	c = l.getc()
	switch {
	default:
		goto yyabort
	case c == 's':
		goto yystate14
	}

yystate14:
	c = l.getc()
	goto yyrule5

yystate15:
	c = l.getc()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate16
	}

yystate16:
	c = l.getc()
	switch {
	default:
		goto yyrule6
	case c >= '0' && c <= '9':
		goto yystate16
	}

yystate17:
	c = l.getc()
	switch {
	default:
		goto yyrule6
	case c == 'x':
		goto yystate18
	case c >= '0' && c <= '9':
		goto yystate16
	}

yystate18:
	c = l.getc()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate19
	}

yystate19:
	c = l.getc()
	switch {
	default:
		goto yyrule6
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate19
	}

yystate20:
	c = l.getc()
	switch {
	default:
		goto yyrule8
	case c == ':':
		goto yystate21
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate21:
	c = l.getc()
	goto yyrule7

yyrule1: // [ \t\r]+
	{

		/* ignore */
		goto yystate0
	}
yyrule2: // #[^\n]*\n
	{
		/* comment */
		return NEWLINE
	}
yyrule3: // \n
	{

		return NEWLINE
	}
yyrule4: // %org
	{

		return ORG

	}
yyrule5: // %words
	{

		return WORDS
	}
yyrule6: // -?{DECIMAL}|{HEX}
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
yyrule7: // {ID}:
	{

		lval.str = string(l.buf[:len(l.buf)-1])
		return LABEL

	}
yyrule8: // {ID}
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
