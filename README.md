# universalsdk

### Project Structure

  1. Model - Same as Entities, A model in Go is a set of data structures and functions, will store any Object’s Struct and its method. Example : Ledger, Account etc.
  2. Services - This layer contains application specific business rules. It encapsulates and implements all of the use cases of the system.
  4. Controller - This layer is a set of adapters that convert data from the format most convenient for the services and models, to the format most convenient for some external interface such as REST API or grpc
  5. Utils  - This layers contains utility functions 
	
	

### RUN using Docker
##### Build Image
docker build -t frankiefinancial/universalsdk:v1.0 -f Dockerfile .
##### Run Image
docker run -p 80:8080 frankiefinancial/universalsdk:v1.0


	

  