# fixer

async io.Writer that prefixes, suffixes and infixes with functions and is thread safe

## prefixing

[examples/prefix/main.go](examples/prefix/main.go)

```
$ go run examples/prefix/main.go
pinging --> PING google.com (172.217.21.142): 56 data bytes
pinging --> 64 bytes from 172.217.21.142: icmp_seq=0 ttl=116 time=24.783 ms
pinging --> 64 bytes from 172.217.21.142: icmp_seq=1 ttl=116 time=26.532 ms
pinging --> 64 bytes from 172.217.21.142: icmp_seq=2 ttl=116 time=25.450 ms
```

## suffixing

[examples/suffix/main.go](examples/suffix/main.go)

```
$ go run examples/suffix/main.go
PING google.com (216.58.211.14): 56 data bytes <-- pinging
64 bytes from 216.58.211.14: icmp_seq=0 ttl=116 time=146.194 ms <-- pinging
64 bytes from 216.58.211.14: icmp_seq=1 ttl=116 time=156.214 ms <-- pinging
64 bytes from 216.58.211.14: icmp_seq=2 ttl=116 time=26.280 ms <-- pinging
```

## prefixing, suffixing and infixing

[examples/allfix/main.go](examples/allfix/main.go)

```
$ go run examples/allfix/main.go
pinging   < PING google.com (172.217. ... >   ponging
pinging   < 64 bytes from 172.217.21. ... >   ponging
pinging   < 64 bytes from 172.217.21. ... >   ponging
pinging   < 64 bytes from 172.217.21. ... >   ponging
pinging   <  >                                ponging
pinging   < --- google.com ping stati ... >   ponging
pinging   < 3 packets transmitted, 3  ... >   ponging
pinging   < round-trip min/avg/max/st ... >   ponging
```
