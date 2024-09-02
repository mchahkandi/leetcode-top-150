func reverseWords(s string) string {

    str := ""
    var i = len(s) - 1

    for i >= 0 {
        if s[i] == ' ' {
            i--
        }else{

        j := i

        for ; j >= 0 && s[j] != ' ' ; j-- {
        }
            str += s[j+1:i+1] + " "
            i = j
        }

    }

    if str[len(str) - 1] == ' ' {
        str = str[:len(str) -1]
    }

    return str
}
