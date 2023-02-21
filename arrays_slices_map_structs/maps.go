package main

import "fmt"

func main() {
	example1()
	example2()
}

func example1() {
	fmt.Println("inside example1.........")

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])

	for k,v := range elements {
		fmt.Println("elements[",k,"]",v) //will not come in same order 
	}

	fmt.Println(elements["Un"]) // trying to access undefined key, nothing will show up

	name, ok := elements["Un"] //right way to do
	fmt.Println(name, ok) //name is empty, ok=false


	name, ok = elements["B"] //right way to do
	fmt.Println(name, ok) //name is Boron, ok=True
	
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println(name, ok)
	}		
}

func example2() {
	fmt.Println("inside example2.........")
	elements := map[string]map[string]string{
	  "H": map[string]string{
		"name":"Hydrogen",
		"state":"gas",
	  },
	  "He": map[string]string{
		"name":"Helium",
		"state":"gas",
	  },
	  "Li": map[string]string{
		"name":"Lithium",
		"state":"solid",
	  },
	  "Be": map[string]string{
		"name":"Beryllium",
		"state":"solid",
	  },
	  "B":  map[string]string{
		"name":"Boron",
		"state":"solid",
	  },
	  "C":  map[string]string{
		"name":"Carbon",
		"state":"solid",
	  },
	  "N":  map[string]string{
		"name":"Nitrogen",
		"state":"gas",
	  },
	  "O":  map[string]string{
		"name":"Oxygen",
		"state":"gas",
	  },
	  "F":  map[string]string{
		"name":"Fluorine",
		"state":"gas",
	  },
	  "Ne":  map[string]string{
		"name":"Neon",
		"state":"gas",
	  },
	}
  
	if el, ok := elements["Li"]; ok {
	  fmt.Println(el["name"], el["state"])
	}
  }