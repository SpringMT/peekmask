# peekmask

A simple Go library for masking parts of strings with full UTF-8 support (including Japanese and emoji).

## Features

- Mask strings while keeping prefix and suffix visible
- Full UTF-8 support (Japanese, emoji, and other multi-byte characters)
- Customizable mask character
- Minimum mask ratio setting (for security)
- Simple and easy-to-use API

## Installation

```bash
go get github.com/SpringMT/peekmask
```

## Usage

### Simple Usage

Mask a string with default settings (use `*` as mask):

```go
package main

import (
    "fmt"
    "github.com/SpringMT/peekmask"
)

func main() {
    masked := peekmask.Mask("1234567890", 2, 3)
    fmt.Println(masked) // Output: 12*****890
    
    masked = peekmask.Mask("メールアドレス", 2, 2)
    fmt.Println(masked) // Output: メー***レス
}
```

### Custom Configuration

Create a `Masker` with custom settings:

```go
package main

import (
    "fmt"
    "github.com/SpringMT/peekmask"
)

func main() {
    // Create a Masker with custom settings
    masker := peekmask.New(
        '■',  // Mask character
        3,    // Number of characters to keep from the start
        3,    // Number of characters to keep from the end
        0.4,  // Minimum mask ratio (40%)
    )
    
    masked := masker.Mask("1234567890")
    fmt.Println(masked) // Output: 123■■■■890
    
    masked = masker.Mask("こんにちは世界")
    fmt.Println(masked) // Output: ■■■■■■■
}
```

## API

### `Mask(s string, prefixLength, suffixLength int) string`

Masks a string using default settings.
- `MaskChar`: `'*'`
- `MinMaskRatio`: `0.3`

### `type Masker`

A struct that holds the configuration for masking.

- `MaskChar`: The character used for masking (e.g., `'*'`, `'■'`)
- `PrefixLength`: Number of characters to keep from the start
- `SuffixLength`: Number of characters to keep from the end
- `MinMaskRatio`: Minimum masking ratio (0.0-1.0)

### `New(maskChar rune, prefixLength, suffixLength int, minMaskRatio float64) *Masker`

Creates a new `Masker` with custom settings.

### `(m *Masker) Mask(s string) string`

Masks a string based on the configuration.

## About Minimum Mask Ratio

`MinMaskRatio` ensures that if the masking ratio falls below the specified value, all characters will be masked instead.

Example:
```go
masker := peekmask.New('*', 4, 4, 0.5)
// For string "1234567890" (10 characters)
// Masked characters: 10 - 4 - 4 = 2 characters
// Mask ratio: 2/10 = 0.2 (20%)
// Since 0.2 < 0.5, all characters will be masked
masked := masker.Mask("1234567890")
fmt.Println(masked) // Output: **********
```

This ensures that even with short strings or configurations that would mask only a few characters, the information is still protected.

## Testing

```bash
go test -v
```

Benchmarks:

```bash
go test -bench=.
```

## License

MIT License

