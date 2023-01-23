# onetab-export-to-json

Export URLs directly from [OneTab](https://chrome.google.com/webstore/detail/onetab/chphlpgkkbolifaimnlloiipkdnihall)'s LevelDB database to JSON.

## Installation

```sh
$ go install github.com/jianyuan/onetab-export-to-json@latest

# go will build the program and install the binary to $GOPATH/bin
$ $GOPATH/bin/onetab-export-to-json
```

## Usage

```
$ onetab-export-to-json
  -i string
        LevelDB database path (shorthand)
  -input string
        LevelDB database path
  -o string
        Output file path ("-" to print to standard output) (default "-")
  -output string
        Output file path ("-" to print to standard output) (default "-")
```

Typical LevelDB locations:

| OS      | Path                                                                                  |
| ------- | ------------------------------------------------------------------------------------- |
| Windows | `C:\Users\{USER}\AppData\Local\Google\Chrome\User Data\Default\Local Storage\leveldb` |

Output to standard output:

```sh
$ onetab-export-to-json -input {PATH}
```

Output to file:

```sh
$ onetab-export-to-json -input {PATH} -output tabs.json
```
