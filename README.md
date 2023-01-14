# Covid-19 GraphQL API

This API serves as a graphql server, that returns data fetched from the covid cases dataset of Buenos Aires ,Argentina.

It has two main endpoints:
- `/query` is the endpoint to perform GraphQL queries to the database
- `/update_dataset` the dataset is being updated weekly, so we must define an endpoint to refresh that dataset. This endpoint should be scheduled to run once a week (scheduled job)
