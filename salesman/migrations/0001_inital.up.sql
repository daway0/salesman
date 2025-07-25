CREATE TABLE Users 
(
    id            uuid PRIMARY KEY,
    first_name    varchar                  NOT NULL,
    last_name     varchar                  NOT NULL,
    NSID          varchar(10)              NOT NULL,
    birthdate     date                     NOT NULL,
    email         varchar                  NOT NULL,
    password      varchar                  NOT NULL,
    created_at    timestamp with time zone NOT NULL,
    last_login_at timestamp with time zone,
    updated_at    timestamp with time zone,
    deleted_at    timestamp with time zone
);


ALTER TABLE Users
    ADD UNIQUE (NSID);

ALTER TABLE Users
    ADD UNIQUE (email);

CREATE OR REPLACE FUNCTION after_insert_users()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO panel_customuser (
        password,
        is_superuser,
        username,
        user_uuid,
        first_name,
        last_name,
        email,
        is_staff,
        is_active,
        date_joined
    )
    VALUES (
        'pbkdf2_sha256$1000000$BV8vF1TzUL8vxeSONtgHP4$zB+T8eKgL3xN/OYHVjUuQgVW3ItRVdDLVZXWY+UYkO4=',
        false,
        NEW.email,
        NEW.id,
        NEW.first_name,
        NEW.last_name,
        NEW.email,
        false,
        true,
        NOW()
    );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER trigger_after_insert_users
AFTER INSERT ON users
FOR EACH ROW
EXECUTE FUNCTION after_insert_users();



CREATE TABLE Companies 
(
    id          uuid PRIMARY KEY,
    title       varchar                  NOT NULL,
    brand_name  varchar,
    CID         varchar(20),
    description text,
    image_url   varchar(255),
    created_at  timestamp with time zone NOT NULL,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone
);

ALTER TABLE Companies
    ADD UNIQUE (CID);


CREATE TABLE Services 
(
    id          uuid PRIMARY KEY,
    company_id  uuid                     NOT NULL REFERENCES Companies (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    title       varchar                  NOT NULL,
    description text,
    price       numeric                  NOT NULL,
    image_url   varchar(255),
    created_at  timestamp with time zone NOT NULL,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone
);

CREATE INDEX ON Services 
    (company_id);

CREATE TYPE Status AS ENUM ('PENDING', 'APPROVED', 'REJECTED', 'CANCELLED');
CREATE TABLE SalesLedger 
(
    id               uuid PRIMARY KEY,
    customer_id      uuid                     NOT NULL REFERENCES Users (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    service_id       uuid                     NOT NULL REFERENCES Services (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    created_by       uuid                     NOT NULL REFERENCES Users (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    referer_id       uuid REFERENCES Users (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    approved_by      uuid REFERENCES Users (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    cancelled_by     uuid REFERENCES Users (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    status           Status                   NOT NULL,
    price            numeric                  NOT NULL,
    sales_price      numeric                  NOT NULL,
    sales_discount   numeric,
    TRN              varchar,
    workflow_history json,
    approved_at      timestamp with time zone,
    cancelled_at     timestamp with time zone,
    created_at       timestamp with time zone NOT NULL,
    updated_at       timestamp with time zone,
    deleted_at       timestamp with time zone
);

ALTER TABLE SalesLedger
    ADD UNIQUE (TRN);

CREATE INDEX ON SalesLedger
    (service_id);
CREATE INDEX ON SalesLedger
    (approved_by);


CREATE TABLE Roles 
(
    id          uuid PRIMARY KEY,
    title       varchar                  NOT NULL,
    description text,
    created_at  timestamp with time zone NOT NULL,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone
);

ALTER TABLE Roles
    ADD UNIQUE (title);



CREATE TABLE Permissions 
(
    id           uuid PRIMARY KEY,
    action       varchar                  NOT NULL,
    description  text
);


CREATE TABLE RolePermissions 
(
    id            uuid PRIMARY KEY,
    role_id       uuid                     NOT NULL REFERENCES Roles (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    permission_id uuid                     NOT NULL REFERENCES Permissions (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    created_at    timestamp with time zone NOT NULL,
    deleted_at    timestamp with time zone
);

CREATE INDEX ON RolePermissions
    (role_id);

CREATE INDEX ON RolePermissions
    (permission_id);


CREATE TABLE UserRoles 
(
    id         uuid PRIMARY KEY,
    user_id    uuid                     NOT NULL REFERENCES Users (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    role_id    uuid                     NOT NULL REFERENCES Roles (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    created_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


CREATE INDEX ON UserRoles
    (role_id);

CREATE INDEX ON UserRoles
    (user_id);

CREATE TABLE Logs 
(
    id                uuid PRIMARY KEY,
    user_id           uuid                     NOT NULL REFERENCES Users (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    role_id           uuid                     NOT NULL REFERENCES Roles (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    permission_id     uuid                     NOT NULL REFERENCES Permissions (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    role_string       varchar                  NOT NULL,
    permission_string varchar                  NOT NULL,
    route             varchar                  NOT NULL,
    request_payload   json                     NOT NULL,
    additional_data   json,
    created_at        timestamp with time zone NOT NULL

);


