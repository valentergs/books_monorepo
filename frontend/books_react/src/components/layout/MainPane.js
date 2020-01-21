import React, { Fragment } from "react";
import Livros from "../livros/Livros";
import Showcase from "../livros/Showcase";
import LivroCurrent from "../livros/LivroCurrent";
import Toolbar from "../livros/Toolbar";

const MainPane = () => {
  return (
    <div>
      <Toolbar />
      <section>
        <Showcase />
      </section>
      <section className="top-container">
        <div className="caixa caixa-lista-livros">
          <Livros />
        </div>
        <div className="caixa caixa-novo-livro">
          <LivroCurrent />
        </div>
      </section>
    </div>
  );
};

export default MainPane;
