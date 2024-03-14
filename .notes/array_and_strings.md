# Array and Strings

## Two Pointers

Start at the edge of the slice and collapse toward the middly until the right pointer meets the left pointer.

```go
func twoPointer(arr []K){
    // Init pointer on edges
    i, j := 0, len(arr) - 1

    // loop until pointers colide
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

## Slinding Window

Add element from the right until the constrsin is broken, then remove element form the left

```go
func slidingWindow(arr []K, k int) K {
    // both pointer start at 0
    left, right int
    // we have a candidate and result. Candidate take over if it beats the best result
    curr, ans K

    for right < len(arr); right++{
        // add element on the righ
        curr += ...
        for curr > k {
            curr -= ...
            left++
        }

        ans = max(ans, curr)
    }

    return ans
}
```