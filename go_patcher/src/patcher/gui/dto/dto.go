package dto

type GeneratePayload struct {
	NewVersion string
	PrecVersion string
	Patch string
	Crc32 string
} 	

type Request struct {
    RequestType string
    Data string
}

