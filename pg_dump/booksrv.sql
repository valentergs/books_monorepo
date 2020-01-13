-- CREATE USER rodrigovalente
-- WITH PASSWORD 'Gustavo2012';

-- CREATE ROLE rodrigovalente
-- WITH
--   LOGIN
--   SUPERUSER
--   INHERIT
--   CREATEDB
--   CREATEROLE
--   NOREPLICATION;

-- CREATE DATABASE booksrv
-- OWNER rodrigovalente;

CREATE TABLE
IF NOT EXISTS livros
(
    livro_id serial PRIMARY KEY,
    titulo VARCHAR
(255) NOT NULL,
    titulo_original VARCHAR
(255) NOT NULL,
    autor VARCHAR
(255) NOT NULL,
    tradutor VARCHAR
(255) NOT NULL,
    isbn VARCHAR
(13) NOT NULL UNIQUE,
    cdd VARCHAR
(50) NOT NULL,
    cdu VARCHAR
(50) NOT NULL,
    ano VARCHAR
(4) NOT NULL,
    tema VARCHAR
(255) NOT NULL,
    editora VARCHAR
(100) NOT NULL,
    paginas VARCHAR
(8) NOT NULL,
    idioma VARCHAR
(50) NOT NULL,
    formato VARCHAR
(50) NOT NULL,
    dono VARCHAR
(50) NOT NULL
);


INSERT INTO livros
VALUES
    (default, 'O Profeta', 'The prophet', 'Khalil Gibran', 'Ana Guadalupe', '9788542215809', '892.73', '', '2019', 'Ficcão libanesa', 'Planeta do Brasil', '144', 'Português', 'Impresso', 'rodrigo_valente');

INSERT INTO livros
VALUES
    (default, 'Não me faça pensar: atualizado: uma abordagem de bom senso à usabilidde web e mobile', 'Dont make me think, Revisted: A common sense approach to web usability', 'Steve Krug', 'Daniel Croce', '9788576088509', '006.7', '004.738.5', '2014', 'Sites da web - Desenvolvimento; Sites da web - Projetos; Usabilidade', 'Alta Books', '212', 'Português', 'Impresso', 'rodrigo_valente');

INSERT INTO livros
VALUES
    (default, 'Medica dos Horrores', 'The Butchering Art: The Joseph Listers Quest to transform the Grisly World of Victorian Medicine', 'Lindsey Fitzharris', 'Vera Ribeiro', '9788551005224', '926.1', '929.616-089.8', '2017', 'Joseph Lister, Cirurgiões, Grã-Bretanha, Biografia', 'Intrinseca', '320', 'Português', 'Impresso', 'rodrigo_valente');    

