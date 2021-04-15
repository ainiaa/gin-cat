package cat

import (
	"strconv"

	"github.com/cat-go/cat"
	"github.com/cat-go/cat/message"
	"github.com/gin-gonic/gin"
)

const (
	CatCtx          = "cat_ctx"
	CatCtxRootTran  = "cat_root_tran"
	CatCtxRedisTran = "cat_redis_tran"
	CatCtxMysqlTran = "cat_mysql_tran"

	TypeHttpRemote    = "HttpRemoteCall"
	TypeMongoDbPrefix = "MongoDb."
	TypeUrlStatus     = "URL.status"

	RemoteCallMethod = "RemoteCall.Method"
	RemoteCallErr    = "RemoteCall.Err"
	RemoteCallStatus = "RemoteCall.Status"
	RemoteCallScheme = "RemoteCall.Scheme"
)

func InitCat(c *Conf) {
	if c.Flag {
		if c.IsDebug {
			cat.DebugOn()
		}
		cat.Init(&cat.Options{
			AppId:      c.AppId,
			Port:       c.Port,
			HttpPort:   c.HttpPort,
			ServerAddr: c.ServerAddr,
		})
	}
}

func Shutdown(c *Conf) {
	if c.Flag {
		cat.Shutdown()
	}
}

//监控与链路追踪
func Cat(init bool, c *Conf) gin.HandlerFunc {
	if init {
		InitCat(c)
	}
	return func(c *gin.Context) {
		if cat.IsEnabled() {
			tran := cat.NewTransaction(cat.TypeUrl, c.Request.URL.Path)
			setTraceId(c, tran)
			tran.LogEvent(cat.TypeUrlMethod, c.Request.Method, c.FullPath())
			tran.LogEvent(cat.TypeUrlClient, c.ClientIP())
			c.Set(CatCtxRootTran, tran)
			defer func() {
				tran.AddData(TypeUrlStatus, strconv.Itoa(c.Writer.Status()))
				tran.Complete()
			}()
		}
		c.Next()
	}
}

func setTraceId(c *gin.Context, tran message.Transactor) {
	var root, parent, child string
	root = c.Request.Header.Get(cat.RootId)
	parent = c.Request.Header.Get(cat.ParentId)
	child = c.Request.Header.Get(cat.ChildId)
	if root == "" {
		root = cat.MessageId()
	}
	if parent == "" {
		parent = root
	}
	if child == "" {
		child = cat.MessageId()
	}
	tran.SetRootMessageId(root)
	tran.SetParentMessageId(parent)
	tran.SetMessageId(child)
}
