# Dump Project

The `htmldump` package is a tool for developers, designed to simplify the visualization of Go data structures by converting them into HTML tables. This functionality is particularly useful during development and debugging, as it provides a clear and organized view of complex data.

## Supported Data Types

The `htmldump` package supports a wide range of Go data types, including:

- Structs (including nested structs, as demonstrated in the `pack` example)
- Slices
- Maps
- Basic types (e.g., strings, integers, floats)

By leveraging the `htmldump` package, developers can quickly inspect and analyze their data in a user-friendly format, enhancing productivity and reducing debugging time.

## Example

The `example/example.go` file provides a complete example of how to use the `htmldump` package. It includes:

## How to Run

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd dump
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the example:
   ```bash
   go run example/example.go
   ```

4. The generated HTML file (`example.html`) will open in your default browser.

## Dependencies

- [gofakeit](https://github.com/brianvoe/gofakeit): Used for generating random data.

## License

This project is licensed under the MIT License. See the LICENSE file for details.