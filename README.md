![Okra logo](img/logo.jpg)
# Okra
Okra is an interpreted, high-level, general-purpose programming language designed with simplicity and extensibility in mind. Taking inspiration from the likes of Python, JavaScript, and Go, Okra emphasizes readability through a digestible syntax. Although largely procedural, Okra support object-oriented programming by means of structs, interfaces, and composition-based inheritance through struct embedding. While Okra may not be "batteries included" like Python, an extensive standard library to cover basic file I/O, mathematics, and key data structures and algorithms are included from the get-go.

To test out the language, visit our playground! Please note that due to size and dependency restrictions from Repl.it, all language features may not be supported.
Playground updated as of version: 1.0.0

Please note that this interpreter was designed and implemented solely for educational purposes. The Okra development team has no intentions of monetization and exists to reinforce the value of open source software and its community.

### Table of Contents
- [Installation](#Installation)
- [Usage](#Usage)
- [Releases](#Releases)
- [Contributing](#Contributing)
- [Credits](#Credits)
- [License](#License)

### Installation

### Usage

##### `okra run [script]`

##### `okra fmt [script]`

### Releases

### Contributing
Although Okra was designed as an educational project, any contributions or suggestions are greatly appreciated! If you would like to contribute to the codebase, please follow these steps:

```
1. Create an issue 
2. Fork the repo
3. Create a branch; please use the issue number and name when creating your branch (i.e. "134-AddTernaryOperator")
4. Make your changes
5. Write unit tests as applicable; as a rule of thumb, ensure that the test suite has coverage over your changes
6. Format the codebase using `go fmt`
7. Ensure that your changes passes all tests using `go test`
8. Squash your changes to as few commits as possible
9. Make a pull request; please follow the same conventions as step #3
```

Failure to adhere to #4 and #5 will cause the commit to be rejected by Travis CI so please double check your work before sending it up for review.

### Credits
This project would not have have been possible without the following resources: 
- [Crafting Interpreters](https://craftinginterpreters.com/) by Bob Nystrom and the Lox language discussed therein
- [Structure and Interpretation of Computer Programs](https://mitpress.mit.edu/sites/default/files/sicp/full-text/book/book.html) by Gerald Jay Sussman and Hal Abelson and the res of the team behind [MIT 6.001](https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-001-structure-and-interpretation-of-computer-programs-spring-2005/)
- Alex Gaynor's [PyCon 2013 talk](https://www.youtube.com/watch?v=LCslqgM48D4) on writing interpreters
- [Composing Programs](https://composingprograms.com/) by John DeNero

### License


TO-DO:
- Fix line and column tracking from scanner

Order of files to clean up:
  - token
  - scanner
  - expression
  - parser
  - okra_error
  - statement
  - parser_expr
  - parser_stmt
  - parser_decl
  - environment
  - interpreter 
  
