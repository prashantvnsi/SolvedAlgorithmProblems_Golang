/*

In this question, each string can be sorted. After the sorting is completed, 
the strings of the same Anagrams must have the same sorting result. 
Save the sorted string as a key in the map. After traversing the array, you can get a map, 
the key is the sorted string, and the value corresponds to the Anagrams string collection 
after the sorted string. Finally, output the string array corresponding to these values.

*/

package leetcode

import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func groupAnagrams(strs []string) [][]string {
	record, res := map[string][]string{}, [][]string{}
	for _, str := range strs {
		sByte  : = [] rune ( str )
		sort.Sort(sortRunes(sByte))
		sstrs := record[string(sByte)]
		sstrs = append(sstrs, str)
		record[string(sByte)] = sstrs
	}
	for  _ , v  : =  range  record {
		res = append(res, v)
	}
	return res
}