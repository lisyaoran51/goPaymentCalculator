package goPaymentCalculator

/*
 ********************************
 *								*
 *								*
 *      discountCalculator      *
 *								*
 *								*
 ********************************
 */

type DiscountCalculatorConfig struct {
	DiscountMap map[MEMBER_PREVILEDGE]float32
}

type DiscountCalculator struct {
	//EmptyCalculator
	// implementing compoistion over inheritance
	EmptyCalculator
	discountMap map[MEMBER_PREVILEDGE]float32
}

func NewDiscountCalculator(config *DiscountCalculatorConfig) *DiscountCalculator {
	discountCalculator := &DiscountCalculator{
		EmptyCalculator: *NewEmptyCalculator(),
		discountMap:     config.DiscountMap,
	}
	discountCalculator.PaymentCalculator = discountCalculator
	return discountCalculator
}

func (d *DiscountCalculator) Calculate(memeber PaymentMember, costPoint, costCoin float32) (point, coin float32, err error) {
	point, coin, _ = d.EmptyCalculator.Calculate(memeber, costPoint, costCoin)
	if d.child != nil {
		point, coin, err = d.child.Calculate(memeber, point, coin)
	}

	if discount, ok := d.discountMap[memeber.GetPriviledge()]; ok {
		point *= discount
		coin *= discount
	}
	return
}
