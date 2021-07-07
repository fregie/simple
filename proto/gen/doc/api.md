# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [simple-interface/interface.proto](#simple-interface/interface.proto)
    - [Config](#interface.Config)
    - [CreateByConfigReq](#interface.CreateByConfigReq)
    - [CreateByConfigRsp](#interface.CreateByConfigRsp)
    - [CreateReq](#interface.CreateReq)
    - [CreateRsp](#interface.CreateRsp)
    - [CustomOptionSchemaReq](#interface.CustomOptionSchemaReq)
    - [CustomOptionSchemaRsp](#interface.CustomOptionSchemaRsp)
    - [DeleteReq](#interface.DeleteReq)
    - [DeleteRsp](#interface.DeleteRsp)
    - [Field](#interface.Field)
    - [GetAllReq](#interface.GetAllReq)
    - [GetAllRsp](#interface.GetAllRsp)
    - [GetAllRsp.AllEntry](#interface.GetAllRsp.AllEntry)
    - [GetReq](#interface.GetReq)
    - [GetRsp](#interface.GetRsp)
    - [GetStatReq](#interface.GetStatReq)
    - [GetStatRsp](#interface.GetStatRsp)
    - [IsSupportPersistenceReq](#interface.IsSupportPersistenceReq)
    - [IsSupportPersistenceRsp](#interface.IsSupportPersistenceRsp)
    - [NameReq](#interface.NameReq)
    - [NameRsp](#interface.NameRsp)
    - [Option](#interface.Option)
    - [SetMetadataReq](#interface.SetMetadataReq)
    - [SetMetadataRsp](#interface.SetMetadataRsp)
    - [Stat](#interface.Stat)
    - [UpdateOptionReq](#interface.UpdateOptionReq)
    - [UpdateOptionRsp](#interface.UpdateOptionRsp)
  
    - [Code](#interface.Code)
    - [ConfigType](#interface.ConfigType)
    - [Type](#interface.Type)
  
    - [Interface](#interface.Interface)
  
- [api/api.proto](#api/api.proto)
    - [CreateSessionReq](#simple.CreateSessionReq)
    - [CreateSessionRsp](#simple.CreateSessionRsp)
    - [DeleteSessionReq](#simple.DeleteSessionReq)
    - [DeleteSessionRsp](#simple.DeleteSessionRsp)
    - [GetAllSessionsReq](#simple.GetAllSessionsReq)
    - [GetAllSessionsRsp](#simple.GetAllSessionsRsp)
    - [GetProtosReq](#simple.GetProtosReq)
    - [GetProtosRsp](#simple.GetProtosRsp)
    - [Session](#simple.Session)
  
    - [Code](#simple.Code)
  
    - [SimpleAPI](#simple.SimpleAPI)
  
- [Scalar Value Types](#scalar-value-types)



<a name="simple-interface/interface.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## simple-interface/interface.proto



<a name="interface.Config"></a>

### Config



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Name | [string](#string) |  |  |
| ConfigType | [ConfigType](#interface.ConfigType) |  |  |
| Config | [bytes](#bytes) |  |  |






<a name="interface.CreateByConfigReq"></a>

### CreateByConfigReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Index | [string](#string) |  |  |
| Config | [Config](#interface.Config) |  |  |
| Opt | [Option](#interface.Option) |  |  |
| CustomOption | [string](#string) |  | in json |






<a name="interface.CreateByConfigRsp"></a>

### CreateByConfigRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [Code](#interface.Code) |  |  |
| Msg | [string](#string) |  |  |






<a name="interface.CreateReq"></a>

### CreateReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Opt | [Option](#interface.Option) |  |  |
| ConfigType | [ConfigType](#interface.ConfigType) |  |  |
| CustomOption | [string](#string) |  | in json |






<a name="interface.CreateRsp"></a>

### CreateRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [Code](#interface.Code) |  |  |
| Msg | [string](#string) |  |  |
| Index | [string](#string) |  |  |
| Config | [Config](#interface.Config) |  |  |






<a name="interface.CustomOptionSchemaReq"></a>

### CustomOptionSchemaReq







<a name="interface.CustomOptionSchemaRsp"></a>

### CustomOptionSchemaRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Fields | [Field](#interface.Field) | repeated |  |






<a name="interface.DeleteReq"></a>

### DeleteReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Index | [string](#string) |  |  |






<a name="interface.DeleteRsp"></a>

### DeleteRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [Code](#interface.Code) |  |  |
| Msg | [string](#string) |  |  |






<a name="interface.Field"></a>

### Field



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Name | [string](#string) |  |  |
| Type | [Type](#interface.Type) |  |  |
| optional | [string](#string) | repeated |  |






<a name="interface.GetAllReq"></a>

### GetAllReq







<a name="interface.GetAllRsp"></a>

### GetAllRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [Code](#interface.Code) |  |  |
| Msg | [string](#string) |  |  |
| All | [GetAllRsp.AllEntry](#interface.GetAllRsp.AllEntry) | repeated | Index, Config |






<a name="interface.GetAllRsp.AllEntry"></a>

### GetAllRsp.AllEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [Config](#interface.Config) |  |  |






<a name="interface.GetReq"></a>

### GetReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Index | [string](#string) |  |  |






<a name="interface.GetRsp"></a>

### GetRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [Code](#interface.Code) |  |  |
| Msg | [string](#string) |  |  |
| Index | [string](#string) |  |  |
| Config | [Config](#interface.Config) |  |  |






<a name="interface.GetStatReq"></a>

### GetStatReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Index | [string](#string) |  |  |






<a name="interface.GetStatRsp"></a>

### GetStatRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [Code](#interface.Code) |  |  |
| Msg | [string](#string) |  |  |
| Stat | [Stat](#interface.Stat) |  |  |






<a name="interface.IsSupportPersistenceReq"></a>

### IsSupportPersistenceReq







<a name="interface.IsSupportPersistenceRsp"></a>

### IsSupportPersistenceRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| IsSupport | [bool](#bool) |  |  |






<a name="interface.NameReq"></a>

### NameReq







<a name="interface.NameRsp"></a>

### NameRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Name | [string](#string) |  |  |






<a name="interface.Option"></a>

### Option



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| SendRateLimit | [uint64](#uint64) |  |  |
| RecvRateLimit | [uint64](#uint64) |  |  |






<a name="interface.SetMetadataReq"></a>

### SetMetadataReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| IP | [string](#string) |  |  |
| Domain | [string](#string) |  |  |






<a name="interface.SetMetadataRsp"></a>

### SetMetadataRsp







<a name="interface.Stat"></a>

### Stat



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| SendTraffic | [uint64](#uint64) |  |  |
| RecvTraffic | [uint64](#uint64) |  |  |






<a name="interface.UpdateOptionReq"></a>

### UpdateOptionReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Index | [string](#string) |  |  |
| Opt | [Option](#interface.Option) |  |  |






<a name="interface.UpdateOptionRsp"></a>

### UpdateOptionRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [Code](#interface.Code) |  |  |
| Msg | [string](#string) |  |  |





 


<a name="interface.Code"></a>

### Code


| Name | Number | Description |
| ---- | ------ | ----------- |
| OK | 0 | 请求成功 |
| NotFound | 10404 |  |
| Fail | 10500 |  |



<a name="interface.ConfigType"></a>

### ConfigType


| Name | Number | Description |
| ---- | ------ | ----------- |
| JSON | 0 |  |
| YAML | 1 |  |
| URL | 2 |  |
| TEXT | 3 |  |



<a name="interface.Type"></a>

### Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| String | 0 |  |
| Number | 1 |  |
| Bool | 2 |  |
| StringArray | 3 |  |


 

 


<a name="interface.Interface"></a>

### Interface


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Name | [NameReq](#interface.NameReq) | [NameRsp](#interface.NameRsp) |  |
| IsSupportPersistence | [IsSupportPersistenceReq](#interface.IsSupportPersistenceReq) | [IsSupportPersistenceRsp](#interface.IsSupportPersistenceRsp) |  |
| CustomOptionSchema | [CustomOptionSchemaReq](#interface.CustomOptionSchemaReq) | [CustomOptionSchemaRsp](#interface.CustomOptionSchemaRsp) |  |
| SetMetadata | [SetMetadataReq](#interface.SetMetadataReq) | [SetMetadataRsp](#interface.SetMetadataRsp) |  |
| Create | [CreateReq](#interface.CreateReq) | [CreateRsp](#interface.CreateRsp) |  |
| CreateByConfig | [CreateByConfigReq](#interface.CreateByConfigReq) | [CreateByConfigRsp](#interface.CreateByConfigRsp) |  |
| Get | [GetReq](#interface.GetReq) | [GetRsp](#interface.GetRsp) |  |
| Delete | [DeleteReq](#interface.DeleteReq) | [DeleteRsp](#interface.DeleteRsp) |  |
| GetAll | [GetAllReq](#interface.GetAllReq) | [GetAllRsp](#interface.GetAllRsp) |  |
| UpdateOption | [UpdateOptionReq](#interface.UpdateOptionReq) | [UpdateOptionRsp](#interface.UpdateOptionRsp) |  |
| GetStat | [GetStatReq](#interface.GetStatReq) | [GetStatRsp](#interface.GetStatRsp) |  |

 



<a name="api/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/api.proto



<a name="simple.CreateSessionReq"></a>

### CreateSessionReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Proto | [string](#string) |  |  |
| ConfigType | [interface.ConfigType](#interface.ConfigType) |  |  |
| Opt | [interface.Option](#interface.Option) |  |  |
| customOpt | [string](#string) |  |  |






<a name="simple.CreateSessionRsp"></a>

### CreateSessionRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [Code](#simple.Code) |  |  |
| Msg | [string](#string) |  |  |
| Config | [string](#string) |  |  |






<a name="simple.DeleteSessionReq"></a>

### DeleteSessionReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






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






<a name="simple.Session"></a>

### Session



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Proto | [string](#string) |  |  |
| ConfigType | [interface.ConfigType](#interface.ConfigType) |  |  |
| Config | [string](#string) |  |  |
| Opt | [interface.Option](#interface.Option) |  |  |





 


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
| DeleteSession | [DeleteSessionReq](#simple.DeleteSessionReq) | [DeleteSessionRsp](#simple.DeleteSessionRsp) |  |
| GetProtos | [GetProtosReq](#simple.GetProtosReq) | [GetProtosRsp](#simple.GetProtosRsp) |  |

 



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

