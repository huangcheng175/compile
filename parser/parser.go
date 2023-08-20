package parser

import (
	"compile/ast"
	"compile/lexer"
	"compile/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram is the entry point for the parser
// It will parse the tokens and return a pointer to a Program node
// The parser will read tokens until it encounters an EOF token
// The parser will return a pointer to a Program node, which contains a slice of Statement nodes
// The parser will call parseStatement() to parse each statement
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}              // create a new Program node
	program.Statements = []ast.Statement{} // initialize the Statements field to an empty slice
	for p.curToken.Type != token.TypeEOF { // while the current token is not an EOF token
		stmt := p.parseStatement() // parse the statement
		if stmt != nil {           // if the statement is not nil
			program.Statements = append(program.Statements, stmt) // append the statement to the Statements field of the Program node
		}
		p.nextToken() // advance the tokens
	}
	return program // return the Program node
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.TypeLet:
		return p.parseLetStatement()
	case token.TypeReturn:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) peekTokenIs(t token.Type) bool {
	return p.peekToken.Type == t
}

func (p *Parser) curTokenIs(t token.Type) bool {
	return p.curToken.Type == t
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekToken.Type == t { // if the next token is the expected type
		p.nextToken() // advance the tokens
		return true
	} else {
		p.peekError(t) // create an error message
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.Type) {
	msg := "expected next token to be %s, got %s instead" // create an error message
	p.errors = append(p.errors, msg)                      // append the error message to the errors slice
}
