
# Viewing the presentation locally using the "present"-tool

You can also run it locally on your own device:
Use the "present"-tool to "run" the presentation.

## Step 1 - Get the present tool

    $ go get golang.org/x/tools/cmd/present
    # present binary should now be in ${GOPATH}/bin 

## Step 2 - Fetch the presentation from github

    go get github.com/MarcGrol/go-training
    or
    mkdir -p ${GOPATH}/src/github.com/MarcGrol
    cd ${GOPATH}/src/github.com/MarcGrol
    git clone https://github.com/MarcGrol/go-training.git
    # presention is located in ${GOPATH}/github.com/MarcGrol/go-training

## Step 3 - Start presentation within the go-training directory

    cd ${GOPATH}/src/github.com/MarcGrol/go-training
    # run present tool in background
    ${GOPATH}/bin/present -http=127.0.0.1:3999 &

## Step 4 - Point your browser at:
http://127.0.0.1:3999

## Step 5 - Build all examples in your local environment

    cd ${GOPATH}/src/github.com/MarcGrol/go-training
    go install ./...
    # all examples should now be in ${GOPATH}/bin as executables
    ls -ltr ${GOPATH}/bin
    

