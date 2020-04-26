# Introduction

Dice roll generators written in Go

[![Build Status](https://travis-ci.org/pconcepcion/dice.svg?branch=develop)](https://travis-ci.org/pconcepcion/dice)

## Source

[Github repo](https://github.com/pconcepcion/dice.git)

## Testing

Test can be run from the [Makefile](Makefile) by running: 

```shell
make test
```
 
Test include a bunch of statistical tests to understand the randomness of the test.
Fuzzing can be done using [dvyukov/go-fuzz](https://github.com/dvyukov/go-fuzz) and can be done with the Makefile.

```shell
make fuzz
```

### TODO 

* [] Fix the cases found with `go-fuzz` that crass the library
### References

* [go-fuzz github.com/arolek/ase](https://medium.com/@dgryski/go-fuzz-github-com-arolek-ase-3c74d5a3150c), quick totorial on how to fuzz using [go-fuzz](https://github.com/dvyukov/go-fuzz) 

## License

This code is released under the [BSD-3 Clause License](http://opensource.org/licenses/BSD-3-Clause)
