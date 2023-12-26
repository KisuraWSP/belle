# SYNTAX
---
### Variables
```
global_declare:
    bVar name bString = "john cena"
```

### Functions
```
global_declare:
    belle_func test() bAny
    belle_func test(b bInt4) bInt4

impl:
    test()
    start
        bReturn "hallo"
    end

    test(b bInt4)
    start
        bReturn b
    end

```


### Procedures
```
global_declare:


impl:
    test()
    start
        bPrint "hallo"
    end

    test(b bInt4)
    start
        bPrint b
    end
```


