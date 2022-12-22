CREATE TABLE IF NOT EXISTS public.author(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    passwordHash VARCHAR(100) NOT NULL,
    salt VARCHAR(100) NOT NULL
);