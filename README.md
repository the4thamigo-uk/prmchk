# prmchk
_Check if a struct contains a field that overwrites a promoted field from an embedded struct_

In go, fields from embedded structs are _promoted_ to the containing struct, as if they were declared
in the containing struct. However, if you declare a field in the containing struct which has the same name,
then promotion is suppressed.

Sometimes, we have structs for which we never want to suppress field promotion, but it is easy to accidentally
do this, especially if the embedding is deeply nested or there are many fields.

This library provides functions to check whether a struct type contains any fields that suppress promotion of
a field of an embedded type.

The idea is to put these checks in your unit tests for structs for which you want to enforce this constraint.
The alternative is to create a linter, but it was thought that if you are designing a struct using this rule,
then it is likely to be a bug if the rule is broken, so this check ought to operate at level of testing, rather
than linting.

See the unit tests for usage.


TODO:

Support embedded struct pointers
