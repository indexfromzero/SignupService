package service_test

import (
	"testing"

	"github.com/microbusinesses/SignupService/domain"
	. "github.com/microbusinesses/SignupService/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Signup method input parameters", func() {
	var (
		service                   SignupService
		emptyEmailAndUsernameUser domain.User
		emptyPasswordUser         domain.User
	)

	BeforeEach(func() {
		service = SignupService{}
		email := "test@email.com"
		username := "test"
		password := "password"

		var emptyEmail string
		var emptyUsername string
		var emptyPassword string

		emptyEmailAndUsernameUser = domain.User{emptyEmail, emptyUsername, password}
		emptyPasswordUser = domain.User{email, username, emptyPassword}
	})

	Context("when empty email and username provided", func() {
		It("should panic", func() {
			Ω(func() { service.Signup(emptyEmailAndUsernameUser) }).Should(Panic())
		})
	})

	Context("when empty password provided", func() {
		It("should panic", func() {
			Ω(func() { service.Signup(emptyPasswordUser) }).Should(Panic())
		})
	})
})

func TestSignup(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Signup method input parameters")
}
