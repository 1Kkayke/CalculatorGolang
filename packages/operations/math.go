package operations

import "strconv"

func add(a string, b string) string {
	num1, err := strconv.Atoi(a)

	if err != nil {
		return "Invalid input"
	}

	num2, err := strconv.Atoi(b)

	if err != nil {
		return "Invalid input"
	}

	realResult := num1 + num2

	return strconv.Itoa(realResult)
}

func sub(a string, b string) string {
	num1, err := strconv.Atoi(a)

	if err != nil {
		return "Invalid input"
	}

	num2, err := strconv.Atoi(b)

	if err != nil {
		return "Invalid input"
	}

	realResult := num1 - num2

	return strconv.Itoa(realResult)

}

func div(a string, b string) string {

	num1, err := strconv.Atoi(a)

	if err != nil {
		return "Invalid input"
	}

	num2, err := strconv.Atoi(b)

	if err != nil {
		return "Invalid input"
	}

	realResult := num1 / num2

	return strconv.Itoa(realResult)

}

func multi(a string, b string) string {

	num1, err := strconv.Atoi(a)

	if err != nil {
		return "Invalid input"
	}

	num2, err := strconv.Atoi(b)

	if err != nil {
		return "Invalid input"
	}

	realResult := num1 * num2

	return strconv.Itoa(realResult)

}
