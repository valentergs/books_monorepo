import React from "react";
import Livros from "../livros/Livros";
import InserirLivros from "../livros/InserirLivros";
import LivroCurrent from "../livros/livroCurrent";

const MainPane = () => {
  return (
    <div>
      <section className="top-container">
        <div className="caixa caixa-lista-livros">
          <Livros />
        </div>
        <div className="caixa caixa-detalhe-livro">
          <LivroCurrent />
        </div>
        <div className="caixa caixa-novo-livro">
          <InserirLivros />
        </div>
      </section>
    </div>
  );
};

export default MainPane;

// import React from "react";
// import Livros from "../livros/Livros";
// import InserirLivros from "../livros/InserirLivros";
// import LivroCurrent from "../livros/livroCurrent";
// import { Grid, Container } from "@material-ui/core";
// import { makeStyles } from "@material-ui/core/styles";

// const useStyles = makeStyles(theme => ({
//   root: {
//     "& .MuiTextField-root": {
//       margin: theme.spacing(1)
//       //width: 200
//     }
//   }
// }));

// const MainPaneRV = () => {
//   const classes = useStyles();

//   return (
//     <div>
//       <Container spacing={2} className={classes.root}>
//         <Grid container spacing={2}>
//           <Grid item xs={12} sm={6}>
//             <LivroCurrent />
//             <InserirLivros />
//           </Grid>
//           <Grid item xs={12} sm={6}>
//             <Livros />
//           </Grid>
//         </Grid>
//       </Container>
//     </div>
//   );
// };

// export default MainPaneRV;
