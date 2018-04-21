package gh

import (
	"github.com/gin-gonic/gin"
	"time"
	"strconv"
	"github.com/quexer/utee"
	"strings"
)

type ParamHelper struct {
	m map[string]string
	c *gin.Context
}

func NewParamHelper(c *gin.Context) *ParamHelper {
	return &ParamHelper{
		m: make(map[string]string),
		c: c,
	}
}

func (p *ParamHelper) Float64(name string) float64 {
	pure:=p.c.Param(name)
	p.m[name] = pure
	v, err := strconv.ParseFloat(pure, 64)
	utee.Chk(err)
	return v
}
func (p *ParamHelper) Bool(name string) bool {
	pure:=p.c.Param(name)
	p.m[name] = pure
	v, err := strconv.ParseBool(pure)
	utee.Chk(err)
	return v
}
func (p *ParamHelper) Int(name string, defaultVs ...int) int {
	pure:=p.c.Param(name)
	p.m[name] = pure
	v, err := strconv.ParseInt(pure, 10, 32)
	if err!=nil && len(defaultVs)>0{
		return defaultVs[0]
	}
	utee.Chk(err)
	return int(v)
}
func (p *ParamHelper) Int64(name string) int64 {
	pure:=p.c.Param(name)
	p.m[name] = pure
	v, err := strconv.ParseInt(pure, 10, 64)
	utee.Chk(err)
	return v
}
func (p *ParamHelper) Uint64(name string) uint64 {
	pure:=p.c.Param(name)
	p.m[name] = pure
	v,err:=strconv.ParseUint(pure, 10, 64)
	utee.Chk(err)
	return v
}
func (p *ParamHelper) Time(name, format string) time.Time {
	pure:=p.c.Param(name)
	p.m[name] = pure
	t, err := time.Parse(format, pure)
	utee.Chk(err)
	return t
}
func (p *ParamHelper) TimeLoc(name, format string, loc *time.Location) time.Time {
	pure:=p.c.Param(name)
	p.m[name] = pure
	t, err := time.ParseInLocation(format, pure, loc)
	utee.Chk(err)
	return t
}
func (p *ParamHelper) String(name string) string {
	pure:=p.c.Param(name)
	p.m[name] = pure
	return pure
}
func (p *ParamHelper) IntArr(name, separate string) []int {
	pure:=strings.TrimSpace(p.c.Param(name))
	p.m[name] = pure
	svs := strings.Split(pure, separate)
	ivs := make([]int, 0)
	for _, v := range svs {
		if v == "" {
			continue
		}
		if iv, err := strconv.ParseInt(v, 10, 32); err != nil {
			ivs = append(ivs, int(iv))
		}
	}
	return ivs
}

func (p *ParamHelper) Params()map[string]string{
	return p.m
}

