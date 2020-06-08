# Abstract Syntax Tree
Abstract Syntax Tree : Create, Type Check, Evaluate

This repository creates a basic Abstract Syntax Tree(AST) with Integer and Boolean Operation. Using this AST, it can type check and evaluate the expression.

## The Language

It is an expression language with Integer and Boolean Operations.

The Operators used in this project are : 

  ``Integer Operator : +, -, *, /``
  
  ``Boolean Operators: ==, !=, &&, ||``

These operations can be performed on Boolean or Integer Literals. 

Example:

`5 + 2`   /*Result : 7*/

`5 + 2 == 3 + 4 || true`   /*Result : true*/

`5 || 2`   /*Result : Error: Operation Not Possible*/

`5 == 2`   /*Result : false*/

`5 - 2 + 1 * 3 / 4`   /*Result : 4*/

## Assumption

- There is no lexer or parser. 

- AST structure as per operator precedence need to be created, following the Node structure.

- Booleans and Integer Operations can be evaluated and type checked.

## Project Structure

This project is written using Go. 
The project has 5 packages, each one has a Single Responsibility and is created following some of the SOLID Design guidelines

- `tokens` : This contain tokens for different Boolean and Integer Operations. This is enum equivalent in Go.

- `ast` : This contains the structure definition of AST. Each node can be defined either as Boolean Expression or Integer Expression or Operators. Parser should be able to structure the expression language into this AST structure. 

- `checker` : This is responsible for Type Checking of AST. `environment.go` contains the Syntax Type for each of the node in AST.
`checker.go` does the type checking. It recursively goes through the AST and check if the expression matched with the Syntax Type as defined by the language.

`checker_test.go` contains the unit test for the Type Checker. It creates the AST of the 5 examples mentioned above using the ast and tokens package. And type checks each of the created AST.

The test can run directly from the file, once you have the project and Go installed.

- `evaluate` : This is responsible for evaluating the AST into Boolean or Integer result. `evaluate.go` has the implementation for evaluation. `evaluate_test.go` creates AST and evaluate the results.

- `visitor` : This is using Visitor Design Pattern to type check and evaluate the AST. `visitor_test.go` contains the test which is bringing it all together - creates the AST, evaluates them and performs type checking. For evaluation, type checking is always performed prior to calling Evaluate, so that if there is any mismatch it can be shared with the Visitable client.
