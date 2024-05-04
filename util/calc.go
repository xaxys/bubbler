package util

// TODO: complete this file, this is a minimal implementation now
// many functions are untested and may not work properly

// ==================== Calc ====================

func Calculate(args ...interface{}) interface{} {
	c := &consumer{
		Args: args,
	}
	return calcTenary(c)
}

// ==================== consumer ====================

type consumer struct {
	Args []interface{}
}

func (c *consumer) Peek() interface{} {
	if len(c.Args) == 0 {
		return nil
	}
	return c.Args[0]
}

func (c *consumer) Next() interface{} {
	if len(c.Args) == 0 {
		return nil
	}
	ret := c.Args[0]
	c.Args = c.Args[1:]
	return ret
}

// ==================== Implementation ====================

func calcTenary(c *consumer) interface{} {
	oprand1 := calcOr(c)
	switch c.Peek() {
	case "?":
		c.Next()
		oprand2 := calcTenary(c)
		c.Next()
		oprand3 := calcTenary(c)
		if ToBool(oprand1) {
			return oprand2
		} else {
			return oprand3
		}
	default:
		return oprand1
	}
}

func calcOr(c *consumer) interface{} {
	oprand1 := calcAnd(c)
	switch c.Peek() {
	case "||", "or":
		c.Next()
		oprand2 := calcOr(c)
		return ToBool(oprand1) || ToBool(oprand2)
	default:
		return oprand1
	}
}

func calcAnd(c *consumer) interface{} {
	oprand1 := calcBitOr(c)
	switch c.Peek() {
	case "&&", "and":
		c.Next()
		oprand2 := calcAnd(c)
		return ToBool(oprand1) && ToBool(oprand2)
	default:
		return oprand1
	}
}

func calcBitOr(c *consumer) interface{} {
	oprand1 := calcBitXor(c)
	switch c.Peek() {
	case "|":
		c.Next()
		oprand2 := calcBitOr(c)
		return ToInt64(oprand1) | ToInt64(oprand2)
	default:
		return oprand1
	}
}

func calcBitXor(c *consumer) interface{} {
	oprand1 := calcBitAnd(c)
	switch c.Peek() {
	case "^":
		c.Next()
		oprand2 := calcBitXor(c)
		return ToInt64(oprand1) ^ ToInt64(oprand2)
	default:
		return oprand1
	}
}

func calcBitAnd(c *consumer) interface{} {
	oprand1 := calcEqual(c)
	switch c.Peek() {
	case "&":
		c.Next()
		oprand2 := calcBitAnd(c)
		return ToInt64(oprand1) & ToInt64(oprand2)
	default:
		return oprand1
	}
}

func calcEqual(c *consumer) interface{} {
	oprand1 := calcRelational(c)
	switch c.Peek() {
	case "==":
		c.Next()
		oprand2 := calcEqual(c)
		return oprand1 == oprand2
	case "!=":
		c.Next()
		oprand2 := calcEqual(c)
		return oprand1 != oprand2
	default:
		return oprand1
	}
}

func calcRelational(c *consumer) interface{} {
	oprand1 := calcShift(c)
	switch c.Peek() {
	case "<":
		c.Next()
		oprand2 := calcRelational(c)
		if IsFloat(oprand1) || IsFloat(oprand2) {
			return ToFloat64(oprand1) < ToFloat64(oprand2)
		}
		return ToInt64(oprand1) < ToInt64(oprand2)
	case ">":
		c.Next()
		oprand2 := calcRelational(c)
		if IsFloat(oprand1) || IsFloat(oprand2) {
			return ToFloat64(oprand1) > ToFloat64(oprand2)
		}
		return ToInt64(oprand1) > ToInt64(oprand2)
	case "<=":
		c.Next()
		oprand2 := calcRelational(c)
		if IsFloat(oprand1) || IsFloat(oprand2) {
			return ToFloat64(oprand1) <= ToFloat64(oprand2)
		}
		return ToInt64(oprand1) <= ToInt64(oprand2)
	case ">=":
		c.Next()
		oprand2 := calcRelational(c)
		if IsFloat(oprand1) || IsFloat(oprand2) {
			return ToFloat64(oprand1) >= ToFloat64(oprand2)
		}
		return ToInt64(oprand1) >= ToInt64(oprand2)
	default:
		return oprand1
	}
}

func calcShift(c *consumer) interface{} {
	oprand1 := calcAdditive(c)
	switch c.Peek() {
	case "<<":
		c.Next()
		oprand2 := calcShift(c)
		return ToInt64(oprand1) << ToInt64(oprand2)
	case ">>":
		c.Next()
		oprand2 := calcShift(c)
		return ToInt64(oprand1) >> ToInt64(oprand2)
	default:
		return oprand1
	}
}

func calcAdditive(c *consumer) interface{} {
	oprand1 := calcMultiplicative(c)
	switch c.Peek() {
	case "+":
		c.Next()
		oprand2 := calcAdditive(c)
		if IsFloat(oprand1) || IsFloat(oprand2) {
			return ToFloat64(oprand1) + ToFloat64(oprand2)
		}
		return ToInt64(oprand1) + ToInt64(oprand2)
	case "-":
		c.Next()
		oprand2 := calcAdditive(c)
		if IsFloat(oprand1) || IsFloat(oprand2) {
			return ToFloat64(oprand1) - ToFloat64(oprand2)
		}
		return ToInt64(oprand1) - ToInt64(oprand2)
	default:
		return oprand1
	}
}

func calcMultiplicative(c *consumer) interface{} {
	oprand1 := calcUnary(c)
	switch c.Peek() {
	case "*":
		c.Next()
		oprand2 := calcMultiplicative(c)
		if IsFloat(oprand1) || IsFloat(oprand2) {
			return ToFloat64(oprand1) * ToFloat64(oprand2)
		}
		return ToInt64(oprand1) * ToInt64(oprand2)
	case "/":
		c.Next()
		oprand2 := calcMultiplicative(c)
		if IsFloat(oprand1) || IsFloat(oprand2) {
			return ToFloat64(oprand1) / ToFloat64(oprand2)
		}
		return ToInt64(oprand1) / ToInt64(oprand2)
	case "%":
		c.Next()
		oprand2 := calcMultiplicative(c)
		return ToInt64(oprand1) % ToInt64(oprand2)
	default:
		return oprand1
	}
}

func calcUnary(c *consumer) interface{} {
	switch c.Peek() {
	case "+":
		c.Next()
		return calcUnary(c)
	case "-":
		c.Next()
		if IsFloat(c.Peek()) {
			return -ToFloat64(calcUnary(c))
		}
		return -ToInt64(calcUnary(c))
	case "~":
		c.Next()
		return ^ToInt64(calcUnary(c))
	case "!":
		c.Next()
		return !ToBool(calcUnary(c))
	default:
		return calcPrimary(c)
	}
}

func calcPrimary(c *consumer) interface{} {
	switch c.Peek() {
	case "(":
		c.Next()
		ret := calcTenary(c)
		if c.Next() != ")" {
			panic("expected ')'")
		}
		return ret
	default:
		return c.Next()
	}
}
