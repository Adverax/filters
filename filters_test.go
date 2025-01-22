package filters

import (
	"testing"
)

func TestBuilder_AllowExact(t *testing.T) {
	filter, err := NewBuilder().
		AllowExact("example").
		Build()
	if err != nil {
		t.Fatalf("Failed to build filter: %v", err)
	}

	if !filter.IsMatch("example") {
		t.Errorf("Expected 'example' to be allowed")
	}
	if filter.IsMatch("test") {
		t.Errorf("Expected 'test' to be denied")
	}
}

func TestBuilder_DenyExact(t *testing.T) {
	filter, err := NewBuilder().
		DenyExact("test").
		Build()
	if err != nil {
		t.Fatalf("Failed to build filter: %v", err)
	}

	if filter.IsMatch("test") {
		t.Errorf("Expected 'test' to be denied")
	}
	if !filter.IsMatch("example") {
		t.Errorf("Expected 'example' to be allowed")
	}
}

func TestBuilder_AllowRegexp(t *testing.T) {
	filter, err := NewBuilder().
		AllowRegexp(`^example.*`).
		Build()
	if err != nil {
		t.Fatalf("Failed to build filter: %v", err)
	}

	if !filter.IsMatch("example123") {
		t.Errorf("Expected 'example123' to be allowed")
	}
	if filter.IsMatch("test123") {
		t.Errorf("Expected 'test123' to be denied")
	}
}

func TestBuilder_DenyRegexp(t *testing.T) {
	filter, err := NewBuilder().
		DenyRegexp(`^test.*`).
		Build()
	if err != nil {
		t.Fatalf("Failed to build filter: %v", err)
	}

	if filter.IsMatch("test123") {
		t.Errorf("Expected 'test123' to be denied")
	}
	if !filter.IsMatch("example123") {
		t.Errorf("Expected 'example123' to be allowed")
	}
}

func TestBuilder_AllowPrefix(t *testing.T) {
	filter, err := NewBuilder().
		AllowPrefix("example").
		Build()
	if err != nil {
		t.Fatalf("Failed to build filter: %v", err)
	}

	if !filter.IsMatch("example123") {
		t.Errorf("Expected 'example123' to be allowed")
	}
	if filter.IsMatch("test123") {
		t.Errorf("Expected 'test123' to be denied")
	}
}

func TestBuilder_DenyPrefix(t *testing.T) {
	filter, err := NewBuilder().
		DenyPrefix("test").
		Build()
	if err != nil {
		t.Fatalf("Failed to build filter: %v", err)
	}

	if filter.IsMatch("test123") {
		t.Errorf("Expected 'test123' to be denied")
	}
	if !filter.IsMatch("example123") {
		t.Errorf("Expected 'example123' to be allowed")
	}
}

func TestBuilder_AllowSuffix(t *testing.T) {
	filter, err := NewBuilder().
		AllowSuffix("example").
		Build()
	if err != nil {
		t.Fatalf("Failed to build filter: %v", err)
	}

	if !filter.IsMatch("123example") {
		t.Errorf("Expected '123example' to be allowed")
	}
	if filter.IsMatch("123test") {
		t.Errorf("Expected '123test' to be denied")
	}
}

func TestBuilder_DenySuffix(t *testing.T) {
	filter, err := NewBuilder().
		DenySuffix("test").
		Build()
	if err != nil {
		t.Fatalf("Failed to build filter: %v", err)
	}

	if filter.IsMatch("123test") {
		t.Errorf("Expected '123test' to be denied")
	}
	if !filter.IsMatch("123example") {
		t.Errorf("Expected '123example' to be allowed")
	}
}
