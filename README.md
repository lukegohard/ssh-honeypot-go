### ssh-honeypot-go

*****

**Dependencies:**

- [gliderlabs/ssh](https://github.com/gliderlabs/ssh)
- [notify-send](https://man.cx/notify-send)

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

**Usage:**

```bash
./ssh-honeypot-go
```

```bash
./ssh-honeybot-go -p <port>
```

```bash
./ssh-honeybot-go -p <port> -n
```

**Example:**

```bash
./ssh-honeypot-go -p 1234 -n
```

****

**Flags:**

- '**-p**' ==> enter the honeypot server port(default: **2222**)
- '**-n**' ==> activate notifier service(default: false)

