# auth-policy

    1. Role based access control.
    2. JWT based authentication.
    3. JWT based authorization.
    4. JWT based token verification.
    5. JWT based token refresh.

# jwt-auth process

    1. (user -> auth-server )login request.
    2. (auth-server -> user) token in response.
    3. (user -> banking-server) request resource with token.
    4. (banking server -> auth server) verify the token.
    5. (auth-server -> banking-server) token verification response.
    6. (bankng-server -> user) resource response.

# routes

    1. GetAllCustomers: GET /customers
    2. GetCustomer: GET /customers/:id
    3  NewAccount: POST /customers/:id/accounts
    4. NewTransaction: POST customers/:id/account/:account_id


# RBAC

    1. Role: admin -> All.
    2. Role: user -> GetCustomer & NewTransaction.

# Verification Process

    1. Validity of the token(include expiry time and signature).
    2. Verify if the role has access to the resource.
    3. vefify if the resource being accessd by same user.


[//] # ( Path: cmd/auth-server/readme.md


