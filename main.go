package main
import(	

	//"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	//"github.com/sauravpd/go-crud/handler"
	//"errors"
)

type book struct {
	ID       string  `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book

	// Bind the JSON request body to the newBook struct
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Append the newBook to the books slice
	books = append(books, newBook)

	// Respond with the newly created book
	c.JSON(http.StatusCreated, newBook)
}

func deleteBook(c *gin.Context) {
	// Get the book ID from the URL parameter
	id := c.Param("id")

	// Find the index of the book with the specified ID
	index := -1
	for i, b := range books {
		if b.ID == id {
			index = i
			break
		}
	}

	// If the book is not found, respond with 404 Not Found
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Remove the book from the slice
	books = append(books[:index], books[index+1:]...)

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}



func main() {
	 router := gin.Default()
	 router.GET("/books", getBooks)
	 router.POST("/books", createBook)
	 router.DELETE("/books/:id", deleteBook)
	 router.Run("localhost:8080")
}

