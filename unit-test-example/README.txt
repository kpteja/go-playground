go test
Tests whole package.

go test PACKAGE_NAME
Tests specified package.
Eg. go test github.com/kpteja/go-playground

go test PACKAGE_NAME -run FUNC_NAME
Tests specified function in the specified package.
FUNC_NAME refers to function in test file.
Eg. go test github.com/kpteja/go-playground -run TestAdd

go test PACKAGE_NAME -run FUNC_NAME_REGEX
FUNC_NAME_REGEX refers to function name regular expression.
Eg. go test github.com/kpteja/go-playground -run ^(TestSubtract)$