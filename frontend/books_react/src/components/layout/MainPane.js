import React from "react";
import Livros from "../livros/Livros";
import InserirLivros from "../livros/InserirLivros";
import Showcase from "../livros/Showcase";

const MainPane = () => {
  return (
    <div>
      <section className="showcase">
        <Showcase />
      </section>
      <section className="top-container">
        <div className="caixa caixa-lista-livros">
          <Livros />
        </div>
        <div className="caixa caixa-novo-livro">
          <InserirLivros />
        </div>
      </section>
    </div>
  );
};

export default MainPane;
