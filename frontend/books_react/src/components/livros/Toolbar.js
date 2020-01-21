import React, { useContext, Fragment } from "react";
import InserirLivros from "./InserirLivros";
import LivroContext from "../../context/livros/livroContext";

const Toolbar = () => {
  const livroContext = useContext(LivroContext);
  const { addBook, setAddBook, closeAddBook } = livroContext;

  return (
    <Fragment>
      <div className="toolbar">
        {addBook == false ? (
          <button
            onClick={() => {
              setAddBook();
            }}
          >
            <i class="fas fa-plus"></i> Novo livro
          </button>
        ) : (
          <button
            onClick={() => {
              closeAddBook();
            }}
          >
            <i class="far fa-times-circle"></i> Fechar formulário
          </button>
        )}
      </div>
      <Fragment>
        {addBook == false ? null : (
          <div>
            <InserirLivros />
          </div>
        )}
      </Fragment>
    </Fragment>
  );
};

export default Toolbar;
