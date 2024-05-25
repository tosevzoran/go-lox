# GoLox

GoLox is an attempt to implement an interprter for the Lox language from the [CraftingInterpreters](https://craftinginterpreters.com/) book.

## Running/Building

To run the code simply execute the

```sh
go run .
```

or

```sh
go run . ./examples/sum.lox
```

to compile from a source file. Similarly, to generate the binary file, run:

```sh
go build .
```

## Grammar for Lox expressions

### Chapter 5

Initial suport for a handful expressions

```
expression     → literal
               | unary
               | binary
               | grouping ;

literal        → NUMBER | STRING | "true" | "false" | "nil" ;
grouping       → "(" expression ")" ;
unary          → ( "-" | "!" ) expression ;
binary         → expression operator expression ;
operator       → "==" | "!=" | "<" | "<=" | ">" | ">="
               | "+"  | "-"  | "*" | "/" ;
```
