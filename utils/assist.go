package utils

// 校验证书
import (
	"errors"
	"fmt"
	"math/big"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bugfan/jsoniter"
	"github.com/gobwas/glob"
)

// 年月日字符串转日期  例如:2019-05-23
func strTimeYMDtoTimeStamp10(str string) int64 {
	if len(str) != 10 {
		return 0
	}
	return TimeStrToTimeStamp10(str + " 15:04:05")
}
func TimeStrToTimeStamp10(t string) int64 {
	if len(t) < 18 {
		return 0
	}
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, t, loc)
	sr := theTime.Unix()
	return sr
}
func TimeStrToTime(t string) (ti time.Time, err error) {
	if len(t) < 18 {
		return ti, errors.New("wrong time str")
	}
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	ti, err = time.ParseInLocation(timeLayout, t, loc)
	return
}
func TimeStampToTimeStr(t int64) string {
	//  yyyy-MM-dd'T'HH:mm:ss'.000Z'
	timeLayout := "2006-01-02 15:04:05"
	dataTimeStr := time.Unix(t, 0).Format(timeLayout)
	// fmt.Println(dataTimeStr)
	return dataTimeStr
}

// 检查是否是域名
func IsDomain(domain string) bool {
	if len(domain) > 255 {
		return false
	}
	reg := `[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+(:\d+)*(\/\w+\.\w+)*`
	r, _ := regexp.Compile(reg)
	return r.MatchString(domain)
}

// 匹配是否在子域
func IsPanDomain(regDomain, domain string) bool {
	g, err := glob.Compile(regDomain, '.')
	if err != nil {
		return false
	}
	match := g.Match(domain)
	return match
}
func InetNtoA(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

func InetAtoN(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

func IPv4ToNumber(ip string) int64 {
	bits := strings.Split(ip, ".")
	if len(bits) < 4 {
		return -1
	}

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

// 通用深层json解析器
/*
*	功能: 加速反序列化 直接读取json
*   参照文档: http://jsoniter.com/migrate-from-go-std.html
 */
func NewIJSON() jsoniter.API {
	return jsoniter.ConfigCompatibleWithStandardLibrary
}

var IJSON = NewIJSON()
