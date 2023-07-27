# go-jisho
Jisho Api Golang wrapper. 
this wrapper is very simple keep that in mind 

## Usage
you can pass to `w.Get()`  romaji/katakana/hiragana/kanji/english  word and recive data in the form of `WordData` struct.
To get raw data use `Search()` function.
```go
var w WordData
w.Get("犬")
translatino := w.FirstTransation() 
fmt.Printf("translatino: \n"translatino)
```
## contirbutino 
feel free to contibuite as you can 
