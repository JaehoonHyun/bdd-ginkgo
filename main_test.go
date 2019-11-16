package main

import (
	"fmt"

	. "github.com/JaehoonHyun/bdd-ginkgo/internal/books"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Book", func() {
	var (
		longBook  Book
		shortBook Book
	)

	BeforeEach(func() {
		longBook = Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  1488,
		}

		shortBook = Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	//Categorizing 하는 시나리오에서
	Describe("Categorizing book length", func() {

		BeforeEach(func() {
			longBook.Pages = 10
		})

		//300 페이지 보다 큰 조건에서
		Context("With more than 300 pages", func() {

			//longBook은 novel일 것이다.
			It("should be a novel", func() {
				fmt.Println(longBook)
				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})

		//300페이지보다 작은 조건에서
		Context("With fewer than 300 pages", func() {
			//longbook은 short story일 것이다.
			It("should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})
	})

})
