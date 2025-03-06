package main

import "strings"

// distributor struct
type Distributor struct {
	Name            string
	Includes        []string
	Excludes        []string
	SubDistributors map[string]*Distributor
}

// stores all thedistributors
var distributors = make(map[string]*Distributor)

// add a main distributor
func AddDistributor(name string, includes, excludes []string) {
	distributors[name] = &Distributor{
		Name:            name,
		Includes:        includes,
		Excludes:        excludes,
		SubDistributors: make(map[string]*Distributor),
	}
}

// add a sub-distributor with parent validation
func AddSubDistributor(parentName, subName string, includes, excludes []string) string {
	parent, exists := distributors[parentName]
	if !exists {
		return "ERROR: Parent distributor not found"
	}

	// ensure sub-distributor does not override parent exclusions
	for _, incl := range includes {
		if isExcludedByParent(parent, incl) {
			return "ERROR: Sub-distributor cannot include " + incl + " as it is excluded by parent."
		}
	}

	// create and register the sub-distributor
	subDistributor := &Distributor{
		Name:            subName,
		Includes:        includes,
		Excludes:        excludes,
		SubDistributors: make(map[string]*Distributor),
	}

	parent.SubDistributors[subName] = subDistributor
	distributors[subName] = subDistributor

	return "Sub-distributor " + subName + " added successfully."
}

// check if a parent or ancestor excluded a region
func isExcludedByParent(distributor *Distributor, region string) bool {
	// Direct exclusion check
	for _, excl := range distributor.Excludes {
		if strings.Contains(region, excl) {
			return true
		}
	}

	// recursively check if an ancestor excluded it
	for _, parent := range distributors {
		if parent.SubDistributors[distributor.Name] != nil {
			if isExcludedByParent(parent, region) {
				return true
			}
		}
	}

	return false
}

// check if a distributor is authorized to distribute in a location
func CanDistribute(distributorName, location string) string {
	dist, exists := distributors[distributorName]
	if !exists {
		return "Distributor not found"
	}

	if isExcludedByParent(dist, location) {
		return "NO"
	}

	for _, incl := range dist.Includes {
		if strings.Contains(location, incl) {
			return "YES"
		}
	}

	return "NO"
}
