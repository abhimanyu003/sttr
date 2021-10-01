![sttr](./media/banner.png)

# sttr

`sttr` is command line software that allows you to quickly run various transformation operations on the string.

```go
// With input prompt
sttr

// Direct string input
echo "Hello World" | sttr md5

// File input
cat file.txt | sttr base64-encode

// Writing output to a file
cat file.yml | sttr yaml-json > file-output.json
```

# :movie_camera: Demo

![sttr demo](./media/demo.gif)


# :battery: Installation

#### Quick install

You can run the below `curl` to install it somewhere in your PATH for easy use.
Ideally it will be installed at `./bin` folder

```go
curl -sfL https://raw.githubusercontent.com/abhimanyu003/sttr/main/install.sh | sh
```

#### Homebrew

If you are on macOS and using Homebrew, you can install `sttr` with the following:

```go
brew tap abhimanyu003/sttr
brew install sttr
```

#### Go 

```go
go install github.com/abhimanyu003/sttr@latest
```

#### Manually

Download the pre-compiled binaries from the [Release!](https://github.com/abhimanyu003/sttr/releases) page and copy them to the desired location.

# :books: Guide

* After installation simply run `sttr` command.

```go
// For interactive menu
sttr 
// Provide your input
// Press two enter to open operation menu
// Press `/` to filter various operations.
// Can also press UP-Down arrows select various operations.
```

* Working with help.

```go
sttr -h
sttr zeropad -h
```

* Working with files input.

```go
cat file-input.jpg | sttr base64-encode
```

* Writing output to file.

```go
cat words.txt | sttr count-chars > count.txt
```

* Taking input from other command.

```go
curl https://jsonplaceholder.typicode.com/users | sttr json-yaml
```

* Chaining the different processor.

```go
echo "Hello World" | sttr base64-encode | sttr md5
```


# :boom: Supported Operations


- [x] **base32-decode** - Decode your base32 text
- [x] **base32-encode** - Encode your text to Base32
- [x] **base64-decode** - Decode your base64 text
- [x] **base64-encode** - Encode your text to Base64
- [x] **bcrypt** - Get the Bcrypt hash of your text
- [x] **camel** - Transform your text to CamelCase
- [x] **completion** - generate the autocompletion script for the specified shell
- [x] **count-chars** - Find the length of your text (including spaces)
- [x] **count-lines** - Count the number of lines in your text
- [x] **count-words** - Count the number of words in your text
- [x] **extract-emails** - Extract emails from given text
- [x] **hex-encode** - Encode your text Hex
- [x] **hex-rgb** - Convert a #hex-color code to RGB
- [x] **html-decode** - Unescape your HTML
- [x] **html-encode** - Escape your HTML
- [x] **interactive** - Use sttr in interactive mode
- [x] **json** - Format your text as JSON
- [x] **json-yaml** - Convert JSON to YAML text
- [x] **kebab** - Transform your text to kebab-case
- [x] **lower** - Transform your text to lower case
- [x] **markdown-html** - Convert Markdown to HTML
- [x] **md5** - Get the MD5 checksum of your text
- [x] **reverse** - Reverse Text ( txeT esreveR )
- [x] **rot13-encode** - Encode your text to ROT13
- [x] **sha1** - Get the SHA1 checksum of your text
- [x] **sha256** - Get the SHA256 checksum of your text
- [x] **sha512** - Get the SHA512 checksum of your text
- [x] **slug** - Transform your text to slug-case
- [x] **snake** - Transform your text to snake_case
- [x] **title** - Transform your text to Title Case
- [x] **upper** - Transform your text to UPPER CASE
- [x] **url-decode** - Decode URL entities
- [x] **url-encode** - Encode URL entities
- [x] **version** - Print the version of sttr
- [x] **yaml-json** - Convert YAML to JSON text
- [x] **zeropad** - Pad a number with zeros
- [x] **and adding more....**

# Contribution

This project welcomes your PR and issues.
For example, refactoring, adding features, correcting English, etc.
If you need any help, you can contact me on [Twitter](https://twitter.com/abhimanyu003).

Thanks to all the people who already contributed!

<a href="https://github.com/abhimanyu003/sttr/graphs/contributors">
  <img src="https://contributors-img.web.app/image?repo=abhimanyu003/sttr" />
</a>

# License

[MIT](./LICENSE)
