# .simp

#### Features
- No Inheritance
- No Polymorphism
- No Encapsulation
- Only Variables ,Functions, Procedures and Methods
- Arrays and Constants exists
- Every variable, function and method must have a type
- Arrays have only one dimension
- Functions can have multiple return types
- Void functions are called procedures
- You cannot import other .simp files into another .simp file
- Class methods cannot be written inside the class body, can only be written on the outside of the class body
- Squirly brackets are allowed for functions, procedures and class methods 
```
Data Types
- simp_int8
- simp_int16
- simp_int32
- simp_float
- simp_double
- simp_long
- simp_string
- simp_char
```
```
Entry Point of Code execution
// similar to main in C
simp()
  start_simp
    // your code goes here
  end_simp
```
```
Declarations
// we can't declare variables, classes, functions and procedures anywhere else only here
// class method implementions must be done like this {line 60}
// how we declare methods and variables
simp global impl
  start_simp
    int8 num comma
    simp_proc display() comma // a void function also known as a procedure
    simp_func add(int8 num1, int8 num2)-> int8 comma // a function with a return value
    simp_func vals()->(int8,int8) comma // a function with multiple return values
    simp_class Person() comma
    simp_int8 simp_arr[10] comma 
  end_simp
yes_mommy | yes_daddy
```
```
Implementations
simp num = 5 yes_mommy | yes_daddy
simp add{simp return num1+num2}yes_mommy | yes_daddy
simp vals{simp return 3,7}yes_mommy | yes_daddy
```
```
Object Oriented Programming
simp Person{
  simp_int8 id comma
  simp_string name comma
  simp_int8 age comma
}yes_mommy | yes_daddy
60: simp Person impl
    start_simp
      proc disp() comma
    end_simp
simp Person::disp{
  simp simp_simpPrint($"id:{0}, name:{1}, age:{2}\n",simp_This.id, simp_This.name, simp_This.age) yes_mommy | yes_daddy
}yes_mommy | yes_daddy
```
