import React, { useEffect, useContext, useState } from "react";
import LivroContext from "../../context/livros/livroContext";

const InserirLivros = () => {
  const livroContext = useContext(LivroContext);
  const { current, inserirLivro, clearCurrent, closeAddBook } = livroContext;

  const [livro, setLivro] = useState({
    titulo: "",
    titulo_original: "",
    autor: "",
    tradutor: "",
    isbn: "",
    cdd: "",
    cdu: "",
    ano: "",
    tema: "",
    editora: "",
    paginas: "",
    idioma: "",
    formato: "",
    dono: ""
  });

  const {
    titulo,
    titulo_original,
    autor,
    tradutor,
    isbn,
    cdd,
    cdu,
    ano,
    tema,
    editora,
    paginas,
    idioma,
    formato,
    dono
  } = livro;

  useEffect(() => {
    setLivro({
      titulo: "",
      titulo_original: "",
      autor: "",
      tradutor: "",
      isbn: "",
      cdd: "",
      cdu: "",
      ano: "",
      tema: "",
      editora: "",
      paginas: "",
      idioma: "",
      formato: "",
      dono: ""
    });
  }, [livroContext, current]);

  const onChange = e => {
    const target = e.target;
    const name = target.name;
    const value = target.value;
    setLivro({ ...livro, [name]: value });
  };

  const onSubmit = e => {
    e.preventDefault();
    inserirLivro(livro);
    limpaFormulário();
    closeAddBook();
  };

  const limpaFormulário = () => {
    clearCurrent();
  };

  return (
    <div>
      <form className="form-livro-inserir" onSubmit={onSubmit} method="post">
        <input
          className="livro-input grande"
          type="text"
          placeholder="Titulo"
          name="titulo"
          value={titulo}
          onChange={onChange}
        />
        <input
          className="livro-input grande"
          type="text"
          placeholder="Titulo Original"
          name="titulo_original"
          value={titulo_original}
          onChange={onChange}
        />
        <input
          className="livro-input"
          type="text"
          placeholder="Autor"
          name="autor"
          value={autor}
          onChange={onChange}
        />
        <input
          className="livro-input"
          type="text"
          placeholder="Tradutor"
          name="tradutor"
          value={tradutor}
          onChange={onChange}
        />
        <input
          className="livro-input"
          type="text"
          placeholder="ISBN"
          name="isbn"
          value={isbn}
          onChange={onChange}
        />
        <input
          className="livro-input"
          type="text"
          placeholder="CDD"
          name="cdd"
          value={cdd}
          onChange={onChange}
        />
        <input
          className="livro-input"
          type="text"
          placeholder="CDU"
          name="cdu"
          value={cdu}
          onChange={onChange}
        />
        <input
          className="livro-input"
          type="text"
          placeholder="Ano"
          name="ano"
          value={ano}
          onChange={onChange}
        />
        <input
          className="livro-input grande"
          type="text"
          placeholder="Tema"
          name="tema"
          value={tema}
          onChange={onChange}
        />
        <input
          className="livro-input"
          type="text"
          placeholder="Editora"
          name="editora"
          value={editora}
          onChange={onChange}
        />
        <input
          className="livro-input"
          type="text"
          placeholder="Páginas"
          name="paginas"
          value={paginas}
          onChange={onChange}
        />
        <input
          className="livro-input"
          type="text"
          placeholder="Idioma"
          name="idioma"
          value={idioma}
          onChange={onChange}
        />
        <input
          className="livro-input"
          type="text"
          placeholder="Formato"
          name="formato"
          value={formato}
          onChange={onChange}
        />
        <input
          className="livro-input"
          type="text"
          placeholder="Dono"
          name="dono"
          value={dono}
          onChange={onChange}
          required
        />
        <input
          className="btn grande"
          type="button"
          value="Salvar"
          onClick={onSubmit}
        />
      </form>
    </div>
  );
};

export default InserirLivros;

// import React, { useEffect, useContext, useState } from "react";
// import {
//   TextField,
//   Button,
//   FormControl,
//   Grid,
//   Container
// } from "@material-ui/core";
// import { makeStyles } from "@material-ui/core/styles";

// //Context
// import LivroContext from "../../context/livros/livroContext";

// const useStyles = makeStyles(theme => ({
//   root: {
//     flexGrow: 1,
//     display: "flex",
//     flexWrap: "wrap"
//   },
//   TextField: {
//     flexGrow: 1
//   }
// }));

// const InserirLivros = () => {
//   const livroContext = useContext(LivroContext);
//   const { current, inserirLivro, clearCurrent } = livroContext;

//   const [livro, setLivro] = useState({
//     titulo: "",
//     titulo_original: "",
//     autor: "",
//     tradutor: "",
//     isbn: "",
//     cdd: "",
//     cdu: "",
//     ano: "",
//     tema: "",
//     editora: "",
//     paginas: "",
//     idioma: "",
//     formato: "",
//     dono: ""
//   });

//   const {
//     titulo,
//     titulo_original,
//     autor,
//     tradutor,
//     isbn,
//     cdd,
//     cdu,
//     ano,
//     tema,
//     editora,
//     paginas,
//     idioma,
//     formato,
//     dono
//   } = livro;

//   useEffect(() => {
//     setLivro({
//       titulo: "",
//       titulo_original: "",
//       autor: "",
//       tradutor: "",
//       isbn: "",
//       cdd: "",
//       cdu: "",
//       ano: "",
//       tema: "",
//       editora: "",
//       paginas: "",
//       idioma: "",
//       formato: "",
//       dono: ""
//     });
//   }, [livroContext, current]);

//   const onChange = e => {
//     const target = e.target;
//     const name = target.name;
//     const value = target.value;
//     setLivro({ ...livro, [name]: value });
//   };

//   const onSubmit = e => {
//     e.preventDefault();
//     inserirLivro(livro);
//     limpaFormulário();
//   };

//   const limpaFormulário = () => {
//     clearCurrent();
//   };

//   const classes = useStyles();

//   //   const idiomas = [
//   //     {
//   //       value: "Português",
//   //       label: "Português"
//   //     },
//   //     {
//   //       value: "Inglês",
//   //       label: "Inglês"
//   //     },
//   //     {
//   //       value: "Espanhol",
//   //       label: "Espanhol"
//   //     }
//   //   ];

//   //   const formatos = [
//   //     {
//   //       value: "Impresso",
//   //       label: "Impresso"
//   //     },
//   //     {
//   //       value: "Digital",
//   //       label: "Digital"
//   //     }
//   //   ];

//   return (
//     <div className={classes.root}>
//       <Container>
//         <FormControl
//           onSubmit={onSubmit}
//           method="post"
//           className={classes.formControl}
//         >
//           <Grid container spacing={1}>
//             <Grid item xs={12} sm={6}>
//               <TextField
//                 className={classes.Input}
//                 label="Titulo"
//                 name="titulo"
//                 value={titulo}
//                 onChange={onChange}
//               />
//             </Grid>
//             <Grid item xs={12} sm={6}>
//               <TextField
//                 className={classes.TextField}
//                 label="Titulo Original"
//                 name="titulo_original"
//                 value={titulo_original}
//                 onChange={onChange}
//               />
//             </Grid>
//             <Grid item xs={12} sm={6}>
//               <TextField
//                 id="standard-helperText"
//                 label="Autor"
//                 name="autor"
//                 value={autor}
//                 onChange={onChange}
//               />
//             </Grid>
//             <Grid item xs={12} sm={6}>
//               <TextField
//                 id="standard-helperText"
//                 label="Tradutor"
//                 name="tradutor"
//                 value={tradutor}
//                 onChange={onChange}
//               />
//             </Grid>
//             <Grid item xs={6} sm={3}>
//               <TextField
//                 id="standard-helperText"
//                 label="ISBN"
//                 name="isbn"
//                 value={isbn}
//                 onChange={onChange}
//               />
//             </Grid>
//             <Grid item xs={6} sm={3}>
//               <TextField
//                 id="standard-helperText"
//                 label="CDD"
//                 name="cdd"
//                 value={cdd}
//                 onChange={onChange}
//               />
//             </Grid>
//             <Grid item xs={6} sm={3}>
//               <TextField
//                 id="standard-helperText"
//                 label="CDU"
//                 name="cdu"
//                 value={cdu}
//                 s
//                 onChange={onChange}
//               />
//             </Grid>
//             <Grid item xs={6} sm={3}>
//               <TextField
//                 id="standard-helperText"
//                 label="Ano"
//                 name="ano"
//                 value={ano}
//                 onChange={onChange}
//               />
//             </Grid>
//             <Grid item xs={12} sm={6}>
//               <TextField
//                 id="standard-helperText"
//                 label="Tema"
//                 name="tema"
//                 value={tema}
//                 onChange={onChange}
//               />
//             </Grid>
//             <Grid item xs={12} sm={6}>
//               <TextField
//                 id="standard-helperText"
//                 label="Editora"
//                 name="editora"
//                 value={editora}
//                 onChange={onChange}
//               />
//             </Grid>
//             <Grid item xs={6} sm={3}>
//               <TextField
//                 id="standard-helperText"
//                 label="Paginas"
//                 name="paginas"
//                 value={paginas}
//                 onChange={onChange}
//               />
//             </Grid>
//             <Grid item xs={6} sm={3}>
//               <TextField
//                 id="standard-helperText"
//                 label="Idioma"
//                 name="idioma"
//                 value={idioma}
//                 onChange={onChange}
//               />
//             </Grid>
//             <Grid item xs={6} sm={3}>
//               <TextField
//                 id="standard-helperText"
//                 label="Formato"
//                 value={formato}
//                 name="formato"
//                 onChange={onChange}
//               />
//             </Grid>
//             <Grid item xs={6} sm={3}>
//               <TextField
//                 id="standard-helperText"
//                 label="Dono"
//                 value={dono}
//                 name="dono"
//                 onChange={onChange}
//               />
//             </Grid>
//             <Grid item xs={12}>
//               <Button
//                 variant="contained"
//                 color="primary"
//                 style={{ marginTop: 20, marginBottom: 20 }}
//                 onClick={onSubmit}
//               >
//                 Salvar
//               </Button>
//             </Grid>
//           </Grid>
//         </FormControl>
//       </Container>
//     </div>
//   );
// };

// export default InserirLivros;
