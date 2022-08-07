## MDS Test
### Summary
A system that helps in managing products for a hypothetical ecommerce  platform. A product should have a unique SKU and could be commercialized in
multiple countries. Each product can then have different stock per country.

### Requirements 
  1. Provide a products API
   
            a). Get a product by SKU
            b). Consume stock from a product. 
                - Should validate if the stock requested is available first, and then decrease it.
  
  2. Provide an API that allows a bulk update of products from a CSV.
        
            a). For each CSV line, the stock update could be positive or negative
            b). If a product doesn’t exist, it should be created.


### Implementation Approach
Two services have been implemented; [products](https://github.com/samuelemwangi/jumia-mds-test/tree/master/services/products) and [bulkupdates](https://github.com/samuelemwangi/jumia-mds-test/tree/master/services/bulkupdates) . The ***products*** service  exposes endpoints to **get a product by SKU**, **consume stock from a product**, **upload a file**, among other endpoints. The ***bulkupdates*** service handles processing of the uploaded files. Once a file is uploaded, the ***products*** service sends a an ***uploadId*** message to a ***kafka topic***. The ***bulkupdates*** service then consumes the ***uploadId*** message and processes the file. The ***bulkupdates*** service also exposes a **get processing status** endpoint to get the status of the processing.



### Running the Services
#### Locally
 
1. `git clone` the source code repository
2. Ensure [golang](https://go.dev/doc/install) is installed on your machine. 
3. Ensure you have a [kafka](https://kafka.apache.org/documentation/#quickstart) and a [mysql](https://www.mysql.com/products/workbench/) database running. (you can run  `docker-compose -f docker-compose-cluster.yml up -d` to start them locally if you have [docker](https://docs.docker.com/engine/install/) installed)
4.  Rename the `example.env` file to `env.env`
5.  Update the environment variables in the `env.env` file with the appropriate values.
6.  `cd` to the `services/products` directory and run `go run main.go`
7.  `cd` to the `services/bulkupdates` directory and run `go run main.go`
8.  Import `postman-collection.json` on [Postman](https://www.postman.com/downloads/) and test the endpoints as follows: 
   
 
        POST:  {{products_base_url}}/upload   - to upload a csv file
        GET :  {{bulk_updates_base_url}}/upload-status/{{upload_id}}   - to get processing status
        GET :  {{products_base_url}}/products/{{sku}}   -  to get product by SKU
        GET :  {{products_base_url}}/products   -  to get all products
        GET :  {{products_base_url}}/countries   -  to get all countries
        POST:  {{products_base_url}}/products/{{sku}}/consume   -  to consume stock from a product

    **Note:** *products_base_url = localhost:8085/api/v1  and bulk_updates_base_url = localhost:8086/api/v1*


#### Using Docker
1. `git clone` the source code repository
2. Ensure you have [docker](https://docs.docker.com/engine/install/) installed
3. Run `docker compose up -d` to start the services.
4. Import `postman-collection.json` on [Postman](https://www.postman.com/downloads/) and test the endpoints as follows: 
        
        POST:  {{gateway_base_url}}/upload   -  to upload a csv file
        GET :  {{gateway_base_url}}/upload-status/{{upload_id}}   -  to get processing status
        GET :  {{gateway_base_url}}/products/{{sku}}   -  to get product by SKU
        GET :  {{gateway_base_url}}/products   -  to get all products
        GET :  {{gateway_base_url}}/countries   -  to get all countries
        POST:  {{gateway_base_url}}/products/{{sku}}/consume   -  to consume stock from a product

    **Note:** *gateway_base_url* = localhost:8088/api/v1 as exposed by the [krekend-ce](https://www.krakend.io/) gateway service
