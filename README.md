# Go Rupiah Currency Converter

This is a small Go library for converting numbers into Indonesian Rupiah (Rp) currency format. It's a convenient tool for formatting monetary values in the Indonesian style, with thousands separators and decimal points.

## Installation

To use this library, you can import it into your Go project like this:

```shell
go get github.com/umjiiii/gopiah
```

## Usage

To convert a number into Rupiah currency format, you can use the gopiah.ConvertToRupiah function:
```go
package main

import (
    "fmt"
    "github.com/umjiiii/gopiah"
)

func main() {
    number := 1234567
    rupiah := gopiah.ConvertToRupiah(number)
    fmt.Println(rupiah) // Output: "Rp. 1.234.567,00"
}
```
## Contributing
If you find any issues or have suggestions for improvements, feel free to create a pull request or open an issue on the GitHub repository. Your contributions are welcome!

## License
This library is licensed under the MIT License. See the [LICENSE](https://github.com/umjiiii/gopiah/blob/main/LICENSE) file for details.

## Author
- umjiii

Enjoy using this library to format your numbers into Rupiah currency! If you have any questions or need assistance, don't hesitate to reach out to the author or the community.
