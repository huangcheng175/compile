package parser

import (
	"compile/ast"
	"compile/token"
)

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken} // create a new LetStatement node

	if !p.expectPeek(token.TypeIdent) { // if the next token is not an identifier
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal} // set the Name field of the LetStatement node to the Identifier node

	if !p.expectPeek(token.TypeAssign) { // if the next token is not an assignment operator
		return nil
	}

	for !p.curTokenIs(token.TypeSemicolon) { // while the current token is not a semicolon
		p.nextToken() // advance the tokens
	}

	return stmt
}
