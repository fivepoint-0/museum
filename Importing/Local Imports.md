To do a local import, let us take the process from the start.

Create a project directory, and initialize a Golang project:
```sh
mkdir localimport
cd localimport
go mod init github.com/user/localimport
touch main.go
mkdir functions
touch functions/add.go
```

The file structure should look like this:
```ts
tutorial/ 
├── functions/ 
│ └── add.go
├── go.mod  
└── main.go
```

Edit the contents of `localimport/functions/add.go`:
```go
package functions

func Add(a int, b int) int {
  return a + b
}
```

Notice how the top of the file says `package functions`. This will be important when importing it into the `main.go` file inside `localimport/`.

Edit the contents of `localimport/main.go`:
```go
package main

import (
    "fmt"
    funcs "github.com/user/localimport/functions"
)

func main() {
    result := funcs.Add(1, 4)
    fmt.Println("1 + 4 =", result)
}
```

As you can see, the local import is doing a couple things:
1. We are using a namespace, `funcs`, for the import. This helps segment out functions to a certain library/module.
2. The import statement `"github.com/user/localimport/functions"` looks at the original `localimport/go.mod` file, recognizes the root package name, `github.com/user/localimport`, and if there is a directory in `localimport` called `functions`, then `"github.com/user/localimport/functions"` would result in `main.go` looking in the subdirectory `functions`. 