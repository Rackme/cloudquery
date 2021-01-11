resource "aws_rds_cluster" "cloudquery" {
  cluster_identifier = "cloudquery"
  engine             = "aurora-mysql"
  engine_version     = "5.7.mysql_aurora.2.07.1"
  engine_mode        = "serverless"
  master_username    = "cloudquery"
  master_password    = random_password.password.result

  scaling_configuration {
    auto_pause               = true
    max_capacity             = 4
    min_capacity             = 1
    seconds_until_auto_pause = 300
    timeout_action           = "ForceApplyCapacityChange"
  }
}

resource "random_password" "password" {
  length           = 16
  special          = true
  override_special = "_%@"
}

resource "aws_ssm_parameter" "cloudquery_master_password" {
  name        = "/cloudquery/database/password/master"
  description = "Master password for cloudquery aurora database"
  type        = "SecureString"
  value       = random_password.password.result
}