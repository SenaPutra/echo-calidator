/**
 * @project echo-validator
 * @author sena on 03/05/20
 */

package main

import (
	"bytes"
	"fmt"
	"html/template"
)

func main(){


	t, _ := template.ParseFiles("test.txt")
	//t, _ := template.New("test").Parse("Hello {{.Name}} , Bonuss ngimpiiii kaliii yeeee !!!!")
	u := Users{
		Name: "",
	}

	buf := new(bytes.Buffer)
	t.Execute(buf,u)

	fmt.Printf(buf.String())
}

type Users struct {
	Name  string `json:"name"`
}
