const Book = require("../models/Book");

class BookController {
    static async allBooks (req, res) {
        try {
            const books = await Book.getAll();
            res.render("index", { books });
        } catch (error) {
            console.log(error);
            res.status(500).send("Houve um erro!");
        }
    };
}

module.exports = BookController;

// Validar nomes de livros em apis externas, usuário digita o 
// nome e pesquisa pelo nome na api e preenche automaticamente os outros campos