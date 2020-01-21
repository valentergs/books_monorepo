import React, { useReducer } from "react";
import axios from "axios";
import LivroContext from "./livroContext";
import LivroReducer from "./livroReducer";
import {
  TODOS_LIVROS,
  INSERIR_LIVRO,
  DELETAR_LIVRO,
  EDITAR_LIVRO,
  CLEAR_CURRENT,
  SET_CURRENT,
  SET_ADD_BOOK,
  CLOSE_ADD_BOOK
} from "../types";

const LivroState = props => {
  const initialState = {
    livroState: [],
    current: null,
    addBook: false
  };

  const [state, dispatch] = useReducer(LivroReducer, initialState);

  const todosLivros = async () => {
    const res = await axios.get("http://localhost:8080");
    dispatch({
      type: TODOS_LIVROS,
      payload: res.data
    });
  };

  const inserirLivro = async livro => {
    const res = await axios.post("http://localhost:8080/inserir", livro);
    dispatch({
      type: INSERIR_LIVRO,
      payload: res.data
    });
  };

  const deletarLivro = livro_id => {
    axios.delete(`http://localhost:8080/deletar/${livro_id}`);
    dispatch({
      type: DELETAR_LIVRO,
      payload: livro_id
    });
  };

  const editarLivro = async livro => {
    const res = await axios.put(
      `http://localhost:8080/editar/${livro.livro_id}`,
      livro
    );
    dispatch({
      type: EDITAR_LIVRO,
      payload: res.data
    });
  };

  const setCurrent = livro => {
    dispatch({ type: SET_CURRENT, payload: livro });
  };

  const clearCurrent = () => {
    dispatch({ type: CLEAR_CURRENT });
  };

  const setAddBook = () => {
    dispatch({ type: SET_ADD_BOOK });
  };

  const closeAddBook = () => {
    dispatch({ type: CLOSE_ADD_BOOK });
  };

  return (
    <LivroContext.Provider
      value={{
        livroState: state.livroState,
        current: state.current,
        addBook: state.addBook,
        todosLivros,
        inserirLivro,
        deletarLivro,
        editarLivro,
        setCurrent,
        clearCurrent,
        setAddBook,
        closeAddBook
      }}
    >
      {props.children}
    </LivroContext.Provider>
  );
};

export default LivroState;
