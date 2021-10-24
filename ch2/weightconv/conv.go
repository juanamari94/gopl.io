package weightconv

func KgToLb(kg Kilograms) Pounds { return Pounds(kg / 2.205) }
func LbToKg(lb Pounds) Kilograms { return Kilograms(lb * 2.205) }
