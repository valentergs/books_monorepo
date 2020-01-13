import React, { useContext, useEffect } from "react";
import LivroContext from "../../context/livros/livroContext";

const Livros = () => {
  const livroContext = useContext(LivroContext);
  const { livroState, todosLivros, setCurrent } = livroContext;

  useEffect(() => {
    todosLivros();
  }, []);

  return (
    <div className="livros">
      <div className="caixa-cabeca">
        <h3 className="caixa-titulo">Livros</h3>
        <div className="caixa-cabeca-itens">
          <a href="#">
            <i class="fas fa-book">{livroState.length}</i>
          </a>
          <a href="#">
            <i class="fas fa-search"></i>
          </a>
          <a href="#">
            <i class="far fa-plus-square" />
          </a>
        </div>
      </div>
      <div className="caixa-corpo">
        <ul>
          {livroState.map(linha => (
            <li key={linha.livro_id} onClick={() => setCurrent(linha)}>
              <h4>{linha.titulo}</h4>
              <p>{linha.autor}</p>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default Livros;
