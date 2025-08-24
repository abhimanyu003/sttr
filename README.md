![sttr](./media/banner.png)

# sttr

[Website](https://abhimanyu003.github.io/sttr/)
| [Install](https://github.com/abhimanyu003/sttr#battery-installation)
| [Getting Started](https://github.com/abhimanyu003/sttr#books-guide)
| [CLI Reference](https://abhimanyu003.github.io/sttr/cli/sttr/)
| [Source Code](https://github.com/abhimanyu003/sttr)

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

# :movie_camera: Demo

![sttr demo](./media/demo.gif)

# :battery: Installation

#### Quick install

You can run the below `curl` to install it somewhere in your PATH for easy use. Ideally it will be installed at `./bin`
folder

```shell
curl -sfL https://raw.githubusercontent.com/abhimanyu003/sttr/main/install.sh | sh
```

#### Homebrew

If you are on macOS and using Homebrew, you can install `sttr` with the following:

```shell
brew install abhimanyu003/sttr/sttr
```

#### Snap

```shell
sudo snap install sttr
```

#### Arch Linux

```shell
yay -S sttr-bin
```

#### Docker

```shell
docker run -it --rm -e TERM=xterm-256color ghcr.io/abhimanyu003/sttr:latest 
```
> You can use docker image in your project from `ghcr.io/abhimanyu003/sttr:latest`


#### Winget

```shell
winget install -e --id abhimanyu003.sttr
```

#### Scoop

```shell
scoop bucket add sttr https://github.com/abhimanyu003/scoop-bucket.git
scoop install sttr
```

#### X-CMD

If you are a user of [x-cmd](https://x-cmd.com), you can run:

```shell
x install sttr
```

#### Webi

**macOS / Linux**

```shell
curl -sS https://webi.sh/sttr | sh
```

**Windows**

```shell
curl.exe https://webi.ms/sttr | powershell
```

See [here](https://webinstall.dev/sttr/)

#### Go

```shell
go install github.com/abhimanyu003/sttr@latest
```

#### Binary

**MacOS**
[Binary](https://github.com/abhimanyu003/sttr/releases/latest/download/sttr_Darwin_all.tar.gz) ( Multi-Architecture )

**Linux (Binaries)**
[amd64](https://github.com/abhimanyu003/sttr/releases/latest/download/sttr_Linux_x86_64.tar.gz) | [arm64](https://github.com/abhimanyu003/sttr/releases/latest/download/sttr_Linux_arm64.tar.gz) | [i386](https://github.com/abhimanyu003/sttr/releases/latest/download/sttr_Linux_i386.tar.gz)

**Windows (Exe)**
[amd64](https://github.com/abhimanyu003/sttr/releases/latest/download/sttr_Windows_x86_64.zip) | [arm64](https://github.com/abhimanyu003/sttr/releases/latest/download/sttr_Windows_arm64.zip) | [i386](https://github.com/abhimanyu003/sttr/releases/latest/download/sttr_Windows_i386.zip)

**FreeBSD (Binaries)**
[amd64](https://github.com/abhimanyu003/sttr/releases/latest/download/sttr_Freebsd_x86_64.tar.gz) | [arm64](https://github.com/abhimanyu003/sttr/releases/latest/download/sttr_Freebsd_arm64.tar.gz) | [i386](https://github.com/abhimanyu003/sttr/releases/latest/download/sttr_Freebsd_i386.tar.gz)

#### Manually

Download the pre-compiled binaries from the [Release!](https://github.com/abhimanyu003/sttr/releases) page and copy them
to the desired location.

# :books: Guide

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

# :boom: Supported Operations

#### Encode/Decode

- [x] **ascii85-encode** - Encode your text to Ascii85
- [x] **ascii85-decode** - Decode your Ascii85 text
- [x] **base32-decode** - Decode your Base32 text
- [x] **base32-encode** - Encode your text to Base32
- [x] **base64-decode** - Decode your Base64 text
- [x] **base64-encode** - Encode your text to Base64
- [x] **base85-encode** - Encode your text to Base85
- [x] **base85-decode** - Decode your Base85 text
- [x] **base64url-decode** - Decode your Base64 URL
- [x] **base64url-encode** - Encode your text to URL
- [x] **html-decode** - Unescape your HTML
- [x] **html-encode** - Escape your HTML
- [x] **rot13-encode** - Encode your text to ROT13
- [x] **url-decode** - Decode URL entities
- [x] **url-encode** - Encode URL entities
- [x] **morse-decode** - Decode your Morse code
- [x] **morse-encode** - Encode your text to Morse code

#### Hash

- [x] **bcrypt** - Get the bcrypt hash of your text
- [x] **md5** - Get the MD5 checksum of your text
- [x] **sha1** - Get the SHA1 checksum of your text
- [x] **sha256** - Get the SHA256 checksum of your text
- [x] **sha512** - Get the SHA512 checksum of your text
- [x] **xxh64** - Get the XXH64 checksum of your text

#### String

- [x] **camel** - Transform your text to camelCase
- [x] **kebab** - Transform your text to kebab-case
- [x] **lower** - Transform your text to lower case
- [x] **pascal** - Transform your text to PascalCase
- [x] **reverse** - Reverse Text ( txeT esreveR )
- [x] **slug** - Transform your text to slug-case
- [x] **snake** - Transform your text to snake_case
- [x] **title** - Transform your text to Title Case
- [x] **upper** - Transform your text to UPPER CASE

#### Lines

- [x] **count-lines** - Count the number of lines in your text
- [x] **reverse-lines** - Reverse lines
- [x] **shuffle-lines** - Shuffle lines randomly
- [x] **sort-lines** - Sort lines alphabetically
- [x] **unique-lines** - Get unique lines from list

#### Spaces

- [x] **remove-spaces** - Remove all spaces + new lines
- [x] **remove-newlines** - Remove all new lines

#### Count

- [x] **count-chars** - Find the length of your text (including spaces)
- [x] **count-lines** - Count the number of lines in your text
- [x] **count-words** - Count the number of words in your text

#### RGB/Hex

- [x] **hex-rgb** - Convert a #hex-color code to RGB
- [x] **hex-encode** - Encode your text Hex
- [x] **hex-decode** - Convert Hexadecimal to String

#### JSON

- [x] **json** - Format your text as JSON
- [x] **json-escape** - JSON Escape
- [x] **json-unescape** - JSON Unescape
- [x] **json-yaml** - Convert JSON to YAML text
- [x] **json-msgpack** - Convert JSON to MSGPACK
- [x] **msgpack-json** - Convert MSGPACK to JSON

#### YAML

- [x] **yaml-json** - Convert YAML to JSON text

#### Markdown

- [x] **markdown-html** - Convert Markdown to HTML

#### Extract

- [x] **extract-emails** - Extract emails from given text
- [x] **extract-ip** - Extract IPv4 and IPv6 from your text
- [x] **extract-urls** - Extract URLs your text ( we don't do ping check )

#### Other

- [x] **escape-quotes** - escape single and double quotes from your text
- [x] **completion** - generate the autocompletion script for the specified shell
- [x] **interactive** - Use sttr in interactive mode
- [x] **version** - Print the version of sttr
- [x] **zeropad** - Pad a number with zeros
- [x] **and adding more...**

# Featured On

These are the few locations where `sttr` was highlighted, many thanks to all of you. 
Please feel free to add any blogs/videos you may have made that discuss `sttr` to the list.

* [YouTube: The Giants of Open Source - DevOps Paradox](https://youtu.be/4nFRKbY_HVE?t=2529?ref=abhimanyu003/sttr)
* [Terminal Trove - Tool of the Week](https://terminaltrove.com/sttr/?ref=abhimanyu003/sttr)
* [nixCraft](https://www.cyberciti.biz/open-source/sttr-awesome-linux-unix-command-tool-for-transformation-string/?ref=abhimanyu003/sttr)

# Contribution

This project welcomes your PR and issues. For example, refactoring, adding features, correcting English, etc.

A quick development guide can be found
on. [Developer-Guides](https://github.com/abhimanyu003/sttr/wiki/Developer-Guides) wiki page.

If you need any help, you can contact me on [Twitter](https://twitter.com/abhimanyu003).

Thanks to all the people who already contributed!

<a href="https://github.com/abhimanyu003/sttr/graphs/contributors">
  <img src="https://contributors-img.web.app/image?repo=abhimanyu003/sttr" />
</a>

# License

[MIT](./LICENSE)

<!-- GitAds-Verify: 1L2K42WZSL98PJ8ZJ9H357E19QNY4KX3 -->
