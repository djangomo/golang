package main

import (
	"math/rand"
	"testing"
)


//standard version
func benchmark_GenerateFullPW(b *testing.B) string{
	pw := make([]rune, b.N)
	for i := range pw {
		pw[i] = chars[rand.Intn(len(chars))]
	}
	return string(pw)
}

//version with constants & remainder
func benchmark_GenerateUpperLowerPW(b *testing.B) string {
	pw := make([]byte, b.N)
	for i := range pw {
		pw[i] = benchmark[rand.Int63() % int64(len(benchmark))]
	}
	return string(pw)
}

//version with constants and masking
func benchmark_GenerateLowerChars(b *testing.B) string {
	pw := make([]byte, b.N)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := b.N-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(benchmark) {
			pw[i] = benchmark[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(pw)
}

//version with improved src
func benchmark_GenerateUpperChar(b *testing.B)string {
	pw := make([]byte, b.N)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := b.N-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(benchmark) {
			pw[i] = benchmark[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(pw)
}



func BenchmarkVersion1(b *testing.B){benchmark_GenerateFullPW(b)}
func BenchmarkVersion2(b *testing.B){benchmark_GenerateUpperLowerPW(b)}
func BenchmarkVersion3(b *testing.B){benchmark_GenerateLowerChars(b)}
func BenchmarkVersion4(b *testing.B){benchmark_GenerateUpperChar(b)}


