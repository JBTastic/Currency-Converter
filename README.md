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
2. Extract the archive via command line:

    tar -xzf currency-converter-linux-amd64.tar.gz

3. Change to the extracted directory if necessary and run the binary:

    ./currency-converter-linux-amd64 -h

> **Note:**  
> The executable permission bit is preserved when unpacking with `tar` on Linux/macOS.  
> If you get a permission denied error, run:

    chmod +x [your-binary-name]

### Windows

1. Download the `.zip` archive.  
2. Extract it using Windows Explorer or any unzip tool.  
3. Run the executable, e.g.:

    currency-converter-windows-amd64.exe -h

---

## Usage Examples

    ./currency-converter-linux-amd64 -amount 100 -from EUR -to USD

or on Windows:

    currency-converter-windows-amd64.exe -amount 100 -from EUR -to USD

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