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
- You cannot import other .hello files into another .hello file
- Class methods cannot be written inside the class body, can only be written on the outside of the class body
- Squirly brackets are allowed for functions, procedures and class methods 
```
Data Types
- int8
- int16
- int32
- float
- double
- long
- string
- char
```
```
Entry Point of Code execution
// similar to main in C
simp()
  start
    // your code goes here
  end
```
```
Declarations
// we can't declare variables, classes, functions and procedures anywhere else only here
// class method implementions must be done like this {line 60}
// how we declare methods and variables
simp global impl
  start
    int8 num comma
    proc display() comma // a void function also known as a procedure
    func add(int8 num1, int8 num2)-> int8 comma // a function with a return value
    func vals()->(int8,int8) comma // a function with multiple return values
    class Person() comma
    int8 arr[10] comma 
  end
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
  int8 id comma
  string name comma
  int8 age comma
}yes_mommy | yes_daddy
60: simp Person impl
    start
      proc disp() comma
    end
simp Person::disp{
  simp simp_simpPrint($"id:{0}, name:{1}, age:{2}\n",simp_This.id, simp_This.name, simp_This.age) yes_mommy | yes_daddy
}yes_mommy | yes_daddy
```
