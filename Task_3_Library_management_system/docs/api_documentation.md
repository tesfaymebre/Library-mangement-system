# Library Management System

## Overview

This document provides an overview of the Library Management System, which demonstrates the use of structs, interfaces, and other Go functionalities such as methods, slices, and maps. This is a console-based application.

## Structs

### Book

- **ID**: int
- **Title**: string
- **Author**: string
- **Status**: string // "Available" or "Borrowed"

### Member

- **ID**: int
- **Name**: string
- **BorrowedBooks**: []Book

## Interface

### LibraryManager

- `AddBook(book Book) error`
- `AddMember(member Member) error`
- `RemoveBook(bookID int) error`
- `BorrowBook(bookID int, memberID int) error`
- `ReturnBook(bookID int, memberID int) error`
- `ListAvailableBooks() ([]Book, error)`
- `ListBorrowedBooks(memberID int) []Book`

## Implementation

The `LibraryManager` interface is implemented in the `Library` struct, which stores all books and members in maps. The methods provide functionality to add, remove, borrow, and return books, as well as list available and borrowed books.

## Console Interaction

The application interacts with the user via the console. The following operations are supported:

- **Add Book**: Adds a new book to the library.
- **Add Member**: Adds a new member to the library.
- **Remove Book**: Removes an existing book from the library by its ID.
- **Borrow Book**: Allows a member to borrow a book if it is available.
- **Return Book**: Allows a member to return a borrowed book.
- **List Available Books**: Lists all available books in the library.
- **List Borrowed Books**: Lists all books borrowed by a specific member.

## Testing

Run the application using the `go run main.go` command. The console will prompt for user input to perform various library management operations.
