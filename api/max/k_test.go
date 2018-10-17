// How to use:
//
// 1. Testing
// (1) > cd %GOPATH%\src\github.com\oneleo\technical-analysis\api\max
// or (1) $ cd $GOPATH/src/github.com/oneleo/technical-analysis/api/max
// (2) $> go test -v
//
// 2. Benchmark:
// (1) > cd %GOPATH%\src\github.com\oneleo\technical-analysis\api\max
// or (1) $ cd $GOPATH/src/github.com/oneleo/technical-analysis/api/max
// (2) $> go test -bench={Mathod Name} -v
// for example: $> go test -bench=BenchmarkNode -v
//
package max

import (
	"testing"

	maxapi "github.com/maicoin/max-exchange-api-go"
	"github.com/maicoin/max-exchange-api-go/types"
)

// Test_MultiRdmFlt64 是測試 MultiRdmFlt64 函數所產生出來的亂數平均值及標準差要在合理範圍內。
func TestMax_GetK1Min(t *testing.T) {

	// 定義測試集 Struct。
	var tests = []struct {
		// ----- input -----
		Market    string
		Timestamp int32
		// ----- output -----
		Open   types.Price
		High   types.Price
		Low    types.Price
		Close  types.Price
		Volume types.Price
	}{
		// 測試 1。
		{"ethtwd", 1519206540, 25000, 25000, 25000, 25000, 0.001},
		// 測試 2。
		{"ethtwd", 1519406520, 27500, 27500, 27500, 27500, 0},
		// 測試 3。
		{"ethtwd", 1519806540, 26000, 26000, 26000, 26000, 0},
	}

	client := maxapi.NewClient()
	defer client.Close()

	for _, test := range tests {
		got, _ := GetK1Min(client, test.Market, test.Timestamp)
		// t.Logf("GetK1Min(client, %s, %d) = %g, %g, %g, %g, %g,\n want %g, %g, %g, %g, %g", test.Market, test.Timestamp, got.Open, got.High, got.Low, got.Close, got.Volume, test.Open, test.High, test.Low, test.Close, test.Volume)
		// 如果產生出的 num 筆亂數不在 [min, max) 內，測試失敗。
		if got.Open != test.Open || got.High != test.High || got.Low != test.Low || got.Close != test.Close || got.Volume != test.Volume {
			t.Errorf("GetK1Min(client, %s, %d) = %g, %g, %g, %g, %g,\n want %g, %g, %g, %g, %g", test.Market, test.Timestamp, got.Open, got.High, got.Low, got.Close, got.Volume, test.Open, test.High, test.Low, test.Close, test.Volume)
		}
	}
}
