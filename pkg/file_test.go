package pkg_test

import (
	"fmt"

	"github.com/Mythra/rules_go_repro/embed/pkg"
)

func TestMe() {
	struc := pkg.MyCoolStruct{
		Name: "sadf",
	}
	fmt.Println(struc.Name)
}
