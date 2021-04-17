package
1. pkg.go.dev (browse it)

vscode extension for go
1. clang-format (vscode extension)
2. go (vscode exyension)

useful command:
cli:
go mod init "module" --> create a module
go mod tidy --> tidy up import export
go mod edit -replace="web module"="local module"
go:
make(chan type) : make channel that can pass "type" data
chanel<-data : send data to chanel
variable<-chanel : wait value to be sent into channel, when that happen, set variable into that value
fmt.Println(<-channel) : wait value to be sent into channel, when that happen, println it


notes
1. export only works for Capital variable or function
2. struct has a lot in common with c++/python object, but becareful it isn't entirely the same
3. object function in go is called receiver function
4. go have pointer, the same with c++
5. in go data stucture, below types is pass by value by default:
    a. int
    b. float
    c. string
    d. bool
    e. struct
6. below types is pass by reference by default:
    a. slices
    b. maps
    c. channels
    d. pointers
    e. functions
7. map is like map in c++ and dict in python
8. every variable in go must be used in order to program to successfully compiled
9. interface is like abstract function in c++, that can be inherit
10. concrete type, is type that can create value directly from it like
    a. string
    b. map
    c. int
    d. struct
11. interface type, is type that can't create value directly from it
    a. interface --> sample interface is in bot
12. http client is located in net/http package
13. composite interface, (see httpResponse body in net/http.Get). beside stating function in interface, interface can also state interface within its definition
14. interface in go will automatically recognize struct that satisfiy its definition (by function or by other interface)
15. go routine is a method to create concurency in golang, concurency is not pararelism, concurency is about jumping between execution routine to speed up the process whereas pararelism is executing at the same time.
16. main function in go automatically has one go routine (Main routine). and we can add it along the way with "go" keyword
17. routine that is created by "go" keyword are called child routine, and it has different behaviour with main routine
18. go will jumping between go routine whenever current go routine has blocking call
19. we can't control how go routine jump/switch. it's already done by operating system
20. usually we spawn go routine whenever blocking call is met
21. main routine control exit of our program which means every unfinished child routine will terminate when main routine exit
22. to "communicate" between routine we use channel
23. Never past by reference in different go routine
24. when need to stop use function literal to create new go routine