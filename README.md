# libiperf-cgo

A minimal example of how to call libiperf using CGO in Go.

> [!TIP]
> Install `libiperf-dev` and refer to the `man libiperf 3` page for example C code demonstrating how to use the library.
> See also https://github.com/esnet/iperf/blob/master/examples/mic.c

```
$ git submodule update --init --recursive
$ cd iperf
$ ./configure && make
```

```
$ make build
```

```
$ ./bin/iperf3 -h

Usage of ./bin/iperf3:
  -host string
        Server hostname to connect to (default "127.0.0.1")
  -port int
        Server port to connect to (default 5201)
```

```
$ ./bin/iperf3

Connecting to host 127.0.0.1, port 5201
[  5] local 127.0.0.1 port 52642 connected to 127.0.0.1 port 5201
[ ID] Interval           Transfer     Bitrate         Retr  Cwnd
[  5]   0.00-1.00   sec  9.52 GBytes  81.7 Gbits/sec    0   1.19 MBytes
[  5]   1.00-2.00   sec  8.64 GBytes  74.2 Gbits/sec    0   1.25 MBytes
[  5]   2.00-3.00   sec  9.67 GBytes  83.1 Gbits/sec    0   1.31 MBytes
^C
```