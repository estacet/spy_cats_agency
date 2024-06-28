CREATE TYPE status AS ENUM ('initiated', 'ongoing', 'completed');

CREATE TABLE "public"."spyCats" (
    id               uuid PRIMARY KEY,
    name             VARCHAR(100)   NOT NULL,
    experience_years INT            NOT NULL,
    breed            VARCHAR(50)    NOT NULL,
    salary           DECIMAL(10, 2) NOT NULL
);

CREATE TABLE "public"."missions" (
     id uuid PRIMARY KEY,
     cat_id uuid,
     status status NOT NULL DEFAULT 1,
     FOREIGN KEY (cat_id) REFERENCES public."spyCats"(ID)
);


CREATE TABLE "public"."targets" (
    id uuid PRIMARY KEY,
    mission_id uuid,
    name VARCHAR(100) NOT NULL,
    country VARCHAR(50) NOT NULL,
    notes TEXT,
    status status NOT NULL DEFAULT 1,
    FOREIGN KEY (mission_id) REFERENCES public.missions(ID)
);