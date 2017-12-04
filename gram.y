%term WORD
%{
package main
%}
%union{
	tree *Tree
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

simple:
	word
|	simple word

word:
	comword

comword:
	WORD
