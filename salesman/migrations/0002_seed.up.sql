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

        VIEW_ALL_PRODUCTS_LIST UUID;
        VIEW_ALL_PRODUCTS UUID;
        CREATE_ALL_PRODUCTS UUID;
        UPDATE_ALL_PRODUCTS UUID;
        DELETE_ALL_PRODUCTS UUID;
        
        VIEW_ALL_CUSTOMERS_LIST UUID;
        VIEW_ALL_CUSTOMERS UUID;
        CREATE_ALL_CUSTOMERS UUID;
        UPDATE_ALL_CUSTOMERS UUID;
        DELETE_ALL_CUSTOMERS UUID;

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
        CUSTOMER_2 UUID;
        CUSTOMER_3 UUID;
        CUSTOMER_4 UUID;

        COMPANY_USER_1 UUID;
        COMPANY_USER_2 UUID;
        COMPANY_USER_3 UUID;
        COMPANY_USER_4 UUID;

        -- PRODUCTS
        PRODUCT_1 UUID;
        PRODUCT_2 UUID;
        PRODUCT_3 UUID;
        PRODUCT_4 UUID;

        -- COMPANY USER PRODUCTS
        COMPANY_USER_PRODUCT_1 UUID;
        COMPANY_USER_PRODUCT_2 UUID;
        COMPANY_USER_PRODUCT_3 UUID;
        COMPANY_USER_PRODUCT_4 UUID;
        
        -- SALES LEDGER
        SALES_LEDGER_1 UUID;
        SALES_LEDGER_2 UUID;
        SALES_LEDGER_3 UUID;

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
        VIEW_ALL_PRODUCTS_LIST = uuid_generate_v4();
        VIEW_ALL_PRODUCTS = uuid_generate_v4();
        CREATE_ALL_PRODUCTS = uuid_generate_v4();
        UPDATE_ALL_PRODUCTS = uuid_generate_v4();
        DELETE_ALL_PRODUCTS = uuid_generate_v4();
        VIEW_ALL_CUSTOMERS_LIST = uuid_generate_v4();
        VIEW_ALL_CUSTOMERS = uuid_generate_v4();
        CREATE_ALL_CUSTOMERS = uuid_generate_v4();
        UPDATE_ALL_CUSTOMERS = uuid_generate_v4();
        DELETE_ALL_CUSTOMERS = uuid_generate_v4();
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

        -- PRODUCTS
        PRODUCT_1 = uuid_generate_v4();
        PRODUCT_2 = uuid_generate_v4();
        PRODUCT_3 = uuid_generate_v4();
        PRODUCT_4 = uuid_generate_v4();

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
        CUSTOMER_2 = uuid_generate_v4();
        CUSTOMER_3 = uuid_generate_v4();
        CUSTOMER_4 = uuid_generate_v4();

        COMPANY_USER_1 = uuid_generate_v4();
        COMPANY_USER_2 = uuid_generate_v4();
        COMPANY_USER_3 = uuid_generate_v4();
        COMPANY_USER_4 = uuid_generate_v4();

        COMPANY_USER_PRODUCT_1 = uuid_generate_v4();
        COMPANY_USER_PRODUCT_2 = uuid_generate_v4();
        COMPANY_USER_PRODUCT_3 = uuid_generate_v4();
        COMPANY_USER_PRODUCT_4 = uuid_generate_v4();
        

        -- SALES LEDGER
        SALES_LEDGER_1 = uuid_generate_v4();
        SALES_LEDGER_2 = uuid_generate_v4();
        SALES_LEDGER_3 = uuid_generate_v4();

-- PERMISSIONS
INSERT INTO Permissions (id, action, created_at) VALUES 
        
        -- INVOICES (SALES LEDGER)
        (CREATE_INVOICE, 'CREATE_INVOICE', now()),
        (VIEW_OWN_INVOICES, 'VIEW_OWN_INVOICES', now()),
        (VIEW_OWN_INVOICES_LIST, 'VIEW_OWN_INVOICES_LIST', now()),
        (VIEW_ALL_INVOICES, 'VIEW_ALL_INVOICES', now()),
        (VIEW_ALL_INVOICES_LIST, 'VIEW_ALL_INVOICES_LIST', now()),
        (DELETE_OWN_INVOICES, 'DELETE_OWN_INVOICES', now()),
        (DELETE_ALL_INVOICES, 'DELETE_ALL_INVOICES', now()),
        (APPROVE_ALL_INVOICES, 'APPROVE_ALL_INVOICES', now()),
        (REJECT_ALL_INVOICES, 'REJECT_ALL_INVOICES', now()),
        (CANCEL_OWN_INVOICES, 'CANCEL_OWN_INVOICES', now()),
        (CANCEL_ALL_INVOICES, 'CANCEL_ALL_INVOICES', now()),

        -- DASHBOARD
        (VIEW_DASHBOARD, 'VIEW_DASHBOARD', now()),
        (VIEW_DASHBOARD_CARD_ALL_TODAY_APPROVED_COUNT, 'VIEW_DASHBOARD_CARD_ALL_TODAY_APPROVED_COUNT', now()),
        (VIEW_DASHBOARD_CARD_ALL_TODAY_APPROVED_AMMOUNT, 'VIEW_DASHBOARD_CARD_ALL_TODAY_APPROVED_AMMOUNT', now()),
        (VIEW_DASHBOARD_CARD_ALL_TODAY_UNAPPROVED_COUNT, 'VIEW_DASHBOARD_CARD_ALL_TODAY_UNAPPROVED_COUNT', now()),
        (VIEW_DASHBOARD_CARD_USER_TODAY_APPROVED_COUNT, 'VIEW_DASHBOARD_CARD_USER_TODAY_APPROVED_COUNT', now()),
        (VIEW_DASHBOARD_CARD_USER_TODAY_APPROVED_AMMOUNT, 'VIEW_DASHBOARD_CARD_USER_TODAY_APPROVED_AMMOUNT', now()),
        (VIEW_DASHBOARD_CARD_USER_TODAY_UNAPPROVED_COUNT, 'VIEW_DASHBOARD_CARD_USER_TODAY_UNAPPROVED_COUNT', now()),
        (VIEW_TODAY_MVP_OPERATOR, 'VIEW_TODAY_MVP_OPERATOR', now()),
        (VIEW_SALES_MVP_LIST, 'VIEW_SALES_MVP_LIST', now()),
        (VIEW_ALL_SALES_HISTORY, 'VIEW_ALL_SALES_HISTORY', now()),
        (VIEW_USER_SALES_HISTORY, 'VIEW_USER_SALES_HISTORY', now()),

        -- USERS 
        (VIEW_ALL_USERS_LIST, 'VIEW_ALL_USERS_LIST', now()),
        (VIEW_ALL_USERS, 'VIEW_ALL_USERS', now()),
        (CREATE_ALL_USERS, 'CREATE_ALL_USERS', now()),
        (UPDATE_ALL_USERS, 'UPDATE_ALL_USERS', now()),
        (DELETE_ALL_USERS, 'DELETE_ALL_USERS', now()),
        (VIEW_USER_ROLES, 'VIEW_USER_ROLES', now()),
        (ADD_USER_ROLE, 'ADD_USER_ROLE', now()),
        (DELETE_USER_ROLE, 'DELETE_USER_ROLE', now()),

        -- CUSTOMERS
        (VIEW_ALL_CUSTOMERS_LIST, 'VIEW_ALL_CUSTOMERS_LIST', now()),
        (VIEW_ALL_CUSTOMERS, 'VIEW_ALL_CUSTOMERS', now()),
        (CREATE_ALL_CUSTOMERS, 'CREATE_ALL_CUSTOMERS', now()),
        (UPDATE_ALL_CUSTOMERS, 'UPDATE_ALL_CUSTOMERS', now()),
        (DELETE_ALL_CUSTOMERS, 'DELETE_ALL_CUSTOMERS', now()),

        -- PRODUCTS
        (VIEW_ALL_PRODUCTS_LIST, 'VIEW_ALL_PRODUCTS_LIST', now()),
        (VIEW_ALL_PRODUCTS, 'VIEW_ALL_PRODUCTS', now()),
        (CREATE_ALL_PRODUCTS, 'CREATE_ALL_PRODUCTS', now()),
        (UPDATE_ALL_PRODUCTS, 'UPDATE_ALL_PRODUCTS', now()),
        (DELETE_ALL_PRODUCTS, 'DELETE_ALL_PRODUCTS', now()),
        
        -- SERVICES
        (VIEW_ALL_SERVICES_LIST, 'VIEW_ALL_SERVICES_LIST', now()),
        (VIEW_ALL_SERVICES, 'VIEW_ALL_SERVICES', now()),
        (CREATE_ALL_SERVICES, 'CREATE_ALL_SERVICES', now()),
        (UPDATE_ALL_SERVICES, 'UPDATE_ALL_SERVICES', now()),
        (DELETE_ALL_SERVICES, 'DELETE_ALL_SERVICES', now()),

        -- COMPANIES
        (VIEW_ALL_COMPANIES_LIST, 'VIEW_ALL_COMPANIES_LIST', now()),
        (VIEW_ALL_COMPANIES, 'VIEW_ALL_COMPANIES', now()),
        (CREATE_ALL_COMPANIES, 'CREATE_ALL_COMPANIES', now()),
        (UPDATE_ALL_COMPANIES, 'UPDATE_ALL_COMPANIES', now()),
        (DELETE_ALL_COMPANIES, 'DELETE_ALL_COMPANIES', now()),

        -- ROLES
        (VIEW_ALL_ROLES_LIST, 'VIEW_ALL_ROLES_LIST', now()),
        (VIEW_ALL_ROLES, 'VIEW_ALL_ROLES', now()),
        (CREATE_ALL_ROLES, 'CREATE_ALL_ROLES', now()),
        (UPDATE_ALL_ROLES, 'UPDATE_ALL_ROLES', now()),
        (DELETE_ALL_ROLES, 'DELETE_ALL_ROLES', now()),
        (VIEW_ROLE_PERMISSIONS, 'VIEW_ROLE_PERMISSIONS', now()),
        (ADD_ROLE_PERMISSION, 'ADD_ROLE_PERMISSION', now()),
        (DELETE_ROLE_PERMISSION, 'DELETE_ROLE_PERMISSION', now()),

        -- LOGS
        (VIEW_ALL_LOGS, 'VIEW_ALL_LOGS', now());

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
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_ALL_CUSTOMERS_LIST, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_ALL_CUSTOMERS, now()),
        (uuid_generate_v4(), ADMINISTRATOR, CREATE_ALL_CUSTOMERS, now()),
        (uuid_generate_v4(), ADMINISTRATOR, UPDATE_ALL_CUSTOMERS, now()),
        (uuid_generate_v4(), ADMINISTRATOR, DELETE_ALL_CUSTOMERS, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_ALL_PRODUCTS_LIST, now()),
        (uuid_generate_v4(), ADMINISTRATOR, VIEW_ALL_PRODUCTS, now()),
        (uuid_generate_v4(), ADMINISTRATOR, CREATE_ALL_PRODUCTS, now()),
        (uuid_generate_v4(), ADMINISTRATOR, UPDATE_ALL_PRODUCTS, now()),
        (uuid_generate_v4(), ADMINISTRATOR, DELETE_ALL_PRODUCTS, now()),
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

-- INSERT INTO Companies (id, title, created_at) VALUES 
--         (COMPANY_1, 'Company 1', now()),
--         (COMPANY_2, 'Company 2', now()),
--         (COMPANY_3, 'Company 3', now()),
--         (COMPANY_4, 'Company 4', now());

-- INSERT INTO Products (id, company_id, title, has_commission, created_at) VALUES 
--         (PRODUCT_1, COMPANY_1, 'Product 1', true, now()),
--         (PRODUCT_2, COMPANY_2, 'Product 2', true, now()),
--         (PRODUCT_3, COMPANY_3, 'Product 3', true, now()),
--         (PRODUCT_4, COMPANY_4, 'Product 4', false, now());


-- INSERT INTO Services (id, product_id, title, price, has_commission, created_at) VALUES
--         (SERVICE_1, PRODUCT_1, 'Service 1', 100, true, now()),
--         (SERVICE_2, PRODUCT_2, 'Service 2', 200, true, now()),
--         (SERVICE_3, PRODUCT_3, 'Service 3', 300, true, now()),
--         (SERVICE_4, PRODUCT_4, 'Service 4', 400, false, now());

INSERT INTO Users (id, first_name,last_name, NSID, birthdate, email, password, created_at) VALUES 
        (ADMINISTRATOR_USER, 'admin@sailsman', '', '1234567890', '2000-08-15', 'admin@sailsman.ir', 'pbkdf2_sha256$1000000$BV8vF1TzUL8vxeSONtgHP4$zB+T8eKgL3xN/OYHVjUuQgVW3ItRVdDLVZXWY+UYkO4=', now());
        -- (OPERATOR_USER, 'Operator', 'User', '1234567891', '2000-01-01', 'operator@example.com', 'pbkdf2_sha256$1000000$BV8vF1TzUL8vxeSONtgHP4$zB+T8eKgL3xN/OYHVjUuQgVW3ItRVdDLVZXWY+UYkO4=', now()),
        -- (OPERATOR2_USER, 'Operator2', 'User', '1234567894', '2000-01-01', 'operator2@example.com', 'pbkdf2_sha256$1000000$BV8vF1TzUL8vxeSONtgHP4$zB+T8eKgL3xN/OYHVjUuQgVW3ItRVdDLVZXWY+UYkO4=', now()),
        -- (OPERATOR3_USER, 'Operator3', 'User', '1234567895', '2000-01-01', 'operator3@example.com', 'pbkdf2_sha256$1000000$BV8vF1TzUL8vxeSONtgHP4$zB+T8eKgL3xN/OYHVjUuQgVW3ItRVdDLVZXWY+UYkO4=', now()),
        -- (OPERATOR4_USER, 'Operator4', 'User', '1234567896', '2000-01-01', 'operator4@example.com', 'pbkdf2_sha256$1000000$BV8vF1TzUL8vxeSONtgHP4$zB+T8eKgL3xN/OYHVjUuQgVW3ItRVdDLVZXWY+UYkO4=', now()),
        -- (OPERATOR5_USER, 'Operator5', 'User', '1234567897', '2000-01-01', 'operator5@example.com', 'pbkdf2_sha256$1000000$BV8vF1TzUL8vxeSONtgHP4$zB+T8eKgL3xN/OYHVjUuQgVW3ItRVdDLVZXWY+UYkO4=', now()),
        -- (FINANCE_USER, 'Finance', 'User', '1234567892', '2000-01-01', 'finance@example.com', 'pbkdf2_sha256$1000000$BV8vF1TzUL8vxeSONtgHP4$zB+T8eKgL3xN/OYHVjUuQgVW3ItRVdDLVZXWY+UYkO4=', now()),
        -- (USER_1, 'User', 'User', '1234567893', '2000-01-01', 'user@example.com', 'pbkdf2_sha256$1000000$BV8vF1TzUL8vxeSONtgHP4$zB+T8eKgL3xN/OYHVjUuQgVW3ItRVdDLVZXWY+UYkO4=', now()),
        
        -- (CUSTOMER_1, 'Customer 1', 'User', '1234564213', '2000-01-01', 'customer1@example.com', 'pbkdf2_sha256$1000000$BV8vF1TzUL8vxeSONtgHP4$zB+T8eKgL3xN/OYHVjUuQgVW3ItRVdDLVZXWY+UYkO4=', now()),
        -- (CUSTOMER_2, 'Customer 2', 'User', '1234564521', '2000-01-01', 'customer2@example.com', 'pbkdf2_sha256$1000000$BV8vF1TzUL8vxeSONtgHP4$zB+T8eKgL3xN/OYHVjUuQgVW3ItRVdDLVZXWY+UYkO4=', now()),
        -- (CUSTOMER_3, 'Customer 3', 'User', '1234564685', '2000-01-01', 'customer3@example.com', 'pbkdf2_sha256$1000000$BV8vF1TzUL8vxeSONtgHP4$zB+T8eKgL3xN/OYHVjUuQgVW3ItRVdDLVZXWY+UYkO4=', now()),
        -- (CUSTOMER_4, 'Customer 4', 'User', '1234564745', '2000-01-01', 'customer4@example.com', 'pbkdf2_sha256$1000000$BV8vF1TzUL8vxeSONtgHP4$zB+T8eKgL3xN/OYHVjUuQgVW3ItRVdDLVZXWY+UYkO4=', now());

-- INSERT INTO CompanyUsers (id, company_id, user_id, subscription_no, created_at) VALUES 
--         (COMPANY_USER_1, COMPANY_1, CUSTOMER_1, '1234567890', now()),
--         (COMPANY_USER_2, COMPANY_2, CUSTOMER_2, '1234567891', now()),
--         (COMPANY_USER_3, COMPANY_3, CUSTOMER_3, '1234567892', now()),
--         (COMPANY_USER_4, COMPANY_4, CUSTOMER_4, '1234567893', now());

-- INSERT INTO CompanyUserProducts (id, company_user_id, product_id, is_active, number, created_at) VALUES 
--         (COMPANY_USER_PRODUCT_1, COMPANY_USER_1, PRODUCT_1, true, '1234567890', now()),
--         (COMPANY_USER_PRODUCT_2, COMPANY_USER_2, PRODUCT_2, true, '1234567891', now()),
--         (COMPANY_USER_PRODUCT_3, COMPANY_USER_3, PRODUCT_3, true, '1234567892', now()),
--         (COMPANY_USER_PRODUCT_4, COMPANY_USER_4, PRODUCT_4, false, '1234567893', now());

INSERT INTO UserRoles (id, user_id, role_id, created_at) VALUES 
        (uuid_generate_v4(), ADMINISTRATOR_USER, ADMINISTRATOR, now());
--         (uuid_generate_v4(), OPERATOR_USER, OPERATOR, now()),
--         (uuid_generate_v4(), OPERATOR2_USER, OPERATOR, now()),
--         (uuid_generate_v4(), OPERATOR3_USER, OPERATOR, now()),
--         (uuid_generate_v4(), OPERATOR4_USER, OPERATOR, now()),
--         (uuid_generate_v4(), FINANCE_USER, FINANCE, now()),
--         (uuid_generate_v4(), USER_1, OPERATOR, now()),
        
--         (uuid_generate_v4(), CUSTOMER_1, CUSTOMER, now()),
--         (uuid_generate_v4(), CUSTOMER_2, CUSTOMER, now()),
--         (uuid_generate_v4(), CUSTOMER_3, CUSTOMER, now()),
--         (uuid_generate_v4(), CUSTOMER_4, CUSTOMER, now());

-- INSERT INTO SalesLedger (id, customer_product_id, created_by, marketer_id, status, payment_method, TRN, workflow_history, created_at) VALUES 
--         (SALES_LEDGER_1, COMPANY_USER_PRODUCT_1, OPERATOR_USER, OPERATOR_USER, 'APPROVED', 'C2C', '1234567890', '[{"step": "create and send to finance", "action_at": "2025-07-29T05:57:38.722164931Z", "comment": "", "actor_id": "00000000-0000-0000-0000-000000000036", "actor_name": "Initiator"}]', now()),
--         (SALES_LEDGER_2, COMPANY_USER_PRODUCT_2, OPERATOR_USER, OPERATOR_USER, 'APPROVED', 'C2C', '1234567891', '[{"step": "create and send to finance", "action_at": "2025-07-29T05:57:38.722164931Z", "comment": "", "actor_id": "00000000-0000-0000-0000-000000000036", "actor_name": "Initiator"}]', now()),
--         (SALES_LEDGER_3, COMPANY_USER_PRODUCT_3, OPERATOR_USER, OPERATOR_USER, 'APPROVED', 'C2C', '1234567892', '[{"step": "create and send to finance", "action_at": "2025-07-29T05:57:38.722164931Z", "comment": "", "actor_id": "00000000-0000-0000-0000-000000000036", "actor_name": "Initiator"}]', now());

-- INSERT INTO LedgerServiceItems (id, ledger_id, service_id, price, created_at) VALUES 
--         (uuid_generate_v4(), SALES_LEDGER_1, SERVICE_1, 100, now()),
--         (uuid_generate_v4(), SALES_LEDGER_1, SERVICE_2, 200, now()),
--         (uuid_generate_v4(), SALES_LEDGER_1, SERVICE_3, 300, now()),
--         (uuid_generate_v4(), SALES_LEDGER_2, SERVICE_1, 100, now()),
--         (uuid_generate_v4(), SALES_LEDGER_2, SERVICE_2, 200, now()),
--         (uuid_generate_v4(), SALES_LEDGER_2, SERVICE_3, 300, now()),
--         (uuid_generate_v4(), SALES_LEDGER_3, SERVICE_1, 100, now()),
--         (uuid_generate_v4(), SALES_LEDGER_3, SERVICE_2, 200, now()),
--         (uuid_generate_v4(), SALES_LEDGER_3, SERVICE_3, 300, now());

END $$;
