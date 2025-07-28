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
        (ADMINISTRATOR_USER, 'Administrator', 'User', '1234567890', '2000-01-01', 'admin@example.com', 'password1', now()),
        (OPERATOR_USER, 'Operator', 'User', '1234567891', '2000-01-01', 'operator@example.com', 'password2', now()),
        (OPERATOR2_USER, 'Operator2', 'User', '1234567894', '2000-01-01', 'operator2@example.com', 'password2', now()),
        (OPERATOR3_USER, 'Operator3', 'User', '1234567895', '2000-01-01', 'operator3@example.com', 'password2', now()),
        (OPERATOR4_USER, 'Operator4', 'User', '1234567896', '2000-01-01', 'operator4@example.com', 'password2', now()),
        (OPERATOR5_USER, 'Operator5', 'User', '1234567897', '2000-01-01', 'operator5@example.com', 'password2', now()),
        (FINANCE_USER, 'Finance', 'User', '1234567892', '2000-01-01', 'finance@example.com', 'password3', now()),
        (USER_1, 'User', 'User', '1234567893', '2000-01-01', 'user@example.com', 'password4', now()),
        (CUSTOMER_1, 'Customer', 'User', '1234564213', '2000-01-01', 'customer@example.com', 'password4', now());

INSERT INTO UserRoles (id, user_id, role_id, created_at) VALUES 
        (uuid_generate_v4(), ADMINISTRATOR_USER, ADMINISTRATOR, now()),
        (uuid_generate_v4(), OPERATOR_USER, OPERATOR, now()),
        (uuid_generate_v4(), OPERATOR2_USER, OPERATOR, now()),
        (uuid_generate_v4(), OPERATOR3_USER, OPERATOR, now()),
        (uuid_generate_v4(), OPERATOR4_USER, OPERATOR, now()),
        (uuid_generate_v4(), FINANCE_USER, FINANCE, now()),
        (uuid_generate_v4(), USER_1, OPERATOR, now()),
        (uuid_generate_v4(), CUSTOMER_1, CUSTOMER, now());

INSERT INTO SalesLedger (id, customer_id, service_id, created_by, price, sales_price, created_at, approved_by, approved_at, cancelled_by, cancelled_at, status, updated_at) VALUES 
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_1, OPERATOR_USER, 100, 110, now() - interval '7 days', FINANCE_USER, now() - interval '6 days', NULL, NULL, 'APPROVED', now() - interval '6 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_2, OPERATOR2_USER, 200, 210, now() - interval '6 days', FINANCE_USER, now() - interval '5 days', NULL, NULL, 'APPROVED', now() - interval '5 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_3, OPERATOR3_USER, 300, 310, now() - interval '5 days', FINANCE_USER, now() - interval '4 days', NULL, NULL, 'APPROVED', now() - interval '4 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_4, OPERATOR4_USER, 400, 410, now() - interval '4 days', FINANCE_USER, now() - interval '3 days', NULL, NULL, 'APPROVED', now() - interval '3 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_1, OPERATOR5_USER, 100, 115, now() - interval '3 days', FINANCE_USER, now() - interval '2 days', NULL, NULL, 'APPROVED', now() - interval '2 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_2, OPERATOR_USER, 200, 220, now() - interval '2 days', FINANCE_USER, now() - interval '1 days', NULL, NULL, 'APPROVED', now() - interval '1 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_3, OPERATOR2_USER, 300, 320, now() - interval '1 days', FINANCE_USER, now(), NULL, NULL, 'APPROVED', now()),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_4, OPERATOR3_USER, 400, 420, now() - interval '6 days', FINANCE_USER, now() - interval '5 days', NULL, NULL, 'APPROVED', now() - interval '5 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_1, OPERATOR4_USER, 100, 130, now() - interval '5 days', FINANCE_USER, now() - interval '4 days', NULL, NULL, 'APPROVED', now() - interval '4 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_2, OPERATOR5_USER, 200, 230, now() - interval '4 days', FINANCE_USER, now() - interval '3 days', NULL, NULL, 'APPROVED', now() - interval '3 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_3, OPERATOR_USER, 300, 310, now() - interval '3 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '3 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_4, OPERATOR2_USER, 400, 410, now() - interval '2 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '2 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_1, OPERATOR3_USER, 100, 120, now() - interval '1 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '1 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_2, OPERATOR4_USER, 200, 220, now(), NULL, NULL, NULL, NULL, 'PENDING', now()),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_3, OPERATOR5_USER, 300, 330, now() - interval '7 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '7 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_4, OPERATOR_USER, 400, 440, now() - interval '6 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '6 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_1, OPERATOR2_USER, 100, 140, now() - interval '5 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '5 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_2, OPERATOR3_USER, 200, 250, now() - interval '4 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '4 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_3, OPERATOR4_USER, 300, 350, now() - interval '3 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '3 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_4, OPERATOR5_USER, 400, 460, now() - interval '2 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '2 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_1, OPERATOR_USER, 100, 110, now() - interval '30 days', FINANCE_USER, now() - interval '29 days', NULL, NULL, 'APPROVED', now() - interval '29 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_2, OPERATOR2_USER, 200, 210, now() - interval '30 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '30 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_3, OPERATOR3_USER, 300, 320, now() - interval '29 days', FINANCE_USER, now() - interval '28 days', NULL, NULL, 'APPROVED', now() - interval '28 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_4, OPERATOR4_USER, 400, 430, now() - interval '29 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '29 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_1, OPERATOR5_USER, 150, 160, now() - interval '28 days', FINANCE_USER, now() - interval '27 days', NULL, NULL, 'APPROVED', now() - interval '27 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_2, OPERATOR_USER, 220, 230, now() - interval '28 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '28 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_3, OPERATOR2_USER, 320, 340, now() - interval '27 days', FINANCE_USER, now() - interval '26 days', NULL, NULL, 'APPROVED', now() - interval '26 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_4, OPERATOR3_USER, 420, 450, now() - interval '27 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '27 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_1, OPERATOR4_USER, 130, 140, now() - interval '26 days', FINANCE_USER, now() - interval '25 days', NULL, NULL, 'APPROVED', now() - interval '25 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_2, OPERATOR5_USER, 230, 250, now() - interval '26 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '26 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_3, OPERATOR_USER, 330, 350, now() - interval '25 days', FINANCE_USER, now() - interval '24 days', NULL, NULL, 'APPROVED', now() - interval '24 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_4, OPERATOR2_USER, 430, 470, now() - interval '25 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '25 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_1, OPERATOR3_USER, 140, 150, now() - interval '24 days', FINANCE_USER, now() - interval '23 days', NULL, NULL, 'APPROVED', now() - interval '23 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_2, OPERATOR4_USER, 240, 260, now() - interval '24 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '24 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_3, OPERATOR5_USER, 340, 360, now() - interval '23 days', FINANCE_USER, now() - interval '22 days', NULL, NULL, 'APPROVED', now() - interval '22 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_4, OPERATOR_USER, 440, 480, now() - interval '23 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '23 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_1, OPERATOR2_USER, 160, 170, now() - interval '22 days', FINANCE_USER, now() - interval '21 days', NULL, NULL, 'APPROVED', now() - interval '21 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_2, OPERATOR3_USER, 260, 280, now() - interval '22 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '22 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_3, OPERATOR4_USER, 360, 380, now() - interval '21 days', FINANCE_USER, now() - interval '20 days', NULL, NULL, 'APPROVED', now() - interval '20 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_4, OPERATOR5_USER, 460, 500, now() - interval '21 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '21 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_1, OPERATOR_USER, 170, 180, now() - interval '20 days', FINANCE_USER, now() - interval '19 days', NULL, NULL, 'APPROVED', now() - interval '19 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_2, OPERATOR2_USER, 270, 290, now() - interval '20 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '20 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_3, OPERATOR3_USER, 370, 390, now() - interval '19 days', FINANCE_USER, now() - interval '18 days', NULL, NULL, 'APPROVED', now() - interval '18 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_4, OPERATOR4_USER, 470, 510, now() - interval '19 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '19 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_1, OPERATOR5_USER, 180, 190, now() - interval '18 days', FINANCE_USER, now() - interval '17 days', NULL, NULL, 'APPROVED', now() - interval '17 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_2, OPERATOR_USER, 280, 300, now() - interval '18 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '18 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_3, OPERATOR2_USER, 380, 400, now() - interval '17 days', FINANCE_USER, now() - interval '16 days', NULL, NULL, 'APPROVED', now() - interval '16 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_4, OPERATOR3_USER, 480, 520, now() - interval '17 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '17 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_1, OPERATOR4_USER, 190, 200, now() - interval '16 days', FINANCE_USER, now() - interval '15 days', NULL, NULL, 'APPROVED', now() - interval '15 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_2, OPERATOR5_USER, 290, 310, now() - interval '16 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '16 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_3, OPERATOR_USER, 390, 410, now() - interval '15 days', FINANCE_USER, now() - interval '14 days', NULL, NULL, 'APPROVED', now() - interval '14 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_4, OPERATOR2_USER, 490, 530, now() - interval '15 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '15 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_1, OPERATOR3_USER, 200, 210, now() - interval '14 days', FINANCE_USER, now() - interval '13 days', NULL, NULL, 'APPROVED', now() - interval '13 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_2, OPERATOR4_USER, 300, 320, now() - interval '14 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '14 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_3, OPERATOR5_USER, 400, 420, now() - interval '13 days', FINANCE_USER, now() - interval '12 days', NULL, NULL, 'APPROVED', now() - interval '12 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_4, OPERATOR_USER, 500, 540, now() - interval '13 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '13 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_1, OPERATOR2_USER, 210, 220, now() - interval '12 days', FINANCE_USER, now() - interval '11 days', NULL, NULL, 'APPROVED', now() - interval '11 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_2, OPERATOR3_USER, 310, 330, now() - interval '12 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '12 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_3, OPERATOR4_USER, 410, 430, now() - interval '11 days', FINANCE_USER, now() - interval '10 days', NULL, NULL, 'APPROVED', now() - interval '10 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_4, OPERATOR5_USER, 510, 550, now() - interval '11 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '11 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_1, OPERATOR_USER, 220, 230, now() - interval '10 days', FINANCE_USER, now() - interval '9 days', NULL, NULL, 'APPROVED', now() - interval '9 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_2, OPERATOR2_USER, 320, 340, now() - interval '10 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '10 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_3, OPERATOR3_USER, 420, 440, now() - interval '9 days', FINANCE_USER, now() - interval '8 days', NULL, NULL, 'APPROVED', now() - interval '8 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_4, OPERATOR4_USER, 520, 560, now() - interval '9 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '9 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_1, OPERATOR5_USER, 230, 240, now() - interval '8 days', FINANCE_USER, now() - interval '7 days', NULL, NULL, 'APPROVED', now() - interval '7 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_2, OPERATOR_USER, 330, 350, now() - interval '8 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '8 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_3, OPERATOR2_USER, 430, 450, now() - interval '7 days', FINANCE_USER, now() - interval '6 days', NULL, NULL, 'APPROVED', now() - interval '6 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_4, OPERATOR3_USER, 530, 570, now() - interval '7 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '7 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_1, OPERATOR4_USER, 240, 250, now() - interval '6 days', FINANCE_USER, now() - interval '5 days', NULL, NULL, 'APPROVED', now() - interval '5 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_2, OPERATOR5_USER, 340, 360, now() - interval '6 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '6 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_3, OPERATOR_USER, 440, 460, now() - interval '5 days', FINANCE_USER, now() - interval '4 days', NULL, NULL, 'APPROVED', now() - interval '4 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_4, OPERATOR2_USER, 540, 580, now() - interval '5 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '5 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_1, OPERATOR3_USER, 250, 260, now() - interval '4 days', FINANCE_USER, now() - interval '3 days', NULL, NULL, 'APPROVED', now() - interval '3 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_2, OPERATOR4_USER, 350, 370, now() - interval '4 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '4 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_3, OPERATOR5_USER, 450, 470, now() - interval '3 days', FINANCE_USER, now() - interval '2 days', NULL, NULL, 'APPROVED', now() - interval '2 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_4, OPERATOR_USER, 550, 590, now() - interval '3 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '3 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_1, OPERATOR2_USER, 260, 270, now() - interval '2 days', FINANCE_USER, now() - interval '1 days', NULL, NULL, 'APPROVED', now() - interval '1 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_2, OPERATOR3_USER, 360, 380, now() - interval '2 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '2 days'),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_3, OPERATOR4_USER, 460, 480, now() - interval '1 days', FINANCE_USER, now(), NULL, NULL, 'APPROVED', now()),
        (uuid_generate_v4(), CUSTOMER_1, SERVICE_4, OPERATOR5_USER, 560, 600, now() - interval '1 days', NULL, NULL, NULL, NULL, 'PENDING', now() - interval '1 days');

END $$;
