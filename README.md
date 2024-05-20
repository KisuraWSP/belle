# Goals
- [ ] A Simple Language i can just use for fun
- [ ] Type Safe
- [ ] Esoteric 
- [ ] Have a standard library

```
TODO
- add features over time and make the documentation of the language and repo easier to use
- make sure the langauge is safe and works properly and gives actual output
```

# Belle
- Build System : GNU Make
  
## Important
- Read the CHANGELOG.txt file for how the progress of the Language is occurring
  
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
# Using via GNU make
make run

# Using the GO compiler
go run .

# Using the executable (Build the program first inorder to do this)
./belle
```

### Display List of Flags
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
## Links
[CONTRIBUTING](https://github.com/KisuraWSP/belle/blob/main/CONTRIBUTING.md)<br>
---
