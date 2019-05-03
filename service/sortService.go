package service

type fn func(interface{}) int

func QuickSort(items []interface{}, fn fn, asc bool) []interface{} {
	var merged, smallerItems, largerItems []interface{}
	if len(items) == 1 {
		return items
	}
	if len(items) > 1 {
		pivotIndex := len(items) / 2
		smallerItems = []interface{}{}
		largerItems = []interface{}{}

		for i := range items {
			if i != pivotIndex {
				if asc {
					if fn(items[i]) < fn(items[pivotIndex]) {
						smallerItems = append(smallerItems, items[i])
					} else {
						largerItems = append(largerItems, items[i])
					}
				} else {
					if fn(items[i]) > fn(items[pivotIndex]) {
						smallerItems = append(smallerItems, items[i])
					} else {
						largerItems = append(largerItems, items[i])
					}
				}
			}
		}

		smallerItems = QuickSort(smallerItems, fn, asc)
		largerItems = QuickSort(largerItems, fn, asc)

		merged = append(append(append(merged, smallerItems...), items[pivotIndex]), largerItems...)

	}

	return merged
}