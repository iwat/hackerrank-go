package queues

type TaleOfTwoStacks struct {
	old_to_new     []int
	new_to_old     []int
	old_to_new_pos int
	new_to_old_pos int
}

func NewTaleOfTwoStacks() *TaleOfTwoStacks {
	return &TaleOfTwoStacks{make([]int, 100000), make([]int, 100000), -1, -1}
}

func (s *TaleOfTwoStacks) Enqueue(n int) {
	s.push_old_to_new(n)
}

func (s *TaleOfTwoStacks) Dequeue() {
	if s.new_to_old_pos == -1 {
		for s.old_to_new_pos >= 0 {
			s.push_new_to_old(s.pop_old_to_new())
		}
	}
	s.pop_new_to_old()
}

func (s *TaleOfTwoStacks) Peek() int {
	if s.new_to_old_pos == -1 {
		for s.old_to_new_pos >= 0 {
			s.push_new_to_old(s.pop_old_to_new())
		}
	}
	return s.new_to_old[s.new_to_old_pos]
}

func (s *TaleOfTwoStacks) push_old_to_new(n int) {
	s.old_to_new_pos++
	s.old_to_new[s.old_to_new_pos] = n
}

func (s *TaleOfTwoStacks) pop_old_to_new() int {
	tmp := s.old_to_new[s.old_to_new_pos]
	s.old_to_new_pos--
	return tmp
}

func (s *TaleOfTwoStacks) push_new_to_old(n int) {
	s.new_to_old_pos++
	s.new_to_old[s.new_to_old_pos] = n
}

func (s *TaleOfTwoStacks) pop_new_to_old() int {
	tmp := s.new_to_old[s.new_to_old_pos]
	s.new_to_old_pos--
	return tmp
}
