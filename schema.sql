CREATE TABLE libraries (
    id serial NOT NULL,
    name character varying(255) NOT NULL,
    taken_places integer DEFAULT 0 NOT NULL,
    total_places integer DEFAULT 0 NOT NULL,
    latitude double precision NOT NULL,
    longitude double precision NOT NULL,
    city character varying(255),
    description text,
    mapping_id character varying(255) NOT NULL,
    contact character varying(255),
    country_code character varying(255) NOT NULL
);
