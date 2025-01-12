<!-- DO NOT EDIT: this file is automatically generated using scw-doc-gen -->
# Documentation for `scw webhosting`
Webhosting API.
  
- [DNS record management commands](#dns-record-management-commands)
  - [Get the DNS records](#get-the-dns-records)
- [Hosting management commands](#hosting-management-commands)
  - [Create a hosting](#create-a-hosting)
  - [Get a hosting](#get-a-hosting)
  - [List all hostings](#list-all-hostings)
  - [Update a hosting](#update-a-hosting)
- [Offer management commands](#offer-management-commands)
  - [List all offers](#list-all-offers)

  
## DNS record management commands

The set of DNS record of a specific domain associated to a hosting.



### Get the DNS records

Get the DNS records of a specified domain.

**Usage:**

```
scw webhosting dns-record get [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| domain | Required | Domain associated to the DNS records |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |


**Examples:**


Get DNS records associated to the given domain
```
scw webhosting dns-record get domain=foo.com
```




## Hosting management commands

A Scaleway web hosting associated with a domain name.



### Create a hosting

Create a hosting.

**Usage:**

```
scw webhosting hosting create [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| offer-id |  | ID of the selected offer for the hosting |
| project-id |  | Project ID to use. If none is passed the default project ID will be used |
| email |  | Contact email of the client for the hosting |
| tags.{index} |  | The tags of the hosting |
| domain |  | The domain name of the hosting |
| option-ids.{index} |  | IDs of the selected options for the hosting |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



### Get a hosting

Get the details of a Hosting with the given ID.

**Usage:**

```
scw webhosting hosting get [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| hosting-id | Required | Hosting ID |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |


**Examples:**


Get a Hosting with the given ID
```
scw webhosting hosting get hosting-id=a3244331-5d32-4e36-9bf9-b60233e201c7
```




### List all hostings

List all hostings.

**Usage:**

```
scw webhosting hosting list [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| order-by | One of: `created_at_asc`, `created_at_desc` | Define the order of the returned hostings |
| tags.{index} |  | Return hostings with these tags |
| statuses.{index} | One of: `unknown_status`, `delivering`, `ready`, `deleting`, `error`, `locked` | Return hostings with these statuses |
| domain |  | Return hostings with this domain |
| project-id |  | Return hostings from this project ID |
| organization-id |  | Return hostings from this organization ID |
| region | Default: `fr-par`<br />One of: `fr-par`, `all` | Region to target. If none is passed will use default region from the config |


**Examples:**


List all hostings of a given project ID
```
scw webhosting hosting list project-id=a3244331-5d32-4e36-9bf9-b60233e201c7 organization-id=a3244331-5d32-4e36-9bf9-b60233e201c7
```




### Update a hosting

Update a hosting.

**Usage:**

```
scw webhosting hosting update [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| hosting-id | Required | Hosting ID |
| email |  | New contact email for the hosting |
| tags.{index} |  | New tags for the hosting |
| option-ids.{index} |  | New options IDs for the hosting |
| offer-id |  | New offer ID for the hosting |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |


**Examples:**


Update the contact email of a given hosting
```
scw webhosting hosting update hosting-id=11111111-1111-1111-1111-111111111111 email=foobar@example.com
```

Overwrite tags of a given hosting
```
scw webhosting hosting update hosting-id=11111111-1111-1111-1111-111111111111 tags.0=foo tags.1=bar
```

Overwrite options of a given hosting
```
scw webhosting hosting update hosting-id=11111111-1111-1111-1111-111111111111 option-ids.0=22222222-2222-2222-2222-222222222222 option-ids.1=33333333-3333-3333-3333-333333333333
```




## Offer management commands

A hosting offer, with a set of features, available for purchase.



### List all offers

List all offers.

**Usage:**

```
scw webhosting offer list [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| order-by | One of: `price_asc` | Define the order of the returned hostings |
| without-options |  | Select only offers, no options |
| only-options |  | Select only options |
| hosting-id |  | Define a specific hosting id (optional) |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |


**Examples:**


List all offers available for purchase
```
scw webhosting offer list hosting-id=a3244331-5d32-4e36-9bf9-b60233e201c7
```

List only offers, no options
```
scw webhosting offer list without-options=true
```

List only options
```
scw webhosting offer list only-options=true
```




