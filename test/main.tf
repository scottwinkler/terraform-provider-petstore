provider "petstore" {
    address = "https://g6pny6dke9.execute-api.us-west-2.amazonaws.com/petstore"
}

resource "petstore_pet" "my_pet" {
    name = "mittens"
    species = "cat"
    age = 3
}