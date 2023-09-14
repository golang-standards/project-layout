begin;

CREATE TYPE role AS ENUM (
    'OWNER',
    'COACH',
    'JUDGE',
    'DECOY'
    );

commit;
