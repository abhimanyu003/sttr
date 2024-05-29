![sttr](https://github.com/abhimanyu003/sttr/raw/main/media/banner.png)

# sttr

`sttr` is command line software that allows you to quickly run various transformation operations on the string.

```shell
// With input prompt
sttr

// Direct input
sttr md5 "Hello World"

// File input
sttr md5 file.text
sttr base64-encode image.jpg

// Reading from different processor like cat, curl, printf etc..
echo "Hello World" | sttr md5
cat file.txt | sttr md5

// Writing output to a file
sttr yaml-json file.yaml > file-output.json
```

# 🎥 Demo

![sttr demo](https://github.com/abhimanyu003/sttr/raw/main/media/demo.gif)


# 🔋 Installation

#### Quick install

You can run the below `curl` to install it somewhere in your PATH for easy use.
Ideally it will be installed at `./bin` folder

```shell
curl -sfL https://raw.githubusercontent.com/abhimanyu003/sttr/main/install.sh | sh
```

#### Homebrew

If you are on macOS and using Homebrew, you can install `sttr` with the following:

```shell
brew tap abhimanyu003/sttr
brew install sttr
```

#### Go

```shell
go install github.com/abhimanyu003/sttr@latest
```

#### Manually

Download the pre-compiled binaries from the [Release!](https://github.com/abhimanyu003/sttr/releases) page and copy them to the desired location.

# 📚 Guide

* After installation simply run `sttr` command.

```shell
// For interactive menu
sttr
// Provide your input
// Press two enter to open operation menu
// Press `/` to filter various operations.
// Can also press UP-Down arrows select various operations.
```

* Working with help.

```shell
sttr -h

// Example
sttr zeropad -h
sttr md5 -h
```

* Working with files input.

```shell
sttr {command-name} {filename}

sttr base64-encode image.jpg
sttr md5 file.txt
sttr md-html Readme.md
```

* Writing output to file.

```shell
sttr yaml-json file.yaml > file-output.json
```

* Taking input from other command.

```shell
curl https://jsonplaceholder.typicode.com/users | sttr json-yaml
```

* Chaining the different processor.

```shell
sttr md5 hello | sttr base64-encode

echo "Hello World" | sttr base64-encode | sttr md5
```


# 💥 Supported Operations

* [sttr base32-decode]({{< relref "sttr_base32-decode.md" >}})	 - Decode your Base32 text
* [sttr base32-encode]({{< relref "sttr_base32-encode.md" >}})	 - Encode your text to Base32
* [sttr base64-decode]({{< relref "sttr_base64-decode.md" >}})	 - Decode your Base64 text
* [sttr base64-encode]({{< relref "sttr_base64-encode.md" >}})	 - Encode your text to Base64
* [sttr bcrypt]({{< relref "sttr_bcrypt.md" >}})	 - Get the bcrypt hash of your text
* [sttr camel]({{< relref "sttr_camel.md" >}})	 - Transform your text to CamelCase
* [sttr completion]({{< relref "sttr_completion.md" >}})	 - generate the autocompletion script for the specified shell
* [sttr count-chars]({{< relref "sttr_count-chars.md" >}})	 - Find the length of your text (including spaces)
* [sttr count-lines]({{< relref "sttr_count-lines.md" >}})	 - Count the number of lines in your text
* [sttr count-words]({{< relref "sttr_count-words.md" >}})	 - Count the number of words in your text
* [sttr extract-emails]({{< relref "sttr_extract-emails.md" >}})	 - Extract emails from given text
* [sttr hex-encode]({{< relref "sttr_hex-encode.md" >}})	 - Encode your text Hex
* [sttr hex-rgb]({{< relref "sttr_hex-rgb.md" >}})	 - Convert a #hex-color code to RGB
* [sttr html-decode]({{< relref "sttr_html-decode.md" >}})	 - Unescape your HTML
* [sttr html-encode]({{< relref "sttr_html-encode.md" >}})	 - Escape your HTML
* [sttr interactive]({{< relref "sttr_interactive.md" >}})	 - Use sttr in interactive mode
* [sttr json]({{< relref "sttr_json.md" >}})	 - Format your text as JSON
* [sttr json-yaml]({{< relref "sttr_json-yaml.md" >}})	 - Convert JSON to YAML text
* [sttr kebab]({{< relref "sttr_kebab.md" >}})	 - Transform your text to kebab-case
* [sttr lower]({{< relref "sttr_lower.md" >}})	 - Transform your text to lower case
* [sttr markdown-html]({{< relref "sttr_markdown-html.md" >}})	 - Convert Markdown to HTML
* [sttr md5]({{< relref "sttr_md5.md" >}})	 - Get the MD5 checksum of your text
* [sttr reverse]({{< relref "sttr_reverse.md" >}})	 - Reverse Text ( txeT esreveR )
* [sttr rot13-encode]({{< relref "sttr_rot13-encode.md" >}})	 - Encode your text to ROT13
* [sttr sha1]({{< relref "sttr_sha1.md" >}})	 - Get the SHA-1 checksum of your text
* [sttr sha256]({{< relref "sttr_sha256.md" >}})	 - Get the SHA-256 checksum of your text
* [sttr sha512]({{< relref "sttr_sha512.md" >}})	 - Get the SHA-512 checksum of your text
* [sttr slug]({{< relref "sttr_slug.md" >}})	 - Transform your text to slug-case
* [sttr snake]({{< relref "sttr_snake.md" >}})	 - Transform your text to snake_case
* [sttr sort-lines]({{< relref "sttr_sort-lines.md" >}})	 - Sort lines alphabetically
* [sttr title]({{< relref "sttr_title.md" >}})	 - Transform your text to Title Case
* [sttr upper]({{< relref "sttr_upper.md" >}})	 - Transform your text to UPPER CASE
* [sttr url-decode]({{< relref "sttr_url-decode.md" >}})	 - Decode URL entities
* [sttr url-encode]({{< relref "sttr_url-encode.md" >}})	 - Encode URL entities
* [sttr version]({{< relref "sttr_version.md" >}})	 - Print the version of sttr
* [sttr yaml-json]({{< relref "sttr_yaml-json.md" >}})	 - Convert YAML to JSON text
* [sttr zeropad]({{< relref "sttr_zeropad.md" >}})	 - Pad a number with zeros

# Contribution

This project welcomes your PR and issues.
For example, refactoring, adding features, correcting English, etc.
If you need any help, you can contact me on [Twitter](https://twitter.com/abhimanyu003).

Thanks to all the people who already contributed!

<a href="https://github.com/abhimanyu003/sttr/graphs/contributors">
  <img src="https://contributors-img.web.app/image?repo=abhimanyu003/sttr" />
</a>

# License

[MIT](https://github.com/abhimanyu003/sttr/blob/main/LICENSE)