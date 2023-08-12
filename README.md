# Address Listener

Address Listener is a Go application that receives webhook data about addresses activity in the blockchain, processes it, and stores it in a MongoDB database. It provides an efficient way to handle incoming data from external sources and store it for further analysis.

## Table of Contents

- [Address Listener](#address-listener)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Configuration](#configuration)
  - [Supported Webhook](#supported-webhook)
  - [Adapting for Other Webhooks](#adapting-for-other-webhooks)

## Features

- Receives webhook data and processes it in real-time.
- Stores processed data in a MongoDB Atlas instance.
- Provides a modular and organized code structure.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/AddressListener.git
   cd AddressListener

2. Build the docker image:

    ```bash
    docker build -t address-listener .

## Usage

1. Create an .env file with the instructions in Configuration.

2. Start the application using Docker:

    ```bash
    docker run -p 8080:8080 address-listener

3. The application will now be running and listening for incoming webhooks on port 8080.

## Configuration

Configuration of the application is done through environment variables. Create a .env file in the root directory with the following content:

    USERNAME=your_mongodb_username
    PASSWORD=your_mongodb_password
    CLUSTER_URL=your_mongodb_cluster_url
    DBNAME = "DB_TEST"
    COLLECTIONNAME = "Collection_TEST"

## Supported Webhook

This webhook is specifically designed to handle the Address Activity Webhook from Alchemy (https://docs.alchemy.com/reference/address-activity-webhook). It processes data according to the provided data model and stores it in the MongoDB database.

## Adapting for Other Webhooks

To adapt this application for other webhooks, you'll need to modify the `models/data_model.go` file according to the data structure of the new webhook. Make sure to update the `WebhookData` struct to match the new payload format.