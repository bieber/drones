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
 
%{
package asm
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
%type	<dat>	nums

%% /* Grammar */

input:		  /*empty*/
			| input NEWLINE
			| input line NEWLINE
;

line:		  label_line
			| non_label
			| label_line non_label
			;

label_line:	  LABEL				{ addLabel($1)				}
			;
		
non_label:	  IDENT				{ addOpcodeSolo($1) 		}
			| IDENT IDENT		{ addOpcodeLabel($1, $2) 	}
			| IDENT NUM			{ addOpcodeConstant($1, $2)	}
			| ORG NUM			{ setOrg($2) 				}
			| WORDS nums		{ addWords($2) 				}
			;

nums:	  /*empty*/				{ $$ = nil 					}
		| nums NUM				{ $$ = append($1, $2) 		}
		;

%%

