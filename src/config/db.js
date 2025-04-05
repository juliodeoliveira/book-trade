const { Pool }= require("pg");

const pool = new Pool({
    user: "julio",
    host: "localhost",
    database: "book_trade",
    password: "DBdojulio", 
    port: 5432
});

module.exports = pool;