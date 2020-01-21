import React, { useContext } from "react";
import LivroContext from "../../context/livros/livroContext";

const Showcase = () => {
  const livroContext = useContext(LivroContext);
  const { livroState, setCurrent } = livroContext;

  return (
    <div className="showcase">
      <ul>
        {livroState.map(livro => (
          <li
            key={livro.livro_id}
            onClick={() => {
              setCurrent(livro);
            }}
          >
            <img src={livro.photourl} alt={livro.isbn} />
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Showcase;
