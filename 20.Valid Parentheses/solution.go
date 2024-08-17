func isValid(s string) bool {

    brackets := map[byte]byte{
        ')' : '(',
        '}' : '{',
        ']' : '[',
    }

    stack := []byte{}

    for _, val := range []byte(s) {
        key, exists := brackets[val]

        if !exists {
            stack = append(stack,val)
            continue
        }

        if len(stack) == 0 {
            return false
        }

        if stack[len(stack) - 1] != key {
            return false
        }

        stack = stack[:len(stack) - 1 ]
    }

    return len(stack) == 0
}
