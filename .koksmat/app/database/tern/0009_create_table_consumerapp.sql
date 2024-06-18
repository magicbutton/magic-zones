/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- sure sild

CREATE TABLE public.consumerapp
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
    ,appid character varying COLLATE pg_catalog."default"  NOT NULL
    ,appversion character varying COLLATE pg_catalog."default"  NOT NULL
    ,appdescription character varying COLLATE pg_catalog."default"  NOT NULL
    ,appurl character varying COLLATE pg_catalog."default"  NOT NULL
    ,appicon character varying COLLATE pg_catalog."default"  NOT NULL
    ,applogo character varying COLLATE pg_catalog."default"  NOT NULL
    ,appowner_id int   NOT NULL
    ,apphomezone_id int   NOT NULL


);

                ALTER TABLE IF EXISTS public.consumerapp
                ADD FOREIGN KEY (appowner_id)
                REFERENCES public.person (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;                ALTER TABLE IF EXISTS public.consumerapp
                ADD FOREIGN KEY (apphomezone_id)
                REFERENCES public.zone (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;                -- lollipop
                CREATE TABLE public.consumerapp_m2m_appSecret (
                id SERIAL PRIMARY KEY,
                created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
                created_by character varying COLLATE pg_catalog."default"  ,
                updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
                updated_by character varying COLLATE pg_catalog."default",
                deleted_at timestamp with time zone
                
                    ,consumerapp_id int  
 
                    ,appSecret_id int  
 

                );
            

                ALTER TABLE public.consumerapp_m2m_appSecret
                ADD FOREIGN KEY (consumerapp_id)
                REFERENCES public.consumerapp (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;

                ALTER TABLE public.consumerapp_m2m_appSecret
                ADD FOREIGN KEY (appSecret_id)
                REFERENCES public.appSecret (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;                -- lollipop
                CREATE TABLE public.consumerapp_m2m_accesspass (
                id SERIAL PRIMARY KEY,
                created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
                created_by character varying COLLATE pg_catalog."default"  ,
                updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
                updated_by character varying COLLATE pg_catalog."default",
                deleted_at timestamp with time zone
                
                    ,consumerapp_id int  
 
                    ,accesspass_id int  
 

                );
            

                ALTER TABLE public.consumerapp_m2m_accesspass
                ADD FOREIGN KEY (consumerapp_id)
                REFERENCES public.consumerapp (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;

                ALTER TABLE public.consumerapp_m2m_accesspass
                ADD FOREIGN KEY (accesspass_id)
                REFERENCES public.accesspass (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;


---- create above / drop below ----
DROP TABLE IF EXISTS public.consumerapp_m2m_appSecret;DROP TABLE IF EXISTS public.consumerapp_m2m_accesspass;
DROP TABLE public.consumerapp;

