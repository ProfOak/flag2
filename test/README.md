`$ go run test.go`

```
===== EXAMPLE ERROR =====
[FlagError] Flag already exists in list of strings: name

===== COMMAND LINE ARGUMENTS =====
[-s -xyz --long -a 12 -e=Equal test 1 --long-equal=Equal test 2 --name ProfOak -- --never foot loose]

===== ALL FLAG KEYS =====
[short-equal long-equal name age zero default short long example help]

===== USAGE =====

--- Bools ---
-s, --short             Test short flag
-l, --long              Test long flag
-x, --example           Test for the multi short flag
-h, --help              Display this message and exit

--- Strings ---
-e, --short-equal       Test short + equal sign in args
-q, --long-equal        Test long + equal sign in args
-n, --name              Store a person's name
-a, --age               Store a person's age
-z, --zero              Failing test for -xyz
-d, --default           Using the default value example

Name is: ProfOak

===== FINAL RESULTS =====
short-equal : Equal test 1
long-equal : Equal test 2
zero : (default) incorrect usage
help : true
short : true
example : true
long : true
age : 12
name : ProfOak
default : Default Value

Args: [--never, foot, loose]
```
