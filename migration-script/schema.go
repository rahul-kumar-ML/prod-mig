package main

// TableName returned will be used to create the table in db
func (KConnectSink) TableName() string {
	return "kconnect_sinks"
}

type KConnectSink struct {
	ID              uint `gorm:"primary_key"`
	Name            string
	ProfileID       string
	Type            string
	ConnectURL      string
	MaxTasks        int
	Config          string `gorm:"type:jsonb"`
	State           string
	TrueState       string
	LastKnownError  string
	UserManaged     bool
	ByteRateQuota   int
	QuotaUsageStats string
}
