--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2
-- Dumped by pg_dump version 15.2

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

--
-- Name: users; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA users;


ALTER SCHEMA users OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: user; Type: TABLE; Schema: users; Owner: postgres
--

CREATE TABLE users."user" (
    user_entity_id integer NOT NULL,
    user_name character varying(36),
    user_email character varying(256),
    user_gender character varying(36)
);


ALTER TABLE users."user" OWNER TO postgres;

--
-- Data for Name: user; Type: TABLE DATA; Schema: users; Owner: postgres
--

INSERT INTO users."user" (user_entity_id, user_name, user_email, user_gender) VALUES (101, 'ica', 'icanisa@kelas.id', 'female');


--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: users; Owner: postgres
--

ALTER TABLE ONLY users."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (user_entity_id);


--
-- PostgreSQL database dump complete
--

