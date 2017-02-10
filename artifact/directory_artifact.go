package artifact

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/pivotal-cf/pcf-backup-and-restore/orchestrator"
)

const TAG = "[artifact]"

type DirectoryArtifact struct {
	orchestrator.Logger
	baseDirName string
}

func (d *DirectoryArtifact) DeploymentMatches(deployment string, instances []orchestrator.Instance) (bool, error) {
	_, err := d.metadataExistsAndIsReadable()
	if err != nil {
		d.Debug(TAG, "Error checking metadata file: %v", err)
		return false, err
	}
	meta, err := readMetadata(d.metadataFilename())
	if err != nil {
		d.Debug(TAG, "Error reading metadata file: %v", err)
		return false, err
	}

	for _, inst := range meta.MetadataForEachInstance {
		present := d.backupInstanceIsPresent(inst, instances)
		if present != true {
			d.Debug(TAG, "Instance %v/%v not found in %v", inst.Name(), inst.Index(), instances)
			return false, nil
		}
	}

	return true, nil
}

func (d *DirectoryArtifact) CreateFile(artifactIdentifer orchestrator.BackupBlobIdentifier) (io.WriteCloser, error) {
	d.Debug(TAG, "Trying to create file %s", fileName(artifactIdentifer))
	return os.Create(path.Join(d.baseDirName, fileName(artifactIdentifer)))
}

func (d *DirectoryArtifact) ReadFile(artifactIdentifer orchestrator.BackupBlobIdentifier) (io.ReadCloser, error) {
	filename := d.instanceFilename(artifactIdentifer)
	d.Debug(TAG, "Trying to open %s", filename)
	file, err := os.Open(filename)
	if err != nil {
		d.Debug(TAG, "Error reading artifact file %s", filename)
		return nil, err
	}

	return file, nil
}

func (d *DirectoryArtifact) FetchChecksum(artifactIdentifer orchestrator.BackupBlobIdentifier) (orchestrator.BackupChecksum, error) {
	metadata, err := readMetadata(d.metadataFilename())

	if err != nil {
		d.Debug(TAG, "Error reading metadata from %s %v", d.metadataFilename(), err)
		return nil, err
	}

	if artifactIdentifer.IsNamed() {
		for _, instanceInMetadata := range metadata.MetadataForEachArtifact {
			if instanceInMetadata.Name() == artifactIdentifer.Name() {
				return instanceInMetadata.Checksum, nil
			}
		}
	} else {
		for _, instanceInMetadata := range metadata.MetadataForEachInstance {
			if instanceInMetadata.Index() == artifactIdentifer.Index() && instanceInMetadata.Name() == artifactIdentifer.Name() {
				return instanceInMetadata.Checksum, nil
			}
		}
	}

	d.Warn(TAG, "Checksum for %s not found in artifact", logName(artifactIdentifer))
	return nil, nil
}
func logName(artifactIdentifer orchestrator.BackupBlobIdentifier) string {
	if artifactIdentifer.IsNamed() {
		return fmt.Sprintf("%s", artifactIdentifer.Name())
	}
	return fmt.Sprintf("%s/%s", artifactIdentifer.Name(), artifactIdentifer.Index())
}

func (d *DirectoryArtifact) CalculateChecksum(inst orchestrator.BackupBlobIdentifier) (orchestrator.BackupChecksum, error) {
	file, err := d.ReadFile(inst)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	gzipedReader, err := gzip.NewReader(file)
	if err != nil {
		d.Debug(TAG, "Cant open gzip for %s %v", logName(inst), err)
		return nil, err
	}
	tarReader := tar.NewReader(gzipedReader)
	checksum := map[string]string{}
	for {
		tarHeader, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			d.Debug(TAG, "Error reading tar for %s %v", logName(inst), err)
			return nil, err
		}
		if tarHeader.FileInfo().IsDir() || tarHeader.FileInfo().Name() == "./" {
			continue
		}

		fileShasum := sha1.New()
		if _, err := io.Copy(fileShasum, tarReader); err != nil {
			d.Debug(TAG, "Error calculating sha for %s %v", logName(inst), err)
			return nil, err
		}
		checksum[tarHeader.Name] = fmt.Sprintf("%x", fileShasum.Sum(nil))
	}

	return checksum, nil
}
func (d *DirectoryArtifact) AddChecksum(artifactIdentifer orchestrator.BackupBlobIdentifier, shasum orchestrator.BackupChecksum) error {
	metadata := metadata{}
	if exists, _ := d.metadataExistsAndIsReadable(); exists {
		var err error
		metadata, err = readMetadata(d.metadataFilename())
		if err != nil {
			d.Debug(TAG, "Error reading metadata from %s %v", d.metadataFilename(), err)
			return err
		}
	}

	if artifactIdentifer.IsNamed() {
		metadata.MetadataForEachArtifact = append(metadata.MetadataForEachArtifact, artifactMetadata{
			ArtifactName: artifactIdentifer.Name(),
			Checksum:     shasum,
		})
	} else {
		metadata.MetadataForEachInstance = append(metadata.MetadataForEachInstance, instanceMetadata{
			InstanceName:  artifactIdentifer.Name(),
			InstanceIndex: artifactIdentifer.Index(),
			Checksum:      shasum,
		})
	}

	return metadata.save(d.metadataFilename())
}

func (d *DirectoryArtifact) SaveManifest(manifest string) error {
	return ioutil.WriteFile(d.manifestFilename(), []byte(manifest), 0666)
}

func (d *DirectoryArtifact) Valid() (bool, error) {
	meta, err := readMetadata(d.metadataFilename())
	if err != nil {
		d.Debug(TAG, "Error reading metadata from %s %v", d.metadataFilename(), err)
		return false, err
	}

	for _, blob := range meta.MetadataForEachArtifact {
		actualBlobChecksum, _ := d.CalculateChecksum(blob)
		if !actualBlobChecksum.Match(blob.Checksum) {
			d.Debug(TAG, "Can't match checksums for %s, in metadata: %v, in actual file: %v", blob.Name(), actualBlobChecksum, blob.Checksum)
			return false, nil
		}
	}

	for _, inst := range meta.MetadataForEachInstance {
		actualInstanceChecksum, err := d.CalculateChecksum(inst)
		if err != nil {
			return false, err
		}
		if !actualInstanceChecksum.Match(inst.Checksum) {
			d.Debug(TAG, "Can't match checksums for %s, in metadata: %v, in actual file: %v", logName(inst), actualInstanceChecksum, inst.Checksum)
			return false, nil
		}

	}
	return true, nil
}

func (d *DirectoryArtifact) backupInstanceIsPresent(backupInstance instanceMetadata, instances []orchestrator.Instance) bool {
	for _, inst := range instances {
		if inst.Index() == backupInstance.InstanceIndex && inst.Name() == backupInstance.InstanceName {
			return true
		}
	}
	return false
}

func (d *DirectoryArtifact) instanceFilename(artifactIdentifer orchestrator.BackupBlobIdentifier) string {
	return path.Join(d.baseDirName, fileName(artifactIdentifer))
}

func (d *DirectoryArtifact) metadataFilename() string {
	return path.Join(d.baseDirName, "metadata")
}

func (d *DirectoryArtifact) manifestFilename() string {
	return path.Join(d.baseDirName, "manifest.yml")
}
func (d *DirectoryArtifact) metadataExistsAndIsReadable() (bool, error) {
	_, err := os.Stat(d.metadataFilename())
	if err != nil {
		return false, err
	}
	return true, nil
}

func fileName(artifactIdentifer orchestrator.BackupBlobIdentifier) string {
	if artifactIdentifer.IsNamed() {
		return artifactIdentifer.Name() + ".tgz"
	}

	return artifactIdentifer.Name() + "-" + artifactIdentifer.Index() + ".tgz"
}
