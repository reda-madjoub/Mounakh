## Continuous Testing
* gotestsum: allow to launch test every time we saved any file in the project (https://github.com/gotestyourself/gotestsum#run-tests-when-a-file-is-saved)
* go install gotest.tools/gotestsum@latest
* gotestsum --watch : watching file every time we saved file in the project
* gotestsum --format (dots|testname|pkgname) => display options