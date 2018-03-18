### Anami Challenges

#### Round 2 sample

from the root of this repo run the following command:

`$ go run ./cmd/usersort/main.go` 

Notes: 

I am not a big fan of the strategy pattern in Go, and I can't remember any time where I've seen someone implement it in a Go library. I went with the builder pattern as I felt it lent itself better to the problem at hand. A few things I didn't do, but would given a little more time, refactor my the students type field `s` to be of type `[]*student` instead of the `map`.  It would make adding grades slightly more cumbersome, but mehhhh, it would make the sorting methods much more readable. Most of the sortings behavior I feel would be better wrap around a type something like the following `type studentList []*student`.