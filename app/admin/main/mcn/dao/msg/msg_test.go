package msg

import (
	"context"
	"testing"

	"github.com/smartystreets/goconvey/convey"

	"gopkg.in/h2non/gock.v1"
)

func TestMsgMutliSendSysMsg(t *testing.T) {
	convey.Convey("MutliSendSysMsg", t, func(ctx convey.C) {
		var (
			c       = context.Background()
			allUids = []int64{}
			mc      = ""
			title   = ""
			context = ""
			ip      = ""
		)
		ctx.Convey("When everything goes positive", func(ctx convey.C) {
			defer gock.OffAll()
			httpMock("POST", d.msgURL).Reply(200).JSON(`{"code":0}`)
			err := d.MutliSendSysMsg(c, allUids, mc, title, context, ip)
			ctx.Convey("Then err should be nil.", func(ctx convey.C) {
				ctx.So(err, convey.ShouldBeNil)
			})
		})
	})
}

func TestMsgSendSysMsg(t *testing.T) {
	convey.Convey("SendSysMsg", t, func(ctx convey.C) {
		var (
			c       = context.Background()
			uids    = []int64{2}
			mc      = "116_3_3"
			title   = "11"
			content = "111"
			ip      = "111.111.111.111"
		)
		ctx.Convey("When everything goes positive", func(ctx convey.C) {
			defer gock.OffAll()
			httpMock("POST", d.msgURL).Reply(200).JSON(`{"code":0}`)
			err := d.SendSysMsg(c, uids, mc, title, content, ip)
			ctx.Convey("Then err should be nil.", func(ctx convey.C) {
				ctx.So(err, convey.ShouldBeNil)
			})
		})
	})
}
