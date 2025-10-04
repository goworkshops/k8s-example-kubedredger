# Exercise: Implement the Configuration Controller

## 1: Understand the API Structure

First, familiarize yourself with the API structure by examining the `Configuration` custom resource definition:

- Check the file `api/v1alpha1/configuration_types.go`
- Review the `ConfigurationSpec` field
- Understand what data the controller will be working with

## 2: Review the File Manager

The controller uses a file manager to handle configuration files on disk:

- Check the `HandleSync` method in `internal/configfile/configfile.go`
- This method will be called to create/update configuration files
- Note that it takes a `ConfigRequest` with filename, content, create flag, and permissions

## 3: Implement the Reconcile Logic

Implement the add/update logic in your reconcile loop, in `internal/controller/configuration_controller.go`

### Get the Configuration Object

To retrieve the Configuration object corresponding to the reconcile request:

```go
conf := &workshopv1alpha1.Configuration{}
err := r.Get(ctx, req.NamespacedName, conf)
```

### Implement Add/Update Logic

After getting the Configuration object, implement the logic to:
1. Call the file manager's `HandleSync` method with the appropriate parameters from `conf.Spec`
2. Update the status based on the result
3. Return appropriate results

## 4: Build and Deploy

Once you've implemented the controller logic:

### Deploy to Kind Cluster

```bash
make deploy-on-kind
```

### Apply a Sample Configuration

```bash
kubectl apply -f config/samples/create.yaml
```

### Validate the File was Created

Check that the configuration file was created inside the control plane container:

```bash
docker exec kubedredger-kind-control-plane cat /tmp/config.d/golab.conf
```

You should see the content from your Configuration resource written to the file!

## 5 (Optional): Add Finalizer Support

If you have time, implement proper deletion handling using a finalizer:

### Define the Finalizer Constant

```go
const Finalizer = "configuration.workshop.example.com/finalizer"
```

### Handle Deletion

Add deletion logic when the object is being deleted:

```go
if !conf.DeletionTimestamp.IsZero() {
	// Deletion
	if controllerutil.ContainsFinalizer(conf, Finalizer) {
		// DELETE LOGIC HERE
		// Use the Delete(fileName string) method from configfile.go
		controllerutil.RemoveFinalizer(conf, Finalizer)
		err = r.Update(ctx, conf)
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}
```

### Add Finalizer on Create/Update

Ensure the finalizer is added when the object is created:

```go
// Add or Update
if !controllerutil.ContainsFinalizer(conf, Finalizer) {
	controllerutil.AddFinalizer(conf, Finalizer)
	err = r.Update(ctx, conf)
	if err != nil {
		return ctrl.Result{}, err
	}
}
```
