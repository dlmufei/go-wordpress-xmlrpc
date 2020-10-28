package example

import (
	"log"

	xmlrpc "github.com/dlmufei/go-wordpress-xmlrpc"
	"github.com/dlmufei/go-wordpress-xmlrpc/wordpress"
)

func demo() {
	c, err := xmlrpc.NewClient(`https://example.com/xmlrpc.php`, xmlrpc.UserInfo{
		`your username`,
		`your password`,
	})
	if err != nil {
		log.Fatalln(err)
	}
	p := wordpress.NewPost(`content`, `title`, ``, 100, []string{`tag1`, `tag2`}, []string{`cate1`, `cate2`})
	blogID, err := c.Call(p)
	if err != nil {
		log.Println(err)
	}
	log.Println(blogID)
}
