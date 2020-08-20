package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"
	"text/tabwriter"

	"github.com/b4nst/cobd/internal/testable"
	"github.com/fatih/color"
)

func main() {
	testables := testable.FromEnv(os.Environ())
	go runTests(testables, os.Stdout)

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/test", testHandler)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 80
	}
	fmt.Println("Starting up on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), nil))
}

func testHandler(w http.ResponseWriter, req *http.Request) {
	testables := testable.FromEnv(os.Environ())
	runTests(testables, w)
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintln(w, "Hostname:", hostname)

	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			fmt.Fprintln(w, "IP:", ip)
		}
	}

	fmt.Fprintln(w, "RemoteAddr:", req.RemoteAddr)
	if err := req.Write(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func runTests(testables []testable.Testable, output io.Writer) {
	var wg sync.WaitGroup
	wg.Add(len(testables))
	w := tabwriter.NewWriter(output, 0, 1, 1, '\t', 0)
	passed, failed := "Passed", "Failed"
	if output == os.Stdout {
		passed = color.GreenString(passed)
		failed = color.RedString(failed)
	}

	fmt.Fprintln(output, "Running", len(testables), "test(s)...")
	for _, t := range testables {
		go func(t testable.Testable) {
			defer wg.Done()
			defer w.Flush()
			t.Test()

			name, err := t.Name(), t.Error()
			if err != nil {
				fmt.Fprintln(w, failed, "\t", name, "\t", err, "\t")
			} else {
				fmt.Fprintln(w, passed, "\t", name, "\t")
			}
		}(t)
	}
	wg.Wait()
	fmt.Fprintln(output, "Done.")
}
