## Okra Tutorial


### Preface
This tutorial assumes two things:
1. You've followed the installation instructions or have a playground instance up and running
   <br><i>If this is not applicable to you:</i> Follow the [installation instructions](https://github.com/cdkini/Okra#installation) or visit the [playground](https://repl.it/@ChetanKini/OkraPlayground)
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
var x: 1;  // The variable 'x' is declared and assigned to the numeric value of 1
var y;     // The variable 'y' is declared but is not initialized with a value. Therefore, 'y' is equal to 'null'
```

### Primitives
Okra supports four base data types or primitives: `numerics`, `strings`, `booleans`, and `null`. Their meaning and usage are as follows:
- `numerics`: Encapsulates both integers and floating point numbers.
- `strings`: Either a single character or a sequence of characters.
- `booleans`: Either `true` or `false` to represent the two truth values of logic and Boolean algebra.
- `null`: Signifies that a variable is empty or has no value.

<i>Code Snippet</i>:
```
var age: 20;         // 'age' is of type numeric
var pi: 3.14;        // 'pi' is of type numeric
var letter: "X";     // 'letter' is of type string
var word: "potato";  // 'word' is of type string
var isTrue: true;    // 'isTrue' is of type boolean
var isFalse: false;  // 'isFalse' is of type boolean
var nothing: null;   // 'nothing' is of type null
var empty;           // 'empty' is of type null
```

### Printing
The `print` keyword is used to display the result of an expression to the console or terminal. As stated before, all statements in Okra need to be concluded with `;`.

<i>Code Snippet</i>:
```
print "Hello, World!";  // Prints "Hello, World!" to the terminal
```

### Operators
Okra supports your traditional arithmetic, comparison, and logical operators.

Not explicity discussed but used before, assignment is done using `:`.

As of release <b>1.0.0</b>, string concatenation through `+` is not supported. As such, these operations are only valid on instances of the `numeric` data type. Incompatible primitives will cause an error to be raised.</i>

##### Arithmetic Operators
- `+`: <i>Add</i>
- `-`: <i>Subtract</i>
- `*`: <i>Multiply</i>
- `/`: <i>Divide</i>

##### Comparison Operators
- `=`: &nbsp;<i>Equal</i>
- `!=`: <i>Not equal</i>
- `>`: &nbsp;<i>Greater than</i>
- `<`: &nbsp;<i>Less than</i>
- `>=`: <i>Greater than or equal to</i>
- `<=`: <i>Less than or equal to</i>

##### Logical Operators
- `&&`: <i>And</i>
- `||`: <i>Or</i>
- `!`: &nbsp;<i>Not</i>

<i>Code Snippet</i>:
```
var a: 3 + 5;         // 'a' is equal to 8
var b: 3 - 5;         // 'b' is equal to -2
var c: 3 * 5;         // 'c' is equal to 15
var d: 3 / 5;         // 'd' is equal to 0.6
var e: 3 = 5;         // 'e' is equal to false
var f: 3 != 5;        // 'f' is equal to true
var g: 3 > 5;         // 'g' is equal to false
var h: 3 < 5;         // 'h' is equal to true
var i: 3 >= 3;        // 'i' is equal to true
var j: 3 <= 3:        // 'j' is equal to true
var k: true && false  // 'k' is equal to false
var l: true || false  // 'l' is equal to true
var m: !true          // 'm' is equal to false

var n: 1 + 'a'        // 'n' raises an error due to incompatible types being addes
```

### Looping
Loops in Okra are done using the `for` keyword. The initialization of a variable must be done outside of the loop while the terminating condition must be written inside the parentheses after the keyword.

<i>As of release <b>1.0.0</b>, traditional looping constructs ala `for (var i: 1; i < 5; i++)` are not available.</i>

<i>Code Snippet</i>:
```
var sum: 0;

var i: 1;
for (i < 5) {
  sum: sum + i;
  i: i + 1;
}

print sum;  // 'sum' is equal to 10
```

### Functions
Functions are defined by the `func` keyword using the following format: `func funcName : funcArgs : {}`.

Function returns are done using the `return` statement akin to other programming languages. Invocation or the calling of the function after definition require the usage of `funcName()`. 

As functions are treated as first class objects, recursion and closures are supported features.

<i>Code Snippet</i>:
```
func hello :: {
   print "Hello!";
}

func adder : x, y : {
   return x + y;
}

func fib : n : {
   if (n <= 2) {
   return 1;
   }
   return fib(n-1) + fib(n-2);
}

func makeClosure :: {
   var local: "local";
   func closure :: {
      print local;
   }
   return closure;
}


hello();                     // Prints "Hello!" to the terminal
var z: adder(3, 5);          // 'z'is equal to 8
var thirdFib: fib(3);        // 'thirdFib' is equal to 2 (fib(2) and fib(1) both evaluate to 1)

var closure: makeClosure();  // Sets 'closure' to the enclosed function in makeClosure
closure();                   // Prints "local" to the terminal
```

### Structures
Structures or structs are Okra's way of allowing the user to define their own objects. 

Defining a struct uses the `struct` keyword. To write the constructor, the `construct` method within the class definition must be written.

Instance variables or attributes are passed to the constructor in a similar fashion to calling a function; the struct name, followed by arguments wrapped by parentheses, will invoke the constructor and create a new struct instance.

Assignment of these variables to the struct instance are done using `this`. These fields must be defined in the `construct` method and can be accessed using dot notation.

<i>As of release <b>1.0.0</b>: 
   - Access modifiers (public, private, protected) are not available. All attributes are public.
   - Static methods are not supported. Invokation of class methods or usage of class attributes requires instantiation .
</i>

<i>Code Snippet</i>:
```
struct Person {
   construct : name, age, job : {
      this.name:     name;
      this.age:      age;
      this.job:      job;
      this.graduate: true;
   }
   
   hello :: {
      print "Hello!";
   }
   
   getAge :: {
      return this.age;
   }
}


var bob: Person("Bob", 42, "Accountant");  // Instantiates an instance of the Person struct
bob.hello();                               // Prints "Hello!" to the teriminal
var bobAge: bob.getAge();                  // Sets 'bobAge' to 42
print bob.age;                             // Prints 42 to the terminal
```

### Interfaces
Interfaces are named collections of method signatures. To implement an interface in Okra, we need to implement all the methods in the interface. If a structure is said to implement an interface but fails to define all the method signatures noted within that interface, an error will be raised.

Interfaces are defined by the `interface` keyword using the following format: `interface interfaceName {}`.

A structure is said to implement an interface when the interface name is noted within square brackets next to the structure name: `struct structName [interfaceName*]`. Structures are allowed to implement multiple interfaces.

<i>Code Snippet</i>:
```
import std.Math;

var pi: Math().Pi();

interface Geometry {
   area();
   perimeter();
}


struct Rectangle [Geometry] {
   construct : width, height : {
      this.width:  width;
      this.height: height;
   }
   
   area :: {
      return this.width * this.height;
   }
   
   perimeter :: {
      return 2*this.width + 2*this.height;
   }
}


struct Circle [Geometry] {
   construct : radius : {
      this.radius: radius;
   }
   
   area :: {
      return pi * this.radius * this.radius;
   }
}


var r: Rectangle(3, 5);  // Successfully implements the Geometry interface
var c: Circle(7);        // Raises an error due to Circle not implementing the perimeter method
```

### Packages
Packages, which are inclusive of functions and structures defined in other files, are imported at the top of a script using the `import` keyword. In importing a particular package, the global scope of the current program gains access to the objects defined in that package.

<i>As of release <b>1.0.0</b>, the only imports that are supported are those for the standard library. Current stdlib packages can be found in the [stdlib directory](https://github.com/cdkini/Okra/tree/master/src/stdlib).</i>

<i>Code Snippet</i>:
```
import std.Stack;


var stack: Stack();

stack.push(3);
stack.push(2);
stack.push(1);

print stack.peek();  // Prints 1 to the terminal
print stack.len();   // Prints 3 to the terminal
```

### Style
Although style is subjective, we've determined a few base guidelines that are necessary for idiomatic Okra. The formatter, which can be run using `okra fmt [script]`, applies the below rules to a source file for you.

##### Structs / Interfaces
- Structures and interfaces should capitalize their first letter
- There should be two lines between structure and interface definitions and the rest of the program
- Public methods should be listed towards the top and private/helper methods should be listed towards the bottom

##### Functions
- Any public functions/methods should capitalize their first letter while any private function/methods should be all lowercase
- There should be a single line between function/method definitions and the rest of the program

##### Miscellaneous
- Any capitalized object should be documented with a docstring
- Imports should always be at the very top of the program and nowhere else
- CamelCase is preferred over snake_case or PascalCase
