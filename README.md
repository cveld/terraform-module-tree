# terraform-module-tree
This cli takes a `workloads.json` file that contains an array with absolute paths to terraform root modules. E.g.

```json
[
    "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\eslz\\Environments\\prd",
    "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\dev",
    "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\local"
]
```

Run it as follows:

```
terraform-module-tree workloads.json
```

There is an npm package published which you can run as follows with only nodejs as a prerequisite:
```
npx -y @cveld/terraform-module-tree@0.0.1-preview.5 -- workloads.json
```


It starts crawling all those root modules and spits out a json with a map of all crawled folders towards a list of all root modules that use this folder.

E.g.
```
{
 "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Modules\\app-configuration": [
  "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\dev",
  "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\local"
 ],
 "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Modules\\application-gateway": [
  "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\dev",
  "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\local"
 ],
 "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Modules\\buildingblock": [
  "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\dev",
  "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\local"
 ],
 "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Modules\\buildingblock-function": [
  "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\dev",
  "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\local"
 ],
 "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Modules\\buildingblock-web": [
  "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\dev",
  "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\local"
 ],
 "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Modules\\azure-applicationinsights\\v1.0.0": [
  "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\dev",
  "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\local"
 ],
 "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Modules\\keyvault": [
  "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\dev",
  "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\local"
 ],
 "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Modules\\shell_group": [
  "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\dev",
  "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\local"
 ],
 "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\dev": [
  "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\dev"
 ],
 "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\local": [
  "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\local"
 ],
 "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Infrastructure": [
  "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\dev",
  "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\app\\Environments\\local"
 ],
 "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\eslz\\Environments\\prd": [
  "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\eslz\\Environments\\prd"
 ],
 "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\eslz\\Infrastructure": [
  "C:\\work\\git\\myproject\\Infrastructure\\main-infrastructure\\Workloads\\eslz\\Environments\\prd"
 ]
}
```

This output can be used in combination with the output of `git diff --name-status HEAD~1 HEAD` to detect which root 
modules require a `terraform plan`.

### Developers
Build the solution with: 

```
go build
```

Publish the solution with:
```
./build.ps1
npm publish
```

### Release notes
Version 0.0.1-preview.5 supports terraform provider functions.
