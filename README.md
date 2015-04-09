# goopenkitchen

This training on the go language describes the language fron ground up using examples.
Most examples can be run and edited from the presentation.

The interactive presentation is available via the following link:

http://go-talks.appspot.com/github.com/MarcGrol/goopenkitchen/openKitchen.slide#1

You can also run it locally on your own device:
Use the "present"-tool to view and "execute" the presentation.

Get the present application:

    go get golang.org/x/tools/present
    # present binary should now be in ${GOPATH}/bin 

Start presentation within the goopenkitchen directory:

    go get github.com/MarcGrol/goopenkitchen
    cd ${GOPATH}/src/github.com/MarcGrol/goopenkitchen
    # run present tool in background
    present -http=127.0.0.1:3999 &

Point your browser at: http://127.0.0.1:3999


