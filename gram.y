%term WORD
%{
package main

import (
	"github.com/lufia/qsh/ast"
)
%}
%union{
	tree *ast.Node
}
%type<tree> line cmdsa
%type<tree> cmd simple word comword
%type<tree> WORD
%%
stmt:
	{
		return 1
	}
|	line '\n'
	{
		ast.Dump($1)
	}

line:
	cmd
|	cmdsa line

cmdsa:
	cmd ';'
|	cmd '&'

/*
cmdsan:
	cmdsa
|	cmd '\n'
*/

cmd:
	{
		$$ = nil
	}
|	simple
	{
		$$ = ast.Simple($1)
	}

simple:
	word
|	simple word
	{
		$$ = ast.New(ast.LIST, $1, $2)
	}

word:
	comword

comword:
	WORD
