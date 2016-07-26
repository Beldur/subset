package subset

import (
    "math/rand"
)

// Subset selects a subset from a list of integers for given clientId.
// It is based on `Deterministic Subsetting`.
// See the book "Site Reliability Engineering" page 238
func Subset(backends []int, clientId, size int) []int {
    subsetCount := len(backends) / size

    // Group client into rounds; each round uses the same shuffled list
    round := clientId / subsetCount

    rand.Seed(int64(round))
    for i := range backends {
        j := rand.Intn(i + 1)
        backends[i], backends[j] = backends[j], backends[i]
    }

    // The subset id corresponding to the current client
    subsetId := clientId % subsetCount

    start := subsetId * size

    return backends[start : start+size]
}
