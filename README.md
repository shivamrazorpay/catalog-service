
# Catalog Service

This Service holds the data for various services.

> **_NOTE:_** The Data of various services is stored in internal/data (in memory store) and 10 more random services are populated at run time (check boot folder). In Real System we can replace the same file with a db layer and use GORM or any other ORM to interact with the DB.


## Get Started

Run the below command in your terminal to clone the repo in your local.

```bash
  git clone git@github.com:shivamrazorpay/service-catalog.git
  cd service-catalog
```

To install the deps
```bash
    go mod tidy
```

Start the server on localhost:8080
```bash
    go run main.go
```

Run Unit Tests
```bash
    make unit-test
```

Get Code Coverage
```bash
    make coverage
```

## API Reference

> **_NOTE:_**  Check env folder for basic auth creds.

#### Get service by serviceId

```http
  GET /services/{serviceId}
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `serviceId` | `string` | **Required**. serviceId |

#### Get all versions of a service

```http
  GET /services/{serviceId}/versions
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `serviceId` | `string` | **Required**. serviceId |

#### List Services with Sorting, Pagination, and Search

```http
  GET /services?sortBy={column.order}&limit={limit}&offset={offset}&search={column.value}
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `sortBy` | `string` | **Not Required**. Sort the services by there column in Ascending(asc) or Descending(desc) order. input is the format: column.order|
| `limit` | `string` | **Not Required**. limit the service response array length. Default is 10. Max is 20|
| `offset` | `string` | **Not Required**. offset the service response array|
| `search` | `string` | **Not Required**. search the services by there column with the input keyword. If Matches if the keyword is present in the field (not exact Match). input is the format: column.value |


#### Create a New Service

```http
  POST /services
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `name` | `string` | **Required**. The name of the service|
| `Description` | `string` | **Required**. A brief description of the service. Description Length Can be max 100 char.|
| `LatestVersion` | `string` | **Required**. The latest version of the service|


#### Update a Service

```http
  POST /services/{serviceId}
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `name` | `string` | **Not Required**. The name of the service|
| `Description` | `string` | **Not Required**. A brief description of the service|
| `LatestVersion` | `string` | **Not Required**. The latest version of the service. While Updating if the latest version and the updated version does not match, the latest version is updated and the previous version is added in version list.|



## Example

**LocalHost Curl to add Service**

Request
``` bash
curl --location 'http://localhost:8080/services' \
--header 'Content-Type: application/json' \
--data '{
    "name":"Test Service",
    "description":"This is a Test Service.",
    "latest_version":"1.0.0"
}'
```
Response
``` json
{
    "id": "qPfAr3fuFh",
    "name": "Test Service",
    "description": "This is a Test Service.",
    "latest_version": "1.0.0",
    "versions": [
        "1.0.0"
    ],
    "created_at": 1730905017,
    "updated_at": 1730905017
}
```

**LocalHost Curl to List Versions**

Request
``` bash
curl --location 'http://localhost:8080/services/3xD2hfOxn7/versions' \
--data ''
```
Response
``` json
{
    "versions": [
        "1.0.0",
        "1.0.1"
    ],
    "count": 2
}
```

**LocalHost Curl to Filter + Pagination**

Request
``` bash
curl --location --request GET 'http://localhost:8080/services' \
--header 'Content-Type: application/json' \
--data '{
    "pagination":{
        "limit":4,
        "offset":0
    },
    "search":{
        "column":"name",
        "value":"Service"
    }
}'
```
Response
``` json
[
    {
        "id": "BbItY3V9rV",
        "name": "Service BbItY3V9rV",
        "description": "lorem ipsum dolor sit amet consectetur adipiscing elit efficitur habitasse tellus risus vitae nam neque lectus hendrerit lacinia eget sollicitudin",
        "latest_version": "1.0.1",
        "versions": [
            "1.0.0",
            "1.0.1"
        ],
        "created_at": 1730916394,
        "updated_at": 1730916394
    },
    {
        "id": "6Zt5mpdDsQ",
        "name": "Service 6Zt5mpdDsQ",
        "description": "lorem ipsum dolor sit amet consectetur adipiscing elit maecenas libero tincidunt fringilla pretium lectus volutpat tempor neque suscipit curabitur aliquam",
        "latest_version": "1.0.1",
        "versions": [
            "1.0.0",
            "1.0.1"
        ],
        "created_at": 1730916394,
        "updated_at": 1730916394
    },
    {
        "id": "ri-AMYrC5o",
        "name": "Service ri-AMYrC5o",
        "description": "lorem ipsum dolor sit amet consectetur adipiscing elit semper potenti ultricies mollis pharetra torquent dignissim placerat aptent sagittis nascetur ornare",
        "latest_version": "1.0.1",
        "versions": [
            "1.0.0",
            "1.0.1"
        ],
        "created_at": 1730916394,
        "updated_at": 1730916394
    },
    {
        "id": "HG9nl6Sx4y",
        "name": "Service HG9nl6Sx4y",
        "description": "lorem ipsum dolor sit amet consectetur adipiscing elit iaculis platea mattis nascetur euismod enim dui hac magnis potenti praesent parturient",
        "latest_version": "1.0.1",
        "versions": [
            "1.0.0",
            "1.0.1"
        ],
        "created_at": 1730916394,
        "updated_at": 1730916394
    }
]
```

