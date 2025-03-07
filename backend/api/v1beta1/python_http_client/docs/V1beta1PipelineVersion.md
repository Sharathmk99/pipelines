# V1beta1PipelineVersion

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** | Output. Unique version ID. Generated by API server. | [optional] 
**name** | **str** | Optional input field. Version name provided by user. | [optional] 
**created_at** | **datetime** | Output. The time this pipeline version is created. | [optional] 
**parameters** | [**list[V1beta1Parameter]**](V1beta1Parameter.md) | Output. The input parameters for this pipeline. | [optional] 
**code_source_url** | **str** | Input. Optional. Pipeline version code source. | [optional] 
**package_url** | [**V1beta1Url**](V1beta1Url.md) |  | [optional] 
**resource_references** | [**list[V1beta1ResourceReference]**](V1beta1ResourceReference.md) | Input. Required. E.g., specify which pipeline this pipeline version belongs to. | [optional] 
**description** | **str** | Input. Optional. Description for the pipeline version. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


