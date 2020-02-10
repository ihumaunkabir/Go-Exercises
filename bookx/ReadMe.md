### Bookx ###

Bookx is a simple RESTful API built in GO and MongoDB.

Requirements:
	Go
	MongoDB
	MongoDriver
	gorilla/mux


Endpoints:

	Insert Raw JSON body [POST] : /api/books
	Insert URL Params [POST]: /api/books/{title}/{author}/{psher}/{year}/{cat}/{bookid}
	Search by BookID [GET]: /api/bookid/{bookid}
	Search By Year [GET]: /api/booky/{year}
	Get all books [GET]: /api/books
	Delete by BookID [DELETE]: /api/bookdel/{bookid}


