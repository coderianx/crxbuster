# CrxBuster

A fast, minimal directory & path brute forcer written in Go.

Point it at a URL and a wordlist; CrxBuster fires a request per wordlist entry
and reports every response that isn't a plain `404`.

> ⚠️ **Disclaimer** — CrxBuster is built for authorized penetration testing,
> CTFs and security research. Only use it against systems you own or have
> **explicit written permission** to test. Unauthorized scanning is illegal in
> most jurisdictions. You are responsible for your own actions.

---

## Features

- Single static binary, zero runtime dependencies
- Concurrency-friendly, configurable HTTP timeout
- Automatic `/` normalization — wordlist entries like `/admin` and `admin/`
  both work
- Skips `404 Not Found` noise, keeps the interesting status codes
- Colored status output grouped by response class
- Built-in help screen

## Installation

### Requirements

- Go 1.21 or newer (module targets `go 1.27.1`)

### Install with `go install` (recommended)

Installs the latest tagged release straight into `$GOBIN` / `$GOPATH/bin`:

```bash
go install github.com/coderianx/crxbuster/cmd/crxbuster@latest
```

Make sure your Go bin directory is on `PATH`:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Then run it from anywhere:

```bash
crxbuster -h
```

To install a specific version:

```bash
go install github.com/coderianx/crxbuster/cmd/crxbuster@v0.1.0
```

### Build from source

```bash
git clone https://github.com/coderianx/crxbuster.git
cd crxbuster
go build -o crxbuster ./cmd/crxbuster
```

Or run it directly without building:

```bash
go run ./cmd/crxbuster -mode 1 -u https://example.com -w wordlist.txt
```

## Usage

```
crxbuster -mode <mode> [options]
```

### Flags

| Flag    | Type     | Default | Description                                  |
| ------- | -------- | ------- | -------------------------------------------- |
| `-mode` | `int`    | `0`     | Scan mode, see **Modes** below               |
| `-u`    | `string` | `""`    | Target URL, e.g. `https://example.com`       |
| `-w`    | `string` | `""`    | Path to the wordlist file                    |
| `-t`    | `int`    | `10`    | HTTP timeout in seconds                      |
| `-v`    | `bool`   | `false` | Print version and exit                       |
| `-h`    | `bool`   | `false` | Show the help screen and exit                |

### Modes

| Mode | Name     | Description                                        |
| ---- | -------- | -------------------------------------------------- |
| `1`  | URL scan | Brute force paths on a single target URL           |

Running without `-mode` prints the help screen.

### Examples

Basic directory scan:

```bash
crxbuster -mode 1 -u https://example.com -w wordlist.txt
```

Aggressive settings with a short timeout:

```bash
crxbuster -mode 1 -u https://example.com -w wordlist.txt -t 5
```

Show help / version:

```bash
crxbuster -h
crxbuster -v
```

### Output legend

| Prefix | Meaning                                      |
| ------ | -------------------------------------------- |
| `[+]`  | `2xx` — success, found                       |
| `[>]`  | `3xx` — redirect, often interesting          |
| `[!]`  | `4xx` — client error, except `404` (skipped) |
| `[X]`  | `5xx` — server error                         |

Example:

```
[INFO] Scan Starting...
[+] 200 https://example.com/admin
[>] 301 https://example.com/login
[!] 403 https://example.com/.env
[X] 500 https://example.com/api/export
```

## Project structure

```
cmd/crxbuster/main.go        CLI entrypoint, flag parsing
internal/commands/url_scan.go URL scan mode implementation
internal/helpers/help.go     Help screen
internal/helpers/version.go  Version string
internal/helpers/modes.go    Mode definitions
```

## Contributing

Issues and pull requests are welcome. Please keep the code `gofmt`-ed and run
`go vet ./...` before opening a PR.

## License

Released under the [MIT License](LICENSE).

## Author

[coderianx](https://github.com/coderianx)
