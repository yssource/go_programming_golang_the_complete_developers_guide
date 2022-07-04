package main

import "fmt"

type BeltSize int

const (
	Small BeltSize = iota
	Medium
	Large
)

func (b BeltSize) String() string {
	// return []string{"Small", "Medium", "Large"}[b] // NOTE: order matters

	switch b {
	case Small:
		return fmt.Sprintf("Small")
	case Medium:
		return fmt.Sprintf("Medium")
	case Large:
		return fmt.Sprintf("Large")
	default:
		panic("invalid type")
	}
}

type Shipping int

const (
	Ground Shipping = iota
	Air
)

func (s Shipping) String() string {
	// return []string{"Ground", "Air"}[b]

	switch s {
	case Ground:
		return fmt.Sprintf("Ground")
	case Air:
		return fmt.Sprintf("Air")
	default:
		panic("invalid type")
	}
}

type Conveyer interface {
	Convey() BeltSize
}

type Shipper interface {
	Ship() Shipping
}

type WarehouseAutomator interface {
	Conveyer
	Shipper
}

type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return fmt.Sprintf("spam mail")
}

func (s *SpamMail) Ship() Shipping {
	return Air
}

func (s *SpamMail) Convey() BeltSize {
	return Small
}

func automate(item WarehouseAutomator) {
	fmt.Printf("Convey %v on %v conveyer\n", item, item.Convey())
	fmt.Printf("Ship %v via %v\n", item, item.Ship())
}

type ToxicWaste struct {
	weight int
}

func (t *ToxicWaste) Ship() Shipping {
	return Ground
}

func main() {
	mail := SpamMail{40000}
	automate(&mail)

	// waste := ToxicWaste{3000}
	// automate(&waste) // Error
}
