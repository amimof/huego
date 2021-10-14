package huego

// StrPtr returns pointer to a string.
func StrPtr(val string) *string {
    return &val
}

// BoolPtr returns pointer to an boolean.
func BoolPtr(val bool) *bool {
    return &val
}

// IntPtr returns pointer to an integer.
func IntPtr(val int) *int {
    return &val
}

// Uint8Ptr returns pointer to an uint8.
func Uint8Ptr(val uint8) *uint8 {
    return &val
}

// Uint16Ptr returns pointer to an uint16.
func Uint16Ptr(val uint16) *uint16 {
    return &val
}

// Float32Ptr returns pointer to a float32.
func Float32Ptr(val float32) *float32 {
    return &val
}

// Float64Ptr returns pointer to a float64.
func Float64Ptr(val float64) *float64 {
    return &val
}
