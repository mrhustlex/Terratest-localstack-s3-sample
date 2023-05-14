##Introduction
This README provides instructions for using Terratest to test the creation of an S3 bucket using LocalStack. Terratest is a testing framework for Terraform that allows you to write automated tests to verify that your infrastructure is functioning correctly.

LocalStack is a tool that allows you to run AWS services locally for testing and development purposes. With LocalStack, you can create a local S3 bucket that you can use to test your Terraform configuration.

##Prerequisites
Before you begin, you should make sure that you have the following tools installed:

Terraform
LocalStack
Terratest

##Run Test
1. in <cd ./test>, you can run <go test -v> for testing on terratest
2. in terraform folder, it defines a sample terraform folder so that you can test with <tflocal>
