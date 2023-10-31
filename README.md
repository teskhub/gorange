# Gorange

Display a range with marked value in go.

## Example

First you define middle value (50) and radius length (5), so the "good" values are between 45 and 55:

```
range := gorange.New(50, 5)
```

Then display your range for given value, for example value of 49 should be close to the middle:

```
fmt.Println(range.GetRangeAsText(49))
```

The output is:

```
-----[---█|----]-----
```

When your value is out of range, but close enough to the border, it's still visible:

```
fmt.Println(range.GetRangeAsText(41))

-█---[----|----]-----
```

When your value is out of 2 x range, it's not visible:

```
fmt.Println(range.GetRangeAsText(37))

-----[----|----]-----
```

## Display options

You can set the result string's width by calling the `SetDisplayWidth` function, however
setting a value other than 5, 9, 13, 17... will result with an error. Default width is 21.

Example:

```
range := gorange.New(50, 5)
err := range.SetDisplayWidth(41)
if err != nil {
    fmt.Errorf("error setting display width: %v", err)
}

fmt.Println(range.GetRangeAsText(52.5))
```

Output:

```
----------[---------|----█----]----------
```