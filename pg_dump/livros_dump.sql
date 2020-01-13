--
-- PostgreSQL database dump
--

-- Dumped from database version 11.5
-- Dumped by pg_dump version 11.5

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: livros; Type: TABLE; Schema: public; Owner: rodrigovalente
--

CREATE TABLE public.livros (
    livro_id integer NOT NULL,
    titulo character varying(255) NOT NULL,
    titulo_original character varying(255) NOT NULL,
    autor character varying(255) NOT NULL,
    tradutor character varying(255) NOT NULL,
    isbn character varying(13) NOT NULL,
    cdd character varying(50) NOT NULL,
    cdu character varying(50) NOT NULL,
    ano character varying(4) NOT NULL,
    tema character varying(255) NOT NULL,
    editora character varying(100) NOT NULL,
    paginas character varying(8) NOT NULL,
    idioma character varying(50) NOT NULL,
    formato character varying(50) NOT NULL,
    dono character varying(50) NOT NULL,
    photourl character varying(256)
);


ALTER TABLE public.livros OWNER TO rodrigovalente;

--
-- Name: livros_livro_id_seq; Type: SEQUENCE; Schema: public; Owner: rodrigovalente
--

CREATE SEQUENCE public.livros_livro_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.livros_livro_id_seq OWNER TO rodrigovalente;

--
-- Name: livros_livro_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rodrigovalente
--

ALTER SEQUENCE public.livros_livro_id_seq OWNED BY public.livros.livro_id;


--
-- Name: livros livro_id; Type: DEFAULT; Schema: public; Owner: rodrigovalente
--

ALTER TABLE ONLY public.livros ALTER COLUMN livro_id SET DEFAULT nextval('public.livros_livro_id_seq'::regclass);


--
-- Data for Name: livros; Type: TABLE DATA; Schema: public; Owner: rodrigovalente
--

COPY public.livros (livro_id, titulo, titulo_original, autor, tradutor, isbn, cdd, cdu, ano, tema, editora, paginas, idioma, formato, dono, photourl) FROM stdin;
21	50 ideias de química que você precisa conhecer	50 Chemistry Ideas You Really Need to Know	Hayley Birch	Helena Londres	9788542213621	540		2015	Quimica, Curioisidades e maravilhas	Planeta	216	Português	Impresso	rodrigo_valente	https://m.media-amazon.com/images/I/71OQqdMsZAL._AC_UY218_ML3_.jpg
22	Data Science para Negócios	Data Science for Business	Foster Provost	Marina Boscato	9788576089728	658.4032	658:004.6	2016	Data Science, Big Data, Mineração de dados, Modelagem de Sistemas, Estratégia, Negócios	Alta Books	408	Português	Impresso	rodrigo_valente	https://m.media-amazon.com/images/I/61qn0sroQFL._AC_UY218_ML3_.jpg
3	Medicina dos Horrores	The Butchering Art: The Joseph Listers Quest to transform the Grisly World of Victorian Medicine	Lindsey Fitzharris	Vera Ribeiro	9788551005224	926.1	929.616-089.8	2017	Joseph Lister, Cirurgiões, Grã-Bretanha, Biografia	Intrinseca	320	Português	Impresso	rodrigo_valente	https://fotoscapadelivros.s3-sa-east-1.amazonaws.com/9788551005224_p.jpg
4	Rei Arthur e os cavaleiros da Távola Redonda	The story of King Arthur and his knights	Howard Pyle	Vivien Kogut Lessa de Sá	9788537810651	823	821.111-3	2013	Arthur, Rei - Ficção, Reis e governantes, Grã-Bretanha - Ficção, Ficção americana	Zahar	324	Português	Impresso	rodrigo_valente	https://fotoscapadelivros.s3-sa-east-1.amazonaws.com/9788537810651_p.jpg
6	Os sofrimentos do jovem Werther	Die Leiden das jungen Werther	Johann Wolfgang von Goethe	Claudia Cavalcanti	9788544000137	833		2014	Romance Alemão	Martin Claret	196	Português	Impresso	rodrigo_valente	https://fotoscapadelivros.s3-sa-east-1.amazonaws.com/9788544000137_p.jpg
2	Não me faça pensar: atualizado: uma abordagem de bom senso à usabilidade web e mobile	Dont make me think, Revisited: A common sense approach to web usability	Steve Krug	Daniel Croce	9788576088509	006.7	004.738.5	2014	Sites da web - Desenvolvimento; Sites da web - Projetos; Usabilidade	Alta Books	212	Português	Impresso	rodrigo_valente	https://fotoscapadelivros.s3-sa-east-1.amazonaws.com/9788576088509_p.jpg
17	Uma Breve História do Tempo	A Brief History of Time	Stephen Hawking	Cassio de Arantes Leite	9788580576467	523.1	524	1988	Física, Espaço e tempo, Cosmologia	Intrínseca	256	Português	Impresso	rodrigo_valente	https://fotoscapadelivros.s3-sa-east-1.amazonaws.com/9788580576467_p.jpg
24	Oscar Wilde para inquietos	El Coaching de Oscar Wilde	Allan Percy	Joana Angélica d'Avila Melo	9788575427705	828	821.111-8	2011	Wilde, Oscar, Aformismos e apotegmas	GMT Editores	128	Português	Impresso	rodrigo_valente	https://m.media-amazon.com/images/I/711pof-JIgL._AC_UY218_ML3_.jpg
13	Factfulness: o hábito libertador de só ter opiniões baseadas em fatos	Factfulness: ten reasons we're wrong about the world	Hans Rosling	Vitor Paolozzi	9788501116529	303.4	316.42	2015	Estatistica, Desenvolvimento Econômico, Aspectos Sociais, Comportamento Humano	Record	359	Português	Impresso	rodrigo_valente	https://fotoscapadelivros.s3-sa-east-1.amazonaws.com/9788501116529.jpg
19	1932: São Paulo em Chamas	1932: São Paulo em Chamas	Luis Octavio de Lima		9788542211269	006.7	004.738.5	2014	Psicologia	Intrinseca	212	Português	Impresso	rodrigo_valente	https://m.media-amazon.com/images/I/81BcnC+VFYL._AC_UY218_ML3_.jpg
25	A Quietude é a chave	Stillness is Key	Ryan Holiday	Maria Luiza Borges	9788551005767	158.1	159.923.2	2019	Motivação (Psicologia), Autoconsciência, Quietude, Paz de Espirito	Intrinseca	288	Português	Impresso	rodrigo_valente	https://m.media-amazon.com/images/I/6115K3wdOwL._AC_UY218_ML3_.jpg
23	O Universo numa casca de noz	The Universe in a Nutshell	Stephen Hawking	Cassio de Arantes Leite	9788580578881	530.12	530.145	2001	Fisica Quântica	Intrinseca	224	Português	Impresso	rodrigo_valente	https://images-na.ssl-images-amazon.com/images/I/81ayVgQt9OL.jpg
27	Origem	Origin	Dan Brown	Alves Calado	9788580417661	813	821.111(73)-3	2017	Ficção Americana	Arqueiro	432	Português	Impresso	rodrigo_valente	https://images-na.ssl-images-amazon.com/images/I/81ehB5iQsPL.jpg
14	Origens: uma grande história de tudo	Origin Story: A Big History of Everything	David Christian	Pedro Maia Soares	9788535932614	909		2018	Civilização - Filosofia, Evolução humana, História do mundo	Schwarcz	406	Português	Impresso	rodrigo_valente	https://fotoscapadelivros.s3-sa-east-1.amazonaws.com/9788535932614_p.jpg
15	Algoritmos para viver: a ciência exata para das decisões humanas	Agorithms to Live: The Computer Science of Human Decisions	Brian Christian	Paulo Geiger	9788535929300	150		2016	Algoritmos de computadores, Decisões, Comportamento humano, Matemática, Aspectos psicológicos, Resolução de problemas, Tomada de decisões	Schwarcz	529	Português	Impresso	rodrigo_valente	https://fotoscapadelivros.s3-sa-east-1.amazonaws.com/9788535929300_p.jpg
16	A Arte da guerra	Sun zu bing fa	Sun Tzu	André da Silva Bueno	9788560018000	355		2010	Arte e ciência militar	Jardim dos livros	125	Português	Impresso	rodrigo_valente	https://fotoscapadelivros.s3-sa-east-1.amazonaws.com/9788560018000_p.jpg
1	O Profeta	The prophet	Khalil Gibran	Ana Guadalupe	9788542215809	892.73		2019	Ficção libanesa	Planeta	144	Português	Impresso	rodrigo_valente	https://fotoscapadelivros.s3-sa-east-1.amazonaws.com/9788542215809_p.jpg
20	Sapiens - Uma Breve História da Humanidade	Sapiens - A Brief History of Humankind	Yuval Noah Hariri	Janaina Marcoantonio	9788525432186	576.8	575.8	2015	Evolução humana	L&PM	464	Português	Impresso	rodrigo_valente	https://m.media-amazon.com/images/I/81BTkpMrLYL._AC_UY218_ML3_.jpg
28	Crime e Castigo	Prestuplenie i Nakazanie	Fiódor Dostoiévski	Paulo Bezerra	9788573262087	891.73		2001	Ficcção russa	Editora 34	568	Português	Impresso	rodrigo_valente	https://images-na.ssl-images-amazon.com/images/I/81C1c9Ec03L.jpg
\.


--
-- Name: livros_livro_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rodrigovalente
--

SELECT pg_catalog.setval('public.livros_livro_id_seq', 28, true);


--
-- Name: livros livros_isbn_key; Type: CONSTRAINT; Schema: public; Owner: rodrigovalente
--

ALTER TABLE ONLY public.livros
    ADD CONSTRAINT livros_isbn_key UNIQUE (isbn);


--
-- Name: livros livros_pkey; Type: CONSTRAINT; Schema: public; Owner: rodrigovalente
--

ALTER TABLE ONLY public.livros
    ADD CONSTRAINT livros_pkey PRIMARY KEY (livro_id);


--
-- PostgreSQL database dump complete
--

