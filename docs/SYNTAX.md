# SYNTAX
---
### Variables
```
global_declare:
    bVar name : bString->"john cena"
```

### Functions
```
global_declare:
    belle_func test() : bAny
    belle_func test(b bInteger) :  bInteger

impl:
    test() bAny
    start
        bReturn "hallo"
    end

    test(b bInteger) : bInteger
    start
        bReturn b
    end

```


### Procedures
```
global_declare:
    belle_proc test()
    belle_proc test(b bInteger)

impl:
    test()
    start
        bPrint "hallo"
    end

    test(b bInteger)
    start
        bPrint b
    end
```


