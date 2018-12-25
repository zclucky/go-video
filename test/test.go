package main
import (
	"encoding/json"
	"fmt"
	"os"
)
func main ( ) {
	type UserCredential struct {
		Username string `json:"user_name"`
		Password string `json:"password"`
	}
	group := ColorGroup {
		ID :     1 ,
		Name :   "Reds" ,
		Colors : [ ] string { "Crimson" , "Red" , "Ruby" , "Maroon" } ,
	}
	b , err := json. Marshal ( group )
	if err != nil {
		fmt. Println ( "error:" , err )
	}
	os. Stdout . Write ( b )
}