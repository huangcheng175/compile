package parser

import (
	"compile/ast"
	"compile/token"
)

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken} // create a new ReturnStatement node
	p.nextToken()
	for !p.expectPeek(token.TypeSemicolon) { // while the next token is not a semicolon
		p.nextToken() // advance the tokens
	}

	return stmt
}
