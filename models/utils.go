package models

// IDsDifferences ids differences
func IDsDifferences(sliceA, sliceB []uint64) ([]uint64, []uint64) {
	// Slice A, B are not empty
	if len(sliceA) > 0 && len(sliceB) > 0 {
		// New differences and intersection slices
		diffA := make([]uint64, 0)
		inter := make([]uint64, 0)
		diffB := make([]uint64, 0)
		// Interate slice A
		for _, idA := range sliceA {
			// New intersection flag
			isInter := false
			// Interate slice B
			for _, idB := range sliceB {
				// Is intersection
				if idA == idB {
					// Set isInter to true
					isInter = true
					// Append to inter
					inter = append(inter, idA)
					break
				}
			}
			if !isInter {
				// Append to diffA
				diffA = append(diffA, idA)
			}
		}

		// Check inter
		if len(inter) == 0 {
			// Intersection is empnty
			diffB = sliceB
		} else if len(inter) != len(sliceB) {
			// Intersection is not empty
			for i, idB := range sliceB {
				// New different flag
				isInter := false
				// New intersect count
				count := 0
				// Interate slice inter
				for _, strI := range inter {
					if idB == strI {
						count++
						isInter = true
						break
					}
				}

				if !isInter {
					// Append to diffB
					diffB = append(diffB, idB)
				} else if count == len(inter) {
					// All intersections be found, append sliceB remaining strs to diffB
					diffB = append(diffB, sliceB[i+1:]...)
					break
				}
			}
		}

		return diffA, diffB
	}

	// 1. Slice A is not empty and slice B is empty
	// 2. Slice A is empty and slice B is not empty
	// 3. Slice A is empty and slice B is empty
	return sliceA, sliceB
}
