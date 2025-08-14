function deleteBook(button) {
    Swal.fire({
        title: "Tem certeza?",
        text: "Você está deletando uma possibilidade. Tem certeza de que não é a certa?",
        icon: "warning",
        showCancelButton: true,
        confirmButtonColor: "#3085d6",
        cancelButtonColor: "#d33",
        confirmButtonText: "Sim, deletar!",
        cancelButtonText: "Cancelar!"
    }).then((result) => {
        if (result.isConfirmed) {

            let bookId = button.dataset.bookId
            
            fetch("/books/delete", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                credentials: "include",
                body: JSON.stringify({ id: bookId })

            }).then(async res => {
                const data = await res.json();

                if (!res.ok) {
                    throw new Error(data.error || "Erro desconhecido");
                }                        

                Swal.fire({
                    title: "Deletado!",
                    text: data.message,
                    icon: "success"
                }).then(() => {
                    window.location.reload();
                });

            }).catch(err => {
                Swal.fire({
                    title: "Erro!",
                    text: err.message,
                    icon: "error"
                });
            });
        }
    })
}