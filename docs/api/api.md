# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/api/api.proto](#proto/api/api.proto)
    - [CreateSessionReq](#simple.CreateSessionReq)
    - [CreateSessionRsp](#simple.CreateSessionRsp)
    - [DeleteSessionReq](#simple.DeleteSessionReq)
    - [DeleteSessionRsp](#simple.DeleteSessionRsp)
    - [GetAllSessionsReq](#simple.GetAllSessionsReq)
    - [GetAllSessionsRsp](#simple.GetAllSessionsRsp)
    - [GetProtosReq](#simple.GetProtosReq)
    - [GetProtosRsp](#simple.GetProtosRsp)
    - [GetSchemaReq](#simple.GetSchemaReq)
    - [GetSchemaRsp](#simple.GetSchemaRsp)
    - [GetSchemaRsp.SchemasEntry](#simple.GetSchemaRsp.SchemasEntry)
    - [GetSessionReq](#simple.GetSessionReq)
    - [GetSessionRsp](#simple.GetSessionRsp)
    - [Schema](#simple.Schema)
    - [Session](#simple.Session)
  
    - [Code](#simple.Code)
  
    - [SimpleAPI](#simple.SimpleAPI)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto/api/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/api/api.proto



<a name="simple.CreateSessionReq"></a>

### CreateSessionReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Proto | [string](#string) |  |  |
| ConfigType | [interface.ConfigType](#interface.ConfigType) |  |  |
| Opt | [interface.Option](#interface.Option) |  |  |
| CustomOpt | [string](#string) |  |  |
| Name | [string](#string) |  |  |






<a name="simple.CreateSessionRsp"></a>

### CreateSessionRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [Code](#simple.Code) |  |  |
| Msg | [string](#string) |  |  |
| ID | [string](#string) |  |  |
| Proto | [string](#string) |  |  |
| ConfigType | [interface.ConfigType](#interface.ConfigType) |  |  |
| Config | [string](#string) |  |  |






<a name="simple.DeleteSessionReq"></a>

### DeleteSessionReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| IDorName | [string](#string) |  |  |






<a name="simple.DeleteSessionRsp"></a>

### DeleteSessionRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [Code](#simple.Code) |  |  |
| Msg | [string](#string) |  |  |






<a name="simple.GetAllSessionsReq"></a>

### GetAllSessionsReq







<a name="simple.GetAllSessionsRsp"></a>

### GetAllSessionsRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [Code](#simple.Code) |  |  |
| Msg | [string](#string) |  |  |
| Sessions | [Session](#simple.Session) | repeated |  |






<a name="simple.GetProtosReq"></a>

### GetProtosReq







<a name="simple.GetProtosRsp"></a>

### GetProtosRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [Code](#simple.Code) |  |  |
| Msg | [string](#string) |  |  |
| Protos | [string](#string) | repeated |  |






<a name="simple.GetSchemaReq"></a>

### GetSchemaReq







<a name="simple.GetSchemaRsp"></a>

### GetSchemaRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [Code](#simple.Code) |  |  |
| Msg | [string](#string) |  |  |
| Schemas | [GetSchemaRsp.SchemasEntry](#simple.GetSchemaRsp.SchemasEntry) | repeated |  |






<a name="simple.GetSchemaRsp.SchemasEntry"></a>

### GetSchemaRsp.SchemasEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [Schema](#simple.Schema) |  |  |






<a name="simple.GetSessionReq"></a>

### GetSessionReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| IDorName | [string](#string) |  |  |






<a name="simple.GetSessionRsp"></a>

### GetSessionRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [Code](#simple.Code) |  |  |
| Msg | [string](#string) |  |  |
| Session | [Session](#simple.Session) |  |  |






<a name="simple.Schema"></a>

### Schema



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Proto | [string](#string) |  |  |
| Fields | [interface.Field](#interface.Field) | repeated |  |






<a name="simple.Session"></a>

### Session



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Proto | [string](#string) |  |  |
| ConfigType | [interface.ConfigType](#interface.ConfigType) |  |  |
| Config | [string](#string) |  |  |
| Opt | [interface.Option](#interface.Option) |  |  |
| Name | [string](#string) |  |  |





 


<a name="simple.Code"></a>

### Code


| Name | Number | Description |
| ---- | ------ | ----------- |
| OK | 0 | 请求成功 |
| InternalError | 10005 |  |


 

 


<a name="simple.SimpleAPI"></a>

### SimpleAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateSession | [CreateSessionReq](#simple.CreateSessionReq) | [CreateSessionRsp](#simple.CreateSessionRsp) |  |
| GetAllSessions | [GetAllSessionsReq](#simple.GetAllSessionsReq) | [GetAllSessionsRsp](#simple.GetAllSessionsRsp) |  |
| GetSession | [GetSessionReq](#simple.GetSessionReq) | [GetSessionRsp](#simple.GetSessionRsp) |  |
| DeleteSession | [DeleteSessionReq](#simple.DeleteSessionReq) | [DeleteSessionRsp](#simple.DeleteSessionRsp) |  |
| GetProtos | [GetProtosReq](#simple.GetProtosReq) | [GetProtosRsp](#simple.GetProtosRsp) |  |
| GetSchema | [GetSchemaReq](#simple.GetSchemaReq) | [GetSchemaRsp](#simple.GetSchemaRsp) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

