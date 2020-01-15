import React, { useContext, useEffect } from "react";
import Dialog from "@material-ui/core/Dialog";
import useMediaQuery from "@material-ui/core/useMediaQuery";
import { useTheme } from "@material-ui/core/styles";

import LivroCurrent from "./LivroCurrent";

import LivroContext from "../../context/livros/livroContext";

const Showcase = () => {
  const [open, setOpen] = React.useState(false);
  const theme = useTheme();
  const fullScreen = useMediaQuery(theme.breakpoints.down("sm"));
  const livroContext = useContext(LivroContext);
  const { livroState, setCurrent } = livroContext;

  const handleClose = () => {
    setOpen(false);
  };

  return (
    <div>
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
        <p>{livroState.length}</p>
      </div>

      <ul>
        {livroState.map(livro => (
          <li
            key={livro.livro_id}
            onClick={() => {
              setCurrent(livro);
              setOpen(true);
            }}
          >
            <img src={livro.photourl} alt={livro.isbn} />
          </li>
        ))}
      </ul>
      <Dialog
        fullScreen={fullScreen}
        open={open}
        onClose={handleClose}
        aria-labelledby="responsive-dialog-title"
      >
        <LivroCurrent />
        {/* <DialogTitle id="responsive-dialog-title">
          {"Use Google's location service?"}
        </DialogTitle>
        <DialogContent>
          <DialogContentText>
            Let Google help apps determine location. This means sending
            anonymous location data to Google, even when no apps are running.
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button autoFocus onClick={handleClose} color="primary">
            Disagree
          </Button>
          <Button onClick={handleClose} color="primary" autoFocus>
            Agree
          </Button>
        </DialogActions> */}
      </Dialog>
    </div>
  );
};

export default Showcase;
