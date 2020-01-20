import React from "react";
import Livros from "../livros/Livros";
import InserirLivros from "../livros/InserirLivros";
import Showcase from "../livros/Showcase";
import LivroCurrent from "../livros/LivroCurrent";
import Toolbar from "../livros/Toolbar";

const MainPane = () => {
  return (
    <div>
      <section className="toolbar">
        <Toolbar />
      </section>
      <section className="showcase">
        <Showcase />
      </section>
      <section className="top-container">
        <div className="caixa caixa-lista-livros">
          <Livros />
        </div>
        <div className="caixa caixa-novo-livro">
          <InserirLivros />
          <LivroCurrent />
        </div>
      </section>
    </div>
  );
};

export default MainPane;
