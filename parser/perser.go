package parser

import (
	"github.com/syuya2036/go-in/ast"
	"github.com/syuya2036/go-in/lexer"
	"github.com/syuya2036/go-in/token"
)

type Perser struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Perser {
	p := &Perser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Perser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Perser) ParseProgram() *ast.Program {
	return nil
}