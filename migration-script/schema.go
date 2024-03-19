package main

import (
    "encoding/json" 
)

// TableName returned will be used to create the table in db
func (KConnectSink) TableName() string {
	return "kconnect_sinks"
}

type KConnectSink struct {
	ID              uint            `gorm:"primary_key"`
	Name            string          `gorm:"column:name"`
	ProfileID       string          `gorm:"column:profileid"`
	Type            string          `gorm:"column:type"`
	ConnectURL      string          `gorm:"column:connect_url"`
	MaxTasks        int             `gorm:"column:max_tasks"`
	Config          json.RawMessage `gorm:"type:jsonb;column:config"`
	State           string          `gorm:"column:state"`
	TrueState       string          `gorm:"column:true_state;default:RUNNING"`
	LastKnownError  string          `gorm:"column:last_known_error;default:''"`
	UserManaged     bool            `gorm:"column:user_managed;default:false"`
	ByteRateQuota   int             `gorm:"column:byte_rate_quota"`
	QuotaUsageStats json.RawMessage `gorm:"type:jsonb;column:quota_usage_stats"`
}
