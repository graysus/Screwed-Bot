package randomscrewed

import "math/big"

func BytesHashPhase(thing []byte, multiplier int64, start_x int64, lim_x int64) *big.Int {
	Base128 := big.NewInt(1)
	Base128.Lsh(Base128, 128).Sub(Base128, big.NewInt(1))
	n := big.NewInt(0)
	for _, i := range thing {
		n.Mul(n, big.NewInt(multiplier))
		n.Add(n, big.NewInt(int64(i)))
		n.And(n, Base128)
	}
	final := big.NewInt(0)

	for value := range lim_x - start_x {
		x := value + start_x
		n2 := big.NewInt(1)
		for range x {
			n2.Mul(n2, n)
			n2.And(n2, Base128)
		}
		n2.Div(n2, big.NewInt(x))
		final.Add(final, n2)
		final.And(final, Base128)
	}
	return final
}

func BytesHash(thing []byte) *big.Int {
	Base128 := big.NewInt(1)
	Base128.Lsh(Base128, 128).Sub(Base128, big.NewInt(1))

	n := big.NewInt(0)
	n.Add(n, BytesHashPhase(thing, 5, 2, 10))
	n.Add(n, BytesHashPhase(thing, 3, 5, 14))
	n.Add(n, BytesHashPhase(thing, 7, 7, 13))
	n.And(n, Base128)
	return n
}
