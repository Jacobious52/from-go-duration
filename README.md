# from-go-duration

tiny helper for parsing go durations in nu shell

## example

```nu
k get node | dc | update AGE {|it| $it.AGE | from-go-duration | complete | get stdout | into duration } | sort-by AGE -r
```