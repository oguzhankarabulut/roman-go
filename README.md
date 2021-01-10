# Roman-go
Simple go package for roman numerals encoding and decoding.

## Usage ##

```go
import "github.com/oguzhankarabulut/roman-go"	
func main()  {

  //Decode roman numeral string to int
	roman.Decode("MMDCCCXXXVII") // returns 2837
  
  //Encode int to roman numreal string
	roman.Encode(2837) //returns "MMDCCCXXXVII"
}
```

## Limitations ##
For correct answer use this package in 0, 4000 range.

## TODO ##

- [ ] Compatitable for greater 4000
