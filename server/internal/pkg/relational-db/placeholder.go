package relationalDB

func NewPlaceHolder(n int) string {
	if n == 0 {
		return ""
	} else if n == 1 {
		return "(?)"
	}
	var placeholder []byte = []byte{'('}
	for i := 0; i < n; i++ {
		placeholder = append(placeholder, '?')
		if i != n-1 {
			placeholder = append(placeholder, ',')
		}
	}
	placeholder = append(placeholder, ')')
	return string(placeholder)
}

func NewNPlaceHolder(nPlaceHolders, nParams int) string {
	if nParams == 0 {
		return ""
	} else if nPlaceHolders == 1 && nParams == 1 {
		return "(?)"
	}

	var placeholders []byte

	for i := 0; i < nPlaceHolders; i++ {
		placeholders = append(placeholders, NewPlaceHolder(nParams)...)
		if i != nPlaceHolders-1 {
			placeholders = append(placeholders, ',')
		}
	}
	return string(placeholders)
}
