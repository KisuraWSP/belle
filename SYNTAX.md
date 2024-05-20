# Types
```rust
positive_number             $ unsigned integers (32 bit)
negative_number             $ signed integers (32 bit)
value                       $ string
letter                      $ char
numbers_precise             $ float
condtion                    $ boolean
```

# Comments
```rust
$ single line comment
```

# Entry point of program
```rust
prog hallo(){
    $ code goes here
}
```

# Function
```rust
prog is_func(){
    $ some random code
}

prog is_func2() & letter{
    output 't';
}
```

# Variables
```rust
t : value = "2332";
```

# Constants
```rust
t : is_fixed value = "te";
```

# Structs
```rust
custom_type Person{
    name : value;
    age := 2;
}
```

# Enums
```rust
number_type fruits{
    Orange;
    Banana = 2; $ overriding the value
    Grape; $ now value is 3
}
```

# Condition Handling
```rust
$ switch case
check(num){
    is 1:
        out("1");
    is 2:
        out("2");
    not: 
        out("not the above cases");
}
```

# Type Alias
```rust
make buffer : (value, positive_number);
age : buffer = ("test", 34);
```

# Import statements
```rust
$ used when calling libraries or files
call_this "nerd"; $ math library
call_this "basic"; $ contains basic functions such as out() and in();

call_this&read_file "test.belle";