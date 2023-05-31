# go-unit-test
## 1. Test(Asert, Equal, Skip, Fatal, Nil, NotNil)
```
   1. use `go test ./...` for conduct all test in this repo
   2. use `go test -v -run=TestHelloworld` for conduct specific test in repo
   3. use `go test -v` for verbose test
   4. use `go test` for basic test
```

## 2. Benchmark

```
   1. use `go test -v -bench` for conduct all benchmark and test in this repo
   2. use `go test -v -run=NotMatchUnitTestName -bench=.` for only conduct all benchmark
   3. use `go test -v -run=NotMatchUnitTestName -bench=SpecificBencmark ` for conduct specific test
   4. use `go test -v -bench=. ./...` for conduct all available banchmark test from the root module
```

