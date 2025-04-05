// This file MUST started everything in project

//const connectDB = require("./src/config/db");
const app = require("./src/config/server");

const PORT = 3000;

app.listen(PORT, () => console.log(`Servidor rodando na porta ${PORT}`));
