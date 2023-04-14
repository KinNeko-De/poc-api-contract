# API Reference

# Table of Contents



- Messages
    - [DocumentGeneratedV1](#documentgeneratedv1)
  





- Messages
    - [GenerateDocumentV1](#generatedocumentv1)
    - [GenerateDocumentV1.Document](#generatedocumentv1document)
    - [GenerateDocumentV1.Document.InvoiceV1](#generatedocumentv1documentinvoicev1)
    - [GenerateDocumentV1.Document.InvoiceV1.Item](#generatedocumentv1documentinvoicev1item)
    - [GenerateDocumentV1.Document.InvoiceV1.Recipient](#generatedocumentv1documentinvoicev1recipient)
  





- Messages
    - [RequestV1](#requestv1)
  





- Messages
    - [Decimal](#decimal)
  





- Messages
    - [Uuid](#uuid)
  





- Messages
    - [Any](#any)
  





- Messages
    - [Api](#api)
    - [Method](#method)
    - [Mixin](#mixin)
  





- Messages
    - [DescriptorProto](#descriptorproto)
    - [DescriptorProto.ExtensionRange](#descriptorprotoextensionrange)
    - [DescriptorProto.ReservedRange](#descriptorprotoreservedrange)
    - [EnumDescriptorProto](#enumdescriptorproto)
    - [EnumDescriptorProto.EnumReservedRange](#enumdescriptorprotoenumreservedrange)
    - [EnumOptions](#enumoptions)
    - [EnumValueDescriptorProto](#enumvaluedescriptorproto)
    - [EnumValueOptions](#enumvalueoptions)
    - [ExtensionRangeOptions](#extensionrangeoptions)
    - [FieldDescriptorProto](#fielddescriptorproto)
    - [FieldOptions](#fieldoptions)
    - [FileDescriptorProto](#filedescriptorproto)
    - [FileDescriptorSet](#filedescriptorset)
    - [FileOptions](#fileoptions)
    - [GeneratedCodeInfo](#generatedcodeinfo)
    - [GeneratedCodeInfo.Annotation](#generatedcodeinfoannotation)
    - [MessageOptions](#messageoptions)
    - [MethodDescriptorProto](#methoddescriptorproto)
    - [MethodOptions](#methodoptions)
    - [OneofDescriptorProto](#oneofdescriptorproto)
    - [OneofOptions](#oneofoptions)
    - [ServiceDescriptorProto](#servicedescriptorproto)
    - [ServiceOptions](#serviceoptions)
    - [SourceCodeInfo](#sourcecodeinfo)
    - [SourceCodeInfo.Location](#sourcecodeinfolocation)
    - [UninterpretedOption](#uninterpretedoption)
    - [UninterpretedOption.NamePart](#uninterpretedoptionnamepart)
  





- Messages
    - [Duration](#duration)
  





- Messages
    - [Empty](#empty)
  





- Messages
    - [FieldMask](#fieldmask)
  





- Messages
    - [SourceContext](#sourcecontext)
  





- Messages
    - [ListValue](#listvalue)
    - [Struct](#struct)
    - [Struct.FieldsEntry](#structfieldsentry)
    - [Value](#value)
  


- Enums
    - [NullValue](#nullvalue)
  




- Messages
    - [Timestamp](#timestamp)
  





- Messages
    - [Enum](#enum)
    - [EnumValue](#enumvalue)
    - [Field](#field)
    - [Option](#option)
    - [Type](#type)
  


- Enums
    - [Field.Cardinality](#fieldcardinality)
    - [Field.Kind](#fieldkind)
    - [Syntax](#syntax)
  




- Messages
    - [BoolValue](#boolvalue)
    - [BytesValue](#bytesvalue)
    - [DoubleValue](#doublevalue)
    - [FloatValue](#floatvalue)
    - [Int32Value](#int32value)
    - [Int64Value](#int64value)
    - [StringValue](#stringvalue)
    - [UInt32Value](#uint32value)
    - [UInt64Value](#uint64value)
  



- [Scalar Value Types](#scalar-value-types)



 <!-- end services -->

# Messages


## DocumentGeneratedV1 {#documentgeneratedv1}
Reply to request GenerateDocumentV1 defined in generate_document.proto

TODO Define results according to Enterprise Integration Patterns

 <!-- end HasFields -->
 <!-- end messages -->

# Enums
 <!-- end Enums -->


 <!-- end services -->

# Messages


## GenerateDocumentV1 {#generatedocumentv1}
Request to generate a document
Is a command message and uses request/reply pattern
Reply will be of type DocumentGeneratedV1 defined generate_document.proto


| Field | Type | Description |
| ----- | ---- | ----------- |
| request | [ kinnekode.restaurant.RequestV1](#kinnekoderestaurantrequestv1) | none |
| requested_documents | [repeated GenerateDocumentV1.Document](#generatedocumentv1document) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## GenerateDocumentV1.Document {#generatedocumentv1document}



| Field | Type | Description |
| ----- | ---- | ----------- |
| output_formats | [repeated GenerateDocumentV1.Document.OutputFormat](#generatedocumentv1documentoutputformat) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) type.invoice | [ GenerateDocumentV1.Document.InvoiceV1](#generatedocumentv1documentinvoicev1) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## GenerateDocumentV1.Document.InvoiceV1 {#generatedocumentv1documentinvoicev1}



| Field | Type | Description |
| ----- | ---- | ----------- |
| delivered_on | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | none |
| currency_code | [ string](#string) | three character currency code as specified in ISO 4217 ( see https://de.wikipedia.org/wiki/ISO_4217 ) |
| recipient | [ GenerateDocumentV1.Document.InvoiceV1.Recipient](#generatedocumentv1documentinvoicev1recipient) | none |
| items | [repeated GenerateDocumentV1.Document.InvoiceV1.Item](#generatedocumentv1documentinvoicev1item) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## GenerateDocumentV1.Document.InvoiceV1.Item {#generatedocumentv1documentinvoicev1item}



| Field | Type | Description |
| ----- | ---- | ----------- |
| description | [ string](#string) | none |
| quantity | [ int64](#int64) | none |
| netAmount | [ kinnekode.protobuf.Decimal](#kinnekodeprotobufdecimal) | none |
| taxation | [ kinnekode.protobuf.Decimal](#kinnekodeprotobufdecimal) | none |
| totalAmount | [ kinnekode.protobuf.Decimal](#kinnekodeprotobufdecimal) | none |
| sum | [ kinnekode.protobuf.Decimal](#kinnekodeprotobufdecimal) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## GenerateDocumentV1.Document.InvoiceV1.Recipient {#generatedocumentv1documentinvoicev1recipient}



| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | none |
| street | [ string](#string) | none |
| city | [ string](#string) | none |
| postCode | [ string](#string) | none |
| country | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->
 <!-- end messages -->

# Enums


## GenerateDocumentV1.Document.OutputFormat {#generatedocumentv1documentoutputformat}


| Name | Number | Description |
| ---- | ------ | ----------- |
| OUTPUT_FORMAT_UNSPECIFIED | 0 | none |
| OUTPUT_FORMAT_PDF | 1 | none |


 <!-- end Enums -->


 <!-- end services -->

# Messages


## RequestV1 {#requestv1}



| Field | Type | Description |
| ----- | ---- | ----------- |
| request_id | [ kinnekode.protobuf.Uuid](#kinnekodeprotobufuuid) | Unique id for the reply_to. must be ensured by the requestor |
| reply_to | [ string](#string) | Pub/Sub Topic where the reply should be send to currently only a assumption until implementation of reply is done |
 <!-- end Fields -->
 <!-- end HasFields -->
 <!-- end messages -->

# Enums
 <!-- end Enums -->


 <!-- end services -->

# Messages


## Decimal {#decimal}
Decimal https://en.wikipedia.org/wiki/IEEE_754#Decimal

 Format:
 The whole units of the amount.
 Followed optional by a dot '.' and the number of nano (10^-9) units of the amount.
 The value must be between 000000000 and 999999999 inclusive.


| Field | Type | Description |
| ----- | ---- | ----------- |
| value | [ string](#string) | Example: 42, -42, 42.0000, -42.0000, 42.0001, -42.0001 |
 <!-- end Fields -->
 <!-- end HasFields -->
 <!-- end messages -->

# Enums
 <!-- end Enums -->


 <!-- end services -->

# Messages


## Uuid {#uuid}
Universally Unique Identifier (https://www.rfc-editor.org/rfc/rfc4122)


| Field | Type | Description |
| ----- | ---- | ----------- |
| value | [ string](#string) | Example: 550e8400-e29b-11d4-a716-446655440000 |
 <!-- end Fields -->
 <!-- end HasFields -->
 <!-- end messages -->

# Enums
 <!-- end Enums -->


 <!-- end services -->

# Messages


## Any {#any}
`Any` contains an arbitrary serialized protocol buffer message along with a
URL that describes the type of the serialized message.

Protobuf library provides support to pack/unpack Any values in the form
of utility functions or additional generated methods of the Any type.

Example 1: Pack and unpack a message in C++.

    Foo foo = ...;
    Any any;
    any.PackFrom(foo);
    ...
    if (any.UnpackTo(&foo)) {
      ...
    }

Example 2: Pack and unpack a message in Java.

    Foo foo = ...;
    Any any = Any.pack(foo);
    ...
    if (any.is(Foo.class)) {
      foo = any.unpack(Foo.class);
    }
    // or ...
    if (any.isSameTypeAs(Foo.getDefaultInstance())) {
      foo = any.unpack(Foo.getDefaultInstance());
    }

Example 3: Pack and unpack a message in Python.

    foo = Foo(...)
    any = Any()
    any.Pack(foo)
    ...
    if any.Is(Foo.DESCRIPTOR):
      any.Unpack(foo)
      ...

Example 4: Pack and unpack a message in Go

     foo := &pb.Foo{...}
     any, err := anypb.New(foo)
     if err != nil {
       ...
     }
     ...
     foo := &pb.Foo{}
     if err := any.UnmarshalTo(foo); err != nil {
       ...
     }

The pack methods provided by protobuf library will by default use
'type.googleapis.com/full.type.name' as the type URL and the unpack
methods only use the fully qualified type name after the last '/'
in the type URL, for example "foo.bar.com/x/y.z" will yield type
name "y.z".

JSON

The JSON representation of an `Any` value uses the regular
representation of the deserialized, embedded message, with an
additional field `@type` which contains the type URL. Example:

    package google.profile;
    message Person {
      string first_name = 1;
      string last_name = 2;
    }

    {
      "@type": "type.googleapis.com/google.profile.Person",
      "firstName": <string>,
      "lastName": <string>
    }

If the embedded message type is well-known and has a custom JSON
representation, that representation will be embedded adding a field
`value` which holds the custom JSON in addition to the `@type`
field. Example (for message [google.protobuf.Duration][]):

    {
      "@type": "type.googleapis.com/google.protobuf.Duration",
      "value": "1.212s"
    }


| Field | Type | Description |
| ----- | ---- | ----------- |
| type_url | [ string](#string) | A URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This string must contain at least one "/" character. The last segment of the URL's path must represent the fully qualified name of the type (as in `path/google.protobuf.Duration`). The name should be in a canonical form (e.g., leading "." is not accepted).

In practice, teams usually precompile into the binary all types that they expect it to use in the context of Any. However, for URLs which use the scheme `http`, `https`, or no scheme, one can optionally set up a type server that maps type URLs to message definitions as follows:

* If no scheme is provided, `https` is assumed. * An HTTP GET on the URL must yield a [google.protobuf.Type][] value in binary format, or produce an error. * Applications are allowed to cache lookup results based on the URL, or have them precompiled into a binary to avoid any lookup. Therefore, binary compatibility needs to be preserved on changes to types. (Use versioned type names to manage breaking changes.)

Note: this functionality is not currently available in the official protobuf release, and it is not used for type URLs beginning with type.googleapis.com.

Schemes other than `http`, `https` (or the empty scheme) might be used with implementation specific semantics. |
| value | [ bytes](#bytes) | Must be a valid serialized protocol buffer of the above specified type. |
 <!-- end Fields -->
 <!-- end HasFields -->
 <!-- end messages -->

# Enums
 <!-- end Enums -->


 <!-- end services -->

# Messages


## Api {#api}
Api is a light-weight descriptor for an API Interface.

Interfaces are also described as "protocol buffer services" in some contexts,
such as by the "service" keyword in a .proto file, but they are different
from API Services, which represent a concrete implementation of an interface
as opposed to simply a description of methods and bindings. They are also
sometimes simply referred to as "APIs" in other contexts, such as the name of
this message itself. See https://cloud.google.com/apis/design/glossary for
detailed terminology.


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | The fully qualified name of this interface, including package name followed by the interface's simple name. |
| methods | [repeated Method](#method) | The methods of this interface, in unspecified order. |
| options | [repeated Option](#option) | Any metadata attached to the interface. |
| version | [ string](#string) | A version string for this interface. If specified, must have the form `major-version.minor-version`, as in `1.10`. If the minor version is omitted, it defaults to zero. If the entire version field is empty, the major version is derived from the package name, as outlined below. If the field is not empty, the version in the package name will be verified to be consistent with what is provided here.

The versioning schema uses [semantic versioning](http://semver.org) where the major version number indicates a breaking change and the minor version an additive, non-breaking change. Both version numbers are signals to users what to expect from different versions, and should be carefully chosen based on the product plan.

The major version is also reflected in the package name of the interface, which must end in `v<major-version>`, as in `google.feature.v1`. For major versions 0 and 1, the suffix can be omitted. Zero major versions must only be used for experimental, non-GA interfaces. |
| source_context | [ SourceContext](#sourcecontext) | Source context for the protocol buffer service represented by this message. |
| mixins | [repeated Mixin](#mixin) | Included interfaces. See [Mixin][]. |
| syntax | [ Syntax](#syntax) | The source syntax of the service. |
 <!-- end Fields -->
 <!-- end HasFields -->


## Method {#method}
Method represents a method of an API interface.


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | The simple name of this method. |
| request_type_url | [ string](#string) | A URL of the input message type. |
| request_streaming | [ bool](#bool) | If true, the request is streamed. |
| response_type_url | [ string](#string) | The URL of the output message type. |
| response_streaming | [ bool](#bool) | If true, the response is streamed. |
| options | [repeated Option](#option) | Any metadata attached to the method. |
| syntax | [ Syntax](#syntax) | The source syntax of this method. |
 <!-- end Fields -->
 <!-- end HasFields -->


## Mixin {#mixin}
Declares an API Interface to be included in this interface. The including
interface must redeclare all the methods from the included interface, but
documentation and options are inherited as follows:

- If after comment and whitespace stripping, the documentation
  string of the redeclared method is empty, it will be inherited
  from the original method.

- Each annotation belonging to the service config (http,
  visibility) which is not set in the redeclared method will be
  inherited.

- If an http annotation is inherited, the path pattern will be
  modified as follows. Any version prefix will be replaced by the
  version of the including interface plus the [root][] path if
  specified.

Example of a simple mixin:

    package google.acl.v1;
    service AccessControl {
      // Get the underlying ACL object.
      rpc GetAcl(GetAclRequest) returns (Acl) {
        option (google.api.http).get = "/v1/{resource=**}:getAcl";
      }
    }

    package google.storage.v2;
    service Storage {
      rpc GetAcl(GetAclRequest) returns (Acl);

      // Get a data record.
      rpc GetData(GetDataRequest) returns (Data) {
        option (google.api.http).get = "/v2/{resource=**}";
      }
    }

Example of a mixin configuration:

    apis:
    - name: google.storage.v2.Storage
      mixins:
      - name: google.acl.v1.AccessControl

The mixin construct implies that all methods in `AccessControl` are
also declared with same name and request/response types in
`Storage`. A documentation generator or annotation processor will
see the effective `Storage.GetAcl` method after inheriting
documentation and annotations as follows:

    service Storage {
      // Get the underlying ACL object.
      rpc GetAcl(GetAclRequest) returns (Acl) {
        option (google.api.http).get = "/v2/{resource=**}:getAcl";
      }
      ...
    }

Note how the version in the path pattern changed from `v1` to `v2`.

If the `root` field in the mixin is specified, it should be a
relative path under which inherited HTTP paths are placed. Example:

    apis:
    - name: google.storage.v2.Storage
      mixins:
      - name: google.acl.v1.AccessControl
        root: acls

This implies the following inherited HTTP annotation:

    service Storage {
      // Get the underlying ACL object.
      rpc GetAcl(GetAclRequest) returns (Acl) {
        option (google.api.http).get = "/v2/acls/{resource=**}:getAcl";
      }
      ...
    }


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | The fully qualified name of the interface which is included. |
| root | [ string](#string) | If non-empty specifies a path under which inherited HTTP paths are rooted. |
 <!-- end Fields -->
 <!-- end HasFields -->
 <!-- end messages -->

# Enums
 <!-- end Enums -->


 <!-- end services -->

# Messages


## DescriptorProto {#descriptorproto}
Describes a message type.


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [optional string](#string) | none |
| field | [repeated FieldDescriptorProto](#fielddescriptorproto) | none |
| extension | [repeated FieldDescriptorProto](#fielddescriptorproto) | none |
| nested_type | [repeated DescriptorProto](#descriptorproto) | none |
| enum_type | [repeated EnumDescriptorProto](#enumdescriptorproto) | none |
| extension_range | [repeated DescriptorProto.ExtensionRange](#descriptorprotoextensionrange) | none |
| oneof_decl | [repeated OneofDescriptorProto](#oneofdescriptorproto) | none |
| options | [optional MessageOptions](#messageoptions) | none |
| reserved_range | [repeated DescriptorProto.ReservedRange](#descriptorprotoreservedrange) | none |
| reserved_name | [repeated string](#string) | Reserved field names, which may not be used by fields in the same message. A given name may only be reserved once. |
 <!-- end Fields -->
 <!-- end HasFields -->


## DescriptorProto.ExtensionRange {#descriptorprotoextensionrange}



| Field | Type | Description |
| ----- | ---- | ----------- |
| start | [optional int32](#int32) | Inclusive. |
| end | [optional int32](#int32) | Exclusive. |
| options | [optional ExtensionRangeOptions](#extensionrangeoptions) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## DescriptorProto.ReservedRange {#descriptorprotoreservedrange}
Range of reserved tag numbers. Reserved tag numbers may not be used by
fields or extension ranges in the same message. Reserved ranges may
not overlap.


| Field | Type | Description |
| ----- | ---- | ----------- |
| start | [optional int32](#int32) | Inclusive. |
| end | [optional int32](#int32) | Exclusive. |
 <!-- end Fields -->
 <!-- end HasFields -->


## EnumDescriptorProto {#enumdescriptorproto}
Describes an enum type.


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [optional string](#string) | none |
| value | [repeated EnumValueDescriptorProto](#enumvaluedescriptorproto) | none |
| options | [optional EnumOptions](#enumoptions) | none |
| reserved_range | [repeated EnumDescriptorProto.EnumReservedRange](#enumdescriptorprotoenumreservedrange) | Range of reserved numeric values. Reserved numeric values may not be used by enum values in the same enum declaration. Reserved ranges may not overlap. |
| reserved_name | [repeated string](#string) | Reserved enum value names, which may not be reused. A given name may only be reserved once. |
 <!-- end Fields -->
 <!-- end HasFields -->


## EnumDescriptorProto.EnumReservedRange {#enumdescriptorprotoenumreservedrange}
Range of reserved numeric values. Reserved values may not be used by
entries in the same enum. Reserved ranges may not overlap.

Note that this is distinct from DescriptorProto.ReservedRange in that it
is inclusive such that it can appropriately represent the entire int32
domain.


| Field | Type | Description |
| ----- | ---- | ----------- |
| start | [optional int32](#int32) | Inclusive. |
| end | [optional int32](#int32) | Inclusive. |
 <!-- end Fields -->
 <!-- end HasFields -->


## EnumOptions {#enumoptions}



| Field | Type | Description |
| ----- | ---- | ----------- |
| allow_alias | [optional bool](#bool) | Set this option to true to allow mapping different tag names to the same value. |
| deprecated | [optional bool](#bool) | Is this enum deprecated? Depending on the target platform, this can emit Deprecated annotations for the enum, or it will be completely ignored; in the very least, this is a formalization for deprecating enums. Default: false |
| deprecated_legacy_json_field_conflicts | [optional bool](#bool) | Enable the legacy handling of JSON field name conflicts. This lowercases and strips underscored from the fields before comparison in proto3 only. The new behavior takes `json_name` into account and applies to proto2 as well. TODO(b/261750190) Remove this legacy behavior once downstream teams have had time to migrate. |
| uninterpreted_option | [repeated UninterpretedOption](#uninterpretedoption) | The parser stores options it doesn't recognize here. See above. |
 <!-- end Fields -->
 <!-- end HasFields -->


## EnumValueDescriptorProto {#enumvaluedescriptorproto}
Describes a value within an enum.


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [optional string](#string) | none |
| number | [optional int32](#int32) | none |
| options | [optional EnumValueOptions](#enumvalueoptions) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## EnumValueOptions {#enumvalueoptions}



| Field | Type | Description |
| ----- | ---- | ----------- |
| deprecated | [optional bool](#bool) | Is this enum value deprecated? Depending on the target platform, this can emit Deprecated annotations for the enum value, or it will be completely ignored; in the very least, this is a formalization for deprecating enum values. Default: false |
| uninterpreted_option | [repeated UninterpretedOption](#uninterpretedoption) | The parser stores options it doesn't recognize here. See above. |
 <!-- end Fields -->
 <!-- end HasFields -->


## ExtensionRangeOptions {#extensionrangeoptions}



| Field | Type | Description |
| ----- | ---- | ----------- |
| uninterpreted_option | [repeated UninterpretedOption](#uninterpretedoption) | The parser stores options it doesn't recognize here. See above. |
 <!-- end Fields -->
 <!-- end HasFields -->


## FieldDescriptorProto {#fielddescriptorproto}
Describes a field within a message.


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [optional string](#string) | none |
| number | [optional int32](#int32) | none |
| label | [optional FieldDescriptorProto.Label](#fielddescriptorprotolabel) | none |
| type | [optional FieldDescriptorProto.Type](#fielddescriptorprototype) | If type_name is set, this need not be set. If both this and type_name are set, this must be one of TYPE_ENUM, TYPE_MESSAGE or TYPE_GROUP. |
| type_name | [optional string](#string) | For message and enum types, this is the name of the type. If the name starts with a '.', it is fully-qualified. Otherwise, C++-like scoping rules are used to find the type (i.e. first the nested types within this message are searched, then within the parent, on up to the root namespace). |
| extendee | [optional string](#string) | For extensions, this is the name of the type being extended. It is resolved in the same manner as type_name. |
| default_value | [optional string](#string) | For numeric types, contains the original text representation of the value. For booleans, "true" or "false". For strings, contains the default text contents (not escaped in any way). For bytes, contains the C escaped value. All bytes >= 128 are escaped. |
| oneof_index | [optional int32](#int32) | If set, gives the index of a oneof in the containing type's oneof_decl list. This field is a member of that oneof. |
| json_name | [optional string](#string) | JSON name of this field. The value is set by protocol compiler. If the user has set a "json_name" option on this field, that option's value will be used. Otherwise, it's deduced from the field's name by converting it to camelCase. |
| options | [optional FieldOptions](#fieldoptions) | none |
| proto3_optional | [optional bool](#bool) | If true, this is a proto3 "optional". When a proto3 field is optional, it tracks presence regardless of field type.

When proto3_optional is true, this field must be belong to a oneof to signal to old proto3 clients that presence is tracked for this field. This oneof is known as a "synthetic" oneof, and this field must be its sole member (each proto3 optional field gets its own synthetic oneof). Synthetic oneofs exist in the descriptor only, and do not generate any API. Synthetic oneofs must be ordered after all "real" oneofs.

For message fields, proto3_optional doesn't create any semantic change, since non-repeated message fields always track presence. However it still indicates the semantic detail of whether the user wrote "optional" or not. This can be useful for round-tripping the .proto file. For consistency we give message fields a synthetic oneof also, even though it is not required to track presence. This is especially important because the parser can't tell if a field is a message or an enum, so it must always create a synthetic oneof.

Proto2 optional fields do not set this flag, because they already indicate optional with `LABEL_OPTIONAL`. |
 <!-- end Fields -->
 <!-- end HasFields -->


## FieldOptions {#fieldoptions}



| Field | Type | Description |
| ----- | ---- | ----------- |
| ctype | [optional FieldOptions.CType](#fieldoptionsctype) | The ctype option instructs the C++ code generator to use a different representation of the field than it normally would. See the specific options below. This option is not yet implemented in the open source release -- sorry, we'll try to include it in a future version! Default: STRING |
| packed | [optional bool](#bool) | The packed option can be enabled for repeated primitive fields to enable a more efficient representation on the wire. Rather than repeatedly writing the tag and type for each element, the entire array is encoded as a single length-delimited blob. In proto3, only explicit setting it to false will avoid using packed encoding. |
| jstype | [optional FieldOptions.JSType](#fieldoptionsjstype) | The jstype option determines the JavaScript type used for values of the field. The option is permitted only for 64 bit integral and fixed types (int64, uint64, sint64, fixed64, sfixed64). A field with jstype JS_STRING is represented as JavaScript string, which avoids loss of precision that can happen when a large value is converted to a floating point JavaScript. Specifying JS_NUMBER for the jstype causes the generated JavaScript code to use the JavaScript "number" type. The behavior of the default option JS_NORMAL is implementation dependent.

This option is an enum to permit additional types to be added, e.g. goog.math.Integer. Default: JS_NORMAL |
| lazy | [optional bool](#bool) | Should this field be parsed lazily? Lazy applies only to message-type fields. It means that when the outer message is initially parsed, the inner message's contents will not be parsed but instead stored in encoded form. The inner message will actually be parsed when it is first accessed.

This is only a hint. Implementations are free to choose whether to use eager or lazy parsing regardless of the value of this option. However, setting this option true suggests that the protocol author believes that using lazy parsing on this field is worth the additional bookkeeping overhead typically needed to implement it.

This option does not affect the public interface of any generated code; all method signatures remain the same. Furthermore, thread-safety of the interface is not affected by this option; const methods remain safe to call from multiple threads concurrently, while non-const methods continue to require exclusive access.

Note that implementations may choose not to check required fields within a lazy sub-message. That is, calling IsInitialized() on the outer message may return true even if the inner message has missing required fields. This is necessary because otherwise the inner message would have to be parsed in order to perform the check, defeating the purpose of lazy parsing. An implementation which chooses not to check required fields must be consistent about it. That is, for any particular sub-message, the implementation must either *always* check its required fields, or *never* check its required fields, regardless of whether or not the message has been parsed.

As of May 2022, lazy verifies the contents of the byte stream during parsing. An invalid byte stream will cause the overall parsing to fail. Default: false |
| unverified_lazy | [optional bool](#bool) | unverified_lazy does no correctness checks on the byte stream. This should only be used where lazy with verification is prohibitive for performance reasons. Default: false |
| deprecated | [optional bool](#bool) | Is this field deprecated? Depending on the target platform, this can emit Deprecated annotations for accessors, or it will be completely ignored; in the very least, this is a formalization for deprecating fields. Default: false |
| weak | [optional bool](#bool) | For Google-internal migration only. Do not use. Default: false |
| debug_redact | [optional bool](#bool) | Indicate that the field value should not be printed out when using debug formats, e.g. when the field contains sensitive credentials. Default: false |
| retention | [optional FieldOptions.OptionRetention](#fieldoptionsoptionretention) | none |
| target | [optional FieldOptions.OptionTargetType](#fieldoptionsoptiontargettype) | none |
| uninterpreted_option | [repeated UninterpretedOption](#uninterpretedoption) | The parser stores options it doesn't recognize here. See above. |
 <!-- end Fields -->
 <!-- end HasFields -->


## FileDescriptorProto {#filedescriptorproto}
Describes a complete .proto file.


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [optional string](#string) | file name, relative to root of source tree |
| package | [optional string](#string) | e.g. "foo", "foo.bar", etc. |
| dependency | [repeated string](#string) | Names of files imported by this file. |
| public_dependency | [repeated int32](#int32) | Indexes of the public imported files in the dependency list above. |
| weak_dependency | [repeated int32](#int32) | Indexes of the weak imported files in the dependency list. For Google-internal migration only. Do not use. |
| message_type | [repeated DescriptorProto](#descriptorproto) | All top-level definitions in this file. |
| enum_type | [repeated EnumDescriptorProto](#enumdescriptorproto) | none |
| service | [repeated ServiceDescriptorProto](#servicedescriptorproto) | none |
| extension | [repeated FieldDescriptorProto](#fielddescriptorproto) | none |
| options | [optional FileOptions](#fileoptions) | none |
| source_code_info | [optional SourceCodeInfo](#sourcecodeinfo) | This field contains optional information about the original source code. You may safely remove this entire field without harming runtime functionality of the descriptors -- the information is needed only by development tools. |
| syntax | [optional string](#string) | The syntax of the proto file. The supported values are "proto2", "proto3", and "editions".

If `edition` is present, this value must be "editions". |
| edition | [optional string](#string) | The edition of the proto file, which is an opaque string. |
 <!-- end Fields -->
 <!-- end HasFields -->


## FileDescriptorSet {#filedescriptorset}
The protocol compiler can output a FileDescriptorSet containing the .proto
files it parses.


| Field | Type | Description |
| ----- | ---- | ----------- |
| file | [repeated FileDescriptorProto](#filedescriptorproto) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## FileOptions {#fileoptions}



| Field | Type | Description |
| ----- | ---- | ----------- |
| java_package | [optional string](#string) | Sets the Java package where classes generated from this .proto will be placed. By default, the proto package is used, but this is often inappropriate because proto packages do not normally start with backwards domain names. |
| java_outer_classname | [optional string](#string) | Controls the name of the wrapper Java class generated for the .proto file. That class will always contain the .proto file's getDescriptor() method as well as any top-level extensions defined in the .proto file. If java_multiple_files is disabled, then all the other classes from the .proto file will be nested inside the single wrapper outer class. |
| java_multiple_files | [optional bool](#bool) | If enabled, then the Java code generator will generate a separate .java file for each top-level message, enum, and service defined in the .proto file. Thus, these types will *not* be nested inside the wrapper class named by java_outer_classname. However, the wrapper class will still be generated to contain the file's getDescriptor() method as well as any top-level extensions defined in the file. Default: false |
| java_generate_equals_and_hash | [optional bool](#bool) | This option does nothing. |
| java_string_check_utf8 | [optional bool](#bool) | If set true, then the Java2 code generator will generate code that throws an exception whenever an attempt is made to assign a non-UTF-8 byte sequence to a string field. Message reflection will do the same. However, an extension field still accepts non-UTF-8 byte sequences. This option has no effect on when used with the lite runtime. Default: false |
| optimize_for | [optional FileOptions.OptimizeMode](#fileoptionsoptimizemode) | none |
| go_package | [optional string](#string) | Sets the Go package where structs generated from this .proto will be placed. If omitted, the Go package will be derived from the following: - The basename of the package import path, if provided. - Otherwise, the package statement in the .proto file, if present. - Otherwise, the basename of the .proto file, without extension. |
| cc_generic_services | [optional bool](#bool) | Should generic services be generated in each language? "Generic" services are not specific to any particular RPC system. They are generated by the main code generators in each language (without additional plugins). Generic services were the only kind of service generation supported by early versions of google.protobuf.

Generic services are now considered deprecated in favor of using plugins that generate code specific to your particular RPC system. Therefore, these default to false. Old code which depends on generic services should explicitly set them to true. Default: false |
| java_generic_services | [optional bool](#bool) | none |
| py_generic_services | [optional bool](#bool) | none |
| php_generic_services | [optional bool](#bool) | none |
| deprecated | [optional bool](#bool) | Is this file deprecated? Depending on the target platform, this can emit Deprecated annotations for everything in the file, or it will be completely ignored; in the very least, this is a formalization for deprecating files. Default: false |
| cc_enable_arenas | [optional bool](#bool) | Enables the use of arenas for the proto messages in this file. This applies only to generated classes for C++. Default: true |
| objc_class_prefix | [optional string](#string) | Sets the objective c class prefix which is prepended to all objective c generated classes from this .proto. There is no default. |
| csharp_namespace | [optional string](#string) | Namespace for generated classes; defaults to the package. |
| swift_prefix | [optional string](#string) | By default Swift generators will take the proto package and CamelCase it replacing '.' with underscore and use that to prefix the types/symbols defined. When this options is provided, they will use this value instead to prefix the types/symbols defined. |
| php_class_prefix | [optional string](#string) | Sets the php class prefix which is prepended to all php generated classes from this .proto. Default is empty. |
| php_namespace | [optional string](#string) | Use this option to change the namespace of php generated classes. Default is empty. When this option is empty, the package name will be used for determining the namespace. |
| php_metadata_namespace | [optional string](#string) | Use this option to change the namespace of php generated metadata classes. Default is empty. When this option is empty, the proto file name will be used for determining the namespace. |
| ruby_package | [optional string](#string) | Use this option to change the package of ruby generated classes. Default is empty. When this option is not set, the package name will be used for determining the ruby package. |
| uninterpreted_option | [repeated UninterpretedOption](#uninterpretedoption) | The parser stores options it doesn't recognize here. See the documentation for the "Options" section above. |
 <!-- end Fields -->
 <!-- end HasFields -->


## GeneratedCodeInfo {#generatedcodeinfo}
Describes the relationship between generated code and its original source
file. A GeneratedCodeInfo message is associated with only one generated
source file, but may contain references to different source .proto files.


| Field | Type | Description |
| ----- | ---- | ----------- |
| annotation | [repeated GeneratedCodeInfo.Annotation](#generatedcodeinfoannotation) | An Annotation connects some span of text in generated code to an element of its generating .proto file. |
 <!-- end Fields -->
 <!-- end HasFields -->


## GeneratedCodeInfo.Annotation {#generatedcodeinfoannotation}



| Field | Type | Description |
| ----- | ---- | ----------- |
| path | [repeated int32](#int32) | Identifies the element in the original source .proto file. This field is formatted the same as SourceCodeInfo.Location.path. |
| source_file | [optional string](#string) | Identifies the filesystem path to the original source .proto. |
| begin | [optional int32](#int32) | Identifies the starting offset in bytes in the generated code that relates to the identified object. |
| end | [optional int32](#int32) | Identifies the ending offset in bytes in the generated code that relates to the identified object. The end offset should be one past the last relevant byte (so the length of the text = end - begin). |
| semantic | [optional GeneratedCodeInfo.Annotation.Semantic](#generatedcodeinfoannotationsemantic) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## MessageOptions {#messageoptions}



| Field | Type | Description |
| ----- | ---- | ----------- |
| message_set_wire_format | [optional bool](#bool) | Set true to use the old proto1 MessageSet wire format for extensions. This is provided for backwards-compatibility with the MessageSet wire format. You should not use this for any other reason: It's less efficient, has fewer features, and is more complicated.

The message must be defined exactly as follows: message Foo { option message_set_wire_format = true; extensions 4 to max; } Note that the message cannot have any defined fields; MessageSets only have extensions.

All extensions of your type must be singular messages; e.g. they cannot be int32s, enums, or repeated messages.

Because this is an option, the above two restrictions are not enforced by the protocol compiler. Default: false |
| no_standard_descriptor_accessor | [optional bool](#bool) | Disables the generation of the standard "descriptor()" accessor, which can conflict with a field of the same name. This is meant to make migration from proto1 easier; new code should avoid fields named "descriptor". Default: false |
| deprecated | [optional bool](#bool) | Is this message deprecated? Depending on the target platform, this can emit Deprecated annotations for the message, or it will be completely ignored; in the very least, this is a formalization for deprecating messages. Default: false |
| map_entry | [optional bool](#bool) | NOTE: Do not set the option in .proto files. Always use the maps syntax instead. The option should only be implicitly set by the proto compiler parser.

Whether the message is an automatically generated map entry type for the maps field.

For maps fields: map<KeyType, ValueType> map_field = 1; The parsed descriptor looks like: message MapFieldEntry { option map_entry = true; optional KeyType key = 1; optional ValueType value = 2; } repeated MapFieldEntry map_field = 1;

Implementations may choose not to generate the map_entry=true message, but use a native map in the target language to hold the keys and values. The reflection APIs in such implementations still need to work as if the field is a repeated message field. |
| deprecated_legacy_json_field_conflicts | [optional bool](#bool) | Enable the legacy handling of JSON field name conflicts. This lowercases and strips underscored from the fields before comparison in proto3 only. The new behavior takes `json_name` into account and applies to proto2 as well.

This should only be used as a temporary measure against broken builds due to the change in behavior for JSON field name conflicts.

TODO(b/261750190) This is legacy behavior we plan to remove once downstream teams have had time to migrate. |
| uninterpreted_option | [repeated UninterpretedOption](#uninterpretedoption) | The parser stores options it doesn't recognize here. See above. |
 <!-- end Fields -->
 <!-- end HasFields -->


## MethodDescriptorProto {#methoddescriptorproto}
Describes a method of a service.


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [optional string](#string) | none |
| input_type | [optional string](#string) | Input and output type names. These are resolved in the same way as FieldDescriptorProto.type_name, but must refer to a message type. |
| output_type | [optional string](#string) | none |
| options | [optional MethodOptions](#methodoptions) | none |
| client_streaming | [optional bool](#bool) | Identifies if client streams multiple client messages Default: false |
| server_streaming | [optional bool](#bool) | Identifies if server streams multiple server messages Default: false |
 <!-- end Fields -->
 <!-- end HasFields -->


## MethodOptions {#methodoptions}



| Field | Type | Description |
| ----- | ---- | ----------- |
| deprecated | [optional bool](#bool) | Is this method deprecated? Depending on the target platform, this can emit Deprecated annotations for the method, or it will be completely ignored; in the very least, this is a formalization for deprecating methods. Default: false |
| idempotency_level | [optional MethodOptions.IdempotencyLevel](#methodoptionsidempotencylevel) | none |
| uninterpreted_option | [repeated UninterpretedOption](#uninterpretedoption) | The parser stores options it doesn't recognize here. See above. |
 <!-- end Fields -->
 <!-- end HasFields -->


## OneofDescriptorProto {#oneofdescriptorproto}
Describes a oneof.


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [optional string](#string) | none |
| options | [optional OneofOptions](#oneofoptions) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## OneofOptions {#oneofoptions}



| Field | Type | Description |
| ----- | ---- | ----------- |
| uninterpreted_option | [repeated UninterpretedOption](#uninterpretedoption) | The parser stores options it doesn't recognize here. See above. |
 <!-- end Fields -->
 <!-- end HasFields -->


## ServiceDescriptorProto {#servicedescriptorproto}
Describes a service.


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [optional string](#string) | none |
| method | [repeated MethodDescriptorProto](#methoddescriptorproto) | none |
| options | [optional ServiceOptions](#serviceoptions) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## ServiceOptions {#serviceoptions}



| Field | Type | Description |
| ----- | ---- | ----------- |
| deprecated | [optional bool](#bool) | Is this service deprecated? Depending on the target platform, this can emit Deprecated annotations for the service, or it will be completely ignored; in the very least, this is a formalization for deprecating services. Default: false |
| uninterpreted_option | [repeated UninterpretedOption](#uninterpretedoption) | The parser stores options it doesn't recognize here. See above. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SourceCodeInfo {#sourcecodeinfo}
Encapsulates information about the original source file from which a
FileDescriptorProto was generated.


| Field | Type | Description |
| ----- | ---- | ----------- |
| location | [repeated SourceCodeInfo.Location](#sourcecodeinfolocation) | A Location identifies a piece of source code in a .proto file which corresponds to a particular definition. This information is intended to be useful to IDEs, code indexers, documentation generators, and similar tools.

For example, say we have a file like: message Foo { optional string foo = 1; } Let's look at just the field definition: optional string foo = 1; ^ ^^ ^^ ^ ^^^ a bc de f ghi We have the following locations: span path represents [a,i) [ 4, 0, 2, 0 ] The whole field definition. [a,b) [ 4, 0, 2, 0, 4 ] The label (optional). [c,d) [ 4, 0, 2, 0, 5 ] The type (string). [e,f) [ 4, 0, 2, 0, 1 ] The name (foo). [g,h) [ 4, 0, 2, 0, 3 ] The number (1).

Notes: - A location may refer to a repeated field itself (i.e. not to any particular index within it). This is used whenever a set of elements are logically enclosed in a single code segment. For example, an entire extend block (possibly containing multiple extension definitions) will have an outer location whose path refers to the "extensions" repeated field without an index. - Multiple locations may have the same path. This happens when a single logical declaration is spread out across multiple places. The most obvious example is the "extend" block again -- there may be multiple extend blocks in the same scope, each of which will have the same path. - A location's span is not always a subset of its parent's span. For example, the "extendee" of an extension declaration appears at the beginning of the "extend" block and is shared by all extensions within the block. - Just because a location's span is a subset of some other location's span does not mean that it is a descendant. For example, a "group" defines both a type and a field in a single declaration. Thus, the locations corresponding to the type and field and their components will overlap. - Code which tries to interpret locations should probably be designed to ignore those that it doesn't understand, as more types of locations could be recorded in the future. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SourceCodeInfo.Location {#sourcecodeinfolocation}



| Field | Type | Description |
| ----- | ---- | ----------- |
| path | [repeated int32](#int32) | Identifies which part of the FileDescriptorProto was defined at this location.

Each element is a field number or an index. They form a path from the root FileDescriptorProto to the place where the definition occurs. For example, this path: [ 4, 3, 2, 7, 1 ] refers to: file.message_type(3) // 4, 3 .field(7) // 2, 7 .name() // 1 This is because FileDescriptorProto.message_type has field number 4: repeated DescriptorProto message_type = 4; and DescriptorProto.field has field number 2: repeated FieldDescriptorProto field = 2; and FieldDescriptorProto.name has field number 1: optional string name = 1;

Thus, the above path gives the location of a field name. If we removed the last element: [ 4, 3, 2, 7 ] this path refers to the whole field declaration (from the beginning of the label to the terminating semicolon). |
| span | [repeated int32](#int32) | Always has exactly three or four elements: start line, start column, end line (optional, otherwise assumed same as start line), end column. These are packed into a single field for efficiency. Note that line and column numbers are zero-based -- typically you will want to add 1 to each before displaying to a user. |
| leading_comments | [optional string](#string) | If this SourceCodeInfo represents a complete declaration, these are any comments appearing before and after the declaration which appear to be attached to the declaration.

A series of line comments appearing on consecutive lines, with no other tokens appearing on those lines, will be treated as a single comment.

leading_detached_comments will keep paragraphs of comments that appear before (but not connected to) the current element. Each paragraph, separated by empty lines, will be one comment element in the repeated field.

Only the comment content is provided; comment markers (e.g. //) are stripped out. For block comments, leading whitespace and an asterisk will be stripped from the beginning of each line other than the first. Newlines are included in the output.

Examples:

 optional int32 foo = 1; // Comment attached to foo. // Comment attached to bar. optional int32 bar = 2;

 optional string baz = 3; // Comment attached to baz. // Another line attached to baz.

 // Comment attached to moo. // // Another line attached to moo. optional double moo = 4;

 // Detached comment for corge. This is not leading or trailing comments // to moo or corge because there are blank lines separating it from // both.

 // Detached comment for corge paragraph 2.

 optional string corge = 5; /* Block comment attached * to corge. Leading asterisks * will be removed. */ /* Block comment attached to * grault. */ optional int32 grault = 6;

 // ignored detached comments. |
| trailing_comments | [optional string](#string) | none |
| leading_detached_comments | [repeated string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## UninterpretedOption {#uninterpretedoption}
A message representing a option the parser does not recognize. This only
appears in options protos created by the compiler::Parser class.
DescriptorPool resolves these when building Descriptor objects. Therefore,
options protos in descriptor objects (e.g. returned by Descriptor::options(),
or produced by Descriptor::CopyTo()) will never have UninterpretedOptions
in them.


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [repeated UninterpretedOption.NamePart](#uninterpretedoptionnamepart) | none |
| identifier_value | [optional string](#string) | The value of the uninterpreted option, in whatever type the tokenizer identified it as during parsing. Exactly one of these should be set. |
| positive_int_value | [optional uint64](#uint64) | none |
| negative_int_value | [optional int64](#int64) | none |
| double_value | [optional double](#double) | none |
| string_value | [optional bytes](#bytes) | none |
| aggregate_value | [optional string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## UninterpretedOption.NamePart {#uninterpretedoptionnamepart}
The name of the uninterpreted option.  Each string represents a segment in
a dot-separated name.  is_extension is true iff a segment represents an
extension (denoted with parentheses in options specs in .proto files).
E.g.,{ ["foo", false], ["bar.baz", true], ["moo", false] } represents
"foo.(bar.baz).moo".


| Field | Type | Description |
| ----- | ---- | ----------- |
| name_part | [required string](#string) | none |
| is_extension | [required bool](#bool) | none |
 <!-- end Fields -->
 <!-- end HasFields -->
 <!-- end messages -->

# Enums


## FieldDescriptorProto.Label {#fielddescriptorprotolabel}


| Name | Number | Description |
| ---- | ------ | ----------- |
| LABEL_OPTIONAL | 1 | 0 is reserved for errors |
| LABEL_REQUIRED | 2 | none |
| LABEL_REPEATED | 3 | none |




## FieldDescriptorProto.Type {#fielddescriptorprototype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_DOUBLE | 1 | 0 is reserved for errors. Order is weird for historical reasons. |
| TYPE_FLOAT | 2 | none |
| TYPE_INT64 | 3 | Not ZigZag encoded. Negative numbers take 10 bytes. Use TYPE_SINT64 if negative values are likely. |
| TYPE_UINT64 | 4 | none |
| TYPE_INT32 | 5 | Not ZigZag encoded. Negative numbers take 10 bytes. Use TYPE_SINT32 if negative values are likely. |
| TYPE_FIXED64 | 6 | none |
| TYPE_FIXED32 | 7 | none |
| TYPE_BOOL | 8 | none |
| TYPE_STRING | 9 | none |
| TYPE_GROUP | 10 | Tag-delimited aggregate. Group type is deprecated and not supported in proto3. However, Proto3 implementations should still be able to parse the group wire format and treat group fields as unknown fields. |
| TYPE_MESSAGE | 11 | Length-delimited aggregate. |
| TYPE_BYTES | 12 | New in version 2. |
| TYPE_UINT32 | 13 | none |
| TYPE_ENUM | 14 | none |
| TYPE_SFIXED32 | 15 | none |
| TYPE_SFIXED64 | 16 | none |
| TYPE_SINT32 | 17 | Uses ZigZag encoding. |
| TYPE_SINT64 | 18 | Uses ZigZag encoding. |




## FieldOptions.CType {#fieldoptionsctype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| STRING | 0 | Default mode. |
| CORD | 1 | none |
| STRING_PIECE | 2 | none |




## FieldOptions.JSType {#fieldoptionsjstype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| JS_NORMAL | 0 | Use the default type. |
| JS_STRING | 1 | Use JavaScript strings. |
| JS_NUMBER | 2 | Use JavaScript numbers. |




## FieldOptions.OptionRetention {#fieldoptionsoptionretention}
If set to RETENTION_SOURCE, the option will be omitted from the binary.
Note: as of January 2023, support for this is in progress and does not yet
have an effect (b/264593489).

| Name | Number | Description |
| ---- | ------ | ----------- |
| RETENTION_UNKNOWN | 0 | none |
| RETENTION_RUNTIME | 1 | none |
| RETENTION_SOURCE | 2 | none |




## FieldOptions.OptionTargetType {#fieldoptionsoptiontargettype}
This indicates the types of entities that the field may apply to when used
as an option. If it is unset, then the field may be freely used as an
option on any kind of entity. Note: as of January 2023, support for this is
in progress and does not yet have an effect (b/264593489).

| Name | Number | Description |
| ---- | ------ | ----------- |
| TARGET_TYPE_UNKNOWN | 0 | none |
| TARGET_TYPE_FILE | 1 | none |
| TARGET_TYPE_EXTENSION_RANGE | 2 | none |
| TARGET_TYPE_MESSAGE | 3 | none |
| TARGET_TYPE_FIELD | 4 | none |
| TARGET_TYPE_ONEOF | 5 | none |
| TARGET_TYPE_ENUM | 6 | none |
| TARGET_TYPE_ENUM_ENTRY | 7 | none |
| TARGET_TYPE_SERVICE | 8 | none |
| TARGET_TYPE_METHOD | 9 | none |




## FileOptions.OptimizeMode {#fileoptionsoptimizemode}
Generated classes can be optimized for speed or code size.

| Name | Number | Description |
| ---- | ------ | ----------- |
| SPEED | 1 | Generate complete code for parsing, serialization, |
| CODE_SIZE | 2 | etc.

Use ReflectionOps to implement these methods. |
| LITE_RUNTIME | 3 | Generate code using MessageLite and the lite runtime. |




## GeneratedCodeInfo.Annotation.Semantic {#generatedcodeinfoannotationsemantic}
Represents the identified object's effect on the element in the original
.proto file.

| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 | There is no effect or the effect is indescribable. |
| SET | 1 | The element is set or otherwise mutated. |
| ALIAS | 2 | An alias to the element is returned. |




## MethodOptions.IdempotencyLevel {#methodoptionsidempotencylevel}
Is this method side-effect-free (or safe in HTTP parlance), or idempotent,
or neither? HTTP based RPC implementation may choose GET verb for safe
methods, and PUT verb for idempotent methods instead of the default POST.

| Name | Number | Description |
| ---- | ------ | ----------- |
| IDEMPOTENCY_UNKNOWN | 0 | none |
| NO_SIDE_EFFECTS | 1 | implies idempotent |
| IDEMPOTENT | 2 | idempotent, but may have side effects |


 <!-- end Enums -->


 <!-- end services -->

# Messages


## Duration {#duration}
A Duration represents a signed, fixed-length span of time represented
as a count of seconds and fractions of seconds at nanosecond
resolution. It is independent of any calendar and concepts like "day"
or "month". It is related to Timestamp in that the difference between
two Timestamp values is a Duration and it can be added or subtracted
from a Timestamp. Range is approximately +-10,000 years.

# Examples

Example 1: Compute Duration from two Timestamps in pseudo code.

    Timestamp start = ...;
    Timestamp end = ...;
    Duration duration = ...;

    duration.seconds = end.seconds - start.seconds;
    duration.nanos = end.nanos - start.nanos;

    if (duration.seconds < 0 && duration.nanos > 0) {
      duration.seconds += 1;
      duration.nanos -= 1000000000;
    } else if (duration.seconds > 0 && duration.nanos < 0) {
      duration.seconds -= 1;
      duration.nanos += 1000000000;
    }

Example 2: Compute Timestamp from Timestamp + Duration in pseudo code.

    Timestamp start = ...;
    Duration duration = ...;
    Timestamp end = ...;

    end.seconds = start.seconds + duration.seconds;
    end.nanos = start.nanos + duration.nanos;

    if (end.nanos < 0) {
      end.seconds -= 1;
      end.nanos += 1000000000;
    } else if (end.nanos >= 1000000000) {
      end.seconds += 1;
      end.nanos -= 1000000000;
    }

Example 3: Compute Duration from datetime.timedelta in Python.

    td = datetime.timedelta(days=3, minutes=10)
    duration = Duration()
    duration.FromTimedelta(td)

# JSON Mapping

In JSON format, the Duration type is encoded as a string rather than an
object, where the string ends in the suffix "s" (indicating seconds) and
is preceded by the number of seconds, with nanoseconds expressed as
fractional seconds. For example, 3 seconds with 0 nanoseconds should be
encoded in JSON format as "3s", while 3 seconds and 1 nanosecond should
be expressed in JSON format as "3.000000001s", and 3 seconds and 1
microsecond should be expressed in JSON format as "3.000001s".


| Field | Type | Description |
| ----- | ---- | ----------- |
| seconds | [ int64](#int64) | Signed seconds of the span of time. Must be from -315,576,000,000 to +315,576,000,000 inclusive. Note: these bounds are computed from: 60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years |
| nanos | [ int32](#int32) | Signed fractions of a second at nanosecond resolution of the span of time. Durations less than one second are represented with a 0 `seconds` field and a positive or negative `nanos` field. For durations of one second or more, a non-zero value for the `nanos` field must be of the same sign as the `seconds` field. Must be from -999,999,999 to +999,999,999 inclusive. |
 <!-- end Fields -->
 <!-- end HasFields -->
 <!-- end messages -->

# Enums
 <!-- end Enums -->


 <!-- end services -->

# Messages


## Empty {#empty}
A generic empty message that you can re-use to avoid defining duplicated
empty messages in your APIs. A typical example is to use it as the request
or the response type of an API method. For instance:

    service Foo {
      rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);
    }

 <!-- end HasFields -->
 <!-- end messages -->

# Enums
 <!-- end Enums -->


 <!-- end services -->

# Messages


## FieldMask {#fieldmask}
`FieldMask` represents a set of symbolic field paths, for example:

    paths: "f.a"
    paths: "f.b.d"

Here `f` represents a field in some root message, `a` and `b`
fields in the message found in `f`, and `d` a field found in the
message in `f.b`.

Field masks are used to specify a subset of fields that should be
returned by a get operation or modified by an update operation.
Field masks also have a custom JSON encoding (see below).

# Field Masks in Projections

When used in the context of a projection, a response message or
sub-message is filtered by the API to only contain those fields as
specified in the mask. For example, if the mask in the previous
example is applied to a response message as follows:

    f {
      a : 22
      b {
        d : 1
        x : 2
      }
      y : 13
    }
    z: 8

The result will not contain specific values for fields x,y and z
(their value will be set to the default, and omitted in proto text
output):


    f {
      a : 22
      b {
        d : 1
      }
    }

A repeated field is not allowed except at the last position of a
paths string.

If a FieldMask object is not present in a get operation, the
operation applies to all fields (as if a FieldMask of all fields
had been specified).

Note that a field mask does not necessarily apply to the
top-level response message. In case of a REST get operation, the
field mask applies directly to the response, but in case of a REST
list operation, the mask instead applies to each individual message
in the returned resource list. In case of a REST custom method,
other definitions may be used. Where the mask applies will be
clearly documented together with its declaration in the API.  In
any case, the effect on the returned resource/resources is required
behavior for APIs.

# Field Masks in Update Operations

A field mask in update operations specifies which fields of the
targeted resource are going to be updated. The API is required
to only change the values of the fields as specified in the mask
and leave the others untouched. If a resource is passed in to
describe the updated values, the API ignores the values of all
fields not covered by the mask.

If a repeated field is specified for an update operation, new values will
be appended to the existing repeated field in the target resource. Note that
a repeated field is only allowed in the last position of a `paths` string.

If a sub-message is specified in the last position of the field mask for an
update operation, then new value will be merged into the existing sub-message
in the target resource.

For example, given the target message:

    f {
      b {
        d: 1
        x: 2
      }
      c: [1]
    }

And an update message:

    f {
      b {
        d: 10
      }
      c: [2]
    }

then if the field mask is:

 paths: ["f.b", "f.c"]

then the result will be:

    f {
      b {
        d: 10
        x: 2
      }
      c: [1, 2]
    }

An implementation may provide options to override this default behavior for
repeated and message fields.

In order to reset a field's value to the default, the field must
be in the mask and set to the default value in the provided resource.
Hence, in order to reset all fields of a resource, provide a default
instance of the resource and set all fields in the mask, or do
not provide a mask as described below.

If a field mask is not present on update, the operation applies to
all fields (as if a field mask of all fields has been specified).
Note that in the presence of schema evolution, this may mean that
fields the client does not know and has therefore not filled into
the request will be reset to their default. If this is unwanted
behavior, a specific service may require a client to always specify
a field mask, producing an error if not.

As with get operations, the location of the resource which
describes the updated values in the request message depends on the
operation kind. In any case, the effect of the field mask is
required to be honored by the API.

## Considerations for HTTP REST

The HTTP kind of an update operation which uses a field mask must
be set to PATCH instead of PUT in order to satisfy HTTP semantics
(PUT must only be used for full updates).

# JSON Encoding of Field Masks

In JSON, a field mask is encoded as a single string where paths are
separated by a comma. Fields name in each path are converted
to/from lower-camel naming conventions.

As an example, consider the following message declarations:

    message Profile {
      User user = 1;
      Photo photo = 2;
    }
    message User {
      string display_name = 1;
      string address = 2;
    }

In proto a field mask for `Profile` may look as such:

    mask {
      paths: "user.display_name"
      paths: "photo"
    }

In JSON, the same mask is represented as below:

    {
      mask: "user.displayName,photo"
    }

# Field Masks and Oneof Fields

Field masks treat fields in oneofs just as regular fields. Consider the
following message:

    message SampleMessage {
      oneof test_oneof {
        string name = 4;
        SubMessage sub_message = 9;
      }
    }

The field mask can be:

    mask {
      paths: "name"
    }

Or:

    mask {
      paths: "sub_message"
    }

Note that oneof type names ("test_oneof" in this case) cannot be used in
paths.

## Field Mask Verification

The implementation of any API method which has a FieldMask type field in the
request should verify the included field paths, and return an
`INVALID_ARGUMENT` error if any path is unmappable.


| Field | Type | Description |
| ----- | ---- | ----------- |
| paths | [repeated string](#string) | The set of field mask paths. |
 <!-- end Fields -->
 <!-- end HasFields -->
 <!-- end messages -->

# Enums
 <!-- end Enums -->


 <!-- end services -->

# Messages


## SourceContext {#sourcecontext}
`SourceContext` represents information about the source of a
protobuf element, like the file in which it is defined.


| Field | Type | Description |
| ----- | ---- | ----------- |
| file_name | [ string](#string) | The path-qualified name of the .proto file that contained the associated protobuf element. For example: `"google/protobuf/source_context.proto"`. |
 <!-- end Fields -->
 <!-- end HasFields -->
 <!-- end messages -->

# Enums
 <!-- end Enums -->


 <!-- end services -->

# Messages


## ListValue {#listvalue}
`ListValue` is a wrapper around a repeated field of values.

The JSON representation for `ListValue` is JSON array.


| Field | Type | Description |
| ----- | ---- | ----------- |
| values | [repeated Value](#value) | Repeated field of dynamically typed values. |
 <!-- end Fields -->
 <!-- end HasFields -->


## Struct {#struct}
`Struct` represents a structured data value, consisting of fields
which map to dynamically typed values. In some languages, `Struct`
might be supported by a native representation. For example, in
scripting languages like JS a struct is represented as an
object. The details of that representation are described together
with the proto support for the language.

The JSON representation for `Struct` is JSON object.


| Field | Type | Description |
| ----- | ---- | ----------- |
| fields | [map Struct.FieldsEntry](#structfieldsentry) | Unordered map of dynamically typed values. |
 <!-- end Fields -->
 <!-- end HasFields -->


## Struct.FieldsEntry {#structfieldsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ Value](#value) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## Value {#value}
`Value` represents a dynamically typed value which can be either
null, a number, a string, a boolean, a recursive struct value, or a
list of values. A producer of value is expected to set one of these
variants. Absence of any variant indicates an error.

The JSON representation for `Value` is JSON value.


| Field | Type | Description |
| ----- | ---- | ----------- |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) kind.null_value | [ NullValue](#nullvalue) | Represents a null value. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) kind.number_value | [ double](#double) | Represents a double value. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) kind.string_value | [ string](#string) | Represents a string value. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) kind.bool_value | [ bool](#bool) | Represents a boolean value. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) kind.struct_value | [ Struct](#struct) | Represents a structured value. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) kind.list_value | [ ListValue](#listvalue) | Represents a repeated `Value`. |
 <!-- end Fields -->
 <!-- end HasFields -->
 <!-- end messages -->

# Enums


## NullValue {#nullvalue}
`NullValue` is a singleton enumeration to represent the null value for the
`Value` type union.

 The JSON representation for `NullValue` is JSON `null`.

| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL_VALUE | 0 | Null value. |


 <!-- end Enums -->


 <!-- end services -->

# Messages


## Timestamp {#timestamp}
A Timestamp represents a point in time independent of any time zone or local
calendar, encoded as a count of seconds and fractions of seconds at
nanosecond resolution. The count is relative to an epoch at UTC midnight on
January 1, 1970, in the proleptic Gregorian calendar which extends the
Gregorian calendar backwards to year one.

All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
second table is needed for interpretation, using a [24-hour linear
smear](https://developers.google.com/time/smear).

The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
restricting to that range, we ensure that we can convert to and from [RFC
3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.

# Examples

Example 1: Compute Timestamp from POSIX `time()`.

    Timestamp timestamp;
    timestamp.set_seconds(time(NULL));
    timestamp.set_nanos(0);

Example 2: Compute Timestamp from POSIX `gettimeofday()`.

    struct timeval tv;
    gettimeofday(&tv, NULL);

    Timestamp timestamp;
    timestamp.set_seconds(tv.tv_sec);
    timestamp.set_nanos(tv.tv_usec * 1000);

Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.

    FILETIME ft;
    GetSystemTimeAsFileTime(&ft);
    UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;

    // A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
    // is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
    Timestamp timestamp;
    timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
    timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));

Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.

    long millis = System.currentTimeMillis();

    Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
        .setNanos((int) ((millis % 1000) * 1000000)).build();

Example 5: Compute Timestamp from Java `Instant.now()`.

    Instant now = Instant.now();

    Timestamp timestamp =
        Timestamp.newBuilder().setSeconds(now.getEpochSecond())
            .setNanos(now.getNano()).build();

Example 6: Compute Timestamp from current time in Python.

    timestamp = Timestamp()
    timestamp.GetCurrentTime()

# JSON Mapping

In JSON format, the Timestamp type is encoded as a string in the
[RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the
format is "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z"
where {year} is always expressed using four digits while {month}, {day},
{hour}, {min}, and {sec} are zero-padded to two digits each. The fractional
seconds, which can go up to 9 digits (i.e. up to 1 nanosecond resolution),
are optional. The "Z" suffix indicates the timezone ("UTC"); the timezone
is required. A proto3 JSON serializer should always use UTC (as indicated by
"Z") when printing the Timestamp type and a proto3 JSON parser should be
able to accept both UTC and other timezones (as indicated by an offset).

For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past
01:30 UTC on January 15, 2017.

In JavaScript, one can convert a Date object to this format using the
standard
[toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
method. In Python, a standard `datetime.datetime` object can be converted
to this format using
[`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with
the time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use
the Joda Time's [`ISODateTimeFormat.dateTime()`](
http://www.joda.org/joda-time/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime%2D%2D
) to obtain a formatter capable of generating timestamps in this format.


| Field | Type | Description |
| ----- | ---- | ----------- |
| seconds | [ int64](#int64) | Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive. |
| nanos | [ int32](#int32) | Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. |
 <!-- end Fields -->
 <!-- end HasFields -->
 <!-- end messages -->

# Enums
 <!-- end Enums -->


 <!-- end services -->

# Messages


## Enum {#enum}
Enum type definition.


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | Enum type name. |
| enumvalue | [repeated EnumValue](#enumvalue) | Enum value definitions. |
| options | [repeated Option](#option) | Protocol buffer options. |
| source_context | [ SourceContext](#sourcecontext) | The source context. |
| syntax | [ Syntax](#syntax) | The source syntax. |
 <!-- end Fields -->
 <!-- end HasFields -->


## EnumValue {#enumvalue}
Enum value definition.


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | Enum value name. |
| number | [ int32](#int32) | Enum value number. |
| options | [repeated Option](#option) | Protocol buffer options. |
 <!-- end Fields -->
 <!-- end HasFields -->


## Field {#field}
A single field of a message type.


| Field | Type | Description |
| ----- | ---- | ----------- |
| kind | [ Field.Kind](#fieldkind) | The field type. |
| cardinality | [ Field.Cardinality](#fieldcardinality) | The field cardinality. |
| number | [ int32](#int32) | The field number. |
| name | [ string](#string) | The field name. |
| type_url | [ string](#string) | The field type URL, without the scheme, for message or enumeration types. Example: `"type.googleapis.com/google.protobuf.Timestamp"`. |
| oneof_index | [ int32](#int32) | The index of the field type in `Type.oneofs`, for message or enumeration types. The first type has index 1; zero means the type is not in the list. |
| packed | [ bool](#bool) | Whether to use alternative packed wire representation. |
| options | [repeated Option](#option) | The protocol buffer options. |
| json_name | [ string](#string) | The field JSON name. |
| default_value | [ string](#string) | The string value of the default value of this field. Proto2 syntax only. |
 <!-- end Fields -->
 <!-- end HasFields -->


## Option {#option}
A protocol buffer option, which can be attached to a message, field,
enumeration, etc.


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | The option's name. For protobuf built-in options (options defined in descriptor.proto), this is the short name. For example, `"map_entry"`. For custom options, it should be the fully-qualified name. For example, `"google.api.http"`. |
| value | [ Any](#any) | The option's value packed in an Any message. If the value is a primitive, the corresponding wrapper type defined in google/protobuf/wrappers.proto should be used. If the value is an enum, it should be stored as an int32 value using the google.protobuf.Int32Value type. |
 <!-- end Fields -->
 <!-- end HasFields -->


## Type {#type}
A protocol buffer message type.


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | The fully qualified message name. |
| fields | [repeated Field](#field) | The list of fields. |
| oneofs | [repeated string](#string) | The list of types appearing in `oneof` definitions in this type. |
| options | [repeated Option](#option) | The protocol buffer options. |
| source_context | [ SourceContext](#sourcecontext) | The source context. |
| syntax | [ Syntax](#syntax) | The source syntax. |
 <!-- end Fields -->
 <!-- end HasFields -->
 <!-- end messages -->

# Enums


## Field.Cardinality {#fieldcardinality}
Whether a field is optional, required, or repeated.

| Name | Number | Description |
| ---- | ------ | ----------- |
| CARDINALITY_UNKNOWN | 0 | For fields with unknown cardinality. |
| CARDINALITY_OPTIONAL | 1 | For optional fields. |
| CARDINALITY_REQUIRED | 2 | For required fields. Proto2 syntax only. |
| CARDINALITY_REPEATED | 3 | For repeated fields. |




## Field.Kind {#fieldkind}
Basic field types.

| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_UNKNOWN | 0 | Field type unknown. |
| TYPE_DOUBLE | 1 | Field type double. |
| TYPE_FLOAT | 2 | Field type float. |
| TYPE_INT64 | 3 | Field type int64. |
| TYPE_UINT64 | 4 | Field type uint64. |
| TYPE_INT32 | 5 | Field type int32. |
| TYPE_FIXED64 | 6 | Field type fixed64. |
| TYPE_FIXED32 | 7 | Field type fixed32. |
| TYPE_BOOL | 8 | Field type bool. |
| TYPE_STRING | 9 | Field type string. |
| TYPE_GROUP | 10 | Field type group. Proto2 syntax only, and deprecated. |
| TYPE_MESSAGE | 11 | Field type message. |
| TYPE_BYTES | 12 | Field type bytes. |
| TYPE_UINT32 | 13 | Field type uint32. |
| TYPE_ENUM | 14 | Field type enum. |
| TYPE_SFIXED32 | 15 | Field type sfixed32. |
| TYPE_SFIXED64 | 16 | Field type sfixed64. |
| TYPE_SINT32 | 17 | Field type sint32. |
| TYPE_SINT64 | 18 | Field type sint64. |




## Syntax {#syntax}
The syntax in which a protocol buffer element is defined.

| Name | Number | Description |
| ---- | ------ | ----------- |
| SYNTAX_PROTO2 | 0 | Syntax `proto2`. |
| SYNTAX_PROTO3 | 1 | Syntax `proto3`. |


 <!-- end Enums -->


 <!-- end services -->

# Messages


## BoolValue {#boolvalue}
Wrapper message for `bool`.

The JSON representation for `BoolValue` is JSON `true` and `false`.


| Field | Type | Description |
| ----- | ---- | ----------- |
| value | [ bool](#bool) | The bool value. |
 <!-- end Fields -->
 <!-- end HasFields -->


## BytesValue {#bytesvalue}
Wrapper message for `bytes`.

The JSON representation for `BytesValue` is JSON string.


| Field | Type | Description |
| ----- | ---- | ----------- |
| value | [ bytes](#bytes) | The bytes value. |
 <!-- end Fields -->
 <!-- end HasFields -->


## DoubleValue {#doublevalue}
Wrapper message for `double`.

The JSON representation for `DoubleValue` is JSON number.


| Field | Type | Description |
| ----- | ---- | ----------- |
| value | [ double](#double) | The double value. |
 <!-- end Fields -->
 <!-- end HasFields -->


## FloatValue {#floatvalue}
Wrapper message for `float`.

The JSON representation for `FloatValue` is JSON number.


| Field | Type | Description |
| ----- | ---- | ----------- |
| value | [ float](#float) | The float value. |
 <!-- end Fields -->
 <!-- end HasFields -->


## Int32Value {#int32value}
Wrapper message for `int32`.

The JSON representation for `Int32Value` is JSON number.


| Field | Type | Description |
| ----- | ---- | ----------- |
| value | [ int32](#int32) | The int32 value. |
 <!-- end Fields -->
 <!-- end HasFields -->


## Int64Value {#int64value}
Wrapper message for `int64`.

The JSON representation for `Int64Value` is JSON string.


| Field | Type | Description |
| ----- | ---- | ----------- |
| value | [ int64](#int64) | The int64 value. |
 <!-- end Fields -->
 <!-- end HasFields -->


## StringValue {#stringvalue}
Wrapper message for `string`.

The JSON representation for `StringValue` is JSON string.


| Field | Type | Description |
| ----- | ---- | ----------- |
| value | [ string](#string) | The string value. |
 <!-- end Fields -->
 <!-- end HasFields -->


## UInt32Value {#uint32value}
Wrapper message for `uint32`.

The JSON representation for `UInt32Value` is JSON number.


| Field | Type | Description |
| ----- | ---- | ----------- |
| value | [ uint32](#uint32) | The uint32 value. |
 <!-- end Fields -->
 <!-- end HasFields -->


## UInt64Value {#uint64value}
Wrapper message for `uint64`.

The JSON representation for `UInt64Value` is JSON string.


| Field | Type | Description |
| ----- | ---- | ----------- |
| value | [ uint64](#uint64) | The uint64 value. |
 <!-- end Fields -->
 <!-- end HasFields -->
 <!-- end messages -->

# Enums
 <!-- end Enums -->
 <!-- end Files -->

# Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <div><h4 id="double" /></div><a name="double" /> double |  | double | double | float |
| <div><h4 id="float" /></div><a name="float" /> float |  | float | float | float |
| <div><h4 id="int32" /></div><a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <div><h4 id="int64" /></div><a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <div><h4 id="uint32" /></div><a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <div><h4 id="uint64" /></div><a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <div><h4 id="sint32" /></div><a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <div><h4 id="sint64" /></div><a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <div><h4 id="fixed32" /></div><a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <div><h4 id="fixed64" /></div><a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <div><h4 id="sfixed32" /></div><a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <div><h4 id="sfixed64" /></div><a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <div><h4 id="bool" /></div><a name="bool" /> bool |  | bool | boolean | boolean |
| <div><h4 id="string" /></div><a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <div><h4 id="bytes" /></div><a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

