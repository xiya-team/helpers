package helpers

import (
	"math"
	"math/cmplx"
	"math/rand"
	"strconv"
	"time"
)

// Abs abs()
func Abs(number float64) float64 {
	return math.Abs(number)
}

// Acos - Arc cosine
func Acos(x complex128) complex128 {
	return cmplx.Acos(x)
}

// Acosh - Inverse hyperbolic cosine
func Acosh(x complex128) complex128 {
	return cmplx.Acosh(x)
}

// Asin - Arc sine
func Asin(x complex128) complex128 {
	return cmplx.Asin(x)
}

// Asinh - Inverse hyperbolic sine
func Asinh(x complex128) complex128 {
	return cmplx.Asinh(x)
}

// Atan2 - Arc tangent of two variables
func Atan2(y, x float64) float64 {
	return math.Atan2(y, x)
}

// Atan - Arc tangent
func Atan(x complex128) complex128 {
	return cmplx.Atan(x)
}

// Atanh - Inverse hyperbolic tangent
func Atanh(x complex128) complex128 {
	return cmplx.Atanh(x)
}

// BaseConvert - Convert a number between arbitrary bases
func BaseConvert(num string, frombase, tobase int) (string, error) {
	i, err := strconv.ParseInt(num, frombase, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, tobase), nil
}

// Ceil - Round fractions up
func Ceil(x float64) float64 {
	return math.Ceil(x)
}

// Cos - Cosine
func Cos(x float64) float64 {
	return math.Cos(x)
}

// Cosh - Hyperbolic cosine
func Cosh(x float64) float64 {
	return math.Cosh(x)
}

// Decbin - Decimal to binary
func Decbin(x int64) string {
	return strconv.FormatInt(x, 2)
}

// Dechex - Decimal to hexadecimal
func Dechex(x int64) string {
	return strconv.FormatInt(x, 16)
}

// Decoct - Decimal to octal
func Decoct(x int64) string {
	return strconv.FormatInt(x, 8)
}

// Exp - Calculates the exponent of e
func Exp(x float64) float64 {
	return math.Exp(x)
}

// Expm1 - Returns exp(number) - 1
// computed in a way that is accurate even when the value of number is close to zero
func Expm1(x float64) float64 {
	return math.Exp(x) - 1
}

// Floor - Round fractions down
func Floor(x float64) float64 {
	return math.Floor(x)
}

// IsFinite - Finds whether a value is a legal finite number
func IsFinite(f float64, sign int) bool {
	return !math.IsInf(f, sign)
}

// IsInfinite - Finds whether a value is infinite
func IsInfinite(f float64, sign int) bool {
	return math.IsInf(f, sign)
}

// IsNan - Finds whether a value is not a number
func IsNan(f float64) bool {
	return math.IsNaN(f)
}

// Log - Natural logarithm
func Log(x float64) float64 {
	return math.Log(x)
}

// Log10 - Base-10 logarithm
func Log10(x float64) float64 {
	return math.Log10(x)
}

// Log1p - Returns log(1 + number)
// computed in a way that is accurate even when the value of number is close to zero
func Log1p(x float64) float64 {
	return math.Log1p(x)
}


// Max max()
func Max(nums ...float64) float64 {
	if len(nums) < 2 {
		panic("nums: the nums length is less than 2")
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		max = math.Max(max, nums[i])
	}
	return max
}

// Min min()
func Min(nums ...float64) float64 {
	if len(nums) < 2 {
		panic("nums: the nums length is less than 2")
	}
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		min = math.Min(min, nums[i])
	}
	return min
}

// Pi - Get value of pi
func Pi() float64 {
	return math.Pi
}

// Pow - Exponential expression
func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

// Rand rand()
// Range: [0, 2147483647]
func Rand(min, max int) int {
	if min > max {
		panic("min: min cannot be greater than max")
	}
	// PHP: getrandmax()
	if int31 := 1<<31 - 1; max > int31 {
		panic("max: max can not be greater than " + strconv.Itoa(int31))
	}
	if min == max {
		return min
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max+1-min) + min
}

// Round - Rounds a float
func Round(x float64) float64 {
	return math.Round(x)
}

// Sin - Sine
func Sin(x float64) float64 {
	return math.Sin(x)
}

// Sinh - Hyperbolic sine
func Sinh(x float64) float64 {
	return math.Sinh(x)
}

// Sqrt - Square root
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Tan - Tangent
func Tan(x float64) float64 {
	return math.Tan(x)
}

// Tanh - Hyperbolic tangent
func Tanh(x float64) float64 {
	return math.Tanh(x)
}
