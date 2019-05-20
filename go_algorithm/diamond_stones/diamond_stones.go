package diamond_stones

func NumJewelsInStones(J string, S string) int {
	count := 0
	if len(J) <= 50 && len(S) <= 50 {
		for _, j := range J {
			for _, s := range S {
				if s == j {
					count++
				}
			}
		}
	}
	return count
}
