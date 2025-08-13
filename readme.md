# gorle
Just an educational project containing implementation of **RLE algorithm decoding** for unicode strings.

## Usage example
```go
	decoded, _ := gorle.Decode("qwe\\45", gorle.WithEscapeChar('\\'), gorle.WithEscapeSeq(true))
	fmt.Println(decoded) // prints "qwe44444"
```
