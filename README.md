# nopadding

Provides:
- A Golang analyzer to show whether your struct has been automatically padded by the Golang compiler
- A Golang analyzer wrapper which can be invoked from a unit test

This package is very useful during system development (e.g. when marshalling syscall returned structs).
