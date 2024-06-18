/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- sure sild

CREATE TABLE public.auditlog
(
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by character varying COLLATE pg_catalog."default"  ,

    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by character varying COLLATE pg_catalog."default" ,

    deleted_at timestamp with time zone
    ,tenant character varying COLLATE pg_catalog."default"  NOT NULL
    ,searchindex character varying COLLATE pg_catalog."default"  NOT NULL
    ,name character varying COLLATE pg_catalog."default"  NOT NULL
    ,description character varying COLLATE pg_catalog."default" 
    ,action character varying COLLATE pg_catalog."default"  NOT NULL
    ,entity character varying COLLATE pg_catalog."default"  NOT NULL
    ,entityid character varying COLLATE pg_catalog."default"  NOT NULL
    ,requesttype character varying COLLATE pg_catalog."default"  NOT NULL
    ,responsetype character varying COLLATE pg_catalog."default"  NOT NULL
    ,actor_id int   NOT NULL
    ,details JSONB  


);

                ALTER TABLE IF EXISTS public.auditlog
                ADD FOREIGN KEY (actor_id)
                REFERENCES public.person (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;


---- create above / drop below ----

DROP TABLE public.auditlog;

