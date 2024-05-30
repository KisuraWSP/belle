# Types
```go
// read this https://www.tutorialspoint.com/clojure/clojure_data_types.htm
// What data structure that LISP uses https://en.wikipedia.org/wiki/S-expression
int // i32
string
uint // u32
float32 
bool
byte // basically char but its actually u8
```
```
we compile to golang code and then generate an executable via the golang 
once we make golang compilation possible then we can start implementing a small standard library with some basic functions
include raylib in the standard library
ship the langauge
```

# Entry Point
```go
(func main []
    (code block))
```

# Variables
```go
(var name "hallo")
```

# Constants
```go
(const PI 3.14)
```

# Functions
```go
(func print_name [name]
(print name))
```