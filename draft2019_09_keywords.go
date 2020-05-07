package main

func LoadDraft2019_09() {
	// core keywords
	RegisterKeyword("$schema", NewSchemaURI)
	RegisterKeyword("$id", NewId)
	RegisterKeyword("description", NewDescription)
	RegisterKeyword("title", NewTitle)
	RegisterKeyword("$comment", NewComment)
	RegisterKeyword("examples", NewExamples)
	RegisterKeyword("readOnly", NewReadOnly)
	RegisterKeyword("writeOnly", NewWriteOnly)
	RegisterKeyword("$ref", NewRef)
	RegisterKeyword("$recursiveRef", NewRecursiveRef)
	RegisterKeyword("$anchor", NewAnchor)
	RegisterKeyword("$recursiveAnchor", NewRecursiveAnchor)
	RegisterKeyword("$defs", NewDefs)
	RegisterKeyword("default", NewDefault)

	SetKeywordOrder("$ref", 0)
	SetKeywordOrder("$recursiveRef", 0)

	// standard keywords
	RegisterKeyword("type", NewType)
	RegisterKeyword("enum", NewEnum)
	RegisterKeyword("const", NewConst)

	// numeric keywords
	RegisterKeyword("multipleOf", NewMultipleOf)
	RegisterKeyword("maximum", NewMaximum)
	RegisterKeyword("exclusiveMaximum", NewExclusiveMaximum)
	RegisterKeyword("minimum", NewMinimum)
	RegisterKeyword("exclusiveMinimum", NewExclusiveMinimum)

	// string keywords
	RegisterKeyword("maxLength", NewMaxLength)
	RegisterKeyword("minLength", NewMinLength)
	RegisterKeyword("pattern", NewPattern)

	// boolean keywords
	RegisterKeyword("allOf", NewAllOf)
	RegisterKeyword("anyOf", NewAnyOf)
	RegisterKeyword("oneOf", NewOneOf)
	RegisterKeyword("not", NewNot)

	// object keywords
	RegisterKeyword("properties", NewProperties)
	RegisterKeyword("patternProperties", NewPatternProperties)
	RegisterKeyword("additionalProperties", NewAdditionalProperties)
	RegisterKeyword("required", NewRequired)
	RegisterKeyword("propertyNames", NewPropertyNames)
	RegisterKeyword("maxProperties", NewMaxProperties)
	RegisterKeyword("minProperties", NewMinProperties)
	RegisterKeyword("dependentSchemas", NewDependentSchemas)
	RegisterKeyword("dependentRequired", NewDependentRequired)

	SetKeywordOrder("properties", 2)
	SetKeywordOrder("additionalProperties", 3)

	// array keywords
	RegisterKeyword("items", NewItems)
	RegisterKeyword("additionalItems", NewAdditionalItems)
	RegisterKeyword("maxItems", NewMaxItems)
	RegisterKeyword("minItems", NewMinItems)
	RegisterKeyword("uniqueItems", NewUniqueItems)
	RegisterKeyword("contains", NewContains)
	RegisterKeyword("maxContains", NewMaxContains)
	RegisterKeyword("minContains", NewMinContains)

	SetKeywordOrder("maxContains", 2)
	SetKeywordOrder("minContains", 2)
	SetKeywordOrder("additionalItems", 3)

	// conditional keywords
	RegisterKeyword("if", NewIf)
	RegisterKeyword("then", NewThen)
	RegisterKeyword("else", NewElse)

	SetKeywordOrder("then", 2)
	SetKeywordOrder("else", 2)

	//optional formats
	RegisterKeyword("format", NewFormat)
}
