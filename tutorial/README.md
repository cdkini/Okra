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
- [Operators](#Operators)
- [Looping](#Looping)
- [Functions](#Functions)
- [Structures](#Structures)
- [Interfaces](#Interfaces)
- [Packages](#Packages)
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
[script.okr]:
print "Hello, World!";

==============================

[Terminal]:
>>> okra run script.okr
"Hello, World!"
```

### Operators
Okra supports your traditional arithmetic, comparison, and logical operators.

Not explicity discussed but used before, assignment is done using `:`.

As of release <b>1.0.0</b>, string concatenation through `+` is not supported. As such, these operations are only valid on instances of the `numeric` data type. Incompatible primitives will cause an error to be raised.</i>

##### Arithmetic Operators
- `+`: Add
- `-`: Subtract
- `*`: Multiply
- `/`: Divide

##### Comparison Operators
- `=`:  Equal
- `!=`: Not equal
- `>`:  Greater than
- `<`:  Less than
- `>=`: Greater than or equal to
- `<=`: Less than or equal to

##### Logical Operators
- `&&`: And
- `||`: Or
- `!`:  Not

<i>Code Snippet</i>:
```
[script.okr]:
print 1 + 1;
var x: 2;
print x * 2;
print "sushi" * 7;

==============================

[Terminal]:
>>> okra run script.okr
2
4
OkraError [4,6]: Invalid usage of "*" on non-numeric operands.
```

### Looping


### Functions


### Structures


### Interfaces


### Packages


### Style


