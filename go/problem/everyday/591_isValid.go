package everyday

import "regexp"

func isValid(code string) bool {
    /*
        1、优先检查以i开始的连续段是否为 CDATA，如果能够匹配到开头，那么尝试匹配到CDATA 的结尾，更新ii；否则，返回false；
        2、判断s[i]是否为 <，再看s[i+1]是不是/，来判断是左还是右：
        （1）左边：将其加入栈，待后续匹配；
        （2）右边：将其与栈顶元素进行匹配，若栈为空或者匹配不上，返回false；
        3、其他情况均是普通字符；
    */

        // regex r1 = regex("<!\\[CDATA\\[.*?\\]\\]>|t") ;
        // regex r2 = regex("<([A-Z]{1,9})>[^<]*</\\1>");
        // code = regex_replace(code, r1 , "-");
        // // cout << "code: " << code << endl;
        // string prev = "";
        // while(code != prev){
        //     prev = code;
        //     code = regex_replace(code, r2, "t");
        //     // cout << "code: " << code << endl;
        // }
        // return code == "t";
    r1 := regexp.MustCompile(`<!\[CDATA\[.*?\]\]>|t`)
    r2 := regexp.MustCompile(`<([A-Z]{1,9})>[^<]*</\\1>`)

    code = r1.ReplaceAllString(code, "-")
    
    fmt.Println("code: ", code)
    prev := ""

    for code != prev {
        prev = code
        code = r2.ReplaceAllString(code, "t")
        fmt.Println("code: ", code)
    }

    return code == "t"
}
