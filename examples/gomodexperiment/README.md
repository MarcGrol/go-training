# Basic startpoint for practicing with "go mod"

see also https://blog.golang.org/using-go-modules

## Move this package outside $GOPATH for go mod to work

    so fromv ${HOME}/go/..
    to 
    ${HOME}/src/gomodexperiment
    
## Prove that module works ok beforehand
    
    cd ${HOME}/src/gomodexperiment
    go test
    
## Create a go module

    go mod init github.com/MarcGrol/gomodexperiment  # Outside GOPATH so need to tell complete import path
    cat go.mod                                       # Initially without external dependencies
    go test                                          # Still passes
    
## Introduce an external dependency: 

update hello.go to the following

    package gomodexperiment

    import "rsc.io/quote"
    
    func Hello() string {
        return quote.Hello()
    }

Trigger new depencency being fetched
    
    go test        # Fetches latest tagged "stable" version of encountered dependencies (including transient deps)
    cat go.mod     # Only direct dependencies end up in go.mod
    go list -m all # Lists all dependencies (direct and transient)
    cat go.sum     # Contains hashes of all direct and transient dependencies (to code detect changes with version updates)
    
    go test        # Second test will not fetch again, because downloaded modules are cached in GOPATH/pkg/mod
    
## Upgrade dependency "golang.org/x/text"

We are using an untagged version of golang.org/x/text, so let's upgrade

    go get golang.org/x/text
    go test
    go list -m all            # golang.org/x/text is upgraded to v0.3.0
    cat go.mod      
 
## Upgrade dependency "go get rsc.io/sampler"

Current version of rsc.io/sampler is v1.3.0.
Let's try upgrading to a newer minor version.
 
    go get rsc.io/sampler
    go list -m all                      # rsc.io/sampler is upgraded to v1.99.99

    go test                             # Ouch, this version of sampler breaks the build
    
    go list -m -versions rsc.io/sampler # Check which versions are available
    
    go get rsc.io/sampler@v1.3.1        # Let's try v1.3.1
    
    go test                             # Things work again
    
    go list -m all                      
         
## Adding a dependency on a new major version    

update hello.go to the following:

    package hello
    
    import (
        "rsc.io/quote"
        quoteV3 "rsc.io/quote/v3"
    )
    
    func Hello() string {
        return quote.Hello()
    }
    
    func Proverb() string {
        return quoteV3.Concurrency()
    }
                  
and add a test to hello_test.go
        
    func TestProverb(t *testing.T) {
        want := "Concurrency is not parallelism."
        if got := Proverb(); got != want {
            t.Errorf("Proverb() = %q, want %q", got, want)
        }
    }    


Test again
    
    go test
    go list -m all

Note that our module now depends on both rsc.io/quote and rsc.io/quote/v3.
Different major version have a different import path, so can co-exist.
    
## Phase out the old version

update hello.go to the following:

    package gomodexperiment
    
    import "rsc.io/quote/v3"
    
    func Hello() string {
        return quote.HelloV3()
    }
    
    func Proverb() string {
        return quote.Concurrency()
    }    
        
   go test
   
The old "rsc.io/quote" still shows up in go list -m all and in our go.mod file.
Why? 
go build or go test, can easily tell when something is missing and needs to be added, but not when something can safely be removed.

    go mod tidy      # Cleanup unused dependencies
    go list -m all   # Now its gone
    cat go.mod                  
    
