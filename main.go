package main

/*
#cgo CFLAGS: -I./iperf/src
#cgo LDFLAGS: -L./iperf/src/.libs -liperf
#include <iperf_api.h>
#include <stdlib.h>
*/
import "C"

import (
	"flag"
	"fmt"
	"os"
	"unsafe"
)

const (
	EXIT_WITH_SUCCESS = iota
	EXIT_WITH_USAGE_ERROR
	EXIT_WITH_LIBRARY_ERROR
)

var (
	remote_host *string
	remote_port *int
)

func init() {
	remote_host = flag.String("host", "127.0.0.1", "Server hostname to connect to")
	remote_port = flag.Int("port", 5201, "Server port to connect to")
}

func main() {
	flag.Parse()

	host := C.CString(*remote_host)
	defer C.free(unsafe.Pointer(host))
	port := C.int(*remote_port)

	test := C.iperf_new_test()
	if test == nil {
		fmt.Fprintf(os.Stderr, "failed to create test")
		os.Exit(EXIT_WITH_LIBRARY_ERROR)
	}

	C.iperf_defaults(test)
	C.iperf_set_test_role(test, C.char('c'))
	C.iperf_set_test_server_hostname(test, host)
	C.iperf_set_test_server_port(test, port)

	if C.iperf_run_client(test) < 0 {
		e := C.GoString(C.iperf_strerror(C.i_errno))
		fmt.Fprintf(os.Stderr, "%s", e)
		os.Exit(EXIT_WITH_LIBRARY_ERROR)
	}

	C.iperf_free_test(test)
}
