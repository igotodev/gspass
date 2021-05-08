# gspass 
Strong Password Generator
## About
Simple, lightweight, easy-to-use package for creating strong passwords or tokens

## Simple install
```
$ go get github.com/igotodev/gspass
```

## Example of usage
func GetPass takes the required password length `(n int)` 
and returns a password of random characters as a string,
lowercase and uppercase letters, numbers and special characters 
with the length of `n` and `error`, 
the maximum value for `n` is 32767
```
func main() {
	strPass, err := gspass.GetPass(20) // max arg 32767
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strPass)
}
```
or func GetPassDL for password/token without special characters 
```
func main() {
	strToken, err := gspass.GetPassDL(20) // max arg 32767
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strToken)
}
```
