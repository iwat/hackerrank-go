package main

func CountAnagram(a, b string) int {
	aMap := makeMap(a)
	bMap := makeMap(b)

	counter := 0
	for aChar, aHit := range aMap {
		if bHit, ok := bMap[aChar]; ok {
			counter += diff(aHit, bHit)
			delete(aMap, aChar)
			delete(bMap, aChar)
		}
	}
	for _, aHit := range aMap {
		counter += aHit
	}
	for _, bHit := range bMap {
		counter += bHit
	}
	return counter
}

func diff(a, b int) int {
	d := a - b
	if d < 0 {
		return -d
	}
	return d
}

func makeMap(s string) map[byte]int {
	out := make(map[byte]int, 24)
	for i := 0; i < len(s); i++ {
		if _, ok := out[s[i]]; ok {
			out[s[i]]++
		} else {
			out[s[i]] = 1
		}
	}
	return out
}
