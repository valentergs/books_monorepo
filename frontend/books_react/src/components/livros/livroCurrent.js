import React, { useContext } from "react";
import LivroContext from "../../context/livros/livroContext";

const LivroCurrent = () => {
  const livroContext = useContext(LivroContext);
  const { current, livroState, editarLivro, clearCurrent } = livroContext;

  // const onSubmit = e => {
  //   e.preventDefault();
  //   editarLivro();
  //   limpaFormulário();
  // };

  // const limpaFormulário = () => {
  //   clearCurrent();
  // };

  return (
    <div className="livro-current">
      <div className="caixa-cabeca">
        <h3 className="caixa-cabeca-titulo">Detalhes da Obra</h3>
        <div className="caixa-cabeca-itens">
          <a
            href="#"
            // onClick={() => {
            //   closeDialogLivroCurrent(false);
            // }}
          >
            <i className="fas fa-times"></i>
          </a>
        </div>
      </div>
      <div className="caixa-corpo">
        {current === null ? (
          <p>Selecione um livro</p>
        ) : (
          <div className="caixa-detalhe-livro">
            <div className="foto">
              <img src={current.photourl} alt={current.isbn} />
            </div>
            <div className="descricao">
              <div className="linha grande">
                <p className="titulo">Titulo</p>
                <p className="texto">{current.titulo}</p>
              </div>
              <div className="linha grande">
                <p className="titulo">Titulo Original</p>
                <p className="texto">{current.titulo_original}</p>
              </div>
              <div className="linha">
                <p className="titulo">Autor</p>
                <p className="texto">{current.autor}</p>
              </div>
              {current.tradutor === "" ? (
                <div className="linha">
                  <p className="titulo">Tradutor</p>
                  <p className="texto">Indisponivel</p>
                </div>
              ) : (
                <div className="linha">
                  <p className="titulo">Tradutor</p>
                  <p className="texto">{current.tradutor}</p>
                </div>
              )}
              <div className="linha">
                <p className="titulo">ISBN</p>
                <p className="texto">{current.isbn}</p>
              </div>
              {current.cdd === "" ? (
                <div className="linha">
                  <p className="titulo">CDD</p>
                  <p className="texto">Indisponivel</p>
                </div>
              ) : (
                <div className="linha">
                  <p className="titulo">CDU</p>
                  <p className="texto">{current.cdd}</p>
                </div>
              )}
              {current.cdu === "" ? (
                <div className="linha">
                  <p className="titulo">CDU</p>
                  <p className="texto">Indisponivel</p>
                </div>
              ) : (
                <div className="linha">
                  <p className="titulo">CDU</p>
                  <p className="texto">{current.cdu}</p>
                </div>
              )}
              <div className="linha">
                <p className="titulo">Ano</p>
                <p className="texto">{current.ano}</p>
              </div>
              <div className="linha grande">
                <p className="titulo">Tema</p>
                <p className="texto">{current.tema}</p>
              </div>
              <div className="linha">
                <p className="titulo">Editora</p>
                <p className="texto">{current.editora}</p>
              </div>
              <div className="linha">
                <p className="titulo">Páginas</p>
                <p className="texto">{current.paginas}</p>
              </div>
              <div className="linha">
                <p className="titulo">Idioma</p>
                <p className="texto">{current.idioma}</p>
              </div>
              <div className="linha">
                <p className="titulo">Formato</p>
                <p className="texto">{current.formato}</p>
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default LivroCurrent;
