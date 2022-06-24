//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project
package main

import (
	"fmt"
	"time"
)

// a member's name
type Name string

// a books title
type Title string

type Member struct {
	name  Name
	books map[Title]LendAudit
}

func (m *Member) audit() {
	fmt.Println()
	for title, audit := range m.books {
		var returnedTime string
		if audit.checkIn.IsZero() {
			returnedTime = "[not returned yet]"
		} else {
			returnedTime = audit.checkIn.String()
		}
		fmt.Println(m.name, ":", title, ":", audit.checkOut.String(), "through",
			returnedTime)
	}
	fmt.Println()
}

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type BookEntry struct {
	// total books owned by the library
	total int
	// total books lended out
	lended int
}

type Library struct {
	books   map[Title]BookEntry
	members map[Name]Member
}

func (l *Library) auditMembers() {
	for _, member := range l.members {
		if len(member.books) > 0 {
			member.audit()
		}
	}
}

func (l *Library) printBooks() {
	fmt.Println()
	for title, entry := range l.books {
		fmt.Println(title, ":\n\t\t\ttotal:\t", entry.total, "\tlended:\t", entry.lended)
	}
	fmt.Println()
}

func (l *Library) checkOutBook(title Title, member *Member) bool {
	book, found := l.books[title]
	if !found {
		fmt.Println(title, "is not part of library")
		return false
	}

	if book.lended == book.total {
		fmt.Println("no more books available to lend")
		return false
	}

	book.lended += 1
	l.books[title] = book

	member.books[title] = LendAudit{
		checkOut: time.Now(),
	}
	return true
}

func (l *Library) returnBook(title Title, member *Member) bool {
	book, found := l.books[title]
	if !found {
		fmt.Println(title, "is not part of library")
		return false
	}

	audit, found := member.books[title]
	if !found {
		fmt.Println(member.name, "did't checkout this book")
		return false
	}

	book.lended -= 1
	l.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit

	return true
}

func main() {
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	library.books["Blindness"] = BookEntry{total: 4, lended: 0}
	library.books["Dune"] = BookEntry{total: 3, lended: 0}
	library.books["Pride and prejudice"] = BookEntry{total: 2, lended: 0}
	library.books["Lolita"] = BookEntry{total: 1, lended: 0}

	library.members["Mustafa"] = Member{name: "Mustafa", books: make(map[Title]LendAudit)}
	library.members["Malena"] = Member{name: "Malena", books: make(map[Title]LendAudit)}
	library.members["Maya"] = Member{name: "Maya", books: make(map[Title]LendAudit)}

	fmt.Println("\nInitial status of library:")
	library.printBooks()
	library.auditMembers()

	member := library.members["Mustafa"]
	if checkedOut := library.checkOutBook("Dune", &member); checkedOut {
		fmt.Println("\nAfter checking out a book:")
		library.printBooks()
		library.auditMembers()
	}

	if returned := library.returnBook("Dune", &member); returned {
		fmt.Println("\nAfter checking in a book:")
		library.printBooks()
		library.auditMembers()
	}
}
