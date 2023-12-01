data "archive_file" "zip" {
  count       = var.needsZip ? 1 : 0
  type        = "zip"
  source_file = "../bin/portfolio_x86"
  output_path = "../bin/portfolio.zip"
}

resource "aws_lambda_function" "portfolio" {
  filename      = "../bin/portfolio.zip"
  function_name = "portfolio-api-${var.branch}"
  role          = aws_iam_role.iam_for_lambda.arn
  handler       = "portfolio_x86"

  source_code_hash = length(data.archive_file.zip) == 1 ? data.archive_file.zip[0].output_base64sha256 : null

  architectures = ["x86_64"]
  runtime       = "go1.x"
  memory_size   = 128
  timeout       = 10

  environment {
    variables = {
      BRANCH_NAME = var.branch
      TABLE_NAME  = aws_dynamodb_table.portfolio.name
    }
  }
}

resource "aws_dynamodb_table" "portfolio" {
  name         = "Portfolio-${var.branch}"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "Id"
  range_key    = "Version"

  attribute {
    name = "Id"
    type = "S"
  }

  attribute {
    name = "Version"
    type = "S"
  }

  tags = {
    Name        = "dynamodb-table-1"
    Environment = "staging"
  }
}
