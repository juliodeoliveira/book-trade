let otherPersonsBookId = null;
let otherPersonsBookName = null;
let yourBookId = null;
let yourBookName = null;

function openPopup(button) {
    document.getElementById("popup").classList.add("active");
    document.getElementById("overlay").classList.add("active");

    // TODO: proxima feat é o chat com websocket para fazer os usuários se comunicarem para trocar o livro
    console.log("Livro a ser trocado (dos seus livros): " + button.dataset.othersBookName);
     // Livro da outra pessoa
    otherPersonsBookId = button.dataset.bookId;
    otherPersonsBookName = button.dataset.othersBookName;

}

function closePopup() {
    document.getElementById("popup").classList.remove("active");
    document.getElementById("overlay").classList.remove("active");

    otherPersonsBookId = null;
    otherPersonsBookName = null;
}

function openConfirmation(cardElement) {
    
    yourBookId = cardElement.dataset.bookId;
    yourBookName = cardElement.dataset.yourBookName;

    console.log("Seu livro:", yourBookId, yourBookName);
    console.log("Livro da outra pessoa:", otherPersonsBookId, otherPersonsBookName);

    // Exibe a confirmação
    Swal.fire({
        title: "Tem certeza?",
        text: `Livro de outra pessoa: ${otherPersonsBookName}, Seu livro: ${yourBookName}`,
        icon: "warning",
        showCancelButton: true,
        confirmButtonColor: "#3085d6",
        cancelButtonColor: "#d33",
        confirmButtonText: "Sim, deletar!",
        cancelButtonText: "Cancelar!"
    }).then((result) => {
        if (result.isConfirmed) {

            // TODO: conexao com websocket

            Swal.fire({
                title: "Troca será iniciada",
                icon: "success"
            }).then(() => {
                window.location.reload();
            });
           
        }
    })

}


function closeConfirmation() {
    const confirmacao = document.getElementById("confirmacao");
    confirmacao.classList.remove("active");

    selectedBook = null;
    
}