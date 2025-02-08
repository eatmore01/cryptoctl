## cryptoctl

cryptoctl is a lightweight CLI tool for simple crypto operations with text and files, it supports encrypting and decrypting base64 encoded strings and files and builded such as superstructure over [this is crypto library](https://github.com/eatmore01/cryptolib)

- [Requirements](#requirements)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)

### Requirements

- `CRYPTOCTL_PASSWORD` environment variable set to the password you want to use for the crypto operations
- `go` language installed for build the binary

### Prerequisites

First u need to set the `CRYPTOCTL_PASSWORD` environment variable or paste `CRYPTOCTL_PASSWORD="qwe"` before running every commands

```bash
export CRYPTOCTL_PASSWORD="123456"
```

##### or u can setup permamently in your shell

```bash
echo "export CRYPTOCTL_PASSWORD="123456"" >> ~/.zshrc
source ~/.zshrc
```

### Installation


#### Build from code
```bash
git clone https://github.com/eatmore01/cryptoctl.git
cd cryptoctl
go build -o cryptoctl main.go && sudo mv cryptoctl /usr/local/bin/
```

#### Install a binary from releases

##### for macos
```bash
curl -L https://github.com/eatmore01/cryptoctl/releases/download/v1.0.0/cryptoctl_darwin_amd64 -o cryptoctl
chmod +x cryptoctl
sudo mv cryptoctl /usr/local/bin/
```

##### for linux 

```bash
curl -L https://github.com/eatmore01/cryptoctl/releases/download/v1.0.0/cryptoctl_linux_amd64 -o cryptoctl
chmod +x cryptoctl
sudo mv cryptoctl /usr/local/bin/
```

#### Install a binary from releases for Windows manualy from here [here](https://github.com/eatmore01/cryptoctl/releases/tag/v1.0.0)



### Usage

```bash
cryptoctl encrypt-str "hello all this message i adress for you"
# or CRYPTOCTL_PASSWORD="123456" cryptoctl encrypt-str "str" if not setup CRYPTOCTL_PASSWORD permamently
```

```bash
cryptoctl decrypt-str "aGVsbG8gYWxsIHRoaXMgbWVzc2FnZSBpIGFkcmVzc2UgZm9yIHlvdQ=="
```

```bash
cryptoctl encrypt-file hello.txt
```

```bash
cryptoctl decrypt-file hello.txt.enc
```

