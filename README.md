# OAO (Operating Account Operators)
> ⚙️ Operating Account Operators (OAO) is a Golang tool to interact with the LDAP protocol to manage account groups, roles, ACLs/ACEs, etc...

<div align="center">
 <img src="https://i.imgur.com/3hxKgOq.jpg" width="850">
</div>

<br>

<p align="center">
    <img src="https://img.shields.io/github/license/oppsec/OAO?color=cyan&logo=github&logoColor=cyan&style=for-the-badge">
    <img src="https://img.shields.io/github/issues/oppsec/OAO?color=cyan&logo=github&logoColor=cyan&style=for-the-badge">
    <img src="https://img.shields.io/github/stars/oppsec/OAO?color=cyan&label=STARS&logo=github&logoColor=cyan&style=for-the-badge">
    <img src="https://img.shields.io/github/forks/oppsec/OAO?color=cyan&logo=github&logoColor=cyan&style=for-the-badge">
    <img src="https://img.shields.io/github/languages/code-size/oppsec/OAO?color=cyan&logo=github&logoColor=cyan&style=for-the-badge">
</p>

___

### 🕵️ What is OAO?
🕵️ **OAO** is a Golang tool to interact with the LDAP protocol to manage account groups, roles, ACLs/ACEs, etc...

<br>

### ⚡ Installing / Getting started

A quick guide of how to install and use OAO.

```shell
1. go install github.com/oppsec/OAO@latest
2. oao -u admin -p 123 -t victim -g "EXCHANGE TRUSTED SUBSYSTEM" -m add/rem
```

You can use `go install github.com/oppsec/OAO@latest` to update the tool

<br>

### ⚙️ Pre-requisites
- [Golang](https://go.dev/dl/) installed on your machine
- An valid user on domain with LDAP access

<br>

### ✨ Features
- Interact direct with LDAP (not malicious)
- Windows shell don't required
- Extremely fast
- Low RAM and CPU usage
- Made in Golang

<br>

### ⚔️ Attack Scenario & Suggestions
First of all, we suggest you to use this tool in combination with BloodHound to easily find exploitable paths. You can find a real attack scenario on our [article](https://twitter.com) that we used an another version to just add a specific user to a group with high privileges and use DSync attack to extract the Domain Admin NTLM hash.

<div align="left">
    <img src="https://i.imgur.com/1xfCaMC.png">
    <br><br>
    <img src="https://i.imgur.com/la3e7vM.png">
    <br><br>
    <img src="https://i.imgur.com/Y7P7HZn.png">
</div>

<br>

### 🔨 Contributing

A quick guide of how to contribute with the project.

```shell
1. Create a fork from OAO repository.
2. Download the project with git clone https://github.com/your/OAO.git
3. cd OAO/
4. Make your changes.
5. Commit and make a git push.
6. Open a pull request.
```

<br>

### ⚠️ Warning
- The developer is not responsible for any malicious use of this tool.
