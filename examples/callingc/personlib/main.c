
#include <stdlib.h>
#include <stdio.h>

#include "person.h"

int main( int argc, char** argv) {
    Person person;

    if( !person_get_on_id( "123", &person) != 0 ){
        fprintf( stderr, "Error fetching person");
        exit(EXIT_FAILURE);
    }

    char printbuffer[64] = "";
    person_as_string(&person, printbuffer, sizeof(printbuffer));
    fprintf(stdout, "%s", printbuffer);

    exit(EXIT_SUCCESS);

}