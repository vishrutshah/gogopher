package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You just browsed page (if blank you're at the new root): %s", r.URL.Path[1:])
	fmt.Fprintf(w, "\nGOROOT: %s", os.Getenv("GOROOT"))
	fmt.Fprintf(w, "\nGOPATH: %s", os.Getenv("GOPATH"))
	fmt.Fprintf(w, "\nGO15VENDOREXPERIMENT: %s", os.Getenv("GO15VENDOREXPERIMENT"))
	// createAppendBlobExample()
	// Create Storage Client
	fmt.Fprintf(w, "\n Just before calling NewBasicService")
	// _, err := storage.NewBasicClient("vishrutrg", "KEY")
	// if err != nil {
	// 	fmt.Fprintf(w, "Error page (if blank you're at the root): %v", err)
	// } else {
	// 	fmt.Fprintf(w, "Not Error page (if blank you're at the root): %v", err)
	// }
	fmt.Fprintf(w, "\n Back from calling NewBasicService")
}

func createAppendBlobExample() {

	// 	// Create Blob Service
	// 	blobService := storageClient.GetBlobService()

	// 	containerName := "container1" // Container must exists beforehand
	// 	blobName := "blob_name"       // blobName name must not exists
	// 	blobContent := []byte("Hello, Gophers!")

	// 	// Create Append Blob & Append blob content
	// 	blobService.PutAppendBlob(containerName, blobName, nil)
	// 	blobService.AppendBlock(containerName, blobName, blobContent, nil)
	// 	if err != nil {
	// 		fmt.Printf("Error %v", err)
	// 	}

	// 	// Read the blob content back
	// 	out, err := blobService.GetBlobRange(containerName, blobName, fmt.Sprintf("%v-%v", 0, len(blobContent)-1))
	// 	if err != nil {
	// 		fmt.Printf("Error %v", err)
	// 	}
	// 	defer out.Close()
	// 	blobContents, err1 := ioutil.ReadAll(out)
	// 	if err1 != nil {
	// 		fmt.Printf("Error %v", err1)
	// 	}
	// 	fmt.Printf("Content %v", string(blobContents))
	// 	out.Close()
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("Port %s", os.Getenv("HTTP_PLATFORM_PORT"))
	http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil)
}
