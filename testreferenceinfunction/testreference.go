package main

import (
	"fmt"

	"github.com/lairmartes/testgo/references"
)

func main() {

	var myReference references.ReferenceTestType

	myReference.Name = "Nome inicial"
	myReference.Address = "Endereco inicial"

	toString("Reference before send to modification", myReference)

	myReference.TestReference("Name to be modified", "Address to be modified")

	toString("Reference after modification", myReference)

	toString("To be sure that data is the same", myReference)
}

func toString(refTitle string, ref references.ReferenceTestType) {

	fmt.Printf("\n================== Reference Title: %s =================\n", refTitle)
	fmt.Printf("Name    : %s\n", ref.Name)
	fmt.Printf("Address : %s\n", ref.Address)
	fmt.Printf("=========================================================================\n")

}
