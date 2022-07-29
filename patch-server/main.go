package main

import (
    "net/http"
	"encoding/json"
	b64 "encoding/base64"
	//"net/http/httputil"
	// "log"
)

type Mutation struct {
	Version string	`json:"apiVersion"`
	Kind string		`json:"kind"`
	Response Auth 	`json:"response"`
}

type Auth struct {
	Uid string				`json:"uid"`
	Allow bool				`json:"allowed"`
	Patch string			`json:"patch"`
	Patchtype string		`json:"patchType"`
}

type Json struct {
	Request struct {
		Uid string 
		Object struct {
			Metadata struct {
				Labels *Labels
			}
		}
	}
}

type Labels struct {}

type Minimal struct {
	Version string			`json:"apiVersion"`
	Kind string				`json:"kind"`
	Response MinimalAuth 	`json:"response"`
}

type MinimalAuth struct {
	Uid string				`json:"uid"`
	Allow bool				`json:"allowed"`
}

func main() {
	http.HandleFunc("/mutate", mutate)
	http.ListenAndServeTLS(":8443", "/ssl/server.crt", "/ssl/server.key", nil)
}

func mutate(w http.ResponseWriter, r *http.Request) {

	var info Json 
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		return
	}
	uid := info.Request.Uid
	// if there is a label we can send a patch, otherwise send k8s a minimal response allowing the creation
	if (info.Request.Object.Metadata.Labels != nil) {
		patch := `[{"op": "add", "path": "/metadata/labels/powered-by", "value": "slixperi says this could have been a total yaml re-write"}]`
		patch = b64.StdEncoding.EncodeToString([]byte(patch))

		mutation := Mutation {
			Version: "admission.k8s.io/v1",
			Kind: "AdmissionReview",
			Response: Auth{
				Uid: uid,
				Allow: true,
				Patch: patch,
				Patchtype: "JSONPatch",
			},
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mutation)

	} else {
		minimalMutation := Minimal {
			Version: "admission.k8s.io/v1",
			Kind: "AdmissionReview",
			Response: MinimalAuth{
				Uid: uid,
				Allow: true,
			},
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(minimalMutation)
	}
}