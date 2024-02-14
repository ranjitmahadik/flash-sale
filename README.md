# Flash Sale

## Overview
This project implements a very simple but effective flash sale system, allowing users to purchase products at discounted prices for a limited time. The system is designed to be set up using Docker for easy deployment and management.

## Setup
### Prerequisites
- Docker Installed: Ensure that Docker is installed on your system. Docker allows you to package and distribute applications and their dependencies as containers.

### Setup Environment Variables

To configure the necessary environment variables for the project, follow these steps:

1. Rename the example environment file provided (`example.env`) to `.env`. This file will contain the environment variables required for the project.

    ```bash
    mv example.env .env
    ```

   This command renames `example.env` to `.env`, ensuring that the environment variables are properly configured for your local environment.

### Initializing PostgreSQL Instance
1. Use the provided Docker Compose file to set up a PostgreSQL database instance.
2. Run the following command in your terminal:

```
docker compose up -d
```
This command will start the PostgreSQL container in detached mode (-d), allowing it to run in the background.

### Running the Project
1. After initializing the PostgreSQL instance, proceed to run the project.
2. Utilize the provided Makefile to simplify the process.
3. Run the following command in your terminal:

```
make run
```
This command will execute the necessary steps to run the flash sale project.
