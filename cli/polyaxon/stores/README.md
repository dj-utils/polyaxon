## Stores

This module includes clients and stores abstraction that can be used to interact with AWS S3, Azure Storage, and Google Cloud Storage.


## S3

### Normal instantiation

```python
from polyaxon.stores.stores.s3_store import S3Store

s3_store = S3Store(endpoint_url=..., 
                   access_key_id=...,
                   secret_access_key=...,
                   session_token=...,
                   region=...)
```


### Using env vars

```bash
export AWS_ENDPOINT_URL=...
export AWS_ACCESS_KEY_ID=...
export AWS_SECRET_ACCESS_KEY=...
export AWS_SECURITY_TOKEN=...
exprot AWS_REGION=...
```

And then you can instantiate the store
```python
from polyaxon.stores.stores.s3_store import S3Store

s3_store = S3Store()
```

### Using a client

```python
from polyaxon.stores.stores.s3_store import S3Store

s3_store = S3Store(client=client)
```

### Important methods

```python
s3_store.list(bucket_name, prefix='', delimiter='', page_size=None, max_items=None, keys=True, prefixes=True)
s3_store.list_prefixes(bucket_name, prefix='', delimiter='', page_size=None, max_items=None)
s3_store.list_keys(bucket_name, prefix='', delimiter='', page_size=None, max_items=None)
s3_store.check_key(key, bucket_name=None)
s3_store.get_key(key, bucket_name=None)
s3_store.read_key(key, bucket_name=None)
s3_store.upload_bytes(bytes_data, key, bucket_name=None, overwrite=False, encrypt=False, acl=None)
s3_store.upload_string(string_data, key, bucket_name=None, overwrite=False, encrypt=False, acl=None, encoding='utf-8')
s3_store.upload_file(filename, key, bucket_name=None, overwrite=False, encrypt=False, acl=None, use_basename=True)
s3_store.download_file(key, local_path, bucket_name=None, use_basename=True)
s3_store.upload_dir(dirname, key, bucket_name=None, overwrite=False, encrypt=False, acl=None, use_basename=True)
s3_store.download_dir(key, local_path, bucket_name=None, use_basename=True)
```


## GCS

### Normal instantiation

```python
from polyaxon.stores.stores.gcs_store import GCSStore

gcs_store = GCSStore(project_id=..., 
                     credentials=...,
                     key_path=...,
                     key_path=...,
                     scopes=...)
```

### Using a client

```python
from polyaxon.stores.stores.gcs_store import GCSStore

gcs_store = GCSStore(client=client)
```

### Important methods

```python
gcs_store.list(key, bucket_name=None, path=None, delimiter='/', blobs=True, prefixes=True)
gcs_store.upload_file(filename, blob, bucket_name=None, use_basename=True)
gcs_store.download_file(blob, local_path, bucket_name=None, use_basename=True)
gcs_store.upload_dir(dirname, blob, bucket_name=None, use_basename=True)
gcs_store.download_dir(blob, local_path, bucket_name=None, use_basename=True)
```

## Azure Storage

### Normal instantiation

```python
from polyaxon.stores.stores.azure_store import AzureStore

az_store = AzureStore(account_name=..., 
                      account_key=...,
                      connection_string=...)
```

### Using env vars

```bash
export AZURE_ACCOUNT_NAME=...
export AZURE_ACCOUNT_KEY=...
export AZURE_CONNECTION_STRING=...
```

And then you can instantiate the store
```python
from polyaxon.stores.stores.azure_store import AzureStore

az_store = AzureStore()
```

### Using a client

```python
from polyaxon.stores.stores.azure_store import AzureStore

az_store = AzureStore(client=client)
```

### Important methods

```python
az_store.list(key, container_name=None, path=None, delimiter='/', blobs=True, prefixes=True)
az_store.upload_file(filename, blob, container_name=None, use_basename=True)
az_store.download_file(blob, local_path, container_name=None, use_basename=True)
az_store.upload_dir(dirname, blob, container_name=None, use_basename=True)
az_store.download_dir(blob, local_path, container_name=None, use_basename=True)
```
