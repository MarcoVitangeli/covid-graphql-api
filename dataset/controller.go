package dataset

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type service interface {
	LoadDataset(ctx context.Context, ls []string) error
}

func HandleDataLoad(s service) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		url := os.Getenv("DATASET_URL")
		if url == "" {
			panic("error: missing dataset url")
		}

		res, err := http.Get(url)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			msg := map[string]string{"message": "error fetching data from dataset"}
			err := json.NewEncoder(writer).Encode(msg)
			if err != nil {
				log.Println(fmt.Errorf("error enconding json body: %w", err))
			}
			return
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			msg := map[string]string{"message": "error reading response body"}
			err := json.NewEncoder(writer).Encode(msg)
			if err != nil {
				log.Println(fmt.Errorf("error enconding json body: %w", err))
			}
			return
		}
		code := http.StatusOK
		ls := strings.Split(string(body), "\n")
		// we ignore the first element of "ls" beacause it is the column names row
		// we also remove the last one, because it is an empty string
		if err := s.LoadDataset(request.Context(), ls[1:len(ls)-1]); err != nil {
			code = http.StatusInternalServerError
			log.Println(fmt.Errorf("error loading dataset: %w", err))
		}
		writer.WriteHeader(code)
	}
}
