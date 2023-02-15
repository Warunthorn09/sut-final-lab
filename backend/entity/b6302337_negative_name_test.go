package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

type Customer struct {
	Name       string `valid:"required~name not bank"`
	Email      string
	CustomerID string
}

func TestCustomer(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Name not bank", func(t *testing.T) {
		c := Customer{
			Name:       "",
			Email:      "nine@gmail.com",
			CustomerID: "L9999999",
		}

		ok, err := govalidator.ValidateStruct(c)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("name not bank"))
	})
}
