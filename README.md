![Okra logo](img/logo.jpg)
# Okra
Okra is an interpreted, high-level, general-purpose programming language designed with simplicity and extensibility in mind. Taking inspiration from the likes of Python, JavaScript, and Go, Okra emphasizes readability through a digestible syntax. Although largely procedural, Okra support object-oriented programming by means of structs, interfaces, and composition-based inheritance through struct embedding. While Okra may not be "batteries included" like Python, an extensive standard library to cover basic file I/O, mathematics, and key data structures and algorithms are included from the get-go. To promote consistency across the language, a built-in formatter akin to `go fmt` is included.

To test out the language, visit our playground! Please note that due to size and dependency restrictions from Repl.it, all language features may not be supported.

Playground updated as of version: **1.0.0**

Please note that this interpreter was designed and implemented solely for educational purposes. The Okra development team has no intentions of monetization or commercialization and exists solely to reinforce the value of open source software and its community.

## Table of Contents
- [Installation](#Installation)
- [Usage](#Usage)
- [Updates](#Updates)
- [Contributing](#Contributing)
- [Credits](#Credits)
- [License](#License)

## Installation

##### Dependencies


## Usage
Upon building the executable and setting up your path, there are two ways in which you can utilize the `okra` command:
##### `okra run [script]`
This command will run the interpreter on your program. All output, including error messages, will be displayed to your terminal.
##### `okra fmt [script/dir]`
This command will run the formatter on your program or all files in a particular directory. Okra strongly encourages usage of the native formatting rules to ensure consistency between file and codebases. If you would like to modify the standard rules, please update `fmt.json` in `okra/formatter` before creating the executable.

Please note that both commands will only work on files with a `.okr` extension. The formatter will ignore invalid files if used on a directory.

## Updates
##### Releases
- 1.0.0: Initial release (TBD)
##### Roadmap
High Priority:

Low Priority:

## Contributing
Although Okra was designed as an educational project, any contributions or suggestions are greatly appreciated! If you would like to contribute to the codebase, please follow these steps:

1. Create an issue 
2. Fork the repo
3. Create a branch; please use the issue number and name when creating your branch (i.e. "134-AddTernaryOperator")
4. Make your changes
5. Write unit tests as applicable; as a rule of thumb, ensure that the test suite has coverage over your changes
6. Format the codebase using `go fmt`
7. Ensure that your changes passes all tests using `go test`
8. Squash your changes to as few commits as possible
9. Make a pull request; please follow the same conventions as step #3

Failure to adhere to #4 and #5 will cause the commit to be rejected by Travis CI so please double check your work. Thanks!

## Credits
This project would not have have been possible without the following resources: 
- [Crafting Interpreters](https://craftinginterpreters.com/) by Bob Nystrom and the Lox language discussed therein
- [Structure and Interpretation of Computer Programs](https://mitpress.mit.edu/sites/default/files/sicp/full-text/book/book.html) by Gerald Jay Sussman and Hal Abelson and the res of the team behind [MIT 6.001](https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-001-structure-and-interpretation-of-computer-programs-spring-2005/)
- Alex Gaynor's [PyCon 2013 talk](https://www.youtube.com/watch?v=LCslqgM48D4) on writing interpreters
- [Composing Programs](https://composingprograms.com/) by John DeNero

## License
The Okra project is licensed under the MIT License Copyright (c) 2020.

See the [LICENSE](https://github.com/cdkini/Okra/blob/master/LICENSE) for information on the history of this software, terms & conditions for usage, and a DISCLAIMER OF ALL WARRANTIES.

All trademarks referenced herein are property of their respective holders.
