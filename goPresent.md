
# Viewing the presentation locally using the "present"-tool

You can also run it locally on your own device:
Use the "present"-tool to "run" the presentation.

## Step 1 - Get the present tool

    $ go get golang.org/x/tools/cmd/present
    # present binary should now be in ${HOME}/go/bin 

## Step 2 - Fetch the presentation from github

    git clone https://github.com/MarcGrol/go-training.git
    # presention is located in ${GOPATH}/github.com/MarcGrol/go-training

## Step 3 - Start presentation within the go-training directory

    cd go-training
    # run present tool in background
    go run golang.org/x/tools/cmd/present -http=:3999 -use_playground=true

## Step 4 - Point your browser at:
http://127.0.0.1:3999


