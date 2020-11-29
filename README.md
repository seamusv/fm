# FM-Integration

### FM
- create the ability to un/marshal structures that will work with FM XML specifications for screens
- define a `fm` tag that will map to the FM fields

### Jobs
- define an input and output JSON structure for a job (ie: AddLine to a PO)
- are given a `Processor` used to submit the job function
- the submitted function takes an `Executor` to  




### Entities
- `Processor` - exposes a `chan` interface
- `Executor` - exposes login, logout and execute methods