Big Example
===

Uses a `cmd` and a `pkg` directory containing multiple packages.

## Expected Result
```
# go-test-only-linter cmd/thing/thing.go ./...

test-only-example-big/pkg/other.OtherNotUsedFunc is not used
test-only-example-big/pkg/stuff.NotUsedFunction is not used
```
