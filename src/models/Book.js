const pool = require("../config/db");

class Book {
    static async getAll() {
        try {
            const result = await pool.query("SELECT * FROM books");
            return result.rows;
        } catch (error) {
            throw new Error("erro aqui cara!");
        }
    }
}

module.exports = Book;