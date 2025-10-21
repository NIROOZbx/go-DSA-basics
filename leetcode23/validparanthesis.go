
type Stack struct {
	items []string
	len   int
}

func (s *Stack) Push(val string) {
	s.items = append(s.items, val)
	s.len++

}
func (s *Stack) Pop() string {
    if s.IsEmpty() {
        return ""
    }
	s.len--
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return top

}

func (s *Stack) Peek() string {
    if s.IsEmpty() {
        return ""
    }
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return s.len == 0
}

func isValid(s string) bool {
	var st Stack
	var Map = map[string]string{}
	Map[")"] = "("
    Map["}"] = "{"
    Map["]"] = "["

	for i := 0; i < len(s); i++ {
        char := s[i]
		if char == '(' || char == '{' || char == '[' {
			st.Push(string(s[i]))
		} else {
			if st.IsEmpty() {
				return false
			}
		res := st.Peek()
        ch:=string(s[i])
		if  ch== ")" && res == Map[ch] || ch == "}" && res == Map[ch] || ch == "]" &&res == Map[ch] {
			st.Pop()
		} else {
			return false
		}
        }
	}
		

	return st.IsEmpty()
}