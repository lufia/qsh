%term IF FOR IN
%term WORD REDIR
%left IF FOR LOAD
%left ANDAND OROR
%left '|'
%{
package main

import (
	"github.com/lufia/qsh/ast"
	"github.com/lufia/qsh/build"
)
%}
%union{
	tree *ast.Node
}
%type<tree> line block body cmdsa cmdsan assign redir
%type<tree> cmd simple first word comword words
%type<tree> WORD REDIR
%%
stmt:
	{
		return 1
	}
|	line '\n'
	{
		if *flagDebug {
			ast.Dump($1)
		}
		build.Compile($1)
		return 0
	}

line:
	cmd
|	cmdsa line
	{
		$$ = ast.New(ast.LIST, $1, $2)
	}

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
		$$ = ast.Assign($1, $3)
	}

cmd:
	{
		$$ = nil
	}
|	IF block block
	{
		$$ = ast.New(ast.IF, $2, $3)
	}
|	FOR word IN words block
	{
		p := ast.New(ast.LIST, $2, $4)
		$$ = ast.New(ast.FOR, p, $5)
	}
|	LOAD word
	{
		$$ = ast.Load($2)
	}
|	simple
	{
		$$ = ast.Simple($1)
	}
|	cmd ANDAND cmd
	{
		$$ = ast.New(ast.ANDAND, $1, $3)
	}
|	cmd OROR cmd
	{
		$$ = ast.New(ast.OROR, $1, $3)
	}
|	cmd '|' cmd
	{
		$$ = ast.New(ast.PIPE, $1, $3)
	}
|	assign

simple:
	first
|	simple word
	{
		$$ = ast.New(ast.LIST, $1, $2)
	}
|	simple redir
	{
		$$ = ast.New(ast.LIST, $1, $2)
	}

redir:
	REDIR word
	{
		$$ = ast.Redirect($1, $2)
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
|	'$' '{' words '}'
	{
		$$ = ast.Module($3)
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
