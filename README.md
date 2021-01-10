# Roman-go
Simple go package for roman numerals encoding and decoding.

## Usage ##

```go
import "github.com/oguzhankarabulut/roman-go"	
func main()  {

  //Decode roman numeral string to int
	roman.Decode("MMDCCCXXXVII") // return 2837
  
  //Encode int to roman numreal string
	roman.Encode(2837) //return "MMDCCCXXXVII"
}
```

## Limitations ##
For correct answer use this package in 0, 4000 range.

## TODO ##

- [ ] Compatitable for higher 4000
