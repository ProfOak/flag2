`$ go run test.go`

```
===== EXAMPLE ERROR =====
Atribute already exists in list of strings: name

===== COMMAND LINE ARGUMENTS =====
[-s -xyz --long -a 12 --name billy -- --never foot loose]

===== ALL FLAG KEYS =====
[age zero default name short long example help]

===== USAGE =====

--- Bools ---
-s, --short     Test short flag
-l, --long      Test long flag
-x, --example   Test for the multi short flag
-h, --help      Display this message and exit

--- Strings ---
-n, --name      Store a person's name
-a, --age       Store a person's age
-z, --zero      Failing test for -xyz
-d, --default   Using the default value example

Name is: billy

===== FINAL RESULTS =====
Options: map[example:true long:true age:12 name:billy default:Default Value zero:bad help:true short:true]
Args: [--never foot loose]
```
