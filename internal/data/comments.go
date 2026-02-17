package data

import (
  "time"
)

// each name begins with uppercase so that they are exportable/public
type Class struct {
    ID int64
    Studio  string 
    Trainer  string 
    Capacity_limit  int64 
    Membership_tier  string
    Name  string
    Status  string
    CreatedAt  time.Time
} 
