package expr

import (
	"errors"
	"github.com/bill-lhr/utils/src/common"
	"github.com/bill-lhr/utils/src/stack"
	"regexp"
)

// CheckExpr 四则运算公式校验 支持 +、-、*、/、(、)  注意：因为参与运算的元素不做限制，所以只校验运算符四则运算合法性
func CheckExpr(expr []string) (bool, error) {
	if len(expr) == 0 {
		return false, errors.New("表达式为空")
	}

	// 先判断括号是否匹配、不是开头的(前必须是运算符、
	parenthesesStack := stack.NewStack()
	for idx, v := range expr {
		switch v {
		case "(":
			parenthesesStack.Push(v)
			// 非开头的 ( 前必须是 +-*/(
			if idx != 0 && !common.InArrayForString(expr[idx-1], []string{"+", "-", "*", "/", "("}) {
				return false, errors.New("括号错误")
			}
		case ")":
			if parenthesesStack.Empty() {
				return false, errors.New("括号不匹配")
			}
			parenthesesStack.Pop()
			// 非结尾的 ) 前必须是 +-*/)
			if idx != len(expr)-1 && !common.InArrayForString(expr[idx+1], []string{"+", "-", "*", "/", ")"}) {
				return false, errors.New("括号错误")
			}
		default:
			continue
		}
	}
	if !parenthesesStack.Empty() {
		return false, errors.New("括号不匹配")
	}

	opts := []string{"+", "-", "*", "/", "(", ")"}
	// 将表达式list转成string方便正则
	exprStr := ""
	for _, v := range expr {
		// 只校验合法性，对运算元素用一位字符代替
		if len(v) == 1 && common.InArrayForString(v, opts) {
			exprStr += v
		} else {
			exprStr += "@" // 用于替代运算元素
		}
	}

	if regexp.MustCompile(`^[+\-*/]`).MatchString(exprStr) {
		return false, errors.New("不能以运算符开头")
	}
	if regexp.MustCompile(`[+\-*/]$`).MatchString(exprStr) {
		return false, errors.New("不能以运算符结尾")
	}
	if regexp.MustCompile(`[+\-*/]{2,}`).MatchString(exprStr) {
		return false, errors.New("存在连续运算符")
	}
	if regexp.MustCompile(`\(\)`).MatchString(exprStr) {
		return false, errors.New("存在空括号")
	}
	if regexp.MustCompile(`\([+\-*/]`).MatchString(exprStr) {
		return false, errors.New("(后不能为运算符")
	}
	if regexp.MustCompile(`[+\-*/]\)`).MatchString(exprStr) {
		return false, errors.New(")前不能为运算符")
	}

	return true, nil
}

// PrefixToPostFix 中缀表达式转后缀
/*	需要一个运算符栈和结果数组
	1.遇到操作元素直接输出到结果
	2.左括号入栈
	3.右括号不入栈，出栈直至遇到左括号
	4.遇到运算符，从栈顶起连续的所有优先级大于等于当前运算符的运算符出栈，当前运算符入栈
	5.出栈所有运算符
*/
func PrefixToPostFix(prefix []string) ([]string, error) {
	postfix := make([]string, 0, len(prefix))
	operators := []string{"+", "-", "*", "/"} // 运算符列表
	opStack := stack.NewStack()

	for _, item := range prefix {
		// 非运算符及括号直接写入结果
		if !common.InArrayForString(item, operators) && !common.InArrayForString(item, []string{"(", ")"}) {
			postfix = append(postfix, item)
			continue
		}
		switch item {
		case "+", "-":
			for !opStack.Empty() && common.InArrayForString(opStack.Top().(string), []string{"+", "-", "*", "/"}) {
				postfix = append(postfix, opStack.Pop().(string))
			}
			opStack.Push(item)
		case "*", "/":
			for !opStack.Empty() && common.InArrayForString(opStack.Top().(string), []string{"*", "/"}) {
				postfix = append(postfix, opStack.Pop().(string))
			}
			opStack.Push(item)
		case "(":
			opStack.Push(item)
		case ")":
			for !opStack.Empty() && opStack.Top() != "(" {
				postfix = append(postfix, opStack.Pop().(string))
			}
			if opStack.Empty() {
				return nil, errors.New("表达式错误")
			}
			opStack.Pop()
		}
	}

	for !opStack.Empty() {
		item := opStack.Pop().(string)
		if item == "(" {
			return nil, errors.New("表达式错误")
		}
		postfix = append(postfix, item)
	}

	return postfix, nil
}
