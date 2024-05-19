--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3
-- Dumped by pg_dump version 16.3

-- Started on 2024-05-19 20:13:23

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
-- TOC entry 5 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- TOC entry 4889 (class 0 OID 0)
-- Dependencies: 5
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


--
-- TOC entry 220 (class 1255 OID 16408)
-- Name: set_updated_at(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.set_updated_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$;


ALTER FUNCTION public.set_updated_at() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 218 (class 1259 OID 16426)
-- Name: premium_purchase; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.premium_purchase (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid,
    premium_type character varying,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.premium_purchase OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 16410)
-- Name: swipes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.swipes (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid,
    target_id uuid,
    is_like boolean,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.swipes OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 16418)
-- Name: swipes_count; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.swipes_count (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid,
    count integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.swipes_count OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 16451)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    email character varying,
    password character varying,
    full_name character varying,
    age character varying,
    gender character varying,
    is_verified boolean DEFAULT false,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    no_swipe_limit boolean DEFAULT false
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 4882 (class 0 OID 16426)
-- Dependencies: 218
-- Data for Name: premium_purchase; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.premium_purchase (id, user_id, premium_type, created_at, updated_at) FROM stdin;
\.


--
-- TOC entry 4880 (class 0 OID 16410)
-- Dependencies: 216
-- Data for Name: swipes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.swipes (id, user_id, target_id, is_like, created_at, updated_at) FROM stdin;
\.


--
-- TOC entry 4881 (class 0 OID 16418)
-- Dependencies: 217
-- Data for Name: swipes_count; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.swipes_count (id, user_id, count, created_at, updated_at) FROM stdin;
\.


--
-- TOC entry 4883 (class 0 OID 16451)
-- Dependencies: 219
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, email, password, full_name, age, gender, is_verified, created_at, updated_at, no_swipe_limit) FROM stdin;
\.


--
-- TOC entry 4730 (class 2606 OID 16434)
-- Name: premium_purchase premium_purchase_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.premium_purchase
    ADD CONSTRAINT premium_purchase_pkey PRIMARY KEY (id);


--
-- TOC entry 4728 (class 2606 OID 16424)
-- Name: swipes_count swipes_count_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.swipes_count
    ADD CONSTRAINT swipes_count_pkey PRIMARY KEY (id);


--
-- TOC entry 4726 (class 2606 OID 16416)
-- Name: swipes swipes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.swipes
    ADD CONSTRAINT swipes_pkey PRIMARY KEY (id);


--
-- TOC entry 4732 (class 2606 OID 16460)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 4735 (class 2620 OID 16435)
-- Name: premium_purchase set_premium_purchase_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER set_premium_purchase_updated_at BEFORE UPDATE ON public.premium_purchase FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();


--
-- TOC entry 4734 (class 2620 OID 16425)
-- Name: swipes_count set_swipes_count_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER set_swipes_count_updated_at BEFORE UPDATE ON public.swipes_count FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();


--
-- TOC entry 4733 (class 2620 OID 16417)
-- Name: swipes set_swipes_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER set_swipes_updated_at BEFORE UPDATE ON public.swipes FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();


--
-- TOC entry 4736 (class 2620 OID 16461)
-- Name: users set_users_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER set_users_updated_at BEFORE UPDATE ON public.users FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();


-- Completed on 2024-05-19 20:13:23

--
-- PostgreSQL database dump complete
--

