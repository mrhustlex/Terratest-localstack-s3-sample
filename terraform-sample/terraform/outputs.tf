output "hello_world" {
  value = "hello_world"
}

output "bucket_id" {
    value = "${module.s3_bucket.bucket_id}"
}