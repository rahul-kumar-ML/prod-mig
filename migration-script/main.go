package main

import (
	"encoding/json"
	"gorm.io/gorm"
	"fmt"
)

func updateDatabaseAzure(db *gorm.DB, sink KConnectSink) error {
	containerName := "archive-" + sink.ProfileID + "-profile"
	
	// Unmarshal the JSON data into a map
	var configMap map[string]interface{}
	if err := json.Unmarshal([]byte(sink.Config), &configMap); err != nil {
		return err
	}
	// Update the specific fields in the map
	configMap["format.class"] = "io.confluent.connect.azureblob.format.orc.OrcFormat"
	configMap["storage.class"] = "io.confluent.connect.azureblob.storage.AzureBlobStorage"
	configMap["connector.class"] = "io.confluent.connect.azureblob.AzureBlobSinkConnector"
	configMap["azureblob.account.key"] = "newAccountKey"
	configMap["azureblob.container.name"] = containerName
	configMap["azureblob.storage.account.name"] = "apmmanagerstorageacc"
	configMap["azureblob.block.size"] = 5242880

	// Convert updateData to JSON
	updateJSON, err := json.Marshal(configMap)
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
			err := updateDatabaseAzure(db, sink)
			if err != nil {
				fmt.Printf("unable to update db for id=?",sink.ID)
			}
		}
	} else if config.CloudType == "AWS" {
		//cloud specific function call
	} else if config.CloudType == "GCP" {
		//cloud specific function call
	} else {
		panic("Unknown cloud type")
	}
}
