func countCharacters(words []string, chars string) int {
	f := make(map[byte]int)
	for i := 0; i < len(chars); i++ {
		f[chars[i]]++
	}
    result := 0
	for _, word := range words {
		good := true
		m := make(map[byte]int)
		for i := 0; i < len(word); i++ {
			m[word[i]]++
		}
		for c := range m {
			freq := m[c]
			//fmt.Println(c)
			if f[c] < freq {
				good = false
				break
			}
		}
		if good {
			result += len(word)
		}

	}
	return result
}
