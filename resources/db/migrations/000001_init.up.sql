--------------------------------------------------------
-- sheet entities --
--------------------------------------------------------

CREATE TYPE application_type AS ENUM ('USER_APP', 'LINE_APP');
CREATE TYPE application_role AS ENUM ('APP_USER', 'APP_ADMIN');

CREATE TABLE sheets
(
    id             SERIAL PRIMARY KEY,
    name           VARCHAR(64)  NOT NULL,
    description    VARCHAR(64)  NOT NULL,
    external_code  VARCHAR(64)  NOT NULL,
    application_id VARCHAR(64)  NOT NULL,
    asset_id       VARCHAR(64)  NOT NULL,
    process_id     VARCHAR(64)  NOT NULL,
    is_active      boolean      NOT NULL default True,
    sections       jsonb,
    namespace      VARCHAR(64)  NOT NULL,
    created_at     TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMP,
    deleted_at     TIMESTAMP,
    created_by     VARCHAR(256) NOT NULL,
    updated_by     VARCHAR(256) NOT NULL,
    deleted_by     VARCHAR(256)
);

--------------------------------------------------------
-- sheet entities end --
--------------------------------------------------------
