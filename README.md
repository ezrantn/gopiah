<p align="center">
  <img src="./assets/gopiah.png" width="180" alt="Gopiah Logo" />
</p>

<p align="center">A small library to convert any number to <a href="https://en.wikipedia.org/wiki/Indonesian_rupiah#:~:text=The%20rupiah%20(symbol%3A%20Rp%3B,referring%20to%20rupiah%20in%20coins.">Indonesian Rupiah currency</a></p>

This is a small Go library for converting numbers into Indonesian Rupiah (Rp) currency format. It's a convenient tool for formatting monetary values in the Indonesian style, with thousands separators and decimal points. Remember that this return value is of type string.

## Installation

To use this library, you can import it into your Go project like this:

```shell
go get github.com/ezrantn/gopiah
```

## Usage

To convert a number into Rupiah currency format, you can use the `gopiah.ConvertToRupiah` function:
```go
package main

import (
    "fmt"
    "github.com/ezrantn/gopiah"
)

func main() {
    number := 1234567
    rupiah := gopiah.ConvertToRupiah(number)
    fmt.Println(rupiah) // Output: "Rp. 1.234.567,00"
}
```
## Contributing
If you find any issues or have suggestions for improvements, feel free to create a pull request or open an issue on the GitHub repository. Your contributions are welcome! 😀

## License
This library is licensed under the MIT License. See the [LICENSE](https://github.com/ezrantn/gopiah/blob/main/LICENSE) file for details.

*Made with ❤️ by Ezra Natanael*
