# PassGen

A fast, secure CLI tool for generating cryptographically strong random passwords with customizable character sets.

## Features

- Cryptographically secure random password generation
- Automatic clipboard integration
- Customizable character sets (numbers, symbols, uppercase, lowercase)
- Generate multiple passwords at once
- Fast and lightweight
- Cross-platform support

## Installation

### From Source
```bash
git clone https://github.com/aecsar/passgen
cd passgen
go build -o passgen
```

<!-- ### Direct Download -->
<!-- Download the latest binary from the [releases page](https://github.com/yourusername/passgen/releases). -->

## Usage

### Basic Usage
```bash
# Generate a 10-character password (default)
passgen

# Generate a 16-character password
passgen 16

# Generate 5 passwords of 12 characters each
passgen 12 --count 5
```

### Excluding Character Types
```bash
# Exclude numbers
passgen 16 --no-nums

# Exclude special characters
passgen 16 --no-symbols

# Exclude uppercase letters
passgen 16 --no-upper

# Exclude lowercase letters
passgen 16 --no-lower

# Combine exclusions (numbers and symbols)
passgen 20 -ns
```

### Multiple Passwords
```bash
# Generate 3 passwords
passgen --count 3

# Generate 5 passwords of 16 characters, no symbols
passgen 16 --no-symbols --count 5
```

## Command Reference

### Synopsis
```
passgen [length] [flags]
```

### Arguments
- `length` - Password length (default: 10)

### Flags
| Flag | Short | Description |
|------|-------|-------------|
| `--no-nums` | `-n` | Exclude numbers from password generation |
| `--no-symbols` | `-s` | Exclude special characters from password generation |
| `--no-upper` | `-u` | Exclude uppercase letters from password generation |
| `--no-lower` | `-l` | Exclude lowercase letters from password generation |
| `--count` | `-c` | Number of passwords to generate (default: 1) |
| `--help` | `-h` | Show help message |

### Character Sets
- **Lowercase**: `abcdefghijklmnopqrstuvwxyz`
- **Uppercase**: `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
- **Numbers**: `0123456789`
- **Special Characters**: `!@#$%^&*()_+-=[]{}|;:,.<>?~\``

## Examples

```bash
# Default 10-character password with all character types
passgen
# Output: rAo'X^S7#=

# 16-character password without special characters
passgen 16 --no-symbols
# Output: aB3kL9mN2pQ7rS1t

# Generate 3 passwords of 8 characters each, numbers only
passgen 8 -usl --count 3
# Output:
# 82947156
# 63051724
# 19386405

# Long password with mixed exclusions
passgen 24 -ns
# Output: AbCdEfGhIjKlMnOpQrStUvWx
```

## Security Notes

- Uses Go's `crypto/rand` package for cryptographically secure random generation
- Passwords are generated entirely in memory and not stored
- Each character is independently randomly selected from the available character set
- No predictable patterns or sequences are used

## Building from Source

### Build
```bash
git clone https://github.com/aecsar/passgen
cd passgen
go mod tidy
go build -o passgen
```

### Install Globally
```bash
go install
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) CLI framework
- Inspired by the need for quick, secure password generation
