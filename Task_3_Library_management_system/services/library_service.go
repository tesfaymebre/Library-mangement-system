package services

import (
	"errors"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book) error
	AddMember(member models.Member) error
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() ([]models.Book, error)
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	books   map[int]models.Book
	members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}

func (library *Library) AddBook(book models.Book) error {
	_, exists := library.books[book.ID]

	if exists {
		return errors.New("book already exists")
	}

	library.books[book.ID] = book
	return nil
}

func (library *Library) AddMember(member models.Member) error {
	_, exists := library.members[member.ID]
	if exists {
		return errors.New("member already exists")
	}

	library.members[member.ID] = member
	return nil
}

func (library *Library) RemoveBook(bookID int) error {
	_, exists := library.books[bookID]
	if !exists {
		return errors.New("book doesn't exist")
	}

	book := library.books[bookID]
	if book.Status == "Borrowed" {
		return errors.New("book is not returned yet")
	}

	delete(library.books, bookID)
	return nil
}

func (library *Library) BorrowBook(bookID int, memberID int) error {
	book, exists := library.books[bookID]

	if !exists {
		return errors.New("book doesn't exist")
	} else if book.Status == "Borrowed" {
		return errors.New("the book is not available now")
	}

	member, exists := library.members[memberID]

	if !exists {
		return errors.New("memeber doesn't exist")
	}

	book.Status = "Borrowed"
	library.books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	library.members[memberID] = member
	return nil
}

func (library *Library) ReturnBook(bookID int, memberID int) error {
	book, exists := library.books[bookID]

	if !exists || book.Status == "Available" {
		return errors.New("this book is not from this library")
	}

	member, exists := library.members[memberID]

	if !exists {
		return errors.New("member doesn't exist")
	}

	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			book.Status = "Available"
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			library.books[bookID] = book
			library.members[memberID] = member
			return nil
		}
	}

	return errors.New("the book isn't available in the list of these member borrowed")
}

func (library *Library) ListAvailableBooks() ([]models.Book, error) {
	var availableBooks []models.Book

	for _, book := range library.books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}

	if len(availableBooks) == 0 {
		return []models.Book{}, errors.New("no books available")
	}

	return availableBooks, nil
}

func (library *Library) ListBorrowedBooks(memberID int) ([]models.Book, error) {
	member, exists := library.members[memberID]

	if !exists {
		return []models.Book{}, errors.New("member doesn't exist")
	}

	if len(member.BorrowedBooks) == 0 {
		return []models.Book{}, errors.New("no books borrowed")
	}

	return member.BorrowedBooks, nil
}
