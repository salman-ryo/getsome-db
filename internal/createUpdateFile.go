package internal

import (
	"fmt"
	"getsome-db/utils"
	"os"
)

func CreateUpdateFile(path string, data []byte) error {

	// Create temp file name for a in memory file with latest data which will replace existing file using rename
	tempFile := fmt.Sprintf("%s.temp.%d", path, utils.RandomUint32())

	// open the file in memory
	memFile, err := os.OpenFile(tempFile, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}

	// Must close the file on unmount, also on error need to remove the created file as it could be corrupted
	defer func() {
		memFile.Close()

		if err != nil {
			os.Remove(tempFile)
		}
	}()

	// proceed to write data to temp file
	if _, err := memFile.Write(data); err != nil {
		return fmt.Errorf("failed to write to temp file: %w", err)
	}

	// Ensure all data is written to the disk, use file sync, since we opened memFile with tempFile name, it's the same now
	if err := memFile.Sync(); err != nil {
		return fmt.Errorf("failed to sync temp file: %w", err)
	}

	// Final step is to rename the temp file to the original, effectively replacing it as the name now points to new file
	if err := os.Rename(tempFile, path); err != nil {
		return fmt.Errorf("failed to rename temp file: %w", err)
	}

	// Everything went well, return nil for error
	return nil
}
