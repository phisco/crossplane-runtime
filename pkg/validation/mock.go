package validation

import (
	"fmt"
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/resource/unstructured/composite"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/apimachinery/pkg/runtime"
	"strings"
)

func MockRequiredFields(res *composite.Unstructured, s *apiextensions.JSONSchemaProps) error {
	o, err := fieldpath.PaveObject(res)
	if err != nil {
		return err
	}
	err = mockRequiredFieldsSchemaProps(s, o, "")
	if err != nil {
		return err
	}
	return runtime.DefaultUnstructuredConverter.FromUnstructured(o.UnstructuredContent(), res)

}

// mockRequiredFieldsSchemaPropos mock required fields for a given schema property
func mockRequiredFieldsSchemaProps(prop *apiextensions.JSONSchemaProps, o *fieldpath.Paved, path string) error {
	if prop == nil {
		return nil
	}
	switch prop.Type {
	case "string":
		if prop.Default == nil {
			return setTypeDefaultValue(o, path, prop.Type)
		}
		v := *prop.Default
		vs, ok := v.(string)
		if !ok {
			return fmt.Errorf("default value for %s is not a string", path)
		}
		return o.SetString(path, vs)
	case "integer", "number":
		if prop.Default == nil {
			return setTypeDefaultValue(o, path, prop.Type)
		}
		v := *prop.Default
		vs, ok := v.(float64)
		if !ok {
			return fmt.Errorf("default value for %s is not an integer", path)
		}
		return o.SetNumber(path, vs)
	case "object":
		for _, s := range prop.Required {
			p := prop.Properties[s]
			err := mockRequiredFieldsSchemaProps(&p, o, strings.TrimLeft(strings.Join([]string{path, s}, "."), "."))
			if err != nil {
				return err
			}
		}
		return nil
	case "array":
		return nil
	}
	return nil
}

// setTypeDefaultValue sets the default value for a given type at a given path
func setTypeDefaultValue(o *fieldpath.Paved, path string, t string) error {
	switch t {
	case "string":
		return o.SetString(path, "default")
	case "integer":
		return o.SetNumber(path, 1)
	}
	return nil
}
