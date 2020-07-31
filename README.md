# prmchk
_Check if a struct contains a field that overwrites a promoted field from an embedded struct_

In go, fields from embedded structs are _promoted_ to the containing struct, as if they were declared
in the containing struct. However, if you declare a field in the containing struct which has the same name,
then promotion is suppressed.

Sometimes, we have structs for which we never want to suppress field promotion, but it is easy to accidentally
do this, especially if the embedding is deeply nested or there are many fields.

This library provides functions to check whether a struct type contains any fields that suppress promotion of
a field of an embedded type.

See the unit tests for usage.


TODO:

Support embedded struct pointers
