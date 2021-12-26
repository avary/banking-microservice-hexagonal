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

    1. Get all customers: GET /customers
    2. Get customer by id: GET /customers/:id
    3. Create new account: POST /customers/:id/accounts
    4. Make a transaction: POST customers/:id/account/:account_id


# RBAC

    1. Role: admin -> can do all.
    2. Role: user -> Get customer by ID & Make a transaction.


[//] # ( Path: cmd/auth-server/readme.md


