package sampleCrypto

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"sync"
	"time"

	mathrand "math/rand"

	"github.com/bou.ke/monkey"
	uuid "github.com/satori/go.uuid"
)

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func main() {

	var mutex sync.Mutex
	var wg sync.WaitGroup
	var sameID int

	tokenMap := make(map[string]int)
	uuidMap := make(map[string]int)
	mathrandMap := make(map[int]int)

	wayback := time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC)
	patch := monkey.Patch(time.Now, func() time.Time { return wayback })

	// https://github.com/satori/go.uuid/blob/master/generator.go#L109
	for j := 0; j < 1000; j++ {
		wg.Add(1)

		go func(wg *sync.WaitGroup) {
			for i := 0; i < 20; i++ {

				uu := uuid.Must(uuid.NewV1()).String()

				// Example: this will give us a 44 byte, base64 encoded output
				token, err := GenerateRandomString(32)
				if err != nil {
					return
				}

				mutex.Lock()
				if _, found := uuidMap[uu]; found {
					sameID = sameID + 1
				}
				mutex.Unlock()

				fmt.Println(
					time.Now().UnixNano(),
					"- Token from secure random ", token,
					"- UUID from Satori => ", uu,
					"- Math rand ", mathrand.Int())

				mutex.Lock()
				tokenMap[token] = 1
				uuidMap[uu] = 1
				mathrandMap[mathrand.Int()] = 1
				mutex.Unlock()
			}
			wg.Done()
		}(&wg)
	}

	wg.Wait()

	patch.Unpatch()
	fmt.Printf("it unpatch time now %s\n", time.Now())

	fmt.Println(len(tokenMap), " == ", len(uuidMap), " == ", len(mathrandMap))
	fmt.Println("Total Same ID ", sameID)
}
