
### Ansible Link
[Ansible 官网](https://www.ansible.com/)  

[Ansible 文档](https://docs.ansible.com/ansible/latest/index.html)  
[Ansible 权威指南](https://ansible-tran.readthedocs.io/en/latest/)  
[Ansible examples](https://github.com/ansible/ansible-examples)  

[Ansible 7小时入坑](https://www.bilibili.com/video/av25424954/)  

### Ansible 是什么
Ansible 的 slogan 从 “IT Automation Software for System Administrators” 变成了 “AUTOMATION FOR EVERYONE”。从一个给系统管理员使用的工具变成了给所有人使用的工具。

Ansible 有三个最吸引人的地方：无客户端、简单易用和日志集中控管。  
Ansible 的作用就是把写 shell 这件事变成自动化、标准化、模块化，方便更好的自动化运维（减少运维工作中的重复）。  

使用 Ansible 来实现自动化运维需求，需要做这三件事情：  
1、定义目标机器的列表：一种被称为 inventory 的类 ini 文件  
2、定义这些机器的配置：使用 YAML 格式的文件来描述你机器的配置  
3、执行 ansible-playbook -i inventory playbook.yml  
```
# inventory 文件
[servers]
111.111.111.111
112.112.112.112

# 如上 ip 的配置，写在 playbook.yml 中
---
- hosts: servers
  tasks: 
    - name: unarchive apache to /usr/local
      unarchive:
          src: /tmp/apache-8.5.15.tar.gz
          dest: /usr/local/
          remote_src: true
...

# 执行命令，同步所有配置的机器
ansible-playbook -i inventory playbook.yml  
```

### Ansible 安装
```bash
# Ubuntu
sudo apt-get install software-properties-common
sudo apt-add-repository ppa:ansible/ansible
sudo apt-get update
sudo apt-get install ansible

# CentOS：
sudo yum install -y ansible

# 查看版本：
ansible --version
```

### Ansible 基本配置
Ansible 执行的时候会按照以下顺序查找配置项，修改的时候要特别注意改的是哪个文件  
> ANSIBLE_CONFIG (环境变量)  
> ansible.cfg (脚本所在当前目录下)  
> \~/.ansible.cfg (用户家目录下，默认没有)  
> /etc/ansible/ansible.cfg（安装后会自动生成）  

配置远程主机地址 Inventory   
```
# 编辑 Ansible 配置文件：/etc/ansible/hosts
[hadoop-host]
192.168.0.1
192.168.0.2
192.168.0.3
```

### Ansible Playbook
playbook（剧本），顾名思义就是需要定义一个脚本或者说配置文件，然后定义好要做什么。之后 ansible 就会根据 playbook 脚本对远程主机进行操作。  
```
# 简单脚本使用 /opt/simple-playbook.yml
- hosts: all
  tasks:
    - name: whoami
      shell: 'whoami > /opt/whoami.txt'

# 脚本测试
ansible-playbook /opt/simple-playbook.yml
```

```
# 修改 hosts /opt/hosts-playbook.yml
- hosts: all
  remote_user: root
  tasks:
    - name: update hosts
      blockinfile: 
        path: /etc/hosts
        block: |
          127.0.0.1     linux01
          127.0.0.1     linux02
          127.0.0.1     linux03

# 脚本测试
ansible-playbook /opt/hosts-playbook.yml
```

```
# 环境部署脚本 /opt/env-playbook.yml
- hosts: all
  remote_user: root
  tasks:
    - name: Disable SELinux at next reboot
      selinux:
        state: disabled
        
    - name: disable firewalld
      command: "{{ item }}"
      with_items:
         - systemctl stop firewalld
         - systemctl disable firewalld
         - setenforce 0

	- name: install-basic
	      command: "{{ item }}"
	      with_items:
	         - yum install -y zip unzip lrzsz git epel-release wget htop deltarpm

	- name: install-docker
	      shell: "{{ item }}"
	      with_items:
	         - yum install -y yum-utils device-mapper-persistent-data lvm2
	         - yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
	         - yum makecache fast
	         - yum install -y docker-ce
	         - systemctl start docker.service
	         - docker run hello-world
	         
	- name: install-docker-compose
	      shell: "{{ item }}"
	      with_items:
	         - curl -L https://github.com/docker/compose/releases/download/1.18.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
	         - chmod +x /usr/local/bin/docker-compose
	         - docker-compose --version
	         - systemctl restart docker.service
	         - systemctl enable docker.service
...

# 脚本测试
ansible-playbook /opt/env-playbook.yml
```

### Ansible 操作命令
```bash
# 运行 Ansible 的 ping 命令，看看配置正确时输出
sudo ansible --private-key ~/.ssh/id_rsa all -m ping

# 让远程所有主机都执行 ps 命令
ansible all -a 'ps'
# 让远程所有 hadoop-host 组的主机都执行 ps 命令
ansible hadoop-host -a 'ps'
```


