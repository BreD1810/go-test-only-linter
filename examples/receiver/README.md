Receiver Example
===

Uses a `cmd` and a `pkg` directory containing multiple packages.
Also uses receivers and pointer receivers

## Expected Result
```
# go-test-only-linter cmd/thing/thing.go ./...

main.notUsedFuncMain is not used
test-only-example-receiver/pkg/other.OtherNotUsedFunc is not used
test-only-example-receiver/pkg/stuff.StuffStruct.NotUsedFunction is not used
test-only-example-receiver/pkg/stuff.(*StuffStruct).NotUsedPointerFunction is not used
```
