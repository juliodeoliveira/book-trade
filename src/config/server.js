// Config related to the server

const express = require("express");
const path = require("path");
// const userRoutes = require("../routes/userRoutes");
const bookRoutes = require("../routes/bookRoutes");

// import BookController from "../controllers/bookController";
const BookController = require("../controllers/bookController");
// const BookController = await import("../controllers/bookController");

const app = express();
app.use(express.json());

app.set("view engine", "ejs");

const livros = [
    { titulo: "1984", autor: "George Orwell" },
    { titulo: "O Senhor dos Anéis", autor: "J.R.R. Tolkien" },
    { titulo: "Dom Casmurro", autor: "Machado de Assis" }
];

let teste = (req, res) => BookController.allBooks(req, res);
console.log(teste);

// app.use("/user", userRoutes);
app.use("/books", bookRoutes);

app.get("/", BookController.allBooks);

module.exports = app;