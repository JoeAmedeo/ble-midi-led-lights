## Example command

`ansible-playbook main.yaml --extra-vars "version=x.x.x"`

## Adjusting for your pi

In order to run this playbook on your own pi, change the inventory file in this directory to follow this format:
```
[all]
<pi IP Address> ansible_user=<pi user name> ansible_ssh_pass=<password for said user>
```