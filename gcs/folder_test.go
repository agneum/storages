package gcs

import (
	"errors"
	"testing"

	gcs "cloud.google.com/go/storage"

	"github.com/wal-g/storages/storage"

	"github.com/stretchr/testify/assert"
)

func TestGSFolder(t *testing.T) {
	t.Skip("Credentials needed to run GCP tests")

	storageFolder, err := ConfigureFolder("gs://x4m-test/walg-bucket",
		nil)

	assert.NoError(t, err)

	storage.RunFolderTest(storageFolder, t)
}

func TestGSExactFolder(t *testing.T) {
	t.Skip("Credentials needed to run GCP tests")

	//os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/Users/x4mmm/Downloads/mdb-tests-d0uble-0b98813b622b.json")
	//os.Setenv("GCS_CONTEXT_TIMEOUT", "1024000000")

	storageFolder, err := ConfigureFolder("gs://x4m-test//walg-bucket////strange_folder",
		map[string]string{
			NormalizePrefix: "false",
		})

	assert.NoError(t, err)

	storage.RunFolderTest(storageFolder, t)
}

type fakeReader struct{}

func (f fakeReader) Read(_ []byte) (int, error) {
	return 0, errors.New("failed to fake read")
}

func TestUploadingReaderFails(t *testing.T) {
	folder := Folder{
		bucket: &gcs.BucketHandle{},
		path:   "path",
	}

	err := folder.PutObject("name", fakeReader{})
	assert.EqualError(t, err, "GCS error : Unable to copy to object: failed to fake read")
}
