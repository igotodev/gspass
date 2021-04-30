# genpass

## About
Simple, easy-to-use package for strong passwords

## Simple install
```
$ go get github.com/igotodev/genpass
```

## Example of usage
func GetStrongPass takes the required password length (n int) 
and returns a password of random characters as a string,
lowercase and uppercase letters of length n and error, 
the maximum value for n is 32767
```
func main() {
	x, err := genpass.GetPass(20) // max arg 32767
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(x)
}
```
