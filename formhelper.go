package gh

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/quexer/utee"
	"strings"
)

type FormHelper struct {
	m map[string]string
	c *gin.Context
}

func NewFormHelper(c *gin.Context) *FormHelper {
	return &FormHelper{
		m: make(map[string]string),
		c: c,
	}
}

func (p *FormHelper) Int(name string, defaultVs ...int) int {
	pure := p.c.PostForm(name)
	p.m[name] = pure
	v, err := strconv.ParseInt(pure, 10, 32)
	if err != nil && len(defaultVs) > 0 {
		return defaultVs[0]
	}
	return int(v)
}

func (p *FormHelper) Int64(name string) int64 {
	pure := p.c.PostForm(name)
	p.m[name] = pure
	v, err := strconv.ParseInt(pure, 10, 64)
	utee.Chk(err)
	return v
}

func (p *FormHelper) Float32(name string, defaultVs ...float32) float32 {
	pure := strings.TrimSpace(p.c.PostForm(name))
	p.m[name] = pure
	v, err := strconv.ParseFloat(pure, 32)
	utee.Chk(err)
	return float32(v)
}
func (p *FormHelper) String(name string) string {
	pure := strings.TrimSpace(p.c.PostForm(name))
	p.m[name] = pure
	return pure
}
func (p *FormHelper) StrArr(name, separate string) []string {
	pure := strings.TrimSpace(p.c.PostForm(name))
	p.m[name] = pure
	return strings.Split(pure, separate)
}
func (p *FormHelper) IntArr(name, separate string) []int {
	pure := strings.TrimSpace(p.c.PostForm(name))
	p.m[name] = pure
	svs := strings.Split(pure, separate)
	ivs := make([]int, 0)
	for _, v := range svs {
		if v == "" {
			continue
		}
		if iv, err := strconv.ParseInt(v, 10, 32); err == nil {
			ivs = append(ivs, int(iv))
		}
	}
	return ivs
}

func (p *FormHelper) PostForms() map[string]string {
	return p.m
}
