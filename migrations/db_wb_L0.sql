--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2
-- Dumped by pg_dump version 15.2

-- Started on 2023-11-18 14:16:20

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

SET default_table_access_method = heap;

--
-- TOC entry 215 (class 1259 OID 41168)
-- Name: delivery; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.delivery (
    order_uid character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    phone character varying(255),
    zip character varying(255),
    city character varying(255),
    address character varying(255),
    region character varying(255),
    email character varying(255)
);


ALTER TABLE public.delivery OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 41192)
-- Name: items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.items (
    order_uid character varying(255) NOT NULL,
    chrt_id bigint NOT NULL,
    track_number character varying(255) NOT NULL,
    price bigint,
    rid character varying(255),
    name character varying(255),
    sale smallint,
    size character varying(255),
    total_price bigint,
    nm_id bigint,
    brand character varying(255),
    status smallint
);


ALTER TABLE public.items OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 41161)
-- Name: orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders (
    order_uid character varying(255) NOT NULL,
    track_number character varying(255) NOT NULL,
    entry character varying(255),
    locale character varying(255),
    internal_signature character varying(255),
    customer_id character varying(255),
    delivery_service character varying(255),
    shardkey character varying(255),
    sm_id bigint,
    date_created character varying(255),
    oof_shred character varying(255)
);


ALTER TABLE public.orders OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 41180)
-- Name: payment; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payment (
    order_uid character varying(255) NOT NULL,
    transaction character varying(255) NOT NULL,
    request_id character varying(255),
    currency character varying(255),
    provider character varying(255),
    amount bigint,
    payment_dt bigint,
    bank character varying(255),
    delivery_cost bigint,
    goods_total bigint,
    custom_fee smallint
);


ALTER TABLE public.payment OWNER TO postgres;

--
-- TOC entry 3338 (class 0 OID 41168)
-- Dependencies: 215
-- Data for Name: delivery; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.delivery (order_uid, name, phone, zip, city, address, region, email) VALUES ('1', 'Test Testov', '+9720000000', '2639809', 'Kiryat Mozkin', 'Ploshad Mira 11', 'Kraiot', 'test@gmail.com');
INSERT INTO public.delivery (order_uid, name, phone, zip, city, address, region, email) VALUES ('2', 'Test Testov', '+9720000000', '2639809', 'Kiryat Mozkin', 'Ploshad Mira 12', 'Kraiot', 'test@gmail.com');
INSERT INTO public.delivery (order_uid, name, phone, zip, city, address, region, email) VALUES ('3', 'Test Testov', '+9720000000', '2639809', 'Kiryat Mozkin', 'Ploshad Mira 13', 'Kraiot', 'test@gmail.com');
INSERT INTO public.delivery (order_uid, name, phone, zip, city, address, region, email) VALUES ('4', 'Test Testov', '+9720000000', '2639809', 'Kiryat Mozkin', 'Ploshad Mira 14', 'Kraiot', 'test@gmail.com');
INSERT INTO public.delivery (order_uid, name, phone, zip, city, address, region, email) VALUES ('5', 'Test Testov', '+9720000000', '2639809', 'Kiryat Mozkin', 'Ploshad Mira 15', 'Kraiot', 'test@gmail.com');


--
-- TOC entry 3340 (class 0 OID 41192)
-- Dependencies: 217
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES ('1', 9934930, 'WBILMTESTTRACK', 1001, 'ab4219087a764ae0btest', 'Mascaras', 30, '0', 317, 2389212, 'Vivienne Sabo', 202);
INSERT INTO public.items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES ('2', 9934930, 'WBILMTESTTRACK', 1002, 'ab4219087a764ae0btest', 'Mascaras', 30, '0', 317, 2389212, 'Vivienne Sabo', 202);
INSERT INTO public.items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES ('3', 9934930, 'WBILMTESTTRACK', 1003, 'ab4219087a764ae0btest', 'Mascaras', 30, '0', 317, 2389212, 'Vivienne Sabo', 202);
INSERT INTO public.items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES ('4', 9934930, 'WBILMTESTTRACK', 1004, 'ab4219087a764ae0btest', 'Mascaras', 30, '0', 317, 2389212, 'Vivienne Sabo', 202);
INSERT INTO public.items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES ('5', 9934930, 'WBILMTESTTRACK', 1005, 'ab4219087a764ae0btest', 'Mascaras', 30, '0', 317, 2389212, 'Vivienne Sabo', 202);


--
-- TOC entry 3337 (class 0 OID 41161)
-- Dependencies: 214
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shred) VALUES ('1', 'WBILMTESTTRACK', 'WBIL', 'en', '', 'test', 'meest', '9', 99, '2021-11-26T06:22:19Z', '1');
INSERT INTO public.orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shred) VALUES ('2', 'WBILMTESTTRACK', 'WBIL', 'en', '', 'test', 'meest', '9', 99, '2021-11-26T06:22:19Z', '1');
INSERT INTO public.orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shred) VALUES ('3', 'WBILMTESTTRACK', 'WBIL', 'en', '', 'test', 'meest', '9', 99, '2021-11-26T06:22:19Z', '1');
INSERT INTO public.orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shred) VALUES ('4', 'WBILMTESTTRACK', 'WBIL', 'en', '', 'test', 'meest', '9', 99, '2021-11-26T06:22:19Z', '1');
INSERT INTO public.orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shred) VALUES ('5', 'WBILMTESTTRACK', 'WBIL', 'en', '', 'test', 'meest', '9', 99, '2021-11-26T06:22:19Z', '1');


--
-- TOC entry 3339 (class 0 OID 41180)
-- Dependencies: 216
-- Data for Name: payment; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.payment (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES ('1', '1', '', 'USD', 'wbpay', 100, 1637907727, 'alpha', 1500, 317, 0);
INSERT INTO public.payment (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES ('2', '2', '', 'USD', 'wbpay', 400, 1637907727, 'alpha', 1500, 317, 0);
INSERT INTO public.payment (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES ('3', '3', '', 'USD', 'wbpay', 900, 1637907727, 'alpha', 1500, 317, 0);
INSERT INTO public.payment (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES ('4', '4', '', 'USD', 'wbpay', 1600, 1637907727, 'alpha', 1500, 317, 0);
INSERT INTO public.payment (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES ('5', '5', '', 'USD', 'wbpay', 2500, 1637907727, 'alpha', 1500, 317, 0);


--
-- TOC entry 3187 (class 2606 OID 41174)
-- Name: delivery delivery_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.delivery
    ADD CONSTRAINT delivery_pkey PRIMARY KEY (order_uid);


--
-- TOC entry 3191 (class 2606 OID 41198)
-- Name: items items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (order_uid, chrt_id);


--
-- TOC entry 3185 (class 2606 OID 41167)
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (order_uid);


--
-- TOC entry 3189 (class 2606 OID 41186)
-- Name: payment payment_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment
    ADD CONSTRAINT payment_pkey PRIMARY KEY (order_uid);


--
-- TOC entry 3192 (class 2606 OID 41175)
-- Name: delivery delivery_order_uid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.delivery
    ADD CONSTRAINT delivery_order_uid_fkey FOREIGN KEY (order_uid) REFERENCES public.orders(order_uid);


--
-- TOC entry 3194 (class 2606 OID 41199)
-- Name: items items_order_uid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_order_uid_fkey FOREIGN KEY (order_uid) REFERENCES public.orders(order_uid);


--
-- TOC entry 3193 (class 2606 OID 41187)
-- Name: payment payment_order_uid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment
    ADD CONSTRAINT payment_order_uid_fkey FOREIGN KEY (order_uid) REFERENCES public.orders(order_uid);


-- Completed on 2023-11-18 14:16:21

--
-- PostgreSQL database dump complete
--

