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

**Hostkey file must be in config directory! And with no password!**

****

**Changing Username and Password:**

Edit the *config/config.json* file.

Default:

```json
{
	"creator": "Ex0dIa-dev",
	"auth": {
		"user":     "root",
		"password": "toor"
	}
}
```

****

**Run on Docker:**

*Build image:*

```bash
//First generate your hostkey and put it in config/ -->(not obligatory,will be auto-generated)
//Then you can build image
docker build -t <image-name> .

//Example
docker build -t ssh-honeypot-go .
```

*Run a container:*

```bash
docker run --rm -v $PWD/config:/app/config -v $PWD/logs:/app/logs -p <host_port>:<honeypot_port> <image_name>
//Example
docker run --rm -v $PWD/config:/app/config -v $PWD/logs:/app/logs -p 22:2222 ssh-honeypot-go

//You can use flags too
docker run --rm -v $PWD/config:/app/config -v $PWD/logs:/app/logs -p <host_port>:<honeypot_port> <image_name> -port <honeypot_port> -log

//Example
docker run --rm -v $PWD/config:/app/config -v $PWD/logs:/app/logs -p 22:1234 ssh-honeypot-go -port 1234 -log
```

**Notification Service doesn't work! Using it will crash the app!**

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
./ssh-honeypot-go -port <port> -log -log-all
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
- '**-notify**' ==> activate notification service(default: false)
- '**-log**' ==> activate logging(logs path: "logs/")
- **'-log-all'** ==> logging all attempts(terminal, and notification), failed too

****

**TODO LIST:**

- [x] *adding a fake shell as sessionHandler(temporary)*
- [x] *adding a Dockerfile*
- [x] *log collected data in a file*

****

**Support Project**
Support the project with a donation
[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/donate?hosted_button_id=Z93ULXU3H2TQC)
