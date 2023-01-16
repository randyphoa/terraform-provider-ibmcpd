# Terraform Provider for IBM Cloud Pak for Data

This repository contains an implementation of a custom provider built using the [Terraform plugin framework](https://developer.hashicorp.com/terraform/plugin/framework).

The current [IBM Cloud Terraform Provider](https://github.com/IBM-Cloud/terraform-provider-ibm) focuses primarily on IBM Cloud infrastructure resources.

This provider focuses on IBM Cloud Pak for Data resources, such as,

- Watson Studio (projects, etc.)
- Watson Machine Learning (models, functions, deployments, etc.)
- Watson OpenScale (service providers, subscriptions, monitor instances, records, etc.)
- AI Factsheets

Both IBM Cloud Pak for Data (on-premise) and IBM Cloud Pak for Data as a Service (IBM Cloud) are supported.

## Related project

<br />

<p align="center">
    <a href="https://github.com/crossdeploy-io/crossdeploy">
        <img src="https://raw.githubusercontent.com/crossdeploy-io/crossdeploy/main/assets/logo.png" alt="crossdeploy" height="60" />
    </a>
</p>

If you prefer to use Python to manage IBM Cloud Pak for Data resources **declaratively**, consider using [crossdeploy](https://github.com/crossdeploy-io/crossdeploy), which is a thin wrapper around Terraform.
