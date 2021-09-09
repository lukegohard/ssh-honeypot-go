### ssh-honeypot-go

[![Go Report Card](https://goreportcard.com/badge/github.com/Ex0dIa-dev/ssh-honeypot-go)](https://goreportcard.com/report/github.com/Ex0dIa-dev/ssh-honeypot-go)
*****

**Dependencies:**

- [gliderlabs/ssh](https://github.com/gliderlabs/ssh)
- [notify-send](https://man.cx/notify-send)
- [openssh](https://www.openssh.com/) (not obligatory, you need this only for **ssh-keygen** used for generate **host private key**)

****

**Build:**

```bash
go build
```

or 

```bash
go build -ldflags="-s -w"
```

for a lightweight binary.

****

**Generate Host Private Key**(not obligatory, default host key will be auto generated):

```bash
ssh-keygen -t <type> -b <bits> -N "" -f config/hostkey_rsa
```

```bash
ssh-keygen -t rsa -b 2048 -N "" -f config/hostkey_rsa
```

**Hostkey file must be in config directory!**

****

**Changing Username and Password:**

Edit the *config.json* file.

Default:

```json
{
	"creator": "Ex0dIa-dev",
	"auth": {
		"user": 	"root",
		"password": "toor"
	}
}
```

****

**Usage:**

```bash
./ssh-honeypot-go
```

```bash
./ssh-honeypot-go -port <port>
```

```bash
./ssh-honeypot-go -port <port> -notify
```

```bash
./ssh-honeypot-go -port <port> -notify -log
```

```bash
./ssh-honeypot-go -port <port> -log-all
```

**Example:**

```bash
./ssh-honeypot-go -port 1234 -notify
```

```bash
./ssh-honeypot-go -port 1234 -notify -log
```

****

**Flags:**

- '**-port**' ==> enter the honeypot server port(default: **2222**)
- '**-notify**' ==> activate notifier service(default: false)
- '**-log**' ==> activate logging(logs path: "logs/")
- **'-log-all'** ==> logging all attempts(terminal, and notification), failed too

****

**TODO LIST:**

- [ ] *adding a fake shell as sessionHandler*
- [x] log collected ip addresses in a file

