# Array and Strings

* Two Pointers

```go
func twoPointer[K comparable](arr []K){
    i, j := 0, len(arr) - 1
    for i < j {
        if ... {
            ...
            i++
        } else {
            ...
            j--
        }
    }
}
```

* Slinding Window