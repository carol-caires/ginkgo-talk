package fibonacci_test

import (
	"ginkgo-talk/fibonacci"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFibonacci(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fibonacci Suite")
}

var _ = Describe("Fibonacci", func() {

	Describe("Dada a sequencia desejada, retornar o número de Fibonacci correspondente", func() {
		var (
			n   int
			f   int
			err error
		)

		BeforeEach(func() {
			n = 6
		})

		JustBeforeEach(func() {
			f, err = fibonacci.F(n)
		})

		Context("quando a sequencia é 6", func() {
			Specify("Deve retornar o sexto número da sequencia de Fibonacci", func() {
				Expect(f).Should(Equal(8))
			})
			It("Deve não retornar erro", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})
		Context("quando a sequencia é 10", func() {
			BeforeEach(func() {
				n = 10
			})
			Specify("Deve retornar o décimo número da sequencia de Fibonacci", func() {
				Expect(f).Should(Equal(55))
			})
			It("Deve não retornar erro", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})
		Context("quando a sequencia é 0", func() {
			BeforeEach(func() {
				n = 0
			})
			Specify("Deve retornar zero", func() {
				Expect(f).Should(BeZero())
			})
			It("Deve retornar erro", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
