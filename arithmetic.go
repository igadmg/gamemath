package gamemath

// GCD вычисляет наибольший общий делитель двух чисел с помощью алгоритма Евклида.
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM вычисляет наименьшее общее кратное двух чисел.
func LCM(rest ...int) int {
	if len(rest) > 2 {
		rest[len(rest)-2] = LCM(rest[len(rest)-1], rest[len(rest)-2])
		return LCM(rest[:len(rest)-1]...)
	}
	if len(rest) == 0 {
		return 0
	}
	if len(rest) == 1 {
		return rest[0]
	}

	a := rest[0]
	b := rest[1]

	// Обрабатываем случай, когда одно из чисел равно нулю
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	// Ищем абсолютное значение, так как НОК - положительное число.
	// Иначе деление на НОД может быть некорректным, если числа отрицательные.
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return (a * b) / GCD(a, b)
}

// LCM вычисляет наименьшее общее кратное двух чисел.
func WeightedLCM(rest ...int) (gdc int, w []float32) {
	gdc = LCM(rest...)
	w = DivideF(gdc, rest...)
	return
}
