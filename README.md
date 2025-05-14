API que usa o GORM para auxiliar nas operações de criar, consultar e apagar livros do banco de dados
rotas: /api ->/create_books (necessário body estruturado por autor, titulo e editora)
            ->delete_book/:id
            ->/get_books/:id
            ->/books
