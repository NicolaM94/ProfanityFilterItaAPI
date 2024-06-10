package filters

func MultiWordFilter(words []string) []FilterResult {

	var collector []FilterResult
	for w := range words {
		collector = append(collector, SingleWordFilter(words[w]))
	}
	return collector
}
