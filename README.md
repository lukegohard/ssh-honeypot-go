### ssh-honeypot-go

*****

**Dependencies:**

[gliderlabs/ssh]: https://github.com/gliderlabs/ssh

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

**Example:**

```bash
./ssh-honeypot-go -p 1234
```

****

**Flags:**

- '**-p**' ==> enter the honeypot server port(default: **2222**)

