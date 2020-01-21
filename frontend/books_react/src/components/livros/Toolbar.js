import React, { useState, Fragment } from "react";
import InserirLivros from "./InserirLivros";

const Toolbar = () => {
  const [addBook, setAddBook] = useState(false);
  return (
    <Fragment>
      <div className="toolbar">
        {addBook == false ? (
          <button
            onClick={() => {
              setAddBook(!addBook);
            }}
          >
            <i class="fas fa-plus"></i> Novo livro
          </button>
        ) : (
          <button
            onClick={() => {
              setAddBook(!addBook);
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
