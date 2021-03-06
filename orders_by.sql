PGDMP                          z         	   orders_by #   12.9 (Ubuntu 12.9-0ubuntu0.20.04.1) #   12.9 (Ubuntu 12.9-0ubuntu0.20.04.1)     ?           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            ?           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            ?           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            ?           1262    16399 	   orders_by    DATABASE     {   CREATE DATABASE orders_by WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';
    DROP DATABASE orders_by;
                postgres    false            ?            1259    16422    items    TABLE     ?   CREATE TABLE public.items (
    item_id integer NOT NULL,
    item_code text NOT NULL,
    description text NOT NULL,
    quantity integer NOT NULL,
    order_id integer
);
    DROP TABLE public.items;
       public         heap    postgres    false            ?            1259    16420    items_item_id_seq    SEQUENCE     ?   CREATE SEQUENCE public.items_item_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.items_item_id_seq;
       public          postgres    false    205            ?           0    0    items_item_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.items_item_id_seq OWNED BY public.items.item_id;
          public          postgres    false    204            ?            1259    16408    orders    TABLE     ?   CREATE TABLE public.orders (
    order_id integer NOT NULL,
    customer_name text,
    ordered_at timestamp with time zone DEFAULT now() NOT NULL
);
    DROP TABLE public.orders;
       public         heap    postgres    false            ?            1259    16406    orders_order_id_seq    SEQUENCE     ?   CREATE SEQUENCE public.orders_order_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 *   DROP SEQUENCE public.orders_order_id_seq;
       public          postgres    false    203            ?           0    0    orders_order_id_seq    SEQUENCE OWNED BY     K   ALTER SEQUENCE public.orders_order_id_seq OWNED BY public.orders.order_id;
          public          postgres    false    202                       2604    16425    items item_id    DEFAULT     n   ALTER TABLE ONLY public.items ALTER COLUMN item_id SET DEFAULT nextval('public.items_item_id_seq'::regclass);
 <   ALTER TABLE public.items ALTER COLUMN item_id DROP DEFAULT;
       public          postgres    false    205    204    205                       2604    16411    orders order_id    DEFAULT     r   ALTER TABLE ONLY public.orders ALTER COLUMN order_id SET DEFAULT nextval('public.orders_order_id_seq'::regclass);
 >   ALTER TABLE public.orders ALTER COLUMN order_id DROP DEFAULT;
       public          postgres    false    202    203    203            ?          0    16422    items 
   TABLE DATA           T   COPY public.items (item_id, item_code, description, quantity, order_id) FROM stdin;
    public          postgres    false    205   ?       ?          0    16408    orders 
   TABLE DATA           E   COPY public.orders (order_id, customer_name, ordered_at) FROM stdin;
    public          postgres    false    203   ;       ?           0    0    items_item_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.items_item_id_seq', 2, true);
          public          postgres    false    204            ?           0    0    orders_order_id_seq    SEQUENCE SET     A   SELECT pg_catalog.setval('public.orders_order_id_seq', 1, true);
          public          postgres    false    202                       2606    16430    items items_pkey 
   CONSTRAINT     S   ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (item_id);
 :   ALTER TABLE ONLY public.items DROP CONSTRAINT items_pkey;
       public            postgres    false    205                       2606    16417    orders orders_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (order_id);
 <   ALTER TABLE ONLY public.orders DROP CONSTRAINT orders_pkey;
       public            postgres    false    203                       2606    16487    items fk_items    FK CONSTRAINT     ?   ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items FOREIGN KEY (order_id) REFERENCES public.orders(order_id) ON DELETE CASCADE;
 8   ALTER TABLE ONLY public.items DROP CONSTRAINT fk_items;
       public          postgres    false    203    205    2844            ?   -   x?3?442?????KU04??4?4?2
? 	s ?c???? #
x      ?   /   x?3?.??NU??N?420??54?5?T02?"3ms?=... ? ?     