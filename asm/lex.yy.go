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
	case c == '%':
		goto yystate4
	case c == '-':
		goto yystate13
	case c == '0':
		goto yystate15
	case c == '\n':
		goto yystate3
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate2
	case c >= '1' && c <= '9':
		goto yystate14
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate18
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
	case c == 'w':
		goto yystate8
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
		goto yyabort
	case c == 'o':
		goto yystate9
	}

yystate9:
	c = l.getc()
	switch {
	default:
		goto yyabort
	case c == 'r':
		goto yystate10
	}

yystate10:
	c = l.getc()
	switch {
	default:
		goto yyabort
	case c == 'd':
		goto yystate11
	}

yystate11:
	c = l.getc()
	switch {
	default:
		goto yyabort
	case c == 's':
		goto yystate12
	}

yystate12:
	c = l.getc()
	goto yyrule4

yystate13:
	c = l.getc()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate14
	}

yystate14:
	c = l.getc()
	switch {
	default:
		goto yyrule5
	case c >= '0' && c <= '9':
		goto yystate14
	}

yystate15:
	c = l.getc()
	switch {
	default:
		goto yyrule5
	case c == 'x':
		goto yystate16
	case c >= '0' && c <= '9':
		goto yystate14
	}

yystate16:
	c = l.getc()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate17
	}

yystate17:
	c = l.getc()
	switch {
	default:
		goto yyrule5
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate17
	}

yystate18:
	c = l.getc()
	switch {
	default:
		goto yyrule7
	case c == ':':
		goto yystate19
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate18
	}

yystate19:
	c = l.getc()
	goto yyrule6

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
yyrule4: // %words
	{

		return WORDS
	}
yyrule5: // -?{DECIMAL}|{HEX}
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
yyrule6: // {ID}:
	{

		lval.str = string(l.buf[:len(l.buf)-1])
		return LABEL

	}
yyrule7: // {ID}
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
