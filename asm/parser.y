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
	"errors"
	"strings"
)
%}

%union{
	num uint16
	str string
	dat []uint16
}

%token 	NEWLINE
%token	NUM
%token	LABEL
%token 	IDENT
%token 	ORG
%token	WORDS
%token	ILLEGAL

%type	<num>	NUM
%type	<str>	LABEL
%type	<str>	IDENT

%% /* Grammar */

input:	  /*empty*/
		| input NEWLINE
		| input line NEWLINE
;

line:	  LABEL				{ addLabel($1) }
		| IDENT				{ addOpcodeConstant($1, 0) }
		| IDENT IDENT		{ addOpcodeLabel($1, $2) }
		| IDENT NUM			{ addOpcodeConstant($1, $2) }
		| ORG NUM			{ setOrg($2) }
;

%%

// Assemble parses an assembly language program from in and returns the
// assembled binary, or an error if something goes wrong.
func Assemble(in *bufio.Reader) ([]uint16, error) {
	if yyParse(newLexer(in)) == 0 {
		return code, nil
	} else {
		return nil, errors.New(strings.Join(errorCodes, "\n")) 
	}
}

