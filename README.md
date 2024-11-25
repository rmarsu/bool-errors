# bool-errors
![image](https://github.com/user-attachments/assets/df60e486-6f7c-4a6e-9915-f9bfddc3b4c3)


# Example Usage
```go
package main

import (
	boolerrors "boolerrors/source"
	"fmt"
)


func main() {
	error1 := boolerrors.NewError()
	error2 := boolerrors.NoError()
	megaError := boolerrors.CombineErrors(error1, error2)
	fmt.Println(megaError.Error())
	fmt.Println(megaError.BoolButCool())
	fmt.Println(megaError.Negative().Error())
	//megaError.Secret()
}
```
