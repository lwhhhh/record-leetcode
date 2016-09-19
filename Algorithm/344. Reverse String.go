func reverseString(s string) string {
    var b []byte
    for index:=len(s)-1;index>=0;index-- {
        b = append(b,s[index])
    }
    return string(b)
}
