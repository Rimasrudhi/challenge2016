package main

import "fmt"

func main() {
	ReadCSV("cities.csv")

	// define main distributor
	AddDistributor("DISTRIBUTOR1",
		[]string{"INDIA", "UNITED STATES"},
		[]string{"KARNATAKA-INDIA", "CHENNAI-TAMILNADU-INDIA"},
	)

	// adding sub-distributors
	fmt.Println(AddSubDistributor("DISTRIBUTOR1", "DISTRIBUTOR2",
		[]string{"INDIA"}, []string{"TAMILNADU-INDIA"}))

	fmt.Println(AddSubDistributor("DISTRIBUTOR2", "DISTRIBUTOR3",
		[]string{"HUBLI-KARNATAKA-INDIA"}, []string{})) // must fail

	// permission test checks
	fmt.Println("Can DISTRIBUTOR2 distribute in MUMBAI-MAHARASHTRA-INDIA?", CanDistribute("DISTRIBUTOR2", "MUMBAI-MAHARASHTRA-INDIA"))
	fmt.Println("Can DISTRIBUTOR2 distribute in CHENNAI-TAMILNADU-INDIA?", CanDistribute("DISTRIBUTOR2", "CHENNAI-TAMILNADU-INDIA"))
	fmt.Println("Can DISTRIBUTOR3 distribute in HUBLI-KARNATAKA-INDIA?", CanDistribute("DISTRIBUTOR3", "HUBLI-KARNATAKA-INDIA"))
}
