import React from "react";

const Toolbar = () => {
  return (
    <div className="toolbar">
      <input type="text" placeholder="Buscar" />
      <p>Exibir: </p>
      <p>
        <input
          type="radio"
          name="showcase_display"
          value="capa"
          checked
        ></input>
        <label htmlFor="capa">Capa</label>
      </p>
      <p>
        <input type="radio" name="showcase_display" value="titulo"></input>
        <label htmlFor="titulo">Titulo</label>
      </p>
    </div>
  );
};

export default Toolbar;
