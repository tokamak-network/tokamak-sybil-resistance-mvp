import os
import subprocess
import requests # will be installed during the github action run

# Environment variables
GIT_TOKEN = os.getenv('GIT_TOKEN')
SLACK_WEBHOOK_URL = os.getenv('SLACK_WEBHOOK_URL')

# Git commands
remote_branches_command = "git branch -r --merged origin/develop"
delete_branch_command = "git push origin --delete {}"

# Execute git command to get merged branches
result = subprocess.run(remote_branches_command.split(), capture_output=True, text=True)
merged_branches = result.stdout.splitlines()

# Filter out specific branches
branches_to_delete = [branch.strip() for branch in merged_branches if branch.strip() not in ['origin/develop', 'origin/main', 'origin/refactor']]

# Delete branches and collect deleted branch names
deleted_branches = []
for branch in branches_to_delete:
    branch_name = branch.replace('origin/', '')
    delete_result = subprocess.run(delete_branch_command.format(branch_name).split(), capture_output=True, text=True)
    if delete_result.returncode == 0:
        deleted_branches.append(branch_name)

# Send Slack message
if deleted_branches:
    message = f"The following branches have been deleted:\n" + "\n".join(deleted_branches)
    payload = {"text": message}
    response = requests.post(SLACK_WEBHOOK_URL, json=payload)
    if response.status_code != 200:
        raise ValueError(f"Request to Slack returned an error {response.status_code}, the response is:\n{response.text}")
else:
    message = "No branches to delete."
    payload = {"text": message}
    response = requests.post(SLACK_WEBHOOK_URL, json=payload)
    if response.status_code != 200:
        raise ValueError(f"Request to Slack returned an error {response.status_code}, the response is:\n{response.text}")