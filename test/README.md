`$ go run test.go`

```
===== COMMAND LINE ARGUMENTS =====
[-s -xyz --long -h -a 12 -e=Equal test 1 --long-equal=Equal test 2 --name ProfOak -- --never foot loose]

===== ALL FLAG KEYS =====
[zero default short-equal long-equal name age help short long example]

===== USAGE FLAG =====

--- Bools ---
-s, --short             Test short flag
-l, --long              Test long flag
-x, --example           Test for the multi short flag
-h, --help              Display this message

--- Strings ---
-d, --default           Using the default value example
-e, --short-equal       Test short + equal sign in args
-q, --long-equal        Test long + equal sign in args
-n, --name              Store a person's name
-a, --age               Store a person's age
-z, --zero              Failing test for -xyz

Name is: ProfOak

===== FINAL RESULTS =====
name : ProfOak
zero : (default) incorrect usage
default : Default Value
help : true
short : true
long : true
long-equal : Equal test 2
example : true
age : 12
short-equal : Equal test 1

Args: [--never, foot, loose]
```
