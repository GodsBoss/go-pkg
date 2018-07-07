package typedconf_test

import (
	"github.com/GodsBoss/go-pkg/typedconf"

	"encoding/xml"
	"fmt"
)

func ExampleNewXMLDecoders_instance() {
	input := []byte(
		`
      <list>
        <item>
          <op type="-">
            <value>-5000</value>
          </op>
        </item>
        <item>
          <op type="+">
            <value>-50</value>
            <value>80</value>
            <value>-5</value>
          </op>
        </item>
      </list>
    `,
	)
	list := &List{}
	err := xml.Unmarshal(input, list)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	for _, item := range list.Items {
		fmt.Printf("%d\n", item.Op.Calc())
	}

	// Output:
	// 5000
	// 25
}

type List struct {
	XMLName xml.Name
	Items   []Item `xml:"item"`
}

type Item struct {
	Op Operation
}

func (it *Item) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	decoder := typedconf.NewXMLDecoders()
	decoder.Register("-", func() interface{} {
		return &Negate{}
	})
	decoder.Register("+", func() interface{} {
		return &Add{}
	})
	inst := decoder.Instance()
	i := &item{
		Inst: inst,
	}
	err := d.Decode(inst)
	if err != nil {
		return err
	}
	it.Op = i.Inst.Value().(Operation)
	d.Skip()
	return nil
}

type item struct {
	Inst typedconf.Instance `xml:"op"`
}

type Operation interface {
	Calc() int
}

type Negate struct {
	Value int `xml:"value"`
}

func (neg Negate) Calc() int {
	return -neg.Value
}

type Add struct {
	Values []int `xml:"value"`
}

func (add Add) Calc() int {
	result := 0
	for _, value := range add.Values {
		result += value
	}
	return result
}
