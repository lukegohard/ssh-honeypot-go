### ssh-honeypot-go

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

**Generate Host Private Key**(not obligatory, default hostkey will be autogenerated):

```bash
ssh-keygen -t <type> -b <bits> -N "" -f <output_filepath>
```

```bash
ssh-keygen -t rsa -b 2048 -N "" -f hostkey_rsa
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

**Example:**

```bash
./ssh-honeypot-go -p 1234 -n
```

```bash
./ssh-honeypot-go -p 1234 -k hostkey_rsa -n
```

****

**Flags:**

- '**-p**' ==> enter the honeypot server port(default: **2222**)
- '**-n**' ==> activate notifier service(default: false)
- '**-k**' ==> enter the filepath of host private key

****

**TODO:**

- [ ] *adding a fake shell as sessionHandler*
