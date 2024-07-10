package articles

import (
	"net/http"

	"github.com/CGSG-2021-AE4/blog/api"

	"github.com/gin-gonic/gin"
)

func mainPageHandler(domain string, svc api.ArticlesService, userSvc api.UserService) gin.HandlerFunc {
	_ = svc
	_ = userSvc
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "main.html", gin.H{
			"Domain": domain,
			"Title":  "Lorem Ipsum",
			"Body":   "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse enim risus, auctor sollicitudin mi at, sagittis viverra felis. In pellentesque varius quam, id imperdiet lectus cursus sit amet. Morbi vitae sem ac urna hendrerit auctor nec et elit. Praesent vestibulum ut purus a laoreet. Aenean ac consectetur justo. Aliquam tempor eu erat nec eleifend. Phasellus tincidunt fringilla elit a elementum. Nulla venenatis bibendum tortor, at rhoncus nisi. Pellentesque ante erat, tincidunt vel ante ut, sagittis dignissim elit. Fusce iaculis libero eget cursus aliquam. Vestibulum quis urna libero. In non tempor magna, at cursus leo. In rutrum eu neque quis venenatis. Donec et velit sapien. Ut venenatis ipsum et sagittis aliquet. Morbi quis efficitur enim, a varius dolor. Pellentesque viverra ornare vehicula. Interdum et malesuada fames ac ante ipsum primis in faucibus. Ut condimentum lacus venenatis mauris ultricies, nec sagittis nisl suscipit. Aenean in malesuada arcu. Donec hendrerit odio a nibh consectetur, a mattis sapien tincidunt. Duis felis lacus, porta in luctus nec, imperdiet eget purus. Cras diam lacus, iaculis nec ligula non, condimentum malesuada felis. Etiam mauris tellus, ornare ut fermentum nec, mollis non mauris. Vivamus aliquam, dolor sit amet gravida tempus, erat enim auctor lacus, in aliquam velit ex nec nunc. Duis at tristique elit, ac congue nulla. Praesent gravida hendrerit enim non feugiat. Donec cursus sagittis eros eget eleifend. Sed non nisi malesuada lectus viverra pellentesque. Mauris pellentesque ultrices tincidunt. In hac habitasse platea dictumst. Etiam id iaculis sapien. Duis nec varius nunc. Nullam interdum nulla nisl, nec ullamcorper turpis semper sit amet. Aliquam turpis nibh, aliquam vel porttitor id, lacinia a ligula. Pellentesque non magna quam. Fusce congue purus vitae dapibus pharetra. Maecenas libero nisi, egestas sit amet justo scelerisque, porta placerat ligula. Donec placerat, nulla vitae posuere convallis, nisl urna bibendum quam, in commodo neque lorem eget nunc. Fusce nec ante gravida, convallis diam vitae, consequat felis. In hac habitasse platea dictumst. Sed vitae libero ut massa fermentum rutrum fermentum in arcu. Nunc luctus fringilla enim, sed posuere nibh tincidunt at. Curabitur posuere augue at sapien aliquam rutrum. Curabitur tempus felis sapien, tristique posuere augue bibendum eu. Ut tristique enim ac posuere sagittis. Nulla lacinia dignissim vulputate. Cras faucibus odio ac lorem luctus, vitae tristique sapien dapibus. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Praesent imperdiet quis ligula ac accumsan. Praesent vulputate purus vitae sapien imperdiet, at tempor dui pretium. Donec faucibus orci in ligula iaculis facilisis. Nam non turpis a urna cursus posuere. Fusce et augue mauris. Ut et sapien vitae enim malesuada bibendum nec quis nunc. Nam vulputate egestas magna. Sed tempor sapien ultricies sapien finibus, tincidunt malesuada purus venenatis. Pellentesque vel convallis magna. In eu mattis mi. Cras eget nibh posuere nibh fermentum consequat nec tristique nulla. Integer sed metus sed lectus eleifend semper. Mauris et consectetur sem. Nullam lacinia lectus dui, et efficitur lorem elementum ut. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Duis feugiat viverra consequat.",
		})
	}
}
