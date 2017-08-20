# Tracer

Tracer is a little package of utility functions for tracking and logging function execution times.

Tracer takes advantage of Go's `defer` keyword.

## Usage:
``` go
func MyLongFunction() {
    defer un(trace("MyLongFunction"))
    // ...Do something
}
``` 