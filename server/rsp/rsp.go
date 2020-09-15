package rsp

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/gocolly/colly"
)

const (
	GET           = "GET"
	POST          = "POST"
	Host          = "http://128.18.1.121"
	CheckLoginUrl = Host + "/rsp/j_spring_security_check"
	CongonSrvUrl  = Host + "/rsp/autoRptPerform?srvId=COGNOS_SRV"
	RptPerformUrl = Host + "/rsp/rptPerform"
	BackendHost   = "http://128.18.1.122:8888"
	CongosHost    = "http://128.18.1.122:8888"
	CongosUrl     = CongosHost + "/cognos/cgi-bin/cognos.cgi"
)

var (
	rptIdMap = map[string]int{
		"FMS-A-01": 957, //FMS-A-01 业务状况表
		"FMS-A-03": 963, //FMS-A-03 利润表
	}
	gRsp *Rsp = nil
)

func NewRsp(branch string) *Rsp {
	rsp := &Rsp{}
	rsp.Init()
	return rsp
}

type Rsp struct {
	c      *colly.Collector
	branch string
}

func (this *Rsp) Init() *Rsp {
	if this.c == nil {
		this.c = colly.NewCollector()
		this.c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko"
	}
	return this
}

func (this *Rsp) Collector() *colly.Collector {
	c := this.c.Clone()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Connection", "Keep-Alive")
	})
	return c
}

func (this *Rsp) form(method string, URL string, values url.Values, pattern string) (retAction string, retValues url.Values, err error) {
	c := this.Collector()

	var errForm error
	if pattern != "" {
		c.OnHTML(pattern, func(el *colly.HTMLElement) {
			retAction = el.Attr("action")
			names := el.ChildAttrs("input", "name")
			values := el.ChildAttrs("input", "value")

			if len(names) == 0 || len(values) == 0 || len(names) != len(values) {
				errForm = errors.New("error")
				return
			}

			retValues = make(url.Values)
			for i, name := range names {
				retValues.Set(name, values[i])
			}

			return
		})
	}

	fmt.Println("Start Request")
	values.Encode()
	fmt.Println(1)
	bValues := []byte(values.Encode())
	fmt.Println(123)
	if err = c.Request(method, URL, bytes.NewReader(bValues), nil, nil); err != nil {
		return
	}
	fmt.Println("End Request")
	if errForm != nil {
		err = errForm
	}

	return
}

func (this *Rsp) Login(user, password string) error {
	fmt.Println("Start Login1")

	values := url.Values{
		"j_username":    []string{user},
		"j_password":    []string{password},
		"isCookieLogin": []string{"1"},
		"submit":        []string{"登录"},
	}

	fmt.Println("Start Login111")
	if _, _, err := this.form(POST, CheckLoginUrl, values, ""); err != nil {
		return err
	}

	fmt.Println("Start Login2")
	var retAction string
	var retValues url.Values

	retAction, retValues, err := this.form(GET, CongonSrvUrl, nil, `form[id='autoForm']`)
	if err != nil {
		return err
	}

	fmt.Println("Start Login3")
	_, _, err = this.form(POST, retAction, retValues, `form[id='autoForm']`)
	if err != nil {
		return err
	}

	return nil
}

func (this *Rsp) Args(rptID string, freq string, strDate string) (url.Values, error) {
	id, ok := rptIdMap[rptID]
	if !ok {
		return nil, errors.New("rptID Error")
	}
	date, err := time.Parse("20060102", strDate)
	if err != nil {
		return nil, err
	}
	return url.Values{
		"rptForm:outputFormatCognos": []string{"HTML"},
		"rptForm:rptId":              []string{strconv.Itoa(id)},
		"rptForm:p_branch_code":      []string{this.branch},
		"rptForm:p_currency":         []string{"CNY"},
		"rptForm:p_frequency_desc":   []string{freq},
		"rptForm:p_date":             []string{date.Format("2006-01-02")},
		"rptForm:p_flag":             []string{"N"},
		"rptForm=rptForm":            []string{"rptForm"},
	}, nil
}

func (this *Rsp) Query(args url.Values) (err error) {
	action, values, err := this.form(POST, Host+"/rsp/rptPerform", args, `form[id='autoForm']`)
	if err != nil {
		return
	}
	action, values, err = this.form(POST, action, values, `form[name="pform"]`)
	if err != nil {
		return err
	}

	action = CongosHost + "/cognos/cgi-bin/cognos.cgi"
	c := this.c.Clone()
	c.OnResponse(func(r *colly.Response) {
	})
	err = c.PostRaw(action, []byte(values.Encode()))

	return
}

func (this *Rsp) initArgs(args url.Values) (err error) {
	action, values, err := this.form(POST, Host+"/rsp/rptPerform", args, `form[id='autoForm']`)
	if err != nil {
		return
	}
	action, values, err = this.form(POST, action, values, `form[name="pform"]`)
	if err != nil {
		return
	}
	return
}

func (this *Rsp) FmsA03(freq string, strDate string) (map[string][]float64, error) {
	args, err := this.Args("FMS-A-03", freq, strDate)
	if err != nil {
		return nil, err
	}

	fmt.Println(args)

	c := this.Collector()
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(string(r.Body))
	})
	err = c.PostRaw(CongosUrl, []byte(args.Encode()))
	if err != nil {
		return nil, err
	}

	return nil, nil
}
