![Okra logo](img/logo.jpg)
<h2 align="center"><i>A Simple Yet Extensible Scripting Language</i></h2>

<p align="center">
  <a href="https://travis-ci.org/github/cdkini/Okra">
    <img alt="Travis" src="https://img.shields.io/travis/cdkini/Okra/master?style=flat-square">
  </a> 
  <a href="https://golang.org/doc/go1.14">
    <img alt="Go" src="https://img.shields.io/github/go-mod/go-version/cdkini/Okra?style=flat-square"
  </a> 
  <a href="https://opensource.org/licenses/MIT">
    <img alt="License" src="https://img.shields.io/github/license/cdkini/Okra?color=red&style=flat-square"
  </a>
  <a>
    <img alt="Okra" src="https://img.shields.io/badge/version-v1.0.0-yellow?style=flat-square"
  </a>
</p>


### Intro
Okra is an interpreted, high-level, general-purpose programming language designed to make scripting easy and maintainable. Taking inspiration from the likes of Python, JavaScript, and Go, Okra emphasizes readability through a digestible syntax. Although largely procedural, Okra supports an object-oriented paradigm by means of structs, interfaces, and struct embedding. While Okra may not be "batteries included" like Python, an extensive standard library to cover basic error handling, mathematics, and key data structures and algorithms is accessible from the get-go. To promote consistency across the language's community, a built-in formatter akin to `go fmt` is included.

<i>Please note that the contents of this repository were designed and implemented solely for educational purposes. The Okra development team has no intentions of monetization or commercialization and exists solely to reinforce the value of open source software and its community.</i>

### Playground
To test out the language, visit the [official Okra playground](https://repl.it/@ChetanKini/OkraPlayground) as hosted on Repl.it! Please note that due to size and dependency restrictions, all language features may not be supported. We've include sample files, which are the same as those noted in the [tutorial](https://github.com/cdkini/Okra/tree/master/tutorial), for your reference.

Playground version: <b><i>1.0.0</i></b>


## Table of Contents
- [Installation](#Installation)
- [Usage](#Usage)
- [Updates](#Updates)
- [Contributing](#Contributing)
- [Credits](#Credits)
- [License](#License)


## Installation


### Dependencies
Okra was written on a 64-bit [Ubuntu 20.04 LTS OS](https://releases.ubuntu.com/20.04/) using [Go 1.14.3](https://golang.org/dl/). The project uses no external dependencies so access to any version of Go 1.1x and bash to run `.sh` files should be all you need. Since Go compiles down to a binary specific to the architecture it's running on, we leave the actual creation of the executable up to you.

### Building Executable
You can either:
1. Run the installation script by running `./INSTALL` or `bash INSTALL` to create the `bin` directory. Your executable is located within.
2. Build the file using the Go compiler by running `go build -o okra` and moving the executable to your PATH or other desired location.


## Usage


### Writing Scripts
To learn more about Okra's syntax and the language's features, please see our dedicated [tutorial](https://github.com/cdkini/Okra/tree/master/tutorial). We've included code snippets therein but we highly recommend testing and breaking the sample files on your own machine or the [playground](https://repl.it/@ChetanKini/OkraPlayground).


### Running Executable
Upon building the executable and setting up your path, there are two ways in which you can utilize the `okra` command:

```
// Run the interpreter on a script
./okra run [script]


// Run the formatter on a script or directory
./okra fmt [script/dir]
```

Okra strongly encourages usage of the native formatting rules to ensure consistency between file and codebases; we recommend reading the ['Style' section](https://github.com/cdkini/Okra/tree/master/tutorial#style) of the tutorial to learn the language's standards.

<i>Please note that both commands will only work on files with a `.okr` extension. The formatter will ignore invalid files if used on a directory.</i>


## Updates
### Releases
- <b>1.0.0</b>: Initial release (9/5/2020)

### Roadmap
##### High Priority:
- Built-ins (hash, casting, range, round, etc.)
- Increment/decrement syntactic sugar (i++/i--)
- Operator syntactic sugar (+=, -=, /=, *=)
- Built-in lists (outside of stdlib import)
- Modulo operator (%)
- Square root operator (**)
##### Low Priority:
- Ternary operator (?)
- Struct inheritance
- break/continue within loops
- Multiline strings and comments

## Contributing
Although Okra was designed as an educational project, any contributions or suggestions are greatly appreciated! If you would like to contribute to the codebase, please follow these steps:

```
1. Create an issue 
2. Fork the repo
3. Create a branch*
4. Make your changes
5. Write unit tests as applicable
6. Format the codebase using 'go fmt'**
7. Ensure that your changes passes all tests using 'go test'**
8. Squash your changes to as few commits as possible*
9. Make a pull request*
```
<i>*Please use the issue number and name when possible to improve clarity and project maintainability (i.e. "134-AddTernaryOperator")<br></i>
<i>**Failure to run commands can cause changes to be rejected by Travis so please double check your work.</i>


## Credits
This project would not have have been possible without the following resources: 
- [Crafting Interpreters](https://craftinginterpreters.com/) by Bob Nystrom and the Lox language discussed therein
- [Structure and Interpretation of Computer Programs](https://mitpress.mit.edu/sites/default/files/sicp/full-text/book/book.html) by Gerald Jay Sussman and Hal Abelson
- Alex Gaynor's [PyCon 2013 talk](https://www.youtube.com/watch?v=LCslqgM48D4) on writing interpreters
- [Composing Programs](https://composingprograms.com/) by John DeNero


## License
The Okra project is licensed under the MIT License Copyright (c) 2020.

See the [LICENSE](https://github.com/cdkini/Okra/blob/master/LICENSE) for information on the history of this software, terms & conditions for usage, and a DISCLAIMER OF ALL WARRANTIES.

All trademarks referenced herein are property of their respective holders.
