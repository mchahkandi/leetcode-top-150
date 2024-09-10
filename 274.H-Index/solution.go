func hIndex(citations []int) int {

    sort.Slice(citations, func(a, b int) bool {
      return citations[b] < citations[a]
    })
    h := 0
    for h < len(citations) && citations[h] >= h + 1 {
        h++
    }

    return h
}
