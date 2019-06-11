func mergesortv1(s []int) {
	len := len(s)

	if len > 1 {
		middle := len / 2

		var wg sync.WaitGroup
		wg.Add(2)

		// First half
		go func() {
			defer wg.Done()
			mergesortv1(s[:middle])
		}()

		// Second half
		go func() {
			defer wg.Done()
			mergesortv1(s[middle:])
		}()

		// Wait that the two goroutines are completed
		wg.Wait()
		merge(s, middle)
	}
}