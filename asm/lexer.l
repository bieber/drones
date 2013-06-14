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
	"strconv"
)

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

#[^\n]*\n	/* comment */
	return NEWLINE

\n
	return NEWLINE

%org
	return ORG
	
%words
	return WORDS

-?{DECIMAL}|{HEX}
	num, err := strconv.ParseInt(string(l.buf), 0, 16)
	if err == nil {
		lval.num = uint16(num)
		return NUM
	} else {
		return ILLEGAL
	}

{ID}:
	lval.str = string(l.buf[:len(l.buf) - 1])
	return LABEL
	
{ID}
	lval.str = string(l.buf)
	return IDENT
%%
	l.empty = true
	return int(c)
}

