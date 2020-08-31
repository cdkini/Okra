## Okra Tutorial


### Preface
This tutorial assumes two things:
1. You've followed the installation instructions or have a playground instance up and running
   <br><i>If this is not applicable to you:</i> Follow [the installation instructions](https://github.com/cdkini/Okra#installation) or visit [the playground](https://repl.it)
2. You have a decent understanding of programming fundamentals (variables, conditionals, loops, structs, etc)
   <br><i>If this is not applicable to you:</i> Learn an established language first! I'd recommend Python 🐍🐍🐍

### Table of Contents
- [Comments](#Comments)
- [Variables](#Variables)
- [Primitives](#Primitives)
- [Operations](#Operations)
- [Conditionals](#Conditionals)
- [Looping](#Looping)
- [Functions](#Functions)
- [Structures](#Structures)
- [Interfaces](#Interfaces)
- [Modules](#Modules)
- [stdlib](#stdlib)
- [Style](#Style)


### Comments
Comments in Okra are preceded by `//`. Any text between `//` and the end of the line is ignored by the Okra interpreter and will not be executed. Multiline comments are not supported.

<i>Code Snippet</i>:
```
// This line will have no impact on the program!
```

### Variables
Variables are declared using the `var` keyword. Variables can either be immediately assigned to a value using `:`, the assignment operator, or left empty; doing the latter will set the variable to `null`. All statements in Okra need to be concluded with `;`.

<i>Code Snippet</i>:
```
var x: 1; // The variable 'x' is declared and assigned to the numeric value of 1
var y;    // The variable 'y' is declared but is not initialized with a value. Therefore, 'y' is equal to 'null'
```

### Primitives
Okra supports four base data types or primitives: `numerics`, `strings`, `booleans`, and `null`. Their meaning and usage are as follows:
- `numerics`: Encapsulates both integers and floating point numbers.
- `strings`: Either a single character or a sequence of characters.
- `booleans`: Either `true` or `false` to represent the two truth values of logic and Boolean algebra.
- `null`: Signifies that a variable is empty or has no value.

<i>Code Snippet</i>:
```
var age: 20;        // 'age' is of type 'numeric'
var pi: 3.14;       // 'pi' is of type 'numeric'
var letter: "X";    // 'letter' is of type 'string'
var word: "potato"; // 'word' is of type 'string'
var isTrue: true;   // 'isTrue' is of type 'boolean'
var isFalse: false; // 'isFalse' is of type 'boolean'
var nothing: null;  // 'nothing' is of type 'null'
var empty;          // 'empty' is of type 'null'
```

### Printing
The `print` keyword is used to display the result of an expression to the console or terminal. As stated before, all statements in Okra need to be concluded with `;`.

<i>Code Snippet</i>:
```
// 0-helloWorld.okr
print "Hello, World!";

==============================

[Terminal]:
>>> okra run 0-helloWorld.okr
"Hello, World!"
```

### Operations


### Conditionals


### Looping


### Functions


### Structures


### Interfaces


### Modules


### stdlib


### Style


