%term IF
%term WORD
%left IF
%{
package main

import (
	"github.com/lufia/qsh/ast"
	"github.com/lufia/qsh/cgen"
)
%}
%union{
	tree *ast.Node
}
%type<tree> line block body cmdsa cmdsan assign
%type<tree> cmd simple first word comword words
%type<tree> WORD
%%
stmt:
	{
		return 1
	}
|	line '\n'
	{
		ast.Dump($1)
		cgen.Compile($1)
		return 0
	}

line:
	cmd
|	cmdsa line

body:
	cmd
|	cmdsan body
	{
		$$ = ast.New(ast.LIST, $1, $2)
	}

cmdsa:
	cmd ';'
|	cmd '&'
	{
		$$ = ast.Async($1)
	}

cmdsan:
	cmdsa
|	cmd '\n'

block:
	'{' body '}'
	{
		$$ = ast.Block($2)
	}

assign:
	first '=' word
	{
		$$ = ast.New(ast.ASSIGN, $1, $3)
	}

cmd:
	{
		$$ = nil
	}
|	IF block block
	{
		$$ = ast.New(ast.IF, $2, $3)
	}
|	simple
	{
		$$ = ast.Simple($1)
	}
|	assign

simple:
	word
|	simple word
	{
		$$ = ast.New(ast.LIST, $1, $2)
	}

first:
	comword

word:
	comword

comword:
	'$' word
	{
		$$ = ast.Var($2)
	}
|	'(' words ')'
	{
		$$ = ast.Tuple($2)
	}
|	WORD

words:
	{
		$$ = nil
	}
|	words word
	{
		if $1 == nil {
			$$ = $2
		} else {
			$$ = ast.New(ast.LIST, $1, $2)
		}
	}
