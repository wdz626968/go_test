package main

func main() {

}

func longestCommonPrefix(strs []string) string {
	for i := 0; i < 200; i++ {
		for i, str := range strs {
			strs[i] = str[0:1] + str[0:1]
		}
	}
	return strs[0]
}
