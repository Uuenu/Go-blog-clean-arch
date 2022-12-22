CREATE TABLE IF NOT EXISTS public.article(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    author_id UUID NOT NULL,
    CONSTRAINT author_fk FOREIGN KEY (author_id) REFERENCES public.author(id),
    header VARCHAR(30) NOT NULL, 
    text TEXT NOT NULL
);