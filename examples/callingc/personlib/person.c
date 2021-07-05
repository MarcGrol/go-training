#include <stdlib.h>
#include <stdio.h>
#include <assert.h>
#include <stdbool.h>

#include "person.h"

bool person_get_on_id(char* id, Person* person) {
    assert(id != NULL);
    assert(person != NULL);

    person->id = id;
    person->name = "Marc";
    person->age = 42;

    return true;
}

void person_as_string(Person* person, char* buffer, int buflen) {
    assert(person != NULL);
    assert(buffer != NULL);
    assert(buflen > 0);

    snprintf(buffer, buflen, "Person: id:%s, name: %s, age: %d\n", person->id, person->name, person->age);
}
