ec2 := "seanwelch@ec2-54-194-167-80.eu-west-1.compute.amazonaws.com"
key := "~/.ssh/farmec.pem"
remote_db := "/home/seanwelch/farmec-v2/database/database.db"
local_db := "database/database.db"

# SSH into EC2
ssh:
    ssh -i {{key}} {{ec2}}

# Terraform plan
tf-plan:
    set -a && source .env && set +a && terraform -chdir=misc/terraform plan

# Terraform apply (run tf-plan and snapshot EBS volume before running this)
tf-apply:
    set -a && source .env && set +a && terraform -chdir=misc/terraform apply

# Provision EC2 server (run db-pull first to ensure you have a local backup)
provision:
    set -a && source .env && set +a && ansible-playbook -i misc/ansible/inventory.yml misc/ansible/playbook.yml -K

# Copy local database up to EC2
db-push:
    scp -i {{key}} {{local_db}} {{ec2}}:{{remote_db}}

# Copy EC2 database down to local
db-pull:
    scp -i {{key}} {{ec2}}:{{remote_db}} {{local_db}}
