// Package Villiage contains utility functions for working with age.
package Villiage

func find_age(yy int,mm int,dd int) int {
	birth_day := time.Date(yy,time.Month(mm),dd,0,0,0,0,time.UTC) 		
	age := time.Since(birth_day)
	return int(( age.Hours() / 24 ) / 365)
}


