## WC Implementiton in Go

- Simple version of the Unix utility wc that displays the number of lines, words, and bytes contained in each input file.

- First stage main go file will print the text by using [ioutil](https://pkg.go.dev/io/ioutil)
After that we open the file and count the line-words-characters by using bufio packages scanner method.