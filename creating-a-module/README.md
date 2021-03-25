# Creating a Go module!

> Source: [Tutorial: Create a Go module](https://golang.org/doc/tutorial/create-module)

## Steps

### Create the module itself

1. Create a folder called greetings
2. Go inside the greetings folder and create a module.  

Syntax:  
`$ go mod init ${parent-folder-name}/${module}`

Executed:  
`$ go mod init creating-a-module/module`

3. Create `greetings.go` file and create some functions inside of it.


### Calling the greetings module from another module

1. From the creating-a-module folder, create another folder called Hello.

2. Inside the hello folder, create a module by running `$ go mod init creating-a-module/hello`

3. Using the import statement, import the recently created module, i.e:

```go
package hello

import (
  "fmt""

  "creating-a-module/greetings"
)

// rest of the code
```

**Important**: In this case, our Greetings module is not a public module, so we have to tell Go where to find this module.

Running `$ go mod edit -replace=creating-a-module/greetings=../greetings` will tell Go to look for the greetings module inside `../greetings` folder (or in this case, inside: **creating-a-module/greetings**).

This will add the `replace` directive inside hello mod.

```go
module creating-a-module/hello

go 1.16

replace creating-a-module/greetings => ../greetings

```

Now, we need to synchronize what's required by the code but it's not in the hello mod yet. We do that using the `$ go mod tidy` command. So, run it on.

Now, we should see the `require` directive added to our hello mod.

```go
module creating-a-module/hello

go 1.16

replace creating-a-module/greetings => ../greetings

require creating-a-module/greetings v0.0.0-00010101000000-000000000000
```

**Require syntax:**

required {module} {module version}

i.e:

required http v1.16.2

### Running hello module


```bash
$ cd hello
$ go run .
Hi, Jonathan. Welcome! 
```
