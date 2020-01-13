import React from "react";

//Components
import Navbar from "./components/layout/Navbar";
import MainPane from "./components/layout/MainPane";

//Context
import LivroState from "./context/livros/livroState";

const App = () => {
  return (
    <LivroState>
      <div className="App container">
        <Navbar />
        <MainPane />
      </div>
    </LivroState>
  );
};

export default App;
