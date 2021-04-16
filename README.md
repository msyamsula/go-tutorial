package
1. pkg.go.dev (browse it)


useful command:
go mod init "module" --> create a module
go mod tidy --> tidy up import export
go mod edit -replace="web module"="local module"

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