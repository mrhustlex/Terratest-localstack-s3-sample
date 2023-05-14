variable "bucket_name" {
  type = string
}

variable "tag_bucket_name" {
  type        = string
  default     = "Tag Bucket"
}

variable "tag_bucket_environment" {
  type        = string
  default     = "Test"
}