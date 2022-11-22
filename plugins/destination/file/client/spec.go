package client

type BackendType string
type FormatType string

const (
	BackendTypeLocal BackendType = "local"
	BackendTypeS3    BackendType = "s3"
	BackendTypeGCS   BackendType = "gcs"

	FormatTypeCSV  FormatType = "csv"
	FormatTypeJSON FormatType = "json"

	batchSize = uint64(250)
)

type Spec struct {
	Directory string      `json:"directory,omitempty"`
	Backend   BackendType `json:"backend,omitempty"`
	Format    FormatType  `json:"format,omitempty"`
	BatchSize uint64      `json:"batchSize,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.Directory == "" {
		s.Directory = "./cq_csv_output"
	}
	if s.Backend == "" {
		s.Backend = BackendTypeLocal
	}
	if s.Format == "" {
		s.Format = FormatTypeCSV
	}
	if s.BatchSize == 0 {
		s.BatchSize = batchSize
	}
}

func (*Spec) Validate() error {
	return nil
}
