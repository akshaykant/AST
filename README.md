# Abstract Syntax Tree
Abstract Syntax Tree : Create, Type Check, Evaluate

This repository creates a basic Abstract Syntax Tree(AST) with basic Integer and Boolean Operation. Using this AST, it can type check and evaluate the expression.

## The Language

It is a basic expression language with Integer and Boolean Operations.

Some of the Operators are : 

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

- AST structure as per operator precedence need to be created, following the Node strucuture.

- Booleans and Integer Operations can be evaluated and type checked.

## Project Structure

This project is written using Go. 
The project has 5 packages, each one given a Single Responsibility and is created following some of the SOLID Design guidelines

- `tokens` : This contains tokens for different Boolean and Integer Operations. This is enum equivalent deginition in Go.

- `ast` : This contains the structure definition of AST. Each node can be defined either in Boolean Expression, Integer Expression and Operators. Parser should be strucutre the expression language into this AST structure. 

- `checker` : This is responsible for Type Checking of AST. `environment.go` contains the Syntax Type for each of the node in AST.
`checker.go` does the type checking. It recursively goes throught the AST and check if the expression matched with the Syntax Type as defined by the language.

`checker_test.go` contains the unit test for the Type Checker. It create the AST of the 5 examples mentioned above using the ast and tokens package. And type checks each of the created AST.

These can be differently run from the file, once you have the project and Go installed.

- `evaluate` : This is responsible for evaluation the AST into Boolean or Integer result. `evaluate.go` has the implementation for evaluation. `evaluate_test.go` creates AST and evaluate the results.

- `visitor` : This is using Visitor Design Pattern to type check and evaluate the AST. `visitor_test.go` contains the test which si bringing it all together - creates the AST, evaluates them and perform type checking. For evaluation, type checking is always performed prior to calling Evaluate, so that if there is any mismatch it can be shared with the Visitable client.
