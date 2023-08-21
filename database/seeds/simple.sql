--
-- PostgreSQL database dump
--

-- Dumped from database version 11.2 (Debian 11.2-1.pgdg90+1)
-- Dumped by pg_dump version 11.2 (Debian 11.2-1.pgdg90+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Data for Name: tenants; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tenants (id, prometheus, jaeger) FROM stdin;
139df444-9839-448f-8d7f-98cb865286a1	http://localhost:9090/	http://localhost:16686
32409b1d-0e61-4874-9e25-a2b1dc52447e	http://localhost:19292/	http://localhost:29293
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, password, token) FROM stdin;
8e0016dd-5bec-470c-a26f-5e13b988801a	test2	$2y$10$Rb6YrKoQtpzEoYQiik2RAe/3Irfk0iCpv2CvptvdduaUnpH0mET6y	5fb17a96-3e38-4211-963a-03d43583fe6d
b8501716-4d8a-4bf8-a2e9-25a147bfe105	test1	$2y$10$AnskbYwIgFvI1K0tSrUSX.IKOToMIz1mA.s9W5QOH0k8/7Jg6knGW	acfa769b-f2a9-4fb8-8fd3-9d8337159d46
\.


--
-- Data for Name: tenants_users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tenants_users (tenant_id, user_id) FROM stdin;
139df444-9839-448f-8d7f-98cb865286a1	b8501716-4d8a-4bf8-a2e9-25a147bfe105
\.


--
-- PostgreSQL database dump complete
--

