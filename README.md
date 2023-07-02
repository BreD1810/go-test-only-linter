go-test-only-linter
===

`golangci-lint` doesn't spot functions that are only used in tests - ideally these should be declared as unused.
This tool will identify this kind of unused code.

**NOTE**: This will highlight exported code in libraries that is not used within the library itself.

## Usage

Run
```shell
make install
```
to install `go-test-only-linter`.

To use the tool on all code:
```shell
go-test-only-linter path/to/main.go ./...
```

## How it Works
1. Builds an unoptimised binary with `go build -o unoptimised-binary -gcflags '-N -l' path/to/main.go`.
    - It must be unoptimised otherwise functions may be inlined, leading to false positives.
2. Crawls over the ast of your go files to determine which functions are declared:
    - Normal functions: `func myFunction(...)`
    - Receiver functions: `func (r receiver) myReceiverFunction(...)`
    - Pointer receiver functions: `func (r *receiver) myPointerReceiverFunction(...)`
3. Prints out functions which are not used.
4. Removes `unoptimised-binary`
