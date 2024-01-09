package global

import (
	"fmt"
	"google.golang.org/grpc"
)

var accountConn *grpc.ClientConn

func GetAccountConn() (*grpc.ClientConn, error) {
	if accountConn == nil {
		return nil, fmt.Errorf("accountConn为空")
	}
	return accountConn, nil
}

func SetAccountConn(cc *grpc.ClientConn) {
	accountConn = cc
}

/*type TokenCredential struct{}

func (tc TokenCredential) GetRequestMetadata(ctx context.Context, url ...string) (map[string]string, error) {
	appId, secret := "e6abbb583d30415cbccc54ce7a305631", "2f0a2ed5d424440b9f7ebc54700109685c81c963982d4d8e853dfc409bdf6392"
	randBytes := make([]byte, 16)
	_, _ = rand.Read(randBytes)
	timestamp := time.Now().Unix()
	randString := hex.EncodeToString(randBytes)

	raw := fmt.Sprintf("%s-%d-%s-%s", appId, timestamp, randString, secret)
	hash := sha256.New()
	hash.Write([]byte(raw))
	sum := hash.Sum(nil)

	return map[string]string{
		"appid":     appId,
		"timestamp": fmt.Sprintf("%d", timestamp),
		"rand":      randString,
		"signature": base64.StdEncoding.EncodeToString(sum),
	}, nil
}

func (tc TokenCredential) RequireTransportSecurity() bool {
	return false
}*/
