Pointers
--------

variables that store value's address instead of the actual value.

e.g.

age := 30

in computer memory: value -> 32, address -> 0xc00138475

agePointer := &age -> address -> 0xc00138475

Why do we need this?
- avoid unnecessary value copies
    - when you pass a variable to a function as parameter, go creates a copy of the variable.
    Later garbage collector comes and cleans. If values are too big/complicated, might be burden.
    But when you pass value as pointer, this won't happen
- directly mutate values
    - pass pointer to the function instead of variable. same reason as above.
    - function can EDIT the value directly, no return value needed. -> Less Code! (=Less understandable code, unexpected result)


Structs
-------

Group data together and also methods and functions
defined by "type"

'''
type Note struct {

}
'''

Interfaces
----------

Also use "type"
'''
type saver interface {

}
'''