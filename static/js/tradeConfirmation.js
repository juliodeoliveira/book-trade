function openConfirmation(cardElement) {
    
    const bookId = cardElement.dataset.bookId;
    console.log("ID do livro selecionado:", bookId);

    // Exibe a confirmação
    const confirmacao = document.getElementById("confirmacao");
    confirmacao.classList.add("active");

    // Guarda o ID selecionado em um atributo temporário do modal
    confirmacao.dataset.bookId = bookId;
}

function closeConfirmation() {
    const confirmacao = document.getElementById("confirmacao");
    confirmacao.classList.remove("active");
}