# Tree-walking interpreter
Implementation of a tree-walking interpreter in Go for a simple language. The interpreter consists of a lexer, parser, AST, internal object system, and an evaluator. Reference implementation: https://interpreterbook.com/

# Language
Language has the following features:
- variable bindings
- integers and booleans
- arithmetic expressions
- built-in functions
- first-class and higher-order functions
- closures
- a string data structure
- an array data structure
- a hash map data structure

Binding values to variables:
```
let age = 1;
let result = 5 + 10;
```

Data structures:
```
let myArray = [1, 2, 3];
myArray[0]; // => 1
let myHashMap = {"name": "Alice", "age": 30};
myHashMap["name"]; // => "Alice"
```

Functions:
```
let add = fn(x, y) { return x + y; };
let implicitReturnAdd = fn(x, y) { x + y; };
add(2, 3); // calling a function
```

Recursive and higher-order functions:
```
let fibonacci = fn(n) {
  if (n == 0) {
    return 0;
  } else {
    if (n == 1) {
      return 1;
    } else {
      return fibonacci(n - 1) + fibonacci(n - 2);
    }
  }
};

// first class functions support: functions can be used as arguments in function calls
let twice = fn(f, x) {
    return f(f(x));
};
let addTwo = fn(x) { return x + 2; };
twice(addTwo, 2); // => 6
```