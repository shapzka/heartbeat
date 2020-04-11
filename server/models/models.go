package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movement struct {

  ID                   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
  Title                string             `json:"title,omitempty"`
  Summary              string              `json:"summary,omitempty"`
  NumberOfParticipants int                 `json:"numberOfParticipants,omitempty"`
  Latitude             float64             `json:"latitude,omitempty"`
  Longitude            float64             `json:"longitude,omitempty"`

}
