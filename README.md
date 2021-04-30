# genpass

## About
Simple, easy-to-use package for strong passwords

## Simple install

```
$ go get github.com/igotodev/genpass
```

## Example of usage

```
func main() {
	x, err := genpass.GetPass(20) // max arg 32767
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(x)
}

```
