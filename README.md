# Currency Converter CLI

A simple CLI tool to convert currencies using the Frankfurter API.

---

## Download & Installation

You can download prebuilt binaries from the [Releases](https://github.com/yourusername/yourrepo/releases) page.

### Supported Platforms

- Linux (amd64)  
- Windows (amd64)  
- macOS (amd64 and arm64)  

---

## Archive Formats

- Linux and macOS binaries are packaged as `.tar.gz`  
- Windows binaries are packaged as `.zip`  

---

## How to Use

### Linux / macOS

1. Download the appropriate `.tar.gz` archive.  
2. Extract the archive using your preferred method, or via command line:

```bash
    tar -xzf currency-converter-linux-amd64.tar.gz
```

3. Change to the extracted directory if necessary and run the binary:

```bash
    ./currency-converter -h
```

> **Note:**  
> The executable permission bit is preserved when unpacking with `tar` on Linux/macOS.  
> If you get a permission denied error, run:

```bash
    chmod +x currency-converter
```

### Windows

1. Download the `.zip` archive.  
2. Extract it using Windows Explorer or any unzip tool.  
3. Run the executable, e.g.:

```powershell
    currency-converter.exe -h
```

---

## How to Compile Yourself

To compile the Currency Converter CLI from source, you need to have Go installed (version 1.22 or higher recommended).

1. Clone the repository:

```bash
    git clone https://github.com/JBTastic/Currency-Converter.git
    cd Currency-Converter
```

2. Build the binary for your platform:

```bash
    # For your current platform:
    go build -o currency-converter

    # Or cross-compile, e.g. for Linux amd64:
    GOOS=linux GOARCH=amd64 go build -o currency-converter-linux-amd64

    # For Windows amd64:
    GOOS=windows GOARCH=amd64 go build -o currency-converter-windows-amd64.exe

    # For macOS amd64:
    GOOS=darwin GOARCH=amd64 go build -o currency-converter-darwin-amd64

    # For macOS arm64:
    GOOS=darwin GOARCH=arm64 go build -o currency-converter-darwin-arm64
```

3. Run the built binary:

```bash
    ./currency-converter -h
```

> **Note:** When cross-compiling, make sure your Go environment supports it and that you use the correct `GOOS` and `GOARCH` values.

---

## Usage Examples

```bash
./currency-converter -amount 100 -from EUR -to USD

```

or on Windows:

```powershell
    currency-converter.exe -amount 100 -from EUR -to USD
```

---

## Flags

- `-amount` float64: Amount to convert (required)  
- `-from` string: Source currency code (e.g. EUR) (required)  
- `-to` string: Target currency code (e.g. USD) (required)  
```powershell
    currency-converter.exe -amount 100 -from EUR -to USD
```

---

## Flags

- `-amount` float64: Amount to convert (required)  
- `-from` string: Source currency code (e.g. EUR) (required)  
- `-to` string: Target currency code (e.g. USD) (required)  
- `-currencies`: List all available currencies  
- `-abbr` string: Search currencies by name substring  

---

## Support

Report issues or feature requests on the GitHub repository.

---
