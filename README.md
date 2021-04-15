# gin-cat
gin cat middleware

base on https://github.com/cat-go/cat 

```go
package main
import (
	"github.com/gin-gonic/gin"
	catmidd "github.com/ainiaa/gin-cat"
)
func initGin(c *catmidd.Conf) *gin.Engine {
    r := gin.New()
    r.Use(catmidd.Cat(true, c))
    
    return r
}

func main() {
    c := &catmidd.Conf{
        Flag:true,
        AppId:"gin-cat",
        Port:2280,
        HttpPort:8080,
        ServerAddr:"127.0.0.1",
        IsDebug:true,
    }
    
    initGin(c)
    
   // shutdown cat
   // catmidd.Shutdown(c)     
}
```


