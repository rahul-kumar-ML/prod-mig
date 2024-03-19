package main

import (
	"encoding/json"
	"gorm.io/gorm"
)

func updateDatabaseAzure(db *gorm.DB, sink KConnectSink) error {
	containerName := "archive-" + sink.ProfileID + "-profile"
	// Define a map to represent the update data
	updateData := map[string]interface{}{
		"format.class":                   "io.confluent.connect.azureblob.format.orc.OrcFormat",
		"storage.class":                  "io.confluent.connect.azureblob.storage.AzureBlobStorage",
		"connector.class":                "io.confluent.connect.azureblob.AzureBlobSinkConnector",
		"azureblob.account.key":          "JWk+PUNVStyk3ogVowgxxJRZVhPwxEFP4ffzABm74u2ZlA4eMICmgNmQw+94udrpGZTHmsXlwzE1+AStBYj7nw==",
		"azureblob.container.name":       containerName,
		"azureblob.storage.account.name": "apmmanagerstorageacc",
		"azureblob.block.size":           5242880,
	}

	// Convert updateData to JSON
	updateJSON, err := json.Marshal(updateData)
	if err != nil {
		return err
	}

	// Perform the update
	result := db.Model(&KConnectSink{}).Where("name LIKE ?", "s3%").Where("id = ?", sink.ID).Update("config", updateJSON)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func main() {
	config := InitConfig()
	db := InitDB(config)
	var sinks []KConnectSink
	db.Model(&KConnectSink{}).Where("name LIKE ?", "s3%").Find(&sinks)
	if config.CloudType == "Azure" {
		for _, sink := range sinks {
			updateDatabaseAzure(db, sink)
		}
	} else if config.CloudType == "AWS" {
		//cloud specific function call
	} else if config.CloudType == "GCP" {
		//cloud specific function call
	} else {
		panic("Unknown cloud type")
	}
}
