# Sybil API Server

This Sybil node API serves as the interface for third-party applications and services to utilize the layer two features of the SYB rollup.

## Base URL
**Base URL:** `http://localhost:8086/v1`

## Endpoints

### Accounts
- **GET** `/accounts/index/:{accountIndex}`
- **GET** `/accounts/address/:{ethAddr}`

### Transaction History
- **GET** `/transactions-history`
    - **Query Parameters:**
        - `EthAddr`
        - `FromEthAddr`
        - `ToEthAddr`
        - `Idx`
        - `FromIdx`
        - `ToIdx`
        - `BatchNum`
        - `TxType`
        - **Pagination:**
            - `FromItem`
            - `Limit`
            - `Order`

### Batches
- **GET** `/batches`
    - **Query Parameters:**
        - `MinBatchNum`
        - `MaxBatchNum`
        - `ForgerAddr`
        - **Pagination:**
            - `FromItem`
            - `Limit`
            - `Order`
- **GET** `/batches/:batchNum`
- **GET** `/full-batches/:batchNum`

## Tests

### Unit Tests
```bash
task test-api
```

### Integration Tests
To run the integration tests, you can use the following services defined in the `docker-compose.yaml`:

1. **Swagger UI Documentation**
   - **Service Name:** `sybil-api-doc`
   - **Port:** `8001`
   - **Description:** Serves the Swagger UI documentation for the API.

2. **Mock API Responses**
   - **Service Name:** `sybil-api-mock`
   - **Port:** `4010`
   - **Description:** Mocks API responses based on the Swagger specification.

3. **Swagger API Editor**
   - **Service Name:** `sybil-api-editor`
   - **Port:** `8002`
   - **Description:** Allows editing of Swagger API specifications.

To start the services, run:
```bash
docker compose up
```