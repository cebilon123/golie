package cfg

import "testing"

func TestConfiguration_OverwriteWithDifferentFields(t *testing.T) {
	type fields struct {
		Path string
	}
	type args struct {
		c2 *Configuration
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected string
	}{
		{name: "c2_configuration_empty_fields", fields: fields{Path: configurationFileName}, args: struct{ c2 *Configuration }{c2: &Configuration{Path: ""}}, expected: configurationFileName},
		{name: "c2_configuration_different_fields", fields: fields{Path: configurationFileName}, args: struct{ c2 *Configuration }{c2: &Configuration{Path: "hello"}}, expected: "hello"},
		{name: "c2_configuration_empty_first_struct_second_struct_set", fields: fields{Path: ""}, args: struct{ c2 *Configuration }{c2: &Configuration{Path: "hello"}}, expected: "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Configuration{
				Path: tt.fields.Path,
			}

			c.OverwriteWithDifferentFields(tt.args.c2)

			if len(c.Path) == 0 || c.Path != tt.expected {
				t.Fatalf("expected: %v \n got: %v", tt.expected, c.Path)
			}
		})
	}
}
