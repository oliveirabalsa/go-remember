package main

type note float64

func (n note) between(start, end float64) bool {
	return float64(n) >= start && float64(n) <= end
}

func noteForConcept(n note) string {
	if n.between(9.0, 10.0) {
		return "A"
	} else if n.between(7.0, 8.99) {
		return "B"
	} else if n.between(5.0, 7.99) {
		return "C"
	} else if n.between(3.0, 4.99) {
		return "D"
	}
	return "E"
}

func sum(a, b int) int {
	return a + b
}

func noteForConceptSwitch(n note) string {
	switch {
	case n.between(9.0, 10.0):
		return "A"
	case n.between(7.0, 8.99):
		return "B"
	case n.between(5.0, 7.99):
		return "C"
	case n.between(3.0, 4.99):
		return "D"
	}
	return "E"
}

func init() {
	println(noteForConceptSwitch(9.8))
	println(noteForConceptSwitch(6.9))
	println(noteForConceptSwitch(2.1))
	println(sum(5, 5))
}
