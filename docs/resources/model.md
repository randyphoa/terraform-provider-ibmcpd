---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ibmcpd_model Resource - ibmcpd"
subcategory: ""
description: |-
  Manages a model on IBM Cloud Pak for Data.
---

# ibmcpd_model (Resource)

Manages a model on IBM Cloud Pak for Data.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `model_path` (String) tar.gz file of model from joblib.
- `name` (String) Name of model.
- `software_spec` (String) Software spec of model.
- `type` (String) Type of model.

### Optional

- `asset_id` (String) Asset ID of model in project. Used when promoting model in project to deployment space.
- `checksum` (String) Checksum of Python object.
- `input_schema` (Attributes List) Training input schema of model. (see [below for nested schema](#nestedatt--input_schema))
- `label_column` (String) Label column of model.
- `project_id` (String) Project ID of model.
- `space_id` (String) Space ID of model.

### Read-Only

- `id` (String) Identifier for model.

<a id="nestedatt--input_schema"></a>
### Nested Schema for `input_schema`

Optional:

- `name` (String)
- `type` (String)

