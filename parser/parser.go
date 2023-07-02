// parser/parser.go

package parser

import (
	"seahorsesyntax.villdyr.dev/ast"
	"seahorsesyntax.villdyr.dev/lexer"
	"seahorsesyntax.villdyr.dev/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {

}
