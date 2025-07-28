DO $$
DECLARE
        -- PERMISSIONS
        CREATE_INVOICE UUID;
        VIEW_OWN_INVOICES UUID;
        VIEW_OWN_INVOICES_LIST UUID;
        VIEW_ALL_INVOICES UUID;
        VIEW_ALL_INVOICES_LIST UUID;
        DELETE_OWN_INVOICES UUID;
        DELETE_ALL_INVOICES UUID;
        APPROVE_ALL_INVOICES UUID;
        REJECT_ALL_INVOICES UUID;
        CANCEL_OWN_INVOICES UUID;
        CANCEL_ALL_INVOICES UUID;
        VIEW_ALL_USERS_LIST UUID;
        VIEW_ALL_USERS UUID;
        CREATE_ALL_USERS UUID;
        UPDATE_ALL_USERS UUID;
        DELETE_ALL_USERS UUID;
        VIEW_USER_ROLES UUID;
        ADD_USER_ROLE UUID;
        DELETE_USER_ROLE UUID;
        VIEW_ALL_SERVICES_LIST UUID;
        VIEW_ALL_SERVICES UUID;
        CREATE_ALL_SERVICES UUID;
        UPDATE_ALL_SERVICES UUID;
        DELETE_ALL_SERVICES UUID;
        VIEW_ALL_COMPANIES_LIST UUID;
        VIEW_ALL_COMPANIES UUID;
        CREATE_ALL_COMPANIES UUID;
        UPDATE_ALL_COMPANIES UUID;
        DELETE_ALL_COMPANIES UUID;
        VIEW_ALL_ROLES_LIST UUID;
        VIEW_ALL_ROLES UUID;
        CREATE_ALL_ROLES UUID;
        UPDATE_ALL_ROLES UUID;
        DELETE_ALL_ROLES UUID;
        VIEW_ROLE_PERMISSIONS UUID;
        ADD_ROLE_PERMISSION UUID;
        DELETE_ROLE_PERMISSION UUID;
        VIEW_ALL_LOGS UUID;
        VIEW_DASHBOARD UUID;
        VIEW_DASHBOARD_CARD_ALL_TODAY_APPROVED_COUNT UUID;
        VIEW_DASHBOARD_CARD_ALL_TODAY_APPROVED_AMMOUNT UUID;
        VIEW_DASHBOARD_CARD_ALL_TODAY_UNAPPROVED_COUNT UUID;
        VIEW_DASHBOARD_CARD_USER_TODAY_APPROVED_COUNT UUID;
        VIEW_DASHBOARD_CARD_USER_TODAY_APPROVED_AMMOUNT UUID;
        VIEW_DASHBOARD_CARD_USER_TODAY_UNAPPROVED_COUNT UUID;
        VIEW_TODAY_MVP_OPERATOR UUID;
        VIEW_SALES_MVP_LIST UUID;
        VIEW_ALL_SALES_HISTORY UUID;
        VIEW_USER_SALES_HISTORY UUID;

        -- ROLES
        ADMINISTRATOR UUID;
        OPERATOR UUID;
        FINANCE UUID;
        CUSTOMER UUID;

        -- COMPANIES
        COMPANY_1 UUID;
        COMPANY_2 UUID;
        COMPANY_3 UUID;
        COMPANY_4 UUID;

        -- SERVICES
        SERVICE_1 UUID;
        SERVICE_2 UUID;
        SERVICE_3 UUID;
        SERVICE_4 UUID;

        -- USERS
        ADMINISTRATOR_USER UUID;
        OPERATOR_USER UUID;
        FINANCE_USER UUID;
        OPERATOR1_USER UUID;
        OPERATOR2_USER UUID;
        OPERATOR3_USER UUID;
        OPERATOR4_USER UUID;
        OPERATOR5_USER UUID;
        USER_1 UUID;
        CUSTOMER_1 UUID;

BEGIN
        CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

        -- PERMISSIONS
        CREATE_INVOICE = uuid_generate_v4();
        VIEW_OWN_INVOICES = uuid_generate_v4();
        VIEW_OWN_INVOICES_LIST = uuid_generate_v4();
        VIEW_ALL_INVOICES = uuid_generate_v4();
        VIEW_ALL_INVOICES_LIST = uuid_generate_v4();
        DELETE_OWN_INVOICES = uuid_generate_v4();
        DELETE_ALL_INVOICES = uuid_generate_v4();
        APPROVE_ALL_INVOICES = uuid_generate_v4();
        REJECT_ALL_INVOICES = uuid_generate_v4();
        CANCEL_OWN_INVOICES = uuid_generate_v4();
        CANCEL_ALL_INVOICES = uuid_generate_v4();
        VIEW_ALL_USERS_LIST = uuid_generate_v4();
        VIEW_ALL_USERS = uuid_generate_v4();
        CREATE_ALL_USERS = uuid_generate_v4();
        UPDATE_ALL_USERS = uuid_generate_v4();
        DELETE_ALL_USERS = uuid_generate_v4();
        VIEW_USER_ROLES = uuid_generate_v4();
        ADD_USER_ROLE = uuid_generate_v4();
        DELETE_USER_ROLE = uuid_generate_v4();
        VIEW_ALL_SERVICES_LIST = uuid_generate_v4();
        VIEW_ALL_SERVICES = uuid_generate_v4();
        CREATE_ALL_SERVICES = uuid_generate_v4();
        UPDATE_ALL_SERVICES = uuid_generate_v4();
        DELETE_ALL_SERVICES = uuid_generate_v4();
        VIEW_ALL_COMPANIES_LIST = uuid_generate_v4();
        VIEW_ALL_COMPANIES = uuid_generate_v4();
        CREATE_ALL_COMPANIES = uuid_generate_v4();
        UPDATE_ALL_COMPANIES = uuid_generate_v4();
        DELETE_ALL_COMPANIES = uuid_generate_v4();
        VIEW_ALL_ROLES_LIST = uuid_generate_v4();
        VIEW_ALL_ROLES = uuid_generate_v4();
        CREATE_ALL_ROLES = uuid_generate_v4();
        UPDATE_ALL_ROLES = uuid_generate_v4();
        DELETE_ALL_ROLES = uuid_generate_v4();
        VIEW_ROLE_PERMISSIONS = uuid_generate_v4();
        ADD_ROLE_PERMISSION = uuid_generate_v4();
        DELETE_ROLE_PERMISSION = uuid_generate_v4();
        VIEW_ALL_LOGS = uuid_generate_v4();
        VIEW_DASHBOARD = uuid_generate_v4();
        VIEW_DASHBOARD_CARD_ALL_TODAY_APPROVED_COUNT = uuid_generate_v4();
        VIEW_DASHBOARD_CARD_ALL_TODAY_APPROVED_AMMOUNT = uuid_generate_v4();
        VIEW_DASHBOARD_CARD_ALL_TODAY_UNAPPROVED_COUNT = uuid_generate_v4();
        VIEW_DASHBOARD_CARD_USER_TODAY_APPROVED_COUNT = uuid_generate_v4();
        VIEW_DASHBOARD_CARD_USER_TODAY_APPROVED_AMMOUNT = uuid_generate_v4();
        VIEW_DASHBOARD_CARD_USER_TODAY_UNAPPROVED_COUNT = uuid_generate_v4();
        VIEW_TODAY_MVP_OPERATOR = uuid_generate_v4();
        VIEW_SALES_MVP_LIST = uuid_generate_v4();
        VIEW_ALL_SALES_HISTORY = uuid_generate_v4();
        VIEW_USER_SALES_HISTORY = uuid_generate_v4();

        -- ROLES
        ADMINISTRATOR = uuid_generate_v4();
        OPERATOR = uuid_generate_v4();
        FINANCE = uuid_generate_v4();
        CUSTOMER = uuid_generate_v4();

        -- COMPANIES
        COMPANY_1 = uuid_generate_v4();
        COMPANY_2 = uuid_generate_v4();
        COMPANY_3 = uuid_generate_v4();
        COMPANY_4 = uuid_generate_v4();

        -- SERVICES
        SERVICE_1 = uuid_generate_v4();
        SERVICE_2 = uuid_generate_v4();
        SERVICE_3 = uuid_generate_v4();
        SERVICE_4 = uuid_generate_v4();

        -- USERS
        ADMINISTRATOR_USER = uuid_generate_v4();
        OPERATOR_USER = '00000000-0000-0000-0000-000000000036';
        FINANCE_USER = uuid_generate_v4();
        OPERATOR1_USER = uuid_generate_v4();
        OPERATOR2_USER = uuid_generate_v4();
        OPERATOR3_USER = uuid_generate_v4();
        OPERATOR4_USER = uuid_generate_v4();
        OPERATOR5_USER = uuid_generate_v4();
        USER_1 = uuid_generate_v4();
        CUSTOMER_1 = uuid_generate_v4();
        

-- PERMISSIONS
INSERT INTO Permissions (id, action) VALUES 
        
        -- INVOICES (SALES LEDGER)
        (CREATE_INVOICE, 'CREATE_INVOICE'),
        (VIEW_OWN_INVOICES, 'VIEW_OWN_INVOICES'),
        (VIEW_OWN_INVOICES_LIST, 'VIEW_OWN_INVOICES_LIST'),
        (VIEW_ALL_INVOICES, 'VIEW_ALL_INVOICES'),
        (VIEW_ALL_INVOICES_LIST, 'VIEW_ALL_INVOICES_LIST'),
        (DELETE_OWN_INVOICES, 'DELETE_OWN_INVOICES'),
        (DELETE_ALL_INVOICES, 'DELETE_ALL_INVOICES'),
        (APPROVE_ALL_INVOICES, 'APPROVE_ALL_INVOICES'),
        (REJECT_ALL_INVOICES, 'REJECT_ALL_INVOICES'),
        (CANCEL_OWN_INVOICES, 'CANCEL_OWN_INVOICES'),
        (CANCEL_ALL_INVOICES, 'CANCEL_ALL_INVOICES'),

        -- DASHBOARD
        (VIEW_DASHBOARD, 'VIEW_DASHBOARD'),
        (VIEW_DASHBOARD_CARD_ALL_TODAY_APPROVED_COUNT, 'VIEW_DASHBOARD_CARD_ALL_TODAY_APPROVED_COUNT'),
        (VIEW_DASHBOARD_CARD_ALL_TODAY_APPROVED_AMMOUNT, 'VIEW_DASHBOARD_CARD_ALL_TODAY_APPROVED_AMMOUNT'),
        (VIEW_DASHBOARD_CARD_ALL_TODAY_UNAPPROVED_COUNT, 'VIEW_DASHBOARD_CARD_ALL_TODAY_UNAPPROVED_COUNT'),
        (VIEW_DASHBOARD_CARD_USER_TODAY_APPROVED_COUNT, 'VIEW_DASHBOARD_CARD_USER_TODAY_APPROVED_COUNT'),
        (VIEW_DASHBOARD_CARD_USER_TODAY_APPROVED_AMMOUNT, 'VIEW_DASHBOARD_CARD_USER_TODAY_APPROVED_AMMOUNT'),
        (VIEW_DASHBOARD_CARD_USER_TODAY_UNAPPROVED_COUNT, 'VIEW_DASHBOARD_CARD_USER_TODAY_UNAPPROVED_COUNT'),
        (VIEW_TODAY_MVP_OPERATOR, 'VIEW_TODAY_MVP_OPERATOR'),
        (VIEW_SALES_MVP_LIST, 'VIEW_SALES_MVP_LIST'),
        (VIEW_ALL_SALES_HISTORY, 'VIEW_ALL_SALES_HISTORY'),
        (VIEW_USER_SALES_HISTORY, 'VIEW_USER_SALES_HISTORY'),

        -- USERS 
        (VIEW_ALL_USERS_LIST, 'VIEW_ALL_USERS_LIST'),
        (VIEW_ALL_USERS, 'VIEW_ALL_USERS'),
        (CREATE_ALL_USERS, 'CREATE_ALL_USERS'),
        (UPDATE_ALL_USERS, 'UPDATE_ALL_USERS'),
        (DELETE_ALL_USERS, 'DELETE_ALL_USERS'),
        (VIEW_USER_ROLES, 'VIEW_USER_ROLES'),
        (ADD_USER_ROLE, 'ADD_USER_ROLE'),
        (DELETE_USER_ROLE, 'DELETE_USER_ROLE'),

        -- SERVICES
        (VIEW_ALL_SERVICES_LIST, 'VIEW_ALL_SERVICES_LIST'),
        (VIEW_ALL_SERVICES, 'VIEW_ALL_SERVICES'),
        (CREATE_ALL_SERVICES, 'CREATE_ALL_SERVICES'),
        (UPDATE_ALL_SERVICES, 'UPDATE_ALL_SERVICES'),
        (DELETE_ALL_SERVICES, 'DELETE_ALL_SERVICES'),

        -- COMPANIES
        (VIEW_ALL_COMPANIES_LIST, 'VIEW_ALL_COMPANIES_LIST'),
        (VIEW_ALL_COMPANIES, 'VIEW_ALL_COMPANIES'),
        (CREATE_ALL_COMPANIES, 'CREATE_ALL_COMPANIES'),
        (UPDATE_ALL_COMPANIES, 'UPDATE_ALL_COMPANIES'),
        (DELETE_ALL_COMPANIES, 'DELETE_ALL_COMPANIES'),

        -- ROLES
        (VIEW_ALL_ROLES_LIST, 'VIEW_ALL_ROLES_LIST'),
        (VIEW_ALL_ROLES, 'VIEW_ALL_ROLES'),
        (CREATE_ALL_ROLES, 'CREATE_ALL_ROLES'),
        (UPDATE_ALL_ROLES, 'UPDATE_ALL_ROLES'),
        (DELETE_ALL_ROLES, 'DELETE_ALL_ROLES'),
        (VIEW_ROLE_PERMISSIONS, 'VIEW_ROLE_PERMISSIONS'),
        (ADD_ROLE_PERMISSION, 'ADD_ROLE_PERMISSION'),
        (DELETE_ROLE_PERMISSION, 'DELETE_ROLE_PERMISSION'),

        -- LOGS
        (VIEW_ALL_LOGS, 'VIEW_ALL_LOGS');

INSERT INTO Roles (id, title, created_at) VALUES 
        (ADMINISTRATOR, 'ADMINISTRATOR', now()),
        (OPERATOR, 'OPERATOR', now()),
        (FINANCE, 'FINANCE', now()),
        (CUSTOMER, 'CUSTOMER', now());
    
INSERT INTO RolePermissions (id, role_id, permission_id, created_at) VALUES 
        -- ADMINISTRATOR
        (uuid_generate_v4(), ADMINISTRATOR, CREATE_INVOICE , now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_OWN_INVOICES , now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_OWN_INVOICES_LIST , now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_ALL_INVOICES, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_ALL_INVOICES_LIST, now()),
        (uuid_generate_v4(), ADMINISTRATOR, DELETE_OWN_INVOICES, now()),
        (uuid_generate_v4(), ADMINISTRATOR, DELETE_ALL_INVOICES, now()),
        (uuid_generate_v4(), ADMINISTRATOR, APPROVE_ALL_INVOICES, now()),
        (uuid_generate_v4(), ADMINISTRATOR, REJECT_ALL_INVOICES, now()),
        (uuid_generate_v4(), ADMINISTRATOR, CANCEL_OWN_INVOICES, now()),
        (uuid_generate_v4(), ADMINISTRATOR, CANCEL_ALL_INVOICES, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_ALL_USERS_LIST, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_ALL_USERS, now()),
        (uuid_generate_v4(), ADMINISTRATOR, CREATE_ALL_USERS, now()),
        (uuid_generate_v4(), ADMINISTRATOR, UPDATE_ALL_USERS, now()),
        (uuid_generate_v4(), ADMINISTRATOR, DELETE_ALL_USERS, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_USER_ROLES, now()),
        (uuid_generate_v4(), ADMINISTRATOR, ADD_USER_ROLE, now()),
        (uuid_generate_v4(), ADMINISTRATOR, DELETE_USER_ROLE, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_ALL_SERVICES_LIST, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_ALL_SERVICES, now()),
        (uuid_generate_v4(), ADMINISTRATOR, CREATE_ALL_SERVICES, now()),
        (uuid_generate_v4(), ADMINISTRATOR, UPDATE_ALL_SERVICES, now()),
        (uuid_generate_v4(), ADMINISTRATOR, DELETE_ALL_SERVICES, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_ALL_COMPANIES_LIST, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_ALL_COMPANIES, now()),
        (uuid_generate_v4(), ADMINISTRATOR, CREATE_ALL_COMPANIES, now()),
        (uuid_generate_v4(), ADMINISTRATOR, UPDATE_ALL_COMPANIES, now()),
        (uuid_generate_v4(), ADMINISTRATOR, DELETE_ALL_COMPANIES, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_ALL_ROLES_LIST, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_ALL_ROLES, now()),
        (uuid_generate_v4(), ADMINISTRATOR, CREATE_ALL_ROLES, now()),
        (uuid_generate_v4(), ADMINISTRATOR, UPDATE_ALL_ROLES, now()),
        (uuid_generate_v4(), ADMINISTRATOR, DELETE_ALL_ROLES, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_ROLE_PERMISSIONS, now()),
        (uuid_generate_v4(), ADMINISTRATOR, ADD_ROLE_PERMISSION, now()),
        (uuid_generate_v4(), ADMINISTRATOR, DELETE_ROLE_PERMISSION, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_ALL_LOGS, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_DASHBOARD, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_DASHBOARD_CARD_ALL_TODAY_APPROVED_COUNT, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_DASHBOARD_CARD_ALL_TODAY_APPROVED_AMMOUNT, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_DASHBOARD_CARD_ALL_TODAY_UNAPPROVED_COUNT, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_DASHBOARD_CARD_USER_TODAY_APPROVED_COUNT, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_DASHBOARD_CARD_USER_TODAY_APPROVED_AMMOUNT, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_DASHBOARD_CARD_USER_TODAY_UNAPPROVED_COUNT, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_TODAY_MVP_OPERATOR, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_SALES_MVP_LIST, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_ALL_SALES_HISTORY, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_USER_SALES_HISTORY, now()),

        -- OPERATOR
        (uuid_generate_v4(), OPERATOR, CREATE_INVOICE , now()),
        (uuid_generate_v4(), OPERATOR, VIEW_OWN_INVOICES , now()),
        (uuid_generate_v4(), OPERATOR, VIEW_OWN_INVOICES_LIST , now()),
        (uuid_generate_v4(), OPERATOR, DELETE_OWN_INVOICES, now()),
        (uuid_generate_v4(), OPERATOR, CANCEL_OWN_INVOICES, now()),
        (uuid_generate_v4(), OPERATOR, VIEW_ALL_USERS_LIST, now()),
        (uuid_generate_v4(), OPERATOR, VIEW_ALL_SERVICES_LIST, now()),
        (uuid_generate_v4(), OPERATOR, VIEW_ALL_COMPANIES_LIST, now()),
        (uuid_generate_v4(), OPERATOR, VIEW_DASHBOARD, now()),
        (uuid_generate_v4(), OPERATOR, VIEW_DASHBOARD_CARD_USER_TODAY_APPROVED_COUNT, now()),
        (uuid_generate_v4(), OPERATOR, VIEW_DASHBOARD_CARD_USER_TODAY_APPROVED_AMMOUNT, now()),
        (uuid_generate_v4(), OPERATOR, VIEW_DASHBOARD_CARD_USER_TODAY_UNAPPROVED_COUNT, now()),
        (uuid_generate_v4(), OPERATOR, VIEW_TODAY_MVP_OPERATOR, now()),
        (uuid_generate_v4(), OPERATOR, VIEW_USER_SALES_HISTORY, now()),

        -- FINANCE
        (uuid_generate_v4(), FINANCE, VIEW_ALL_INVOICES, now()),
        (uuid_generate_v4(), FINANCE, VIEW_ALL_INVOICES_LIST, now()),
        (uuid_generate_v4(), FINANCE, DELETE_ALL_INVOICES, now()),
        (uuid_generate_v4(), FINANCE, APPROVE_ALL_INVOICES, now()),
        (uuid_generate_v4(), FINANCE, REJECT_ALL_INVOICES, now()),
        (uuid_generate_v4(), FINANCE, CANCEL_ALL_INVOICES, now()),
        (uuid_generate_v4(), FINANCE, VIEW_ALL_SERVICES_LIST, now()),
        (uuid_generate_v4(), FINANCE, VIEW_ALL_COMPANIES_LIST, now()),
        (uuid_generate_v4(), FINANCE, VIEW_DASHBOARD, now()),
        (uuid_generate_v4(), FINANCE, VIEW_DASHBOARD_CARD_ALL_TODAY_APPROVED_COUNT, now()),
        (uuid_generate_v4(), FINANCE, VIEW_DASHBOARD_CARD_ALL_TODAY_APPROVED_AMMOUNT, now()),
        (uuid_generate_v4(), FINANCE, VIEW_DASHBOARD_CARD_ALL_TODAY_UNAPPROVED_COUNT, now()),
        (uuid_generate_v4(), FINANCE, VIEW_TODAY_MVP_OPERATOR, now()),
        (uuid_generate_v4(), FINANCE, VIEW_SALES_MVP_LIST, now()),
        (uuid_generate_v4(), FINANCE, VIEW_ALL_SALES_HISTORY, now());

INSERT INTO Companies (id, title, created_at) VALUES 
        (COMPANY_1, 'Company 1', now()),
        (COMPANY_2, 'Company 2', now()),
        (COMPANY_3, 'Company 3', now()),
        (COMPANY_4, 'Company 4', now());

INSERT INTO Services (id, company_id, title, price, created_at, type) VALUES 
        (SERVICE_1, COMPANY_1, 'Service 1', 100, now(), 'NORMAL'),
        (SERVICE_2, COMPANY_2, 'Service 2', 200, now(), 'NORMAL'),
        (SERVICE_3, COMPANY_3, 'Service 3', 300, now(), 'NORMAL'),
        (SERVICE_4, COMPANY_4, 'Service 4', 400, now(), 'PREMIUM');

INSERT INTO Users (id, first_name,last_name, NSID, birthdate, email, password, created_at) VALUES 
        (ADMINISTRATOR_USER, 'Administrator', 'User', '1234567890', '2000-01-01', 'admin@example.com', 'pbkdf2_sha256$1000000$BV8vF1TzUL8vxeSONtgHP4$zB+T8eKgL3xN/OYHVjUuQgVW3ItRVdDLVZXWY+UYkO4=', now()),
        (OPERATOR_USER, 'Operator', 'User', '1234567891', '2000-01-01', 'operator@example.com', 'pbkdf2_sha256$1000000$BV8vF1TzUL8vxeSONtgHP4$zB+T8eKgL3xN/OYHVjUuQgVW3ItRVdDLVZXWY+UYkO4=', now()),
        (OPERATOR2_USER, 'Operator2', 'User', '1234567894', '2000-01-01', 'operator2@example.com', 'pbkdf2_sha256$1000000$BV8vF1TzUL8vxeSONtgHP4$zB+T8eKgL3xN/OYHVjUuQgVW3ItRVdDLVZXWY+UYkO4=', now()),
        (OPERATOR3_USER, 'Operator3', 'User', '1234567895', '2000-01-01', 'operator3@example.com', 'pbkdf2_sha256$1000000$BV8vF1TzUL8vxeSONtgHP4$zB+T8eKgL3xN/OYHVjUuQgVW3ItRVdDLVZXWY+UYkO4=', now()),
        (OPERATOR4_USER, 'Operator4', 'User', '1234567896', '2000-01-01', 'operator4@example.com', 'pbkdf2_sha256$1000000$BV8vF1TzUL8vxeSONtgHP4$zB+T8eKgL3xN/OYHVjUuQgVW3ItRVdDLVZXWY+UYkO4=', now()),
        (OPERATOR5_USER, 'Operator5', 'User', '1234567897', '2000-01-01', 'operator5@example.com', 'pbkdf2_sha256$1000000$BV8vF1TzUL8vxeSONtgHP4$zB+T8eKgL3xN/OYHVjUuQgVW3ItRVdDLVZXWY+UYkO4=', now()),
        (FINANCE_USER, 'Finance', 'User', '1234567892', '2000-01-01', 'finance@example.com', 'pbkdf2_sha256$1000000$BV8vF1TzUL8vxeSONtgHP4$zB+T8eKgL3xN/OYHVjUuQgVW3ItRVdDLVZXWY+UYkO4=', now()),
        (USER_1, 'User', 'User', '1234567893', '2000-01-01', 'user@example.com', 'pbkdf2_sha256$1000000$BV8vF1TzUL8vxeSONtgHP4$zB+T8eKgL3xN/OYHVjUuQgVW3ItRVdDLVZXWY+UYkO4=', now()),
        (CUSTOMER_1, 'Customer', 'User', '1234564213', '2000-01-01', 'customer@example.com', 'pbkdf2_sha256$1000000$BV8vF1TzUL8vxeSONtgHP4$zB+T8eKgL3xN/OYHVjUuQgVW3ItRVdDLVZXWY+UYkO4=', now());

INSERT INTO UserRoles (id, user_id, role_id, created_at) VALUES 
        (uuid_generate_v4(), ADMINISTRATOR_USER, ADMINISTRATOR, now()),
        (uuid_generate_v4(), OPERATOR_USER, OPERATOR, now()),
        (uuid_generate_v4(), OPERATOR2_USER, OPERATOR, now()),
        (uuid_generate_v4(), OPERATOR3_USER, OPERATOR, now()),
        (uuid_generate_v4(), OPERATOR4_USER, OPERATOR, now()),
        (uuid_generate_v4(), FINANCE_USER, FINANCE, now()),
        (uuid_generate_v4(), USER_1, OPERATOR, now()),
        (uuid_generate_v4(), CUSTOMER_1, CUSTOMER, now());

END $$;
