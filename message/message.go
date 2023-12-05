package message

import "errors"

// Arith 是包含远程调用方法的类型
type Arith int

// Args 是传递给远程方法的参数
type Args struct {
    A, B int
}

// Quotient 是远程方法的返回类型
type Quotient struct {
    Quo, Rem int
}

// Multiply 是远程方法，计算两个整数的乘积
func (t *Arith) Multiply(args *Args, reply *int) error {
    *reply = args.A * args.B
    return nil
}

// Divide 是远程方法，计算两个整数的商和余数
func (t *Arith) Divide(args *Args, quo *Quotient) error {
    if args.B == 0 {
        return errors.New("divide by zero")
    }
    quo.Quo = args.A / args.B
    quo.Rem = args.A % args.B
    return nil
}

