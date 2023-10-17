package collect

import (
	desc "gint/descAntlr"
	"github.com/antlr4-go/antlr"
	"github.com/jinzhu/copier"
	"os"
)

// CollectData 从接口描述文件收集数据
type DescData struct {
	Dtos      []desc.Dto
	ApiModule desc.ApiModule
}

func CollectData(descFile string) (data DescData, err error) {
	b, err := os.ReadFile(descFile)
	if err != nil {
		return DescData{}, err
	}

	input := string(b)

	// 创建输入流
	is := antlr.NewInputStream(input)
	// 创建词法分析器
	lexer := desc.NewdescLexer(is)
	/// 创建词法错误处理器
	//lexer.RemoveErrorListeners()
	//lexer.AddErrorListener(&ErrorListener{})

	// token流
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// 创建语法分析器
	parser := desc.NewdescParser(tokens)
	//parser.AddErrorListener(&ErrorListener{})
	tree := parser.File()
	// listener实例
	listener := &desc.BasedescListener{}
	// 遍历语法树，并调用listener对应方法
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	if err := copier.Copy(&data, listener); err != nil {

		return DescData{}, err
	}

	return
}
