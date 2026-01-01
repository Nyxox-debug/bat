Let’s zoom in on the *engine room* of this parser: how prefix and infix precedence cooperate to build the AST. This is a Pratt parser, and once the trick clicks, it feels almost unfairly elegant.

I’ll explain it as a *controlled tug-of-war* between operators.

---

## The mental model first (this matters)

Think of `parseExpression(precedence)` as saying:

> “Parse **as much of an expression as you are allowed to**, given how strong the operators are around you.”

Prefix parsing **starts** an expression.
Infix parsing **extends** an expression *only if it deserves to*.

Precedence is the bouncer at the door.

---

## Prefix parsing: how expressions begin

This line is the first critical move:

```go
prefix := p.prefixParseFns[p.curToken.Type]
leftExp := prefix()
```

This answers the question:

> “Given the current token, how can an expression *start*?”

Examples:

* `5` → integer literal
* `x` → identifier
* `-5` → prefix expression
* `!true` → prefix expression

### Example: parsing `-5`

Tokens:

```
- 5
```

Call:

```go
parseExpression(LOWEST)
```

* `curToken = '-'`
* Prefix function = `parsePrefixExpression`
* It builds:

  ```
  (- ?)
  ```
* Then it calls:

  ```go
  parseExpression(PREFIX)
  ```

  This is the key move: **prefix operators bind tighter than almost everything**.

That `PREFIX` precedence prevents weaker operators from sneaking inside.

Result:

```
(- 5)
```

Prefix parsing *does not care* about what comes after yet. It only establishes the left-hand expression.

---

## Infix parsing: how expressions grow

Now the real magic:

```go
for !p.peekTokenIs(token.SEMICOLON) &&
    precedence < p.peekPrecedence() {
```

This loop asks one question repeatedly:

> “Is the *next operator* stronger than what I’m currently allowed to consume?”

If yes → extend the expression
If no → stop and return control upward

### This is the entire precedence system.

No tables, no recursion spaghetti. Just that comparison.

---

## Concrete example: `1 + 2 * 3`

Let’s walk it slowly.

### Step 1: parse `1`

Call:

```go
parseExpression(LOWEST)
```

* `curToken = 1`
* Prefix parse → `IntegerLiteral(1)`
* `leftExp = 1`

Peek token is `+`.

Check:

```
precedence (LOWEST) < peekPrecedence(+) (SUM)
```

True → we may extend.

---

### Step 2: parse `+`

* Move to `+`
* Call:

  ```go
  parseInfixExpression(leftExp = 1)
  ```

Inside:

```go
precedence := p.curPrecedence() // SUM
p.nextToken()
expression.Right = p.parseExpression(precedence)
```

Important: the **right-hand side is parsed with SUM precedence**, not LOWEST.

This is what enforces left-to-right grouping.

---

### Step 3: parse RHS of `+` → now at `2`

Call:

```go
parseExpression(SUM)
```

* Prefix parse → `2`
* Peek token is `*`

Check:

```
precedence (SUM) < peekPrecedence(*) (PRODUCT)
```

True → multiplication is stronger, so it belongs *inside* the right-hand side.

---

### Step 4: parse `*`

* Move to `*`
* Call `parseInfixExpression(left = 2)`
* Store precedence = PRODUCT
* Parse right with:

  ```go
  parseExpression(PRODUCT)
  ```

Right becomes `3`.

Peek token is now EOF or `;`, loop stops.

Result of RHS:

```
(2 * 3)
```

Return to `+`.

---

### Final AST

```
    +
   / \
  1   *
     / \
    2   3
```

Exactly what you want.
No special-case logic. No grammar rules. Just precedence comparison.

---

## Why `PREFIX` precedence matters

Look at this line:

```go
expression.Right = p.parseExpression(PREFIX)
```

This ensures:

```
-5 * 3
```

Parses as:

```
(-5) * 3
```

and **not**:

```
-(5 * 3)
```

Because `*` has lower precedence than `PREFIX`, it cannot invade the prefix expression.

That one constant quietly prevents a whole category of bugs.

---

## The key invariant to remember

> Prefix functions **create** expressions.
> Infix functions **extend** expressions *only if precedence allows it*.

Everything else is bookkeeping.

---

## How to truly internalize this (self-study path)

1. **Trace by hand**
   Take expressions like:

   ```
   1 + 2 + 3
   1 + 2 * 3 + 4
   -a * b
   a == b + c * d
   ```

   Write down:

   * current precedence
   * peek precedence
   * why the loop continues or stops

2. **Break it on purpose**

   * Swap `<` to `<=` in the loop
   * Change `PREFIX` to `LOWEST`
     Observe how the AST collapses incorrectly.

3. **Add a new operator**
   Implement exponentiation `**` with higher precedence than `*`.
   This forces you to reason about associativity (right vs left).

4. **Read Pratt’s original paper**
   Not for code — for mindset.
   You’ll recognize this parser instantly once it clicks.

---

Once this model locks in, parsing stops feeling mystical. It becomes mechanical, almost boring — which is exactly what you want from something as dangerous as a parser.
