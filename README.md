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


- [x] Base64 Encode
- [x] Base64 Decode
- [x] URL Encode
- [x] URL Decode
- [x] ROT13 Encode
- [x] String To Title
- [x] String To lower
- [x] String To UPPER
- [x] String To snake_case
- [x] String To Kebab
- [x] String To Slug
- [x] String To Camel
- [x] String Reverse
- [x] Count Number Characters
- [x] Count Words
- [x] Count Lines
- [x] MD5 Encode
- [x] SHA1 Encode
- [x] SHA256 Encode
- [x] SHA512 Encode
- [x] Format JSON
- [x] JSON To YAML
- [x] YAML To JSON
- [x] Hex To RGB
- [x] Hexadecimal To String
- [x] String to Hexadecimal
- [x] Sort Lines
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
