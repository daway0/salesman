CREATE TABLE Users
(
    id            uuid PRIMARY KEY,
    first_name    varchar                  NOT NULL,
    last_name     varchar                  NOT NULL,
    NSID          varchar(10)              NOT NULL UNIQUE,
    birthdate     date                     NOT NULL,
    email         varchar                  NOT NULL UNIQUE,
    password      varchar                  NOT NULL,
    last_login_at timestamp with time zone,

    created_at    timestamp with time zone NOT NULL,
    updated_at    timestamp with time zone,
    deleted_at    timestamp with time zone
);

CREATE OR REPLACE FUNCTION after_insert_users()
    RETURNS TRIGGER AS
$$
BEGIN
    INSERT INTO panel_customuser (password,
                                  is_superuser,
                                  username,
                                  user_uuid,
                                  first_name,
                                  last_name,
                                  email,
                                  is_staff,
                                  is_active,
                                  date_joined)
    VALUES (NEW.password,
            false,
            NEW.email,
            NEW.id,
            NEW.first_name,
            NEW.last_name,
            NEW.email,
            false,
            true,
            NOW());
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER trigger_after_insert_users
    AFTER INSERT
    ON users
    FOR EACH ROW
EXECUTE FUNCTION after_insert_users();



CREATE TABLE Companies
(
    id         uuid PRIMARY KEY,
    title      varchar                  NOT NULL UNIQUE,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);

CREATE TABLE CompanyUsers
(
    id              uuid PRIMARY KEY,
    company_id      uuid                     NOT NULL REFERENCES Companies (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    user_id         uuid                     NOT NULL REFERENCES Users (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    subscription_no varchar                  NOT NULL,
    created_at      timestamp with time zone NOT NULL,
    updated_at      timestamp with time zone,
    deleted_at      timestamp with time zone
);

CREATE INDEX ON CompanyUsers (company_id);
CREATE INDEX ON CompanyUsers (user_id);
CREATE INDEX ON CompanyUsers (subscription_no);

CREATE TABLE Products
(
    id             uuid PRIMARY KEY,
    company_id     uuid                     NOT NULL REFERENCES Companies (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    title          varchar                  NOT NULL,
    has_commission boolean                  NOT NULL,
    created_at     timestamp with time zone NOT NULL,
    updated_at     timestamp with time zone,
    deleted_at     timestamp with time zone
);

CREATE INDEX ON Products (company_id);

CREATE TABLE CompanyUserProducts
(
    id              uuid PRIMARY KEY,
    company_user_id uuid                     NOT NULL REFERENCES CompanyUsers (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    product_id      uuid                     NOT NULL REFERENCES Products (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    is_active       boolean                  NOT NULL,
    number          varchar                  NOT NULL,
    created_at      timestamp with time zone NOT NULL,
    updated_at      timestamp with time zone,
    deleted_at      timestamp with time zone
);

CREATE INDEX ON CompanyUserProducts (company_user_id);
CREATE INDEX ON CompanyUserProducts (product_id);
CREATE INDEX ON CompanyUserProducts (number);

CREATE TABLE Services
(
    id             uuid PRIMARY KEY,
    product_id     uuid                     NOT NULL REFERENCES Products (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    title          varchar                  NOT NULL,
    description    text                     ,
    price          numeric                  NOT NULL,
    has_commission boolean                  NOT NULL,
    created_at     timestamp with time zone NOT NULL,
    updated_at     timestamp with time zone,
    deleted_at     timestamp with time zone
);

CREATE INDEX ON Services (product_id);
CREATE INDEX ON Services (title);


CREATE TYPE ServiceStatus AS ENUM ('PENDING', 'APPROVED', 'REJECTED', 'CANCELLED');
CREATE TYPE PaymentType AS ENUM ('C2C', 'POS', 'SHEBA');
CREATE TYPE CommissionLevel AS ENUM ('NORMAL', 'PREMIUM');
CREATE TABLE SalesLedger
(
    id                  uuid PRIMARY KEY,
    customer_product_id uuid                     NOT NULL REFERENCES CompanyUserProducts (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    created_by          uuid                     NOT NULL REFERENCES Users (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    approved_by         uuid REFERENCES Users (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    cancelled_by        uuid REFERENCES Users (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    marketer_id         uuid NOT NULL REFERENCES Users (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    status              ServiceStatus            NOT NULL,
    payment_method      PaymentType              NOT NULL,
    TRN                 varchar                  NOT NULL,
    workflow_history    json                     NOT NULL,
    commission_level    CommissionLevel,
    approved_at         timestamp with time zone,
    cancelled_at        timestamp with time zone,
    created_at          timestamp with time zone NOT NULL,
    updated_at          timestamp with time zone,
    deleted_at          timestamp with time zone
);

CREATE INDEX ON SalesLedger (customer_product_id);
CREATE INDEX ON SalesLedger (created_by);
CREATE INDEX ON SalesLedger (approved_by);
CREATE INDEX ON SalesLedger (cancelled_by);
CREATE INDEX ON SalesLedger (marketer_id);


CREATE TABLE LedgerServiceItems
(
    id         uuid PRIMARY KEY,
    ledger_id    uuid                     NOT NULL REFERENCES SalesLedger (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    service_id uuid                     NOT NULL REFERENCES Services (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    price      numeric,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);

CREATE INDEX ON LedgerServiceItems (ledger_id);
CREATE INDEX ON LedgerServiceItems (service_id);



CREATE TABLE Roles
(
    id          uuid PRIMARY KEY,
    title       varchar                  NOT NULL UNIQUE,
    description text,
    created_at  timestamp with time zone NOT NULL,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone
);



CREATE TABLE Permissions
(
    id          uuid PRIMARY KEY,
    action      varchar                  NOT NULL,
    description text,
    created_at  timestamp with time zone NOT NULL,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone
);


CREATE TABLE RolePermissions
(
    id            uuid PRIMARY KEY,
    role_id       uuid                     NOT NULL REFERENCES Roles (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    permission_id uuid                     NOT NULL REFERENCES Permissions (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    created_at    timestamp with time zone NOT NULL,
    updated_at    timestamp with time zone,
    deleted_at    timestamp with time zone
);

CREATE INDEX ON RolePermissions (role_id);
CREATE INDEX ON RolePermissions (permission_id);

CREATE TABLE UserRoles
(
    id         uuid PRIMARY KEY,
    user_id    uuid                     NOT NULL REFERENCES Users (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    role_id    uuid                     NOT NULL REFERENCES Roles (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


CREATE INDEX ON UserRoles (role_id);
CREATE INDEX ON UserRoles (user_id);

CREATE TABLE Logs
(
    id                uuid PRIMARY KEY,
    user_id           uuid                     NOT NULL REFERENCES Users (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    role_id           uuid                     NOT NULL REFERENCES Roles (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    permission_id     uuid                     NOT NULL REFERENCES Permissions (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    role_string       varchar                  NOT NULL,
    permission_string varchar                  NOT NULL,
    route             varchar,
    request_payload   json,
    additional_data   json,
    created_at        timestamp with time zone NOT NULL

);

CREATE INDEX ON Logs (user_id);
CREATE INDEX ON Logs (role_id);
CREATE INDEX ON Logs (permission_id);


CREATE VIEW ServicePath AS
SELECT services.id, concat(services.title || ' - ' || p.title || ' - ' || c.title) as service_path FROM services
INNER JOIN products p on services.product_id = p.id
INNER JOIN companies c on c.id = p.company_id;


CREATE VIEW UserProductPath AS
SELECT companyuserproducts.id, concat(users.first_name || ' ' || users.last_name || ' - ' || companies.title || ' - ' || products.title || ' - ' || companyuserproducts.number) as user_product_path FROM companyuserproducts
INNER JOIN companyusers on companyuserproducts.company_user_id = companyusers.id
INNER JOIN users on companyusers.user_id = users.id
INNER JOIN products on companyuserproducts.product_id = products.id
INNER JOIN companies on products.company_id = companies.id;


CREATE VIEW InvoiceTotalPrice AS
SELECT  ledgerserviceitems.ledger_id, sum(ledgerserviceitems.price) as total_price FROM ledgerserviceitems
GROUP BY ledgerserviceitems.ledger_id;


CREATE VIEW SalesLedgerPrice AS 
SELECT coalesce(sum(lsi.price), 0) as price, sl.id, sl.marketer_id from ledgerserviceitems lsi
inner join salesledger sl on lsi.ledger_id = sl.id
inner join services s on lsi.service_id = s.id
where s.has_commission = true
group by sl.id, sl.marketer_id;