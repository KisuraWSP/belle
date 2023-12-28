# Belle
- An esoteric programming language

---
- Build System : GNU Make

## Requirements to run the program
- Must have the latest version of the GO compiler installed in your machine
- For windows users to use the build system they must use WSL to use GNU make
- If you don't want to use GNU make, you can still use the Go Compiler and do the commands

### How to Build Program
```
# Using via GNU make
make build

# Using the GO compiler
go build
```

### How to run Program
```
# (You will actually see no output from the program)
# (Thats how we intended it to be anyways)
# Using via GNU make
make run

# Using the GO compiler
go run .

# Using the executable
./belle
```

### Display List of Commands
```
# Using via GNU make
make help

# Using the GO compiler
go run . help

# Using the executable
./belle help
```

### Display Version of the Langauge
```
# Using via GNU make
make version

# Using the GO compiler
go run . version

# Using the executable 
./belle version
```

### Display the Executable Name
```
# (GNU make only)
# For Users who want to know the programs executable file name 
# This can be used by windows users, to set it as an environment variable
make exec_name
```
## Links
[CONTRIBUTING.md](https://github.com/KisuraWSP/simp_/blob/main/CONTRIBUTING.md)<br>
[SYNTAX.md](https://github.com/KisuraWSP/simp_/blob/main/SYNTAX.md)

# References
[AST](https://en.wikipedia.org/wiki/Abstract_syntax_tree)<br>
[uc](https://github.com/mewmew/uc)<br>
[llvm ir for go ](https://github.com/llir/llvm)<br>
[gogo](https://github.com/shivansh/gogo)<br>
[gocc](https://github.com/goccmack/gocc)<br>
---
