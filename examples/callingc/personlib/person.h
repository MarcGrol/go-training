#ifndef _personlib_h
#define _personlib_h

#include <stdbool.h>

typedef  struct{
    char* id;
    char* name;
    int age;
} Person;

bool person_get_on_id(char* id, Person* person);
void person_as_string(Person* person, char* buffer, int buflen);

#endif
