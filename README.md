![sttr](./media/banner.png)

# sttr

`sttr` is command line software that allows you to quickly run various transformation operations on the string.

```go
// With input prompt
sttr

// Direct string input
sttr -i "your string"
```

# Demo

![sttr demo](./media/demo.gif)


# Installation

## Quick install

You can run the below `curl` to install it somewhere in your PATH for easy use.
Ideally it will be installed at `./bin` folder

```go
curl -sfL https://raw.githubusercontent.com/abhimanyu003/sttr/main/install.sh | sh
```

## Homebrew

If you are on macOS and using Homebrew, you can install `sttr` with the following:

```go
brew tap abhimanyu003/sttr
brew install sttr
```

## Go 

```go
go install github.com/abhimanyu003/sttr@latest
```

## Manually

Download the pre-compiled binaries from the [Release!](https://github.com/abhimanyu003/sttr/releases) page and copy them to the desired location.

# Guide

* After installation simply run `sttr` command.

```go
// With input prompt

sttr 
// ( Press two enter to open operation menu )

// You can also provide string directly without any prompt.
sttr -i "your string"
```

* Press `/` to filter various operations.
* Can also press UP-Down arrows select various operations.


# Supported Operations

* Base64Encode
* Base64Decode
* URLEncode
* URLDecode
* ROT13Encode
* StringToTitle
* StringToLower
* StringToUpper
* StringToSnakeCase
* StringToKebab
* StringToSlug
* StringToCamel
* StringReverse
* CountNumberCharacters
* CountWords
* CountLines
* MD5Encode
* SHA1Encode
* SHA256Encode
* SHA512Encode
* FormatJSON
* JSONToYAML
* YAMLToJSON
* HexToRGB
* SortLines
* and adding more....

# License

[MIT](./LICENSE)