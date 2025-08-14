function openPopup(button) {
    document.getElementById("popup").classList.add("active");
    document.getElementById("overlay").classList.add("active");

    // TODO: proxima feat é o chat com websocket para fazer os usuários se comunicarem para trocar o livro
    console.log("Livro a ser trocado: " + button.dataset.bookId);
}

function closePopup() {
    document.getElementById("popup").classList.remove("active");
    document.getElementById("overlay").classList.remove("active");
}