##Dynamodb table

resource "aws_dynamodb_table" "sampleapi-table" {
  name             = "sampleapi-table"
  hash_key         = "filename"
  billing_mode     = "PAY_PER_REQUEST"

  attribute {
    name = "filename"
    type = "S"
  }
}