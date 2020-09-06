# fixer

async io.Writer that prefixes, suffixes and infixes and is thread safe

## prefixing

```
package main

import (
	"os"
	"os/exec"

	"github.com/matti/fixer"
)

func main() {
	prefixer := fixer.Fixer{
		Writer: os.Stdout,
		PrefixFunc: func(s string) string {
			return "pinging --> "
		},
	}
	cmd := exec.Command("ping", "-c", "3", "google.com")
	cmd.Stdout = prefixer
	cmd.Run()
}
```

```
$ go run examples/prefix/main.go
pinging --> PING google.com (172.217.21.142): 56 data bytes
pinging --> 64 bytes from 172.217.21.142: icmp_seq=0 ttl=116 time=24.783 ms
pinging --> 64 bytes from 172.217.21.142: icmp_seq=1 ttl=116 time=26.532 ms
pinging --> 64 bytes from 172.217.21.142: icmp_seq=2 ttl=116 time=25.450 ms
```

## suffixing

```
package main

import (
	"os"
	"os/exec"

	"github.com/matti/fixer"
)

func main() {
	suffixer := fixer.Fixer{
		Writer: os.Stdout,
		SuffixFunc: func(s string) string {
			return " <-- pinging"
		},
	}
	cmd := exec.Command("ping", "-c", "3", "google.com")
	cmd.Stdout = suffixer
	cmd.Run()
}
```

```
$ go run examples/suffix/main.go
PING google.com (216.58.211.14): 56 data bytes <-- pinging
64 bytes from 216.58.211.14: icmp_seq=0 ttl=116 time=146.194 ms <-- pinging
64 bytes from 216.58.211.14: icmp_seq=1 ttl=116 time=156.214 ms <-- pinging
64 bytes from 216.58.211.14: icmp_seq=2 ttl=116 time=26.280 ms <-- pinging
```

## allfixing

```
package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/matti/fixer"
)

func main() {
	prefixer := fixer.Fixer{
		Writer: os.Stdout,
		PrefixFunc: func(s string) string {
			return "pinging   "
		},
		InfixFunc: func(s string) string {
			if len(s) > 25 {
				return "< " + s[:25] + " ... >"
			}

			return "< " + s + " >"
		},
		SuffixFunc: func(s string) string {
			padding := ""
			if amount := 36 - len(s); amount > 0 {
				padding = strings.Repeat(" ", amount)
			}
			return padding + "ponging"
		},
	}
	cmd := exec.Command("ping", "-c", "3", "google.com")
	cmd.Stdout = prefixer
	cmd.Run()
}
```

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
