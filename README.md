# flip-tech-test

This is a simple Go HTTP server implementation according to the requirements set out in the [challenge specification](./Backend_Code_challenge.pdf).

The repo includes a simple React frontend, a Go backend API, and a PostgreSQL database. Since the frontend aspect is not examined, it has been kept minimal with just enough functionality to work with the API.

The project is run entirely through Docker, so there are no external dependencies or restrictions. The `.env` files have been committed to the repo to make it easier to run, especially since this is not a production application.

**I've chosen to address requirements 1 and 2**, but I believe my design also partially satisfies requirement 4, as it could easily be extended to integrate with other services.

## Running the project
The project can be run directly through Docker:
```sh
docker compose up -d
```

The frontend will then be exposed at [localhost:3000](localhost:3000) and the backend at [localhost:8080](localhost:8080) to allow for direct requests if necessary for testing.

## Unit Tests
As per the challenge specification, test coverage is not a concern so I've only made one test file, to demonstrate my understanding and approach to unit tests.

I have chosen to test the `GetProducts` function which is the handler for the `/products` endpoint. This allows me to write a complete unit test for a function and also show how I mock dependencies, without needing too many dependencies or test cases, which would make the implementation harder to review. The database operations have been mocked as they are not relevant to the scope of the unit test, which should only cover the logic of that specific package/function.

If we wanted to cover a more extensive flow including database operations, a component test (or integration test) would be more appropriate.

To run unit tests, we can use:
```sh
go test ./... --cover
```
This would run any test files in the repo (assuming we had more than one) and also show the coverage for each package.

## Backend Info
The API is a simple HTTP REST server that interfaces with the frontend. Concurrency scenarios and race conditions are handled through transactions at the database level, and through channels in the actual Go code where necessary. A graceful shutdown has been implemented to prevent data loss or memory leaks through dangling connections.

The backend API includes two endpoints:

### `/products`
Retrieves a list of products available for purchase. An error message is sent if the request fails.

### `/purchase`
Facilitates a purchase. The request body must contain a JSON object like below:
```json
{
    "items": [
        {
            "sku": "234234",
            "quantity": 1
        }
    ]
}
```
The total price, and a list of updated items, is returned upon a successful transaction (used by the frontend to update the catalogue state). The assumption is that the price would be used to process payment, and the list of updated items used by the frontend to update the catalogue. This has the limitation that another user looking at the catalogue before the purchase is made will still see the old values, but I determined that is outside of the scope of this task, as it could be handled by the frontend.

If the transaction fails, any database changes are rolled back, and the appropriate error is returned.

## Final Comments
Thank you for taking the time to review this submission. Please feel free to reach out with any questions or comments, and I'd be happy to discuss it further over a call if necessary.
