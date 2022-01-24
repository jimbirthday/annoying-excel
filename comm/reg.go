package comm

// RegUserName 字母开头，长度5-16字节，允许字母数字下划线
const RegUserName = `^[a-zA-Z][a-zA-Z0-9_]{4,15}$`

//RegPassword 以字母开头，长度在6~18之间，只能包含字母、数字和下划线
const RegPassword = `^[a-zA-Z]\w{5,17}$`

//RegNick 3~20任意字符
const RegNick = `^.{3,20}$`
