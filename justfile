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
[confirm]
tf-apply:
    set -a && source .env && set +a && terraform -chdir=misc/terraform apply

# Provision EC2 server (run db-pull first to ensure you have a local backup)
[confirm]
provision:
    set -a && source .env && set +a && ansible-playbook -i misc/ansible/inventory.yml misc/ansible/playbook.yml -K

# Copy local database up to EC2
[confirm]
db-push:
    scp -i {{key}} {{local_db}} {{ec2}}:{{remote_db}}

# Copy EC2 database down to local
db-pull:
    scp -i {{key}} {{ec2}}:{{remote_db}} {{local_db}}

# Show recent errors from Docker container logs
logs:
    ssh -i {{key}} {{ec2}} "docker logs \$(docker ps -q) 2>&1 | grep -E '(Error|Exception|Traceback|500)' | tail -50"

# Follow live Docker container logs
logs-live:
    ssh -i {{key}} {{ec2}} "docker logs \$(docker ps -q) -f 2>&1"
