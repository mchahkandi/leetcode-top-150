func maxArea(height []int) int {
    x, y := 0, len(height) - 1
    maxArea := 0

    for x < y {
        currentArea := calcArea(x, y, height)
        if currentArea > maxArea {
            maxArea = currentArea
        }

        if height[x] < height[y] {
            x++
        } else {
            y--
        }
    }

    return maxArea
}

func calcArea(left int, right int, height []int) int {
    minHeight := min(height[left], height[right])
    return minHeight * (right - left)
}
