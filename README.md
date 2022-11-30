# from-go-duration

tiny helper for parsing go durations in nu shell


# wrap in a nu func

```
def "from go-duration" [] {
    $in | from-go-duration | complete | get stdout | into duration 
} 
```


## example

```nu
k get node | detect columns | update AGE {|it| $it.AGE | from go-duration } | sort-by AGE -r
```
