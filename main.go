package main

import (
	"encoding/json"
	schema "github.com/Jamshid90/fhir-schema"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/resource", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
			case "POST":
				CreateResource(w, r)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}

func CreateResource(w http.ResponseWriter, r *http.Request)  {

	result  := Result{Success:false}
	appointment := schema.Appointment{}

	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		result.Message = err.Error()
		json.NewEncoder(w).Encode(result)
		return
	}

	result.Success = true
	result.Data = appointment

	json.NewEncoder(w).Encode(result)
	return
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Result struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
	Errors map[string]string `json:"errors"`
}


/*
{
  "resourceType": "Appointment",
  "id": "examplereq",
  "text": {
    "status": "generated",
    "div": "<div xmlns=\"http://www.w3.org/1999/xhtml\">Brian MRI results discussion</div>"
  },
  "identifier": [
    {
      "system": "http://example.org/sampleappointment-identifier",
      "value": "123"
    }
  ],
  "status": "proposed",
  "serviceCategory": [
    {
      "coding": [
        {
          "system": "http://example.org/service-category",
          "code": "gp",
          "display": "General Practice"
        }
      ]
    }
  ],
  "specialty": [
    {
      "coding": [
        {
          "system": "http://snomed.info/sct",
          "code": "394814009",
          "display": "General practice"
        }
      ]
    }
  ],
  "appointmentType": {
    "coding": [
      {
        "system": "http://terminology.hl7.org/CodeSystem/v2-0276",
        "code": "WALKIN",
        "display": "A previously unscheduled walk-in visit"
      }
    ]
  },
  "reasonCode": [
    {
      "coding": [
        {
          "system": "http://snomed.info/sct",
          "code": "413095006"
        }
      ],
      "text": "Clinical Review"
    }
  ],
  "priority": 5,
  "description": "Discussion on the results of your recent MRI",
  "minutesDuration": 15,
  "slot": [
    {
      "reference": "Slot/example"
    }
  ],
  "created": "2015-12-02",
  "comment": "Further expand on the results of the MRI and determine the next actions that may be appropriate.",
  "participant": [
    {
      "actor": {
        "reference": "Patient/example",
        "display": "Peter James Chalmers"
      },
      "required": "required",
      "status": "needs-action"
    },
    {
      "type": [
        {
          "coding": [
            {
              "system": "http://terminology.hl7.org/CodeSystem/v3-ParticipationType",
              "code": "ATND"
            }
          ]
        }
      ],
      "required": "required",
      "status": "needs-action"
    },
    {
      "actor": {
        "reference": "Location/1",
        "display": "South Wing, second floor"
      },
      "required": "required",
      "status": "accepted"
    }
  ],
  "requestedPeriod": [
    {
      "start": "2016-06-02",
      "end": "2016-06-09"
    }
  ]
}
*/