variable environment {
  type        = string
  default     = "staging"
  description = "environment to deploy to"
}

variable app_bucket_name {
  type        = string
  default     = "portfolio-data"
  description = "name of bucket for the application"
}

variable branch {
  type        = string
  description = "branch name used for adding a qualifier to the lambda name"
}

variable needsZip {
  type = bool
  default = true
  description = "determines whether to include a zip file on the lambda"
}
