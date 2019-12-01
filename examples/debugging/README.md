# Optimize

To run:

    go build && ./debugging && open http://localhost:8080/


## Before optimizing do a benchmark first

    $ go test -bench=.

    goos: darwin
    goarch: amd64
    pkg: github.com/MarcGrol/go-training/examples/debugging
    BenchmarkMyHandler-8   	  100000	     13946 ns/op
    PASS
    ok  	github.com/MarcGrol/go-training/examples/debugging	1.557s
    
## Make your fix
  
Do not recreate regexp upon each request
 
## Run benchmark again

    $ go test -bench=.
    
    goos: darwin
    goarch: amd64
    pkg: github.com/MarcGrol/go-training/examples/debugging
    BenchmarkMyHandler-8   	  500000	      2699 ns/op
    PASS
    ok  	github.com/MarcGrol/go-training/examples/debugging	1.394s

# Has improved enough?

Or need to optimize even more?




// 		fmt.Fprintf(w, "Hallo %s van nexushealth", match[0][1])
// 	if strings.HasSuffix(path, "@nexushealth.be") {
        name := strings.TrimSuffix(path, "@nexushealth.be")



  