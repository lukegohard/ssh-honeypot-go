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
ssh-keygen -t <type> -b <bits> -N "" -f <output_filepath>
```

```bash
ssh-keygen -t rsa -b 2048 -N "" -f hostkey_rsa
```

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
./ssh-honeypot-go -p <port>
```

```bash
./ssh-honeypot-go -p <port> -n
```

```bash
./ssh-honeypot-go -p <port> -n -k <host private key filepath>
```

```bash
./ssh-honeypot-go -p <port> -n -li
```

```bash
./ssh-honeypot-go -p <port> -la
```

**Example:**

```bash
./ssh-honeypot-go -p 1234 -n
```

```bash
./ssh-honeypot-go -p 1234 -k hostkey_rsa -n
```

```bash
./ssh-honeypot-go -p 1234 -k hostkey_rsa -n -li
```



****

**Flags:**

- '**-p**' ==> enter the honeypot server port(default: **2222**)
- '**-n**' ==> activate notifier service(default: false)
- '**-k**' ==> enter the filepath of host private key
- '**-li**' ==> activate ip address logging(logs path: "./logs/ip-address")
- **'-la'** ==> logging all attempts, failed too

****

**TODO LIST:**

- [ ] *adding a fake shell as sessionHandler*
- [x] log collected ip addresses in a file

