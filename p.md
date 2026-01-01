Here's what I get 

When a parser hits a bang "!", it calls the parsePrefixExpression because of this

	p.registerPrefixFns(token.BANG, p.parsePrefixExpression)


func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)

	return expression
}

How does the ast look like
