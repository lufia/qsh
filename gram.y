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
	{
		$$ = $1
	}

comword:
	WORD
