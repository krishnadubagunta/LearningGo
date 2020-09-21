package Map

type StringInteger map[string]int
type IntegerString map[int]string

// Create type for Map StringInteger which maps a string to an integer and create a function for the type to return -1 if no value is present for a key.

func (s StringInteger) GetValueAdd(forKey string) *int {
	if v, ok := s[forKey]; ok {
		return &v
	}
	return nil
}

func (s IntegerString) GetValueAdd(forKey int) *string {
	if v, ok := s[forKey]; ok {
		return &v
	}
	return nil
}