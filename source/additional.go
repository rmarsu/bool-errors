package boolerrors


func combineErrors(err1, err2 BoolError) BoolError {
	return err1 || err2 
}

func CombineErrors(errors ...BoolError) BoolError {
	var result BoolError = errors[0]
     for _, err := range errors {
          result = combineErrors(result, err)
     }
     return result
}