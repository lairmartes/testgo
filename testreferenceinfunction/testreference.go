package main

import (
	"fmt"

	"github.com/lairmartes/testgo/references"
)

func main() {

	var myReference references.ReferenceTestType

	myReference.Name = "Nome inicial"
	myReference.Address = "Endereco inicial"

	toString("Reference before sending to modification", myReference)

	myReference.ReferenceUpdater("Name to be modified", "Address to be modified")

	toString("Reference after modification", myReference)

	toString("To be sure that data are the same", myReference)
}

func toString(refTitle string, ref references.ReferenceTestType) {

	fmt.Printf("\n========= Reference Title: %s ===========\n", refTitle)
	fmt.Printf("Name    : %s\n", ref.Name)
	fmt.Printf("Address : %s\n", ref.Address)
	fmt.Printf("=========================================================================\n")

}
