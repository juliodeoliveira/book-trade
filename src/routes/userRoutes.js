const express = require("express");
const router = express.Router();

router.get("/", () => console.log("Chegou a uma página! (user)"));