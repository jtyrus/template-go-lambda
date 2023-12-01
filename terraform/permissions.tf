resource "aws_iam_role" "iam_for_lambda" {
  name = "portfolio-lambda-${var.branch}"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy" "function_logging_policy" {
  name   = "portfolio-lambda-logging-policy-${var.branch}"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        Action : [
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ],
        Effect : "Allow",
        Resource : "arn:aws:logs:*:*:*"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "function_logging_policy_attachment" {
  role = aws_iam_role.iam_for_lambda.id
  policy_arn = aws_iam_policy.function_logging_policy.arn
}

resource "aws_iam_policy" "portfolio_specific" {
  name   = "portfolio-lambda-policy-${var.branch}"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        Action : [
          "dynamodb:GetItem",
        ],
        Effect : "Allow",
        Resource : "arn:aws:dynamodb:*:*:table/Portfolio-*"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "portfolio_specific" {
  role = aws_iam_role.iam_for_lambda.id
  policy_arn = aws_iam_policy.portfolio_specific.arn
}
