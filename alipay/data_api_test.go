package alipay

import (
	"testing"

	"github.com/chaclee/gopay"
	"github.com/chaclee/gopay/pkg/xlog"
)

func TestClient_DataBillDownloadUrlQuery(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("bill_type", "trade").
		Set("bill_date", "2016-04-05")
	rsp, err := client.DataBillDownloadUrlQuery(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("rsp:", rsp)
}
