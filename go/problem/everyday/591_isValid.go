package everyday

import (
    "strings"
    "unicode"
)

// func isValid(code string) bool {
//     /*
//         1、优先检查以i开始的连续段是否为 CDATA，如果能够匹配到开头，那么尝试匹配到CDATA 的结尾，更新ii；否则，返回false；
//         2、判断s[i]是否为 <，再看s[i+1]是不是/，来判断是左还是右：
//         （1）左边：将其加入栈，待后续匹配；
//         （2）右边：将其与栈顶元素进行匹配，若栈为空或者匹配不上，返回false；
//         3、其他情况均是普通字符；
//     */

//         // regex r1 = regex("<!\\[CDATA\\[.*?\\]\\]>|t") ;
//         // regex r2 = regex("<([A-Z]{1,9})>[^<]*</\\1>");
//         // code = regex_replace(code, r1 , "-");
//         // // cout << "code: " << code << endl;
//         // string prev = "";
//         // while(code != prev){
//         //     prev = code;
//         //     code = regex_replace(code, r2, "t");
//         //     // cout << "code: " << code << endl;
//         // }
//         // return code == "t";
//     // r1 := regexp.MustCompile(`<!\[CDATA\[.*?\]\]>|t`)
//     // r2 := regexp.MustCompile(`<([A-Z]{1,9})>[^<]*</\\1>`)

//     // code = r1.ReplaceAllString(code, "-")
    
//     // fmt.Println("code: ", code)
//     // prev := ""

//     // for code != prev {
//     //     prev = code
//     //     code = r2.ReplaceAllString(code, "t")
//     //     fmt.Println("code: ", code)
//     // }

//     // return code == "t"
// }


func isValid2(code string) bool {
    /*
        1、判断s[i]是否为 <:
            (1)如果s[i+1]是/，那么我们遇到了一个结束标签，我们定位>的位置jj，此时code[i+2,j-1]就是这个结束标签的名字；
                a.将其与栈顶元素进行匹配，若栈为空或者匹配不上，返回false；
                b.如果匹配了，说明标签已经闭合，我们将当前栈顶的名称弹出；
            (2)如果下一个字符为!，那么我们遇到了一个cdata，我们需要继续往后读7个字符，判断是否为[CDATA[，
            随后我们定位下一个]]>的位置j，此时code[i+9,j-1]就是cdata的内容；
            (3)如果下一个字符为大写字母，那么我们遇到了一个开始标签，我们定位下一个>的位置jj，此时code[i+2,j-1]就是
            这个开始标签的名字，判断其是否是合法的，合法则放入栈顶；
            (4)除此以外，如果不存在下一个字符，或者不符合上面几种情况，那么code是不合法的；
        2、其他字符，需要栈中存在一个开放标签；
    */
    tags := []string{}
    for code != "" {
        if code[0] != '<' {
            if len(tags) == 0 {
                return false
            }
            // 不断更新code
            code = code[1:]
            continue
        }
        // <标签之后肯定不止一个字符
        if len(code) == 1 {
            return false
        }
        // 遇到了结束标签
        if code[1] == '/' {
            j := strings.IndexByte(code, '>')
            // 没找到下一个>，不合法
            if j == -1 {
                return false
            }
            // 这时候需要和栈顶的标签相匹配
            if len(tags) == 0 || tags[len(tags)-1] != code[2:j] {
                return false
            }
            // 匹配上之后，更新栈顶元素（出栈）
            tags = tags[:len(tags)-1]
            // 更新code
            code = code[j+1:]
            // 如果这时候栈为空，且code还有剩余，非法（需要被一个闭合的标签包围起来）
            if len(tags) == 0 && code != "" {
                return false
            }
        } else if code[1] == '!' {  // 这时候应该是一个cdata
            // cdata肯定是被包围的，所以栈非空，这时候判断后面是否为[cdata[
            if len(tags) == 0 || len(code) < 9 || code[2:9] != "[CDATA[" {
                return false
            }
            // 后面需要匹配结尾]]>
            j := strings.Index(code, "]]>")
            if j == -1 {
                return false
            }
            // 更新code
            code = code[j+1:]
        } else {
            // 这时候匹配闭合标签的结尾（开头那个标签<DIV>）
            j := strings.IndexByte(code, '>')
            if j == -1 {
                return false
            }
            // 获取tagName
            tagName := code[1:j]
            // 判断tagName是否合法：长度1～9，都是大写
            if tagName == "" || len(tagName) > 9 {
                return false
            }
            for _, ch := range tagName {
                if !unicode.IsUpper(ch) {
                    return false
                }
            }
            // 将开头的标签放入栈里面，更新code
            tags = append(tags, tagName)
            code = code[j+1:]
        }
    }
    // 最后看栈是否为空即可
    return len(tags) == 0
}
