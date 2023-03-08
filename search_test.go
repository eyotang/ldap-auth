package auth

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSearch_Users(t *testing.T) {
	client, err := setup(t)
	Convey("test Search functions", t, func() {
		So(err, ShouldBeNil)

		user, err := client.SearchUser("tangyongqiang" + client.cfg.EmailSuffix)
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.Username, ShouldEqual, "tangyongqiang")

		user2, err := client.SearchUser("tangyongqiang2" + client.cfg.EmailSuffix)
		So(err, ShouldBeNil)
		So(user2, ShouldBeNil)
	})
}
